// Implemenentation of partial create/update/delete functionality for arbitrary Go
// structures of any depth using relative path selectors.
package optikon

import (
	"encoding/json"
	"reflect"
	"strconv"
)

// UpdateJSON will try to perform the operation identified by opType on the source
// object dataIn, given traversal path and a json for a new/updated object.
func UpdateJSON(dataIn interface{}, path []string, dataJSON json.RawMessage, opType OpType) error {
	ptrSrcVal := reflect.ValueOf(dataIn)
	srcVal := reflect.Indirect(ptrSrcVal)
	srcValType := srcVal.Type()

	//fmt.Println("++++++> ", path, ptrSrcVal, srcVal, srcValType, srcValType.Kind(), "canSet=", srcVal.CanSet())

	// Update myself if there's no subpath (recursion termination).
	if len(path) == 0 {
		// Create a suitable concrete object to unmarshal to.
		dstVal := reflect.New(srcValType)
		// Try to unmarshal.
		if err := json.Unmarshal(dataJSON, dstVal.Interface()); err != nil {
			return err
		}
		// Unmarshalled successfully, update the source object.
		srcVal.Set(dstVal.Elem())
		return nil
	}

	// Otherwise we need to traverse into the first path element.

	switch srcValType.Kind() {
	case reflect.Map:
		return traverseMap(srcVal, path, dataJSON, opType)
	case reflect.Struct:
		return traverseStruct(srcVal, path, dataJSON, opType)
	case reflect.Array, reflect.Slice:
		return traverseArraySlice(srcVal, path, dataJSON, opType)
	case reflect.Ptr, reflect.Interface:
		// srcVal is already dereferenced, just call recursively.
		return UpdateJSON(srcVal.Interface(), path, dataJSON, opType)
	default:
		return &KeyNotTraversableError{path[0]}
	}
}

// traverseStruct will try to find if there's a struct field indexed by the first path element.
// If there is, and there are more path elements, it will call UpdateJSON for recursive update.
// If this is the last path element, it will try to handle it according to the OpType requested.
func traverseStruct(srcVal reflect.Value, path []string, dataJSON json.RawMessage, opType OpType) error {
	srcValType := srcVal.Type()
	subPath := path[0]
	// Iterate over object fields and see if there's a field whose json tag
	// matches the first element in the path.
	for i := 0; i < srcValType.NumField(); i++ {
		fieldVal := srcVal.Field(i)
		fieldMeta := srcValType.Field(i)
		fieldKind := fieldMeta.Type.Kind()
		tag := fieldMeta.Tag.Get("json")
		if tag == "" {
			tag = fieldMeta.Name // if no json tag found use field name
		}
		if tag == subPath { // matches the first path element
			if len(path) == 1 { // last element in the path
				switch opType {
				case CreateOp:
					// We can append to a slice.
					if fieldKind == reflect.Slice {
						sliceVal := reflect.New(fieldVal.Type())
						// Update the newly created element.
						if err := UpdateJSON(sliceVal.Interface(), path[1:], dataJSON, opType); err != nil {
							return err
						}
						fieldVal.Set(reflect.AppendSlice(fieldVal, sliceVal.Elem()))
						return nil
					}
					// Otherwise we cannot create a struct field.
					return &KeyExistsError{subPath}
				case SetOp:
					// We can set to a map.
					if fieldKind == reflect.Map {
						mapVal := reflect.New(fieldVal.Type())
						// Update the newly created element.
						if err := UpdateJSON(mapVal.Interface(), path[1:], dataJSON, opType); err != nil {
							return err
						}
						mergeMaps(fieldVal, mapVal.Elem())
						return nil
					}
					// Otherwise fall through and try to update.
				case DeleteOp:
					// We cannot delete a struct field.
					return &OperationForbiddenError{key: subPath, keyType: srcValType, operation: opType}
				}
			}
			if fieldVal.CanAddr() {
				return UpdateJSON(fieldVal.Addr().Interface(), path[1:], dataJSON, opType)
			}
			// Try to update not addressable field directly.
			return UpdateJSON(fieldVal.Interface(), path[1:], dataJSON, opType)
		}
	}
	return &KeyNotFoundError{subPath}
}

