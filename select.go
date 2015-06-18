package optikon

import (
	"reflect"
	"strconv"
)

// Select will try to extract an internal part of dataIn identified
// by hierarchical path elements.
func Select(dataIn interface{}, path []string) (interface{}, error) {

	// Return myself if there's no subpath anymore (recursion termination).
	if len(path) == 0 {
		return dataIn, nil
	}

	// Otherwise we need to traverse into the first path element.
	subPath := path[0]
	typ := reflect.TypeOf(dataIn)
	val := reflect.ValueOf(dataIn)

	// Bail out if cannot traverse.
	if !isTraversable(typ.Kind()) {
		return nil, &KeyNotTraversableError{subPath}
	}

	switch typ.Kind() {
	case reflect.Map:
		// Get the value from the map keyed by the first path element.
		mapVal := val.MapIndex(reflect.ValueOf(subPath))
		if mapVal.IsValid() { // value found - subpath matched
			return Select(mapVal.Interface(), path[1:])
		}
	case reflect.Struct:
		// Iterate over object fields and see if there's a field whose json tag
		// matches the first element in the path.
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			tag := field.Tag.Get("json") // get json tag of this field
			if tag == "" {
				tag = field.Name // if no json tag found use field name
			}
			if tag == subPath { // subpath matched
				return Select(val.Field(i).Interface(), path[1:])
			}
		}
	case reflect.Ptr, reflect.Interface:
		if val.Elem().IsValid() {
			return Select(val.Elem().Interface(), path) // dereference and call again
		}
	case reflect.Array, reflect.Slice:
		// Here subPath must be an integer and a valid array index.
		if i, err1 := strconv.Atoi(subPath); err1 == nil {
			if i >= 0 && i < val.Len() { // valid index
				return Select(val.Index(i).Interface(), path[1:])
			}
		}
	}

	// subPath was not matched.
	return nil, &KeyNotFoundError{subPath}
}
