package optikon

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePrimitive(t *testing.T) {
	strVal := "strVal1"
	intVal := 5
	td := &TypeDeep{}

	data, _ := json.Marshal(strVal)
	err := UpdateJSON(td, []string{"strVal"}, data, CreateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyExistsError{}, err)
		assert.EqualError(t, err, "key exists: strVal")
		assert.Equal(t, "strVal", err.(*KeyExistsError).Key())
	}

	err = UpdateJSON(td, []string{"mapVal", "key1"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.Equal(t, strVal, td.MapVal["key1"])
	}

	data, _ = json.Marshal(intVal)
	err = UpdateJSON(td, []string{"intVal"}, data, CreateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyExistsError{}, err)
	}
}

func TestUpdatePrimitive(t *testing.T) {
	strVal1 := "strVal1"
	strVal2 := "strVal2"
	intVal1 := 5
	intVal2 := 15
	td := &TypeDeep{
		StrVal:   strVal1,
		NoTagVal: strVal1,
		IntVal:   intVal1,
		MapVal: map[string]string{
			"key1": strVal1,
		},
		MapIntVal: map[string]int{
			"key1": intVal1,
		},
		SliceVal: []string{strVal1},
	}

	data, _ := json.Marshal(strVal2)
	err := UpdateJSON(td, []string{"strVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.Equal(t, strVal2, td.StrVal)
	}

	err = UpdateJSON(td, []string{"NoTagVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.Equal(t, strVal2, td.NoTagVal)
	}

	err = UpdateJSON(td, []string{"mapVal", "key1"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.Equal(t, strVal2, td.MapVal["key1"])
	}

	err = UpdateJSON(td, []string{"sliceVal", "0"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.Equal(t, strVal2, td.SliceVal[0])
	}

	data, _ = json.Marshal(intVal2)
	err = UpdateJSON(td, []string{"intVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.Equal(t, intVal2, td.IntVal)
	}

	err = UpdateJSON(td, []string{"mapIntVal", "key1"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.Equal(t, intVal2, td.MapIntVal["key1"])
	}
}

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
	err := UpdateJSON(td, []string{"mapVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapVal, td.MapVal)
	}

	err = UpdateJSON(td, []string{"ptrMapVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &mapVal, td.PtrMapVal)
	}

	err = UpdateJSON(td, []string{"mapPtrVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrVal, td.MapPtrVal)
	}

	data, _ = json.Marshal(mapMapVal)
	err = UpdateJSON(td, []string{"mapMapVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapVal, td.MapMapVal)
	}

	err = UpdateJSON(td, []string{"mapPtrMapVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrMapVal, td.MapPtrMapVal)
	}

	data, _ = json.Marshal(mapSliceVal)
	err = UpdateJSON(td, []string{"mapSliceVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapSliceVal, td.MapSliceVal)
	}

	err = UpdateJSON(td, []string{"mapArrVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapArrVal, td.MapArrVal)
	}

	data, _ = json.Marshal(mapPtrSliceVal)
	err = UpdateJSON(td, []string{"mapPtrSliceVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrSliceVal, td.MapPtrSliceVal)
	}

	// Second level.
	newMapVal := map[string]string{
		"key1": "newStrVal1",
		"key2": "newStrVal2",
	}
	data, _ = json.Marshal(newMapVal)
	err = UpdateJSON(td, []string{"mapMapVal", "key1"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, newMapVal, td.MapMapVal["key1"])
	}

	err = UpdateJSON(td, []string{"mapPtrMapVal", "key1"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &newMapVal, td.MapPtrMapVal["key1"])
	}
}

func TestUpdateSlice(t *testing.T) {
	strVal0 := "strVal0"
	strVal1 := "strVal1"
	strVal2 := "strVal2"

	sliceVal0 := []string{strVal0}
	sliceVal := []string{strVal1, strVal2}
	arrVal := [2]string{strVal1, strVal2}
	slicePtrVal := []*string{&strVal1, &strVal2}

	mapVal := map[string]string{
		"key1": strVal1,
		"key2": strVal2,
	}

	mapSliceVal := map[string][]string{"key1": sliceVal0}
	mapPtrSliceVal := map[string]*[]string{"key1": &sliceVal0}

	sliceMapVal := []map[string]string{mapVal, mapVal}
	slicePtrMapVal := []*map[string]string{&mapVal, &mapVal}
	sliceSliceVal := [][]string{sliceVal, sliceVal}
	slicePtrSliceVal := []*[]string{&sliceVal, &sliceVal}

	td := &TypeDeep{
		MapSliceVal:    mapSliceVal,
		MapPtrSliceVal: mapPtrSliceVal,
	}

	data, _ := json.Marshal(arrVal)
	err := UpdateJSON(td, []string{"arrVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrVal, td.ArrVal)
	}

	data, _ = json.Marshal(sliceVal)
	err = UpdateJSON(td, []string{"sliceVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceVal, td.SliceVal)
	}

	err = UpdateJSON(td, []string{"slicePtrVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrVal, td.SlicePtrVal)
	}

	err = UpdateJSON(td, []string{"mapSliceVal", "key1"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceVal, td.MapSliceVal["key1"])
	}

	err = UpdateJSON(td, []string{"mapPtrSliceVal", "key1"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &sliceVal, td.MapPtrSliceVal["key1"])
	}

	data, _ = json.Marshal(sliceMapVal)
	err = UpdateJSON(td, []string{"sliceMapVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceMapVal, td.SliceMapVal)
	}

	err = UpdateJSON(td, []string{"slicePtrMapVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrMapVal, td.SlicePtrMapVal)
	}

	data, _ = json.Marshal(sliceSliceVal)
	err = UpdateJSON(td, []string{"sliceSliceVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceSliceVal, td.SliceSliceVal)
	}

	err = UpdateJSON(td, []string{"slicePtrSliceVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrSliceVal, td.SlicePtrSliceVal)
	}
}

func TestUpdateIntf(t *testing.T) {
	strVal := "strVal1"
	intVal := 5.0

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
	err := UpdateJSON(td, []string{"arrIntfVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrIntfVal, td.ArrIntfVal)
	}

	data, _ = json.Marshal(sliceIntfVal)
	err = UpdateJSON(td, []string{"sliceIntfVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceIntfVal, td.SliceIntfVal)
	}

	data, _ = json.Marshal(mapIntfVal)
	err = UpdateJSON(td, []string{"mapIntfVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapIntfVal, td.MapIntfVal)
	}

	data, _ = json.Marshal(mapArrIntfVal)
	err = UpdateJSON(td, []string{"mapArrIntfVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapArrIntfVal, td.MapArrIntfVal)
	}

	data, _ = json.Marshal(mapSliceIntfVal)
	err = UpdateJSON(td, []string{"mapSliceIntfVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapSliceIntfVal, td.MapSliceIntfVal)
	}

	data, _ = json.Marshal(mapPtrSliceIntfVal)
	err = UpdateJSON(td, []string{"mapPtrSliceIntfVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrSliceIntfVal, td.MapPtrSliceIntfVal)
	}

	data, _ = json.Marshal(mapMapIntfVal)
	err = UpdateJSON(td, []string{"mapMapIntfVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapIntfVal, td.MapMapIntfVal)
	}

	data, _ = json.Marshal(mapPtrMapIntfVal)
	err = UpdateJSON(td, []string{"mapPtrMapIntfVal"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrMapIntfVal, td.MapPtrMapIntfVal)
	}

	// Second level.
	newSliceIntfVal := []interface{}{6.0, 7.0, "newStrVal"}
	data, _ = json.Marshal(newSliceIntfVal)
	err = UpdateJSON(td, []string{"mapSliceIntfVal", "key1"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, newSliceIntfVal, td.MapSliceIntfVal["key1"])
	}

	err = UpdateJSON(td, []string{"mapPtrSliceIntfVal", "key1"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &newSliceIntfVal, td.MapPtrSliceIntfVal["key1"])
	}

	newMapIntfVal := map[string]interface{}{
		"key3": 6.0,
		"key4": 7.0,
		"key5": "newStrVal",
	}
	data, _ = json.Marshal(newMapIntfVal)
	err = UpdateJSON(td, []string{"mapMapIntfVal", "key1"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, newMapIntfVal, td.MapMapIntfVal["key1"])
	}

	err = UpdateJSON(td, []string{"mapPtrMapIntfVal", "key1"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &newMapIntfVal, td.MapPtrMapIntfVal["key1"])
	}

}

func TestUpdateDeep(t *testing.T) {
	strVal1 := "strVal1"
	strVal2 := "strVal2"
	intVal := 5

	mapVal := map[string]string{
		"key1": strVal1,
		"key2": strVal2,
	}

	td1 := TypeDeep{
		StrVal:    strVal1,
		IntVal:    intVal,
		MapVal:    mapVal,
		PtrMapVal: &mapVal,
	}

	sliceDeep := []TypeDeep{td1, td1}
	arrDeep := [2]TypeDeep{td1, td1}
	slicePtrDeep := []*TypeDeep{&td1, &td1}
	arrPtrDeep := [2]*TypeDeep{&td1, &td1}

	mapDeep := map[string]TypeDeep{
		"key1": td1,
		"key2": td1,
	}
	mapPtrDeep := map[string]*TypeDeep{
		"key1": &td1,
		"key2": &td1,
	}
	mapArrDeep := map[string][2]TypeDeep{
		"key1": arrDeep,
		"key2": arrDeep,
	}
	mapSlicePtrDeep := map[string][]*TypeDeep{
		"key1": slicePtrDeep,
		"key2": slicePtrDeep,
	}
	mapMapDeep := map[string]map[string]TypeDeep{
		"key1": mapDeep,
		"key2": mapDeep,
	}
	mapPtrMapDeep := map[string]*map[string]TypeDeep{
		"key1": &mapDeep,
		"key2": &mapDeep,
	}
	mapMapPtrDeep := map[string]map[string]*TypeDeep{
		"key1": mapPtrDeep,
		"key2": mapPtrDeep,
	}

	td := &TypeDeep{
		PtrDeep:         &td1,
		SliceDeep:       sliceDeep,
		SlicePtrDeep:    slicePtrDeep,
		ArrPtrDeep:      arrPtrDeep,
		MapDeep:         mapDeep,
		MapPtrDeep:      mapPtrDeep,
		MapArrDeep:      mapArrDeep,
		MapSlicePtrDeep: mapSlicePtrDeep,
		MapMapDeep:      mapMapDeep,
		MapPtrMapDeep:   mapPtrMapDeep,
		MapMapPtrDeep:   mapMapPtrDeep,
	}

	// One level

	td1.StrVal = "strNew"
	data, _ := json.Marshal(td1)
	err := UpdateJSON(td, []string{"ptrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep)
	}

	// Two levels

	sliceDeep = []TypeDeep{td1, td1}
	data, _ = json.Marshal(sliceDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "sliceDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceDeep, td.PtrDeep.SliceDeep)
	}

	slicePtrDeep = []*TypeDeep{&td1, &td1}
	data, _ = json.Marshal(slicePtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "slicePtrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrDeep, td.PtrDeep.SlicePtrDeep)
	}

	arrPtrDeep = [2]*TypeDeep{&td1, &td1}
	data, _ = json.Marshal(arrPtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "arrPtrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrPtrDeep, td.PtrDeep.ArrPtrDeep)
	}

	mapDeep = map[string]TypeDeep{
		"key3": td1,
		"key4": td1,
	}
	data, _ = json.Marshal(mapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapDeep, td.PtrDeep.MapDeep)
	}

	mapPtrDeep = map[string]*TypeDeep{
		"key3": &td1,
		"key4": &td1,
	}
	data, _ = json.Marshal(mapPtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapPtrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrDeep, td.PtrDeep.MapPtrDeep)
	}

	mapArrDeep = map[string][2]TypeDeep{
		"key3": arrDeep,
		"key4": arrDeep,
	}
	data, _ = json.Marshal(mapArrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapArrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapArrDeep, td.PtrDeep.MapArrDeep)
	}

	mapSlicePtrDeep = map[string][]*TypeDeep{
		"key3": slicePtrDeep,
		"key4": slicePtrDeep,
	}
	data, _ = json.Marshal(mapSlicePtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapSlicePtrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapSlicePtrDeep, td.PtrDeep.MapSlicePtrDeep)
	}

	mapMapDeep = map[string]map[string]TypeDeep{
		"key3": mapDeep,
		"key4": mapDeep,
	}
	data, _ = json.Marshal(mapMapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapMapDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapDeep, td.PtrDeep.MapMapDeep)
	}

	mapPtrMapDeep = map[string]*map[string]TypeDeep{
		"key3": &mapDeep,
		"key4": &mapDeep,
	}
	data, _ = json.Marshal(mapPtrMapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapPtrMapDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrMapDeep, td.PtrDeep.MapPtrMapDeep)
	}

	mapMapPtrDeep = map[string]map[string]*TypeDeep{
		"key3": mapPtrDeep,
		"key4": mapPtrDeep,
	}
	data, _ = json.Marshal(mapMapPtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapMapPtrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapPtrDeep, td.PtrDeep.MapMapPtrDeep)
	}

	// Three levels

	td1.StrVal = "strNewNew"
	data, _ = json.Marshal(td1)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, td1, td.PtrDeep.MapDeep["key3"])
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapPtrDeep", "key3"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapPtrDeep["key3"])
	}

	arrDeep = [2]TypeDeep{td1, td1}
	data, _ = json.Marshal(arrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapArrDeep", "key3"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrDeep, td.PtrDeep.MapArrDeep["key3"])
	}

	slicePtrDeep = []*TypeDeep{&td1, &td1}
	data, _ = json.Marshal(slicePtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapSlicePtrDeep", "key3"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrDeep, td.PtrDeep.MapSlicePtrDeep["key3"])
	}

	mapDeep = map[string]TypeDeep{
		"key3": td1,
		"key4": td1,
	}
	data, _ = json.Marshal(mapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapMapDeep", "key3"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapDeep, td.PtrDeep.MapMapDeep["key3"])
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapPtrMapDeep", "key3"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &mapDeep, td.PtrDeep.MapPtrMapDeep["key3"])
	}

	mapPtrDeep = map[string]*TypeDeep{
		"key3": &td1,
		"key4": &td1,
	}
	err = UpdateJSON(td, []string{"ptrDeep", "mapMapPtrDeep", "key3"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrDeep, td.PtrDeep.MapMapPtrDeep["key3"])
	}

	// Four levels

	td1.StrVal = "strNewNewNew"

	data, _ = json.Marshal(td1)
	err = UpdateJSON(td, []string{"ptrDeep", "sliceDeep", "0", "ptrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.SliceDeep[0].PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "slicePtrDeep", "0", "ptrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.SlicePtrDeep[0].PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "arrPtrDeep", "0", "ptrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.ArrPtrDeep[0].PtrDeep)
	}

	data, _ = json.Marshal(td1)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "ptrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapDeep["key3"].PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapPtrDeep", "key3", "ptrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapPtrDeep["key3"].PtrDeep)
	}

	sliceDeep = []TypeDeep{td1, td1}
	data, _ = json.Marshal(sliceDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "sliceDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceDeep, td.PtrDeep.MapDeep["key3"].SliceDeep)
	}

	slicePtrDeep = []*TypeDeep{&td1, &td1}
	data, _ = json.Marshal(slicePtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "slicePtrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrDeep, td.PtrDeep.MapDeep["key3"].SlicePtrDeep)
	}

	arrPtrDeep = [2]*TypeDeep{&td1, &td1}
	data, _ = json.Marshal(arrPtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "arrPtrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrPtrDeep, td.PtrDeep.MapDeep["key3"].ArrPtrDeep)
	}

	mapDeep = map[string]TypeDeep{
		"key5": td1,
		"key6": td1,
	}
	data, _ = json.Marshal(mapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "mapDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapDeep, td.PtrDeep.MapDeep["key3"].MapDeep)
	}

	mapPtrDeep = map[string]*TypeDeep{
		"key3": &td1,
		"key4": &td1,
	}
	data, _ = json.Marshal(mapPtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "mapPtrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrDeep, td.PtrDeep.MapDeep["key3"].MapPtrDeep)
	}

	arrDeep = [2]TypeDeep{td1, td1}
	mapArrDeep = map[string][2]TypeDeep{
		"key3": arrDeep,
		"key4": arrDeep,
	}
	data, _ = json.Marshal(mapArrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "mapArrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapArrDeep, td.PtrDeep.MapDeep["key3"].MapArrDeep)
	}

	slicePtrDeep = []*TypeDeep{&td1, &td1}
	data, _ = json.Marshal(slicePtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "slicePtrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrDeep, td.PtrDeep.MapDeep["key3"].SlicePtrDeep)
	}

	mapMapDeep = map[string]map[string]TypeDeep{
		"key3": mapDeep,
		"key4": mapDeep,
	}
	data, _ = json.Marshal(mapMapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "mapMapDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapDeep, td.PtrDeep.MapDeep["key3"].MapMapDeep)
	}

	mapPtrMapDeep = map[string]*map[string]TypeDeep{
		"key3": &mapDeep,
		"key4": &mapDeep,
	}
	data, _ = json.Marshal(mapPtrMapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "mapPtrMapDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrMapDeep, td.PtrDeep.MapDeep["key3"].MapPtrMapDeep)
	}

	mapMapPtrDeep = map[string]map[string]*TypeDeep{
		"key3": mapPtrDeep,
		"key4": mapPtrDeep,
	}
	data, _ = json.Marshal(mapMapPtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "mapMapPtrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapPtrDeep, td.PtrDeep.MapDeep["key3"].MapMapPtrDeep)
	}

	// Five levels

	td1.StrVal = "strNewNewNewNew"
	data, _ = json.Marshal(td1)
	err = UpdateJSON(td, []string{"ptrDeep", "sliceDeep", "0", "ptrDeep", "ptrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.SliceDeep[0].PtrDeep.PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapArrDeep", "key3", "0", "ptrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapArrDeep["key3"][0].PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapSlicePtrDeep", "key3", "0", "ptrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapSlicePtrDeep["key3"][0].PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapMapDeep", "key3", "key4", "ptrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapMapDeep["key3"]["key4"].PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapMapPtrDeep", "key3", "key4", "ptrDeep"}, data, UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapMapPtrDeep["key3"]["key4"].PtrDeep)
	}

}

func TestUpdateFails(t *testing.T) {
	strVal1 := "strVal1"
	strVal2 := "strVal2"
	mapVal := map[string]string{
		"key1": strVal1,
		"key2": strVal2,
	}
	td := &TypeDeep{}

	err := UpdateJSON(td, []string{"bogus"}, nil, UpdateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"strVal", "bogus"}, nil, UpdateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotTraversableError{}, err)
	}

	err = UpdateJSON(td, []string{"strVal"}, json.RawMessage("strVal"), UpdateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &json.SyntaxError{}, err) // TODO
	}

	err = UpdateJSON(td, []string{"sliceVal", "x"}, nil, UpdateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceVal", "10"}, nil, UpdateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	// TODO: update on empty map to result in KeyNotFoundError? Currently works as Create.
	//data, _ := json.Marshal(mapVal)
	//err = UpdateJSON(td, []string{"mapVal"}, data, UpdateOp)
	//if assert.Error(t, err) {
	//	assert.IsType(t, &KeyNotFoundError{}, err)
	//}

	err = UpdateJSON(td, []string{"mapVal", "bogus"}, nil, UpdateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	td.MapVal = mapVal
	err = UpdateJSON(td, []string{"mapVal", "key1"}, json.RawMessage("bogus"), UpdateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &json.SyntaxError{}, err)
	}

}

func TestCreateFails(t *testing.T) {
	strVal0 := "strVal0"
	strVal1 := "strVal1"
	strVal2 := "strVal2"
	sliceVal0 := []string{strVal0}
	sliceSliceVal0 := [][]string{sliceVal0}
	mapVal0 := map[string]string{
		"key1": strVal1,
		"key2": strVal2,
	}
	sliceMapVal0 := []map[string]string{mapVal0}
	mapSliceVal0 := map[string][]string{"key1": sliceVal0}
	mapMapVal0 := map[string]map[string]string{
		"key0": mapVal0,
	}

	td := &TypeDeep{
		SliceSliceVal: sliceSliceVal0,
		SliceMapVal:   sliceMapVal0,
		MapSliceVal:   mapSliceVal0,
		MapMapVal:     mapMapVal0,
	}

	err := UpdateJSON(td, []string{"bogus"}, nil, CreateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"strVal"}, nil, CreateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyExistsError{}, err)
	}

	err = UpdateJSON(td, []string{"strVal", "bogus"}, nil, CreateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotTraversableError{}, err)
	}

	err = UpdateJSON(td, []string{"mapVal", "new"}, json.RawMessage("bogus"), CreateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &json.SyntaxError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceVal", "x"}, nil, CreateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	data, _ := json.Marshal([]int{5, 6})
	err = UpdateJSON(td, []string{"sliceVal"}, data, CreateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &json.UnmarshalTypeError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceSliceVal", "0"}, data, CreateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &json.UnmarshalTypeError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceMapVal", "0"}, data, CreateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapSliceVal", "key1"}, data, CreateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &json.UnmarshalTypeError{}, err)
	}

	err = UpdateJSON(td, []string{"mapMapVal", "key0"}, data, CreateOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyExistsError{}, err)
	}
}

func TestCreateSuccessful(t *testing.T) {
	strVal0 := "strVal0"
	strVal1 := "strVal1"
	strVal2 := "strVal2"
	mapVal := map[string]string{
		"key1": strVal1,
		"key2": strVal2,
	}

	sliceVal0 := []string{strVal0}
	sliceVal := []string{strVal1, strVal2}

	slicePtrVal0 := []*string{&strVal0}
	slicePtrVal := []*string{&strVal1, &strVal2}

	arrVal := [2]string{strVal1, strVal2}

	sliceMapVal0 := []map[string]string{mapVal}
	sliceMapVal := []map[string]string{mapVal, mapVal}

	slicePtrMapVal0 := []*map[string]string{&mapVal}
	slicePtrMapVal := []*map[string]string{&mapVal, &mapVal}

	sliceSliceVal0 := [][]string{sliceVal0}
	sliceSliceVal := [][]string{sliceVal, sliceVal}

	slicePtrSliceVal0 := []*[]string{&sliceVal}
	slicePtrSliceVal := []*[]string{&sliceVal, &sliceVal}

	sliceIntfVal0 := []interface{}{strVal0, 1}
	sliceIntfVal := []interface{}{strVal1, 5.5}

	var intfVal interface{} = sliceIntfVal

	mapSliceVal0 := map[string][]string{"key1": sliceVal0}
	mapPtrSliceVal0 := map[string]*[]string{"key1": &sliceVal0}

	td := &TypeDeep{
		SliceVal:         sliceVal0,
		SlicePtrVal:      slicePtrVal0,
		SliceMapVal:      sliceMapVal0,
		SlicePtrMapVal:   slicePtrMapVal0,
		SliceSliceVal:    sliceSliceVal0,
		SlicePtrSliceVal: slicePtrSliceVal0,
		SliceIntfVal:     sliceIntfVal0,
		MapSliceVal:      mapSliceVal0,
		MapPtrSliceVal:   mapPtrSliceVal0,
	}

	// One level.

	data, _ := json.Marshal(sliceVal)
	err := UpdateJSON(td, []string{"sliceVal"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, append(sliceVal0, sliceVal...), td.SliceVal)
	}

	err = UpdateJSON(td, []string{"slicePtrVal"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, append(slicePtrVal0, slicePtrVal...), td.SlicePtrVal)
	}

	data, _ = json.Marshal(sliceMapVal)
	err = UpdateJSON(td, []string{"sliceMapVal"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, append(sliceMapVal0, sliceMapVal...), td.SliceMapVal)
	}

	err = UpdateJSON(td, []string{"slicePtrMapVal"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, append(slicePtrMapVal0, slicePtrMapVal...), td.SlicePtrMapVal)
	}

	data, _ = json.Marshal(sliceSliceVal)
	err = UpdateJSON(td, []string{"sliceSliceVal"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, append(sliceSliceVal0, sliceSliceVal...), td.SliceSliceVal)
	}

	err = UpdateJSON(td, []string{"slicePtrSliceVal"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, append(slicePtrSliceVal0, slicePtrSliceVal...), td.SlicePtrSliceVal)
	}

	data, _ = json.Marshal(sliceIntfVal)
	err = UpdateJSON(td, []string{"sliceIntfVal"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, append(sliceIntfVal0, sliceIntfVal...), td.SliceIntfVal)
	}

	// Two levels.

	data, _ = json.Marshal(sliceVal)
	err = UpdateJSON(td, []string{"sliceSliceVal", "0"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, append(sliceVal0, sliceVal...), td.SliceSliceVal[0])
	}

	/* TODO: handle pointers to slices.
	err = UpdateJSON(td, []string{"slicePtrSliceVal", "0"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, append(sliceVal0, sliceVal...), td.SlicePtrSliceVal[0])
	}
	*/

	data, _ = json.Marshal(mapVal)
	err = UpdateJSON(td, []string{"mapMapVal", "new"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapVal, td.MapMapVal["new"])
	}

	err = UpdateJSON(td, []string{"mapPtrMapVal", "new"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &mapVal, td.MapPtrMapVal["new"])
	}

	data, _ = json.Marshal(sliceVal)
	err = UpdateJSON(td, []string{"mapSliceVal", "key1"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, append(sliceVal0, sliceVal...), td.MapSliceVal["key1"])
	}

	/* TODO: handle pointers to slices.
	err = UpdateJSON(td, []string{"mapPtrSliceVal", "key1"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &sliceVal, td.MapPtrSliceVal["key1"])
	}
	*/

	err = UpdateJSON(td, []string{"mapArrVal", "new"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrVal, td.MapArrVal["new"])
	}

	data, _ = json.Marshal(intfVal)
	err = UpdateJSON(td, []string{"mapIntfVal", "new"}, data, CreateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, intfVal, td.MapIntfVal["new"])
	}

}

func TestSetSuccessful(t *testing.T) {
	strVal0 := "strVal0"
	strVal1 := "strVal1"
	strVal2 := "strVal2"

	mapVal0 := map[string]string{
		"key0": strVal0,
	}
	mapVal := map[string]string{
		"key1": strVal1,
		"key2": strVal2,
	}

	sliceMapVal0 := []map[string]string{mapVal0}
	mapMapVal0 := map[string]map[string]string{
		"key0": mapVal0,
	}

	td := &TypeDeep{
		MapVal:      mapVal0,
		SliceMapVal: sliceMapVal0,
		MapMapVal:   mapMapVal0,
	}

	data, _ := json.Marshal(mapVal)
	err := UpdateJSON(td, []string{"mapVal"}, data, SetOp)
	if assert.NoError(t, err) && assert.Equal(t, 3, len(td.MapVal)) {
		assert.Equal(t, mapVal0["key0"], td.MapVal["key0"])
		assert.Equal(t, mapVal["key1"], td.MapVal["key1"])
		assert.Equal(t, mapVal["key2"], td.MapVal["key2"])
	}

	err = UpdateJSON(td, []string{"sliceMapVal", "0"}, data, SetOp)
	if assert.NoError(t, err) && assert.Equal(t, 3, len(td.SliceMapVal[0])) {
		assert.Equal(t, mapVal0["key0"], td.SliceMapVal[0]["key0"])
		assert.Equal(t, mapVal["key1"], td.SliceMapVal[0]["key1"])
		assert.Equal(t, mapVal["key2"], td.SliceMapVal[0]["key2"])
	}

	err = UpdateJSON(td, []string{"mapMapVal", "key0"}, data, SetOp)
	if assert.NoError(t, err) && assert.Equal(t, 3, len(td.MapMapVal["key0"])) {
		assert.Equal(t, mapVal0["key0"], td.MapMapVal["key0"]["key0"])
		assert.Equal(t, mapVal["key1"], td.MapMapVal["key0"]["key1"])
		assert.Equal(t, mapVal["key2"], td.MapMapVal["key0"]["key2"])
	}
}

func TestSetFails(t *testing.T) {
	strVal1 := "strVal1"
	strVal2 := "strVal2"
	mapVal := map[string]string{
		"key1": strVal1,
		"key2": strVal2,
	}
	sliceMapVal0 := []map[string]string{mapVal}
	mapMapVal0 := map[string]map[string]string{
		"key0": mapVal,
	}

	td := &TypeDeep{
		MapVal:      mapVal,
		SliceMapVal: sliceMapVal0,
		MapMapVal:   mapMapVal0,
	}

	data, _ := json.Marshal([]int{5, 6})
	err := UpdateJSON(td, []string{"mapVal"}, data, SetOp)
	if assert.Error(t, err) {
		assert.IsType(t, &json.UnmarshalTypeError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceMapVal", "0"}, data, SetOp)
	if assert.Error(t, err) {
		assert.IsType(t, &json.UnmarshalTypeError{}, err)
	}

	err = UpdateJSON(td, []string{"mapMapVal", "key0"}, data, SetOp)
	if assert.Error(t, err) {
		assert.IsType(t, &json.UnmarshalTypeError{}, err)
	}
}

func TestDeleteFails(t *testing.T) {
	strVal1 := "strVal1"
	strVal2 := "strVal2"
	intVal := 5
	mapVal := map[string]string{
		"key1": strVal1,
		"key2": strVal2,
	}
	sliceVal := []string{strVal1, strVal2}
	mapSliceIntfVal := map[string][]interface{}{
		"key1": []interface{}{strVal1, intVal},
		"key2": []interface{}{strVal2, intVal},
	}
	mapMapIntfVal := map[string]map[string]interface{}{
		"key1": map[string]interface{}{
			"key2": strVal1,
			"key3": sliceVal,
		},
	}
	td1 := TypeDeep{}
	mapDeep := map[string]TypeDeep{
		"key1": td1,
		"key2": td1,
	}
	td := &TypeDeep{
		StrVal:          strVal1,
		IntVal:          intVal,
		MapVal:          mapVal,
		PtrMapVal:       &mapVal,
		SliceVal:        sliceVal,
		MapSliceIntfVal: mapSliceIntfVal,
		MapMapIntfVal:   mapMapIntfVal,
		MapDeep:         mapDeep,
	}

	err := UpdateJSON(td, []string{"bogus"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"strVal"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
		assert.EqualError(t, err, "forbidden operation Delete on key strVal of type optikon.TypeDeep")
		assert.Equal(t, "strVal", err.(*OperationForbiddenError).Key())
		assert.Equal(t, "optikon.TypeDeep", err.(*OperationForbiddenError).KeyType().String())
		assert.Equal(t, DeleteOp, err.(*OperationForbiddenError).Operation())
	}

	err = UpdateJSON(td, []string{"strVal", "bogus"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotTraversableError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceVal"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceVal", "x"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceVal", "10"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceVal", "0"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapVal", "bogus"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"ptrDeep"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapDeep"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapDeep", "key1", "bogus"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"mapDeep", "key1", "strVal"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapDeep", "key1", "sliceVal"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapDeep", "key1", "sliceVal", "0"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	// TODO: IntfVal - check delete if actual data permits it?

	err = UpdateJSON(td, []string{"mapSliceIntfVal", "key1", "0"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapMapIntfVal", "key1", "key3", "0"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapMapIntfVal", "key1", "key3", "3"}, nil, DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}
}

func TestDeleteSuccessful(t *testing.T) {
	strVal1 := "strVal1"
	strVal2 := "strVal2"
	intVal := 5
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
	mapPtrSliceVal := map[string]*[]string{
		"key1": &sliceVal,
		"key2": &sliceVal,
	}
	slicePtrVal := []*string{&strVal1, &strVal2}
	sliceMapVal := []map[string]string{mapVal, mapVal}
	slicePtrMapVal := []*map[string]string{&mapVal, &mapVal}
	sliceSliceVal := [][]string{sliceVal, sliceVal}
	slicePtrSliceVal := []*[]string{&sliceVal, &sliceVal}
	sliceIntfVal := []interface{}{strVal1, &mapSliceVal}
	td1 := TypeDeep{
		MapVal:    mapVal,
		PtrMapVal: &mapVal,
		MapPtrVal: mapPtrVal,
		MapMapVal: mapMapVal,
	}
	mapDeep := map[string]TypeDeep{
		"key1": td1,
		"key2": td1,
	}
	mapPtrDeep := map[string]*TypeDeep{
		"key1": &td1,
		"key2": &td1,
	}
	td := &TypeDeep{
		StrVal:           strVal1,
		IntVal:           intVal,
		MapVal:           mapVal,
		PtrMapVal:        &mapVal,
		MapPtrVal:        mapPtrVal,
		MapMapVal:        mapMapVal,
		MapPtrMapVal:     mapPtrMapVal,
		MapSliceVal:      mapSliceVal,
		MapPtrSliceVal:   mapPtrSliceVal,
		SliceVal:         sliceVal,
		SlicePtrVal:      slicePtrVal,
		SliceMapVal:      sliceMapVal,
		SlicePtrMapVal:   slicePtrMapVal,
		SliceSliceVal:    sliceSliceVal,
		SlicePtrSliceVal: slicePtrSliceVal,
		SliceIntfVal:     sliceIntfVal,
		//MapIntfVal
		//PtrMapIntfVal
		//MapSliceIntfVal
		//MapPtrSliceIntfVal
		//MapMapIntfVal
		//MapPtrMapIntfVal
		PtrDeep: &td1,
		//SliceDeep
		//SlicePtrDeep
		MapDeep:    mapDeep,
		MapPtrDeep: mapPtrDeep,
		//MapSlicePtrDeep
		//MapMapDeep
		//MapPtrMapDeep
		//MapMapPtrDeep
	}

	// One level.

	err := UpdateJSON(td, []string{"mapVal", "key1"}, nil, DeleteOp)
	if assert.NoError(t, err) {
		_, ok := td.MapVal["key1"]
		assert.False(t, ok)
	}

	err = UpdateJSON(td, []string{"ptrMapVal", "key2"}, nil, DeleteOp)
	if assert.NoError(t, err) {
		_, ok := (*td.PtrMapVal)["key2"]
		assert.False(t, ok)
	}

	mapVal["key1"] = strVal1
	mapVal["key2"] = strVal2

	err = UpdateJSON(td, []string{"mapPtrVal", "key1"}, nil, DeleteOp)
	if assert.NoError(t, err) {
		_, ok := td.MapPtrVal["key1"]
		assert.False(t, ok)
	}

	err = UpdateJSON(td, []string{"mapMapVal", "key1"}, nil, DeleteOp)
	if assert.NoError(t, err) {
		_, ok := td.MapMapVal["key1"]
		assert.False(t, ok)
	}

	err = UpdateJSON(td, []string{"mapPtrMapVal", "key1"}, nil, DeleteOp)
	if assert.NoError(t, err) {
		_, ok := td.MapPtrMapVal["key1"]
		assert.False(t, ok)
	}

	// Two levels.

	// TODO: can actually handle deleting indexed elements in the slice, but the
	// delete should probably be done by value and not by index.
	//err = UpdateJSON(td, []string{"sliceIntfVal", "0"}, nil, DeleteOp)
	//if assert.NoError(t, err) {
	//	assert.Equal(t, 1, len(td.SliceIntfVal))
	//}

	err = UpdateJSON(td, []string{"mapDeep", "key1"}, nil, DeleteOp)
	if assert.NoError(t, err) {
		_, ok := td.MapDeep["key1"]
		assert.False(t, ok)
	}

	err = UpdateJSON(td, []string{"mapPtrDeep", "key1"}, nil, DeleteOp)
	if assert.NoError(t, err) {
		_, ok := td.MapPtrDeep["key1"]
		assert.False(t, ok)
	}

	// Three levels.

	err = UpdateJSON(td, []string{"ptrDeep", "mapPtrVal", "key2"}, nil, DeleteOp)
	if assert.NoError(t, err) {
		_, ok := td.PtrDeep.MapPtrVal["key2"]
		assert.False(t, ok)
	}

	err = UpdateJSON(td, []string{"mapMapVal", "key2", "key1"}, nil, DeleteOp)
	if assert.NoError(t, err) {
		_, ok := td.MapMapVal["key2"]["key1"]
		assert.False(t, ok)
	}

	err = UpdateJSON(td, []string{"mapPtrMapVal", "key2", "key2"}, nil, DeleteOp)
	if assert.NoError(t, err) {
		_, ok := (*td.MapPtrMapVal["key2"])["key2"]
		assert.False(t, ok)
	}

	// Four levels.
	mapVal["key1"] = strVal1
	mapVal["key2"] = strVal2
	err = UpdateJSON(td, []string{"mapDeep", "key2", "mapVal", "key1"}, nil, DeleteOp)
	if assert.NoError(t, err) {
		_, ok := td.MapDeep["key2"].MapVal["key1"]
		assert.False(t, ok)
	}

	err = UpdateJSON(td, []string{"mapPtrDeep", "key2", "mapVal", "key2"}, nil, DeleteOp)
	if assert.NoError(t, err) {
		_, ok := td.MapPtrDeep["key2"].MapVal["key2"]
		assert.False(t, ok)
	}
}
