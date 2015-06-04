package optikon

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func isTraversable(fieldKind reflect.Kind) bool {
	return fieldKind == reflect.Map || fieldKind == reflect.Struct ||
		fieldKind == reflect.Array || fieldKind == reflect.Slice ||
		fieldKind == reflect.Interface || fieldKind == reflect.Ptr
}

// UpdateJSON will try to perform the operation identified by opType on the source object dataIn, given traversal path and
// a json for a new/updated object.
func UpdateJSON(dataIn interface{}, path []string, dataJSON string, opType OpType) (err error) {

	ptrSrcVal := reflect.ValueOf(dataIn)
	srcVal := reflect.Indirect(ptrSrcVal)
	srcValType := srcVal.Type()

	//fmt.Println("++++++> ", ptrSrcVal, srcVal, srcValType, srcValType.Kind())

	// Update myself if there's no subpath (recursion termination).
	if len(path) == 0 {
		// Create a suitable concrete object to unmarshal to.
		dstVal := reflect.New(srcValType)
		// Try to unmarshal.
		err = json.Unmarshal([]byte(dataJSON), dstVal.Interface())
		if err != nil {
			return err
		}
		// Unmarshalled successfully, update the source object and leave.
		srcVal.Set(dstVal.Elem())
		return nil
	}

	// There's a subpath, see if we can drill down.
	subPath := path[0]
	subPathVal := reflect.ValueOf(subPath)

	// Identify the field indexed by subPath and try to drill down into it.
	// We can traverse into a struct, a map or an array.
	switch srcValType.Kind() {
	case reflect.Map:
		if srcVal.IsNil() { // uninited map
			if opType == DeleteOp {
				// Nothing to delete.
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
					// Otherwise, we cannot create an existing key.
					return &KeyExistsError{subPath}
				} else if opType == DeleteOp {
					// Alright, delete the entry and leave.
					srcVal.SetMapIndex(subPathVal, reflect.Value{})
					return nil
				}
			}
			// See if we can traverse into the value.
			if isTraversable(elKind) {
				// Drill down and update mapVal recursively.
				return UpdateJSON(mapVal.Interface(), path[1:], dataJSON, opType)
			}
		} else { // no such key in map
			if len(path) == 1 { // last element in the path
				// On this stage, we can only create a new map entry.
				// Updating and deleting will cause KeyNotFoundError.
				if opType == CreateOp {
					elType := srcValType.Elem()
					// Create a new map element.
					mapVal := reflect.New(elType.Elem())
					// Create a map if needed.
					if elType.Kind() == reflect.Map {
						mapVal.Set(reflect.MakeMap(elType))
					}
					// Update the newly created element.
					err := UpdateJSON(mapVal.Interface(), path[1:], dataJSON, opType)
					if err != nil {
						return err
					}
					// Alright, update the original map with the new element and leave.
					srcVal.SetMapIndex(subPathVal, mapVal)
					return nil
				}
			}
		}
	case reflect.Struct:
		// Iterate over object fields and see if there's a field whose json tag
		// matches the first element in the path.
		for i := 0; i < srcValType.NumField(); i++ {
			field := srcValType.Field(i)
			fieldKind := field.Type.Kind()
			tag := field.Tag.Get("json")
			if tag == "" {
				tag = field.Name // if no json tag found use field name
			}
			if tag == subPath { // matches the first path element
				if len(path) == 1 { // last element in the path
					if opType == CreateOp {
						// We can only append to a slice.
						if fieldKind == reflect.Slice {
							return UpdateJSON(srcVal.Field(i).Addr().Interface(), path, dataJSON, opType)
						}
						// Return error because we cannot create a struct field.
						return &KeyExistsError{subPath}
					} else if opType == DeleteOp {
						// Return error because we cannot delete a struct field.
						return fmt.Errorf("cannot delete a struct field")
					}
				}
				// Otherwise see if we can traverse into the value.
				if isTraversable(fieldKind) {
					return UpdateJSON(srcVal.Field(i).Addr().Interface(), path[1:], dataJSON, opType)
				}
			}
		}
	case reflect.Array, reflect.Slice:
		if srcVal.IsNil() { // uninited slice
			if opType == DeleteOp {
				// Nothing to delete.
				return &KeyNotFoundError{subPath}
			}
			// Otherwise, create an empty slice and continue.
			srcVal.Set(reflect.MakeSlice(srcValType, 0, 1))
		}
		if len(path) == 1 { // last element in the path
			if opType == CreateOp {
				// Can only append new elements to a slice.
				if srcValType.Kind() == reflect.Slice {
					sliceVal := reflect.New(srcValType.Elem())
					// Update the newly created element.
					err := UpdateJSON(sliceVal.Interface(), path[1:], dataJSON, opType)
					if err != nil {
						return err
					}
					// Append and replace with appended slice and leave.
					srcVal.Set(reflect.Append(srcVal, sliceVal.Elem()))
					return nil
				} else {
					return fmt.Errorf("cannot create array element")
				}
			} else if opType == DeleteOp {
				return fmt.Errorf("cannot delete array/slice")
			}
		}
		// Here subPath must be an integer and a valid array index.
		i, err1 := strconv.Atoi(subPath)
		if err1 != nil {
			return fmt.Errorf("path element [%s] must be an integer index into an array", subPath)
		}
		if i >= 0 && i < srcVal.Len() { // valid index
			return UpdateJSON(srcVal.Index(i).Addr().Interface(), path[1:], dataJSON, opType)
		}
	default:
		return fmt.Errorf("field referenced by [%s] must be traversable, %s given", subPath, srcValType.Kind())
	}

	return &KeyNotFoundError{subPath}
}
