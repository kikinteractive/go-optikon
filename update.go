package optikon

import (
	"encoding/json"
	"reflect"
	"strconv"
)

// UpdateJSON will try to perform the operation identified by opType on the source
// object dataIn, given traversal path and a json for a new/updated object.
func UpdateJSON(dataIn interface{}, path []string, dataJSON string, opType OpType) error {

	ptrSrcVal := reflect.ValueOf(dataIn)
	srcVal := reflect.Indirect(ptrSrcVal)
	srcValType := srcVal.Type()

	//fmt.Println("++++++> ", path, ptrSrcVal, srcVal, srcValType, srcValType.Kind(), "canSet=", srcVal.CanSet())

	// Update myself if there's no subpath (recursion termination).
	if len(path) == 0 {
		// Create a suitable concrete object to unmarshal to.
		dstVal := reflect.New(srcValType)
		// Try to unmarshal.
		if err := json.Unmarshal([]byte(dataJSON), dstVal.Interface()); err != nil {
			return err
		}
		// Unmarshalled successfully, update the source object.
		srcVal.Set(dstVal.Elem())
		return nil
	}

	// Otherwise we need to traverse into the first path element.

	// Bail out if cannot traverse.
	if !isTraversable(srcValType.Kind()) {
		return &KeyNotTraversableError{path[0]}
	}

	switch srcValType.Kind() {
	case reflect.Map:
		return traverseMap(srcVal, path, dataJSON, opType)
	case reflect.Struct:
		return traverseStruct(srcVal, path, dataJSON, opType)
	case reflect.Array, reflect.Slice:
		return traverseArraySlice(srcVal, path, dataJSON, opType)
	case reflect.Ptr, reflect.Interface:
		// srcVal already dereferenced, just call recursively.
		return UpdateJSON(srcVal.Interface(), path, dataJSON, opType)
	}

	return nil
}

func traverseStruct(srcVal reflect.Value, path []string, dataJSON string, opType OpType) error {
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
				if opType == CreateOp {
					// We can only append to a slice.
					if fieldKind == reflect.Slice {
						if fieldVal.CanAddr() {
							return UpdateJSON(fieldVal.Addr().Interface(), path[1:], dataJSON, opType)
						}
						return UpdateJSON(fieldVal.Interface(), path[1:], dataJSON, opType)
					}
					// We cannot create a struct field.
					return &KeyExistsError{subPath}
				} else if opType == DeleteOp {
					// We cannot delete a struct field.
					return &OperationForbiddenError{key: subPath, keyType: srcValType, operation: opType}
				}
			}
			// Otherwise see if we can traverse into the value.
			if isTraversable(fieldKind) {
				if fieldVal.CanAddr() {
					return UpdateJSON(fieldVal.Addr().Interface(), path[1:], dataJSON, opType)
				}
				// Cannot address this field (a map)
				return UpdateJSON(fieldVal.Interface(), path[1:], dataJSON, opType)
			}
		}
	}
	return &KeyNotFoundError{subPath}
}

func traverseArraySlice(srcVal reflect.Value, path []string, dataJSON string, opType OpType) error {
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
	if len(path) == 1 { // last element in the path
		if opType == CreateOp {
			// Can only append new elements to a slice.
			if srcValType.Kind() == reflect.Slice {
				sliceVal := reflect.New(srcValType.Elem())
				// Update the newly created element.
				if err := UpdateJSON(sliceVal.Interface(), path[1:], dataJSON, opType); err != nil {
					return err
				}
				// Alright, append and replace with appended slice.
				srcVal.Set(reflect.Append(srcVal, sliceVal.Elem()))
				return nil
			}
			return &OperationForbiddenError{key: subPath, keyType: srcValType, operation: opType}
		} else if opType == DeleteOp {
			return &OperationForbiddenError{key: subPath, keyType: srcValType, operation: opType}
		}
	}
	return UpdateJSON(srcVal.Index(i).Addr().Interface(), path[1:], dataJSON, opType)
}

func traverseMap(srcVal reflect.Value, path []string, dataJSON string, opType OpType) error {
	srcValType := srcVal.Type()
	subPath := path[0]
	subPathVal := reflect.ValueOf(subPath)
	if srcVal.IsNil() { // uninited map
		if opType == DeleteOp || opType == UpdateOp {
			// Nothing to delete, bail out.
			return &KeyNotFoundError{subPath}
		}
		// Otherwise, create an empty map and continue.
		srcVal.Set(reflect.MakeMap(srcValType))
	}
	// Check if the first path element exists as a key in this map.
	mapVal := srcVal.MapIndex(subPathVal)
	if mapVal.IsValid() { // key exists in map
		elKind := reflect.Indirect(mapVal).Kind()
		if len(path) == 1 { // last element in the path
			if opType == CreateOp {
				if elKind == reflect.Slice {
					// DOES NOT GET HERE!!!!!! TEST!!!!!!!
					return UpdateJSON(mapVal.Interface(), path, dataJSON, opType)
				}
				// We cannot create an existing key.
				return &KeyExistsError{subPath}
			} else if opType == DeleteOp {
				// Alright, delete the entry and leave.
				srcVal.SetMapIndex(subPathVal, reflect.Value{})
				return nil
			} else { // update
				// Cannot set map entry value directly (error: "Set using unaddressable value").
				// Instead, create a new map value and fill it, then replace the old one.
				newMapVal := reflect.New(srcValType.Elem())
				// Update the newly created element.
				if err := UpdateJSON(newMapVal.Interface(), path[1:], dataJSON, opType); err != nil {
					return err
				}
				// Replace the original map with the new element.
				srcVal.SetMapIndex(subPathVal, newMapVal.Elem())
				return nil
			}
		}
		// See if we can traverse into the value.
		if isTraversable(elKind) {
			// Drill down and update mapVal recursively.
			if mapVal.CanSet() {
				return UpdateJSON(mapVal.Interface(), path[1:], dataJSON, opType)
			}
			// Otherwise, if the map element is not settable, create a new one, update and replace.
			newMapVal := reflect.New(srcValType.Elem())
			newMapVal.Elem().Set(mapVal)
			if err := UpdateJSON(newMapVal.Interface(), path[1:], dataJSON, opType); err != nil {
				return err
			}
			srcVal.SetMapIndex(subPathVal, newMapVal.Elem())
			return nil
		}
	} else { // no such key in map
		if len(path) == 1 { // last element in the path
			// On this stage, we can only create a new map entry.
			// Updating and deleting will cause KeyNotFoundError.
			if opType == CreateOp {
				elType := srcValType.Elem()
				// Create a new map element.
				mapVal := reflect.New(elType)
				// Update the newly created element.
				if err := UpdateJSON(mapVal.Interface(), path[1:], dataJSON, opType); err != nil {
					return err
				}
				// Alright, update the original map with the new element.
				srcVal.SetMapIndex(subPathVal, mapVal.Elem())
				return nil
			}
		}
	}
	return &KeyNotFoundError{subPath}
}
