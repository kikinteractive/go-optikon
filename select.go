package optikon

import (
	"fmt"
	"reflect"
	"strconv"
)

// Select will try to extract an internal part of dataIn identified by hierarchical path elements.
func Select(dataIn interface{}, path []string) (partOut interface{}, err error) {

	// Return myself if there's no subpath anymore (recursion termination).
	if len(path) == 0 {
		partOut = dataIn
		return
	}

	subPath := path[0]
	typ := reflect.TypeOf(dataIn)
	val := reflect.ValueOf(dataIn)

	// There's a subpath, so we need to drill down. See if the value is traversable.
	switch typ.Kind() {
	case reflect.Map:
		// Get the value from the map keyed by the first path element.
		mapVal := val.MapIndex(reflect.ValueOf(subPath))
		if mapVal.IsValid() { // value found - subpath matched
			return Select(mapVal.Interface(), path[1:])
		}
	case reflect.Struct:
		// Iterate over object fields and see if there's a field whose json tag matches the first element in the path.
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			tag := field.Tag.Get("json") // get json tag of this field
			if tag == subPath {          // subpath matched
				return Select(val.Field(i).Interface(), path[1:])
			}
		}
	case reflect.Ptr, reflect.Interface:
		if !val.Elem().IsValid() { // cannot traverse further
			err = fmt.Errorf("field referenced by [%s] is invalid pointer/interface", subPath)
			return
		}
		return Select(val.Elem().Interface(), path) // dereference and call recursively
	case reflect.Array, reflect.Slice:
		// Here subPath must be an integer and a valid array index.
		i, err1 := strconv.Atoi(subPath)
		if err1 != nil {
			err = fmt.Errorf("path element [%s] must be an integer index into an array", subPath)
			return
		}
		if i >= 0 && i < val.Len() { // valid index
			return Select(val.Index(i).Interface(), path[1:])
		}
	default:
		err = fmt.Errorf("field referenced by [%s] must be traversable (struct/map/array/slice/ptr/interface), %s given", subPath, typ.Kind())
		return
	}

	// subPath was not matched or could not traverse into the corresponding field.
	err = &KeyNotFoundError{subPath}
	return
}