// traverseArraySlice will try to find if there's an array element indexed by the first path element.
// If there is, and there are more path elements, it will call UpdateJSON for recursive update.
// If this is the last path element, it will try to handle it according to the OpType requested.
func traverseArraySlice(srcVal reflect.Value, path []string, dataJSON json.RawMessage, opType OpType) error {
	srcValType := srcVal.Type()
	subPath := path[0]
	if srcValType.Kind() == reflect.Slice && srcVal.IsNil() { // uninited slice
		if opType == DeleteOp || opType == UpdateOp {
			// Nothing to delete or update, bail out.
			return &KeyNotFoundError{subPath}
		}
		// Create an empty slice and continue.
		srcVal.Set(reflect.MakeSlice(srcValType, 0, 1))
	}
	// Check that subPath is an integer and a valid array index.
	var i int
	var err error
	if i, err = strconv.Atoi(subPath); err != nil {
		return &KeyNotFoundError{key: subPath}
	}
	if i < 0 || i >= srcVal.Len() {
		return &KeyNotFoundError{key: subPath}
	}
	sliceElem := srcVal.Index(i)
	if len(path) == 1 { // last element in the path
		switch opType {
		case CreateOp:
			// We can append to a slice.
			if sliceElem.Type().Kind() == reflect.Slice {
				sliceVal := reflect.New(srcValType.Elem())
				// Update the newly created element.
				if err := UpdateJSON(sliceVal.Interface(), path[1:], dataJSON, opType); err != nil {
					return err
				}
				// Alright, append and replace with appended slice.
				sliceElem.Set(reflect.AppendSlice(sliceElem, sliceVal.Elem()))
				return nil
			}
			return &OperationForbiddenError{key: subPath, keyType: srcValType, operation: opType}
		case SetOp:
			// We can set to a map.
			if sliceElem.Type().Kind() == reflect.Map {
				mapVal := reflect.New(srcValType.Elem())
				// Update the newly created element.
				if err := UpdateJSON(mapVal.Interface(), path[1:], dataJSON, opType); err != nil {
					return err
				}
				mergeMaps(sliceElem, mapVal.Elem())
				return nil
			}
			// Otherwise fall through and try to update.
		case DeleteOp:
			return &OperationForbiddenError{key: subPath, keyType: srcValType, operation: opType}
		}
	}
	return UpdateJSON(sliceElem.Addr().Interface(), path[1:], dataJSON, opType)
}

// traverseMap will try to find if there's an map value indexed by the first path element.
// If there is, and there are more path elements, it will call UpdateJSON for recursive update.
// If this is the last path element, it will try to handle it according to the OpType requested.
func traverseMap(srcVal reflect.Value, path []string, dataJSON json.RawMessage, opType OpType) error {
	srcValType := srcVal.Type()
	subPath := path[0]
	subPathVal := reflect.ValueOf(subPath)
	if srcVal.IsNil() { // uninited map
		if opType == DeleteOp || opType == UpdateOp {
			// Nothing to delete or update, bail out.
			return &KeyNotFoundError{subPath}
		}
		// Otherwise, create an empty map and continue.
		srcVal.Set(reflect.MakeMap(srcValType))
	}
	// Check if the first path element exists as a key in this map.
	mapVal := srcVal.MapIndex(subPathVal)
	if mapVal.IsValid() { // key exists in map
		if len(path) == 1 { // last element in the path
			switch opType {
			case CreateOp:
				// We can append to a slice.
				if mapVal.Kind() == reflect.Slice {
					sliceVal := reflect.New(mapVal.Type())
					// Update the newly created element.
					if err := UpdateJSON(sliceVal.Interface(), path[1:], dataJSON, opType); err != nil {
						return err
					}
					// Alright, append and replace with appended slice.
					newSlice := reflect.AppendSlice(reflect.Indirect(mapVal), reflect.Indirect(sliceVal.Elem()))
					newMapVal := reflect.New(srcValType.Elem())
					newMapVal.Elem().Set(newSlice)
					// Replace the original map with the new element.
					srcVal.SetMapIndex(subPathVal, newMapVal.Elem())
					return nil
				}
				// We cannot create an existing key.
				return &KeyExistsError{subPath}
			case SetOp:
				// We can set to a map.
				if mapVal.Kind() == reflect.Map {
					newMapVal := reflect.New(mapVal.Type())
					// Update the newly created element.
					if err := UpdateJSON(newMapVal.Interface(), path[1:], dataJSON, opType); err != nil {
						return err
					}
					mergeMaps(mapVal, newMapVal.Elem())
					return nil
				}
				// Otherwise fall through and try to update.
			case DeleteOp:
				// Alright, delete the entry and leave.
				srcVal.SetMapIndex(subPathVal, reflect.Value{})
				return nil
			}
		}
		// Drill down and update mapVal recursively.
		// The map element is not settable, create a new one, update and replace.
		newMapVal := reflect.New(srcValType.Elem())
		newMapVal.Elem().Set(mapVal)
		if err := UpdateJSON(newMapVal.Interface(), path[1:], dataJSON, opType); err != nil {
			return err
		}
		srcVal.SetMapIndex(subPathVal, newMapVal.Elem())
		return nil
	}
	// No such key in map.
	if len(path) == 1 { // last element in the path
		// On this stage, we can only create a new map entry.
		// Updating and deleting will cause KeyNotFoundError.
		if opType == CreateOp || opType == SetOp {
			// Create a new map element.
			newMapVal := reflect.New(srcValType.Elem())
			// Update the newly created element.
			if err := UpdateJSON(newMapVal.Interface(), path[1:], dataJSON, opType); err != nil {
				return err
			}
			// Alright, update the original map with the new element.
			srcVal.SetMapIndex(subPathVal, newMapVal.Elem())
			return nil
		}
	}
	return &KeyNotFoundError{subPath}
}

// mergeMaps merges the second map into the first one. If the second map has the same key,
// its value will replace the original value in the first map. Otherwise, a new key/value
// pair will be created in the first map.
func mergeMaps(m1 reflect.Value, m2 reflect.Value) reflect.Value {
	for _, k := range m2.MapKeys() {
		m1.SetMapIndex(k, m2.MapIndex(k))
	}
	return m1
}
