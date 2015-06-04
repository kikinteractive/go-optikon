package optikon

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdatePrimitive(t *testing.T) {
	//td := &TypeDeep{
	//	StrVal: "strVal",
	//	IntVal: 5,
	//}

	//err := UpdateJSON(td, []string{"strVal"}, "strVal1", UpdateOp)
	//if assert.NoError(t, err) {
	//	assert.Equal(t, "strVal1", td.StrVal)
	//}
}

/*
	MapMapVal      map[string]map[string]string  `json:"mapMapVal"`
	MapPtrMapVal   map[string]*map[string]string `json:"mapPtrMapVal"`
	MapSliceVal    map[string][]string           `json:"mapSliceVal"`
	MapArrVal      map[string][2]string          `json:"mapArrVal"`
	MapPtrSliceVal map[string]*[]string          `json:"mapPtrSliceVal"`
*/
func TestUpdateMap(t *testing.T) {
	strVal1 := "strVal1"
	strVal2 := "strVal2"
	mapVal := map[string]string{
		"key1": strVal1,
		"key2": strVal2,
	}
	mapPtrVal := map[string]*string{
		"key1": &strVal1,
		"key2": &strVal2,
	}
	mapMapVal := map[string]map[string]string{
		"key1": mapVal,
		"key2": mapVal,
	}
	mapPtrMapVal := map[string]*map[string]string{
		"key1": &mapVal,
		"key2": &mapVal,
	}
	sliceVal := []string{strVal1, strVal2}
	mapSliceVal := map[string][]string{
		"key1": sliceVal,
		"key2": sliceVal,
	}
	arrVal := [2]string{strVal1, strVal2}
	mapArrVal := map[string][2]string{
		"key1": arrVal,
		"key2": arrVal,
	}
	mapPtrSliceVal := map[string]*[]string{
		"key1": &sliceVal,
		"key2": &sliceVal,
	}

	td := &TypeDeep{}

	data, _ := json.Marshal(mapVal)
	err := UpdateJSON(td, []string{"mapVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapVal, td.MapVal)
	}

	err = UpdateJSON(td, []string{"ptrMapVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &mapVal, td.PtrMapVal)
	}

	err = UpdateJSON(td, []string{"mapPtrVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrVal, td.MapPtrVal)
	}

	data, _ = json.Marshal(mapMapVal)
	err = UpdateJSON(td, []string{"mapMapVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapVal, td.MapMapVal)
	}

	err = UpdateJSON(td, []string{"mapPtrMapVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrMapVal, td.MapPtrMapVal)
	}

	data, _ = json.Marshal(mapSliceVal)
	err = UpdateJSON(td, []string{"mapSliceVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapSliceVal, td.MapSliceVal)
	}

	err = UpdateJSON(td, []string{"mapArrVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapArrVal, td.MapArrVal)
	}

	data, _ = json.Marshal(mapPtrSliceVal)
	err = UpdateJSON(td, []string{"mapPtrSliceVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrSliceVal, td.MapPtrSliceVal)
	}

}

func TestUpdateSlice(t *testing.T) {
	strVal1 := "strVal1"
	strVal2 := "strVal2"

	sliceVal := []string{strVal1, strVal2}
	arrVal := [2]string{strVal1, strVal2}
	slicePtrVal := []*string{&strVal1, &strVal2}

	mapVal := map[string]string{
		"key1": strVal1,
		"key2": strVal2,
	}
	sliceMapVal := []map[string]string{mapVal, mapVal}
	slicePtrMapVal := []*map[string]string{&mapVal, &mapVal}
	sliceSliceVal := [][]string{sliceVal, sliceVal}
	slicePtrSliceVal := []*[]string{&sliceVal, &sliceVal}

	td := &TypeDeep{}

	data, _ := json.Marshal(sliceVal)
	err := UpdateJSON(td, []string{"sliceVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceVal, td.SliceVal)
	}

	err = UpdateJSON(td, []string{"arrVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrVal, td.ArrVal)
	}

	err = UpdateJSON(td, []string{"slicePtrVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrVal, td.SlicePtrVal)
	}

	data, _ = json.Marshal(sliceMapVal)
	err = UpdateJSON(td, []string{"sliceMapVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceMapVal, td.SliceMapVal)
	}

	err = UpdateJSON(td, []string{"slicePtrMapVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrMapVal, td.SlicePtrMapVal)
	}

	data, _ = json.Marshal(sliceSliceVal)
	err = UpdateJSON(td, []string{"sliceSliceVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceSliceVal, td.SliceSliceVal)
	}

	err = UpdateJSON(td, []string{"slicePtrSliceVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrSliceVal, td.SlicePtrSliceVal)
	}
}

func TestUpdateIntf(t *testing.T) {
	strVal := "strVal1"
	intVal := 5

	arrIntfVal := [2]interface{}{strVal, intVal}
	sliceIntfVal := []interface{}{intVal, strVal}
	mapIntfVal := map[string]interface{}{
		"key1": strVal,
		"key2": intVal,
	}
	mapArrIntfVal := map[string][2]interface{}{
		"key1": arrIntfVal,
		"key2": arrIntfVal,
	}
	mapSliceIntfVal := map[string][]interface{}{
		"key1": sliceIntfVal,
		"key2": sliceIntfVal,
	}
	mapPtrSliceIntfVal := map[string]*[]interface{}{
		"key1": &sliceIntfVal,
		"key2": &sliceIntfVal,
	}
	mapMapIntfVal := map[string]map[string]interface{}{
		"key1": mapIntfVal,
		"key2": mapIntfVal,
	}
	mapPtrMapIntfVal := map[string]*map[string]interface{}{
		"key1": &mapIntfVal,
		"key2": &mapIntfVal,
	}

	td := &TypeDeep{}

	data, _ := json.Marshal(arrIntfVal)
	err := UpdateJSON(td, []string{"arrIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrIntfVal, td.ArrIntfVal)
	}

	data, _ = json.Marshal(sliceIntfVal)
	err = UpdateJSON(td, []string{"sliceIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceIntfVal, td.SliceIntfVal)
	}

	data, _ = json.Marshal(mapIntfVal)
	err = UpdateJSON(td, []string{"mapIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapIntfVal, td.MapIntfVal)
	}

	data, _ = json.Marshal(mapArrIntfVal)
	err = UpdateJSON(td, []string{"mapArrIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapArrIntfVal, td.MapArrIntfVal)
	}

	data, _ = json.Marshal(mapSliceIntfVal)
	err = UpdateJSON(td, []string{"mapSliceIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		//assert.EqualValues(t, mapSliceIntfVal, td.MapSliceIntfVal)
	}

	data, _ = json.Marshal(mapPtrSliceIntfVal)
	err = UpdateJSON(td, []string{"mapPtrSliceIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		//assert.EqualValues(t, mapPtrSliceIntfVal, td.MapPtrSliceIntfVal)
	}

	data, _ = json.Marshal(mapMapIntfVal)
	err = UpdateJSON(td, []string{"mapMapIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		//assert.EqualValues(t, mapMapIntfVal, td.MapMapIntfVal)
	}

	data, _ = json.Marshal(mapPtrMapIntfVal)
	err = UpdateJSON(td, []string{"mapPtrMapIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		//assert.EqualValues(t, mapPtrMapIntfVal, td.MapPtrMapIntfVal)
	}

}
