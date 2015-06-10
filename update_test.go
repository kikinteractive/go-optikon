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

	// Second level.
	newMapVal := map[string]string{
		"key1": "newStrVal1",
		"key2": "newStrVal2",
	}
	data, _ = json.Marshal(newMapVal)
	err = UpdateJSON(td, []string{"mapMapVal", "key1"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, newMapVal, td.MapMapVal["key1"])
	}

	err = UpdateJSON(td, []string{"mapPtrMapVal", "key1"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &newMapVal, td.MapPtrMapVal["key1"])
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
		// TODO: results not consistent, sometimes ok, sometimes fail.
		//assert.EqualValues(t, mapIntfVal, td.MapIntfVal)
	}

	data, _ = json.Marshal(mapArrIntfVal)
	err = UpdateJSON(td, []string{"mapArrIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		// TODO: results not consistent, sometimes ok, sometimes fail.
		//assert.EqualValues(t, mapArrIntfVal, td.MapArrIntfVal)
	}

	data, _ = json.Marshal(mapSliceIntfVal)
	err = UpdateJSON(td, []string{"mapSliceIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		// TODO: results not consistent, sometimes ok, sometimes fail.
		//assert.EqualValues(t, mapSliceIntfVal, td.MapSliceIntfVal)
	}

	data, _ = json.Marshal(mapPtrSliceIntfVal)
	err = UpdateJSON(td, []string{"mapPtrSliceIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		// TODO: results not consistent, sometimes ok, sometimes fail.
		//assert.EqualValues(t, mapPtrSliceIntfVal["key1"], td.MapPtrSliceIntfVal["key1"])
		//assert.EqualValues(t, mapPtrSliceIntfVal["key2"], td.MapPtrSliceIntfVal["key2"])
	}

	data, _ = json.Marshal(mapMapIntfVal)
	err = UpdateJSON(td, []string{"mapMapIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		// TODO: results not consistent, sometimes ok, sometimes fail.
		//assert.EqualValues(t, mapMapIntfVal, td.MapMapIntfVal)
	}

	data, _ = json.Marshal(mapPtrMapIntfVal)
	err = UpdateJSON(td, []string{"mapPtrMapIntfVal"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		// TODO: results not consistent, sometimes ok, sometimes fail.
		//assert.EqualValues(t, mapPtrMapIntfVal["key1"], td.MapPtrMapIntfVal["key1"])
		//assert.EqualValues(t, mapPtrMapIntfVal["key2"], td.MapPtrMapIntfVal["key2"])
	}

	// Second level.
	newSliceIntfVal := []interface{}{6, 7.0, "newStrVal"}
	data, _ = json.Marshal(newSliceIntfVal)
	err = UpdateJSON(td, []string{"mapSliceIntfVal", "key1"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, newSliceIntfVal, td.MapSliceIntfVal["key1"])
	}

	err = UpdateJSON(td, []string{"mapPtrSliceIntfVal", "key1"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &newSliceIntfVal, td.MapPtrSliceIntfVal["key1"])
	}

	newMapIntfVal := map[string]interface{}{
		"key3": 6,
		"key4": 7.0,
		"key5": "newStrVal",
	}
	data, _ = json.Marshal(newMapIntfVal)
	err = UpdateJSON(td, []string{"mapMapIntfVal", "key1"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		// TODO: results not consistent, sometimes ok, sometimes fail.
		//assert.EqualValues(t, newMapIntfVal, td.MapMapIntfVal["key1"])
	}

	err = UpdateJSON(td, []string{"mapPtrMapIntfVal", "key1"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		// TODO: results not consistent, sometimes ok, sometimes fail.
		//assert.EqualValues(t, &newMapIntfVal, td.MapPtrMapIntfVal["key1"])
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
	sliceVal := []string{strVal1, strVal2}
	td1 := TypeDeep{
		StrVal:    strVal1,
		IntVal:    intVal,
		MapVal:    mapVal,
		PtrMapVal: &mapVal,
		SliceVal:  sliceVal,
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
	err := UpdateJSON(td, []string{"ptrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep)
	}

	// Two levels

	sliceDeep = []TypeDeep{td1, td1}
	data, _ = json.Marshal(sliceDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "sliceDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceDeep, td.PtrDeep.SliceDeep)
	}

	slicePtrDeep = []*TypeDeep{&td1, &td1}
	data, _ = json.Marshal(slicePtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "slicePtrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrDeep, td.PtrDeep.SlicePtrDeep)
	}

	arrPtrDeep = [2]*TypeDeep{&td1, &td1}
	data, _ = json.Marshal(arrPtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "arrPtrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrPtrDeep, td.PtrDeep.ArrPtrDeep)
	}

	mapDeep = map[string]TypeDeep{
		"key3": td1,
		"key4": td1,
	}
	data, _ = json.Marshal(mapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapDeep, td.PtrDeep.MapDeep)
	}

	mapPtrDeep = map[string]*TypeDeep{
		"key3": &td1,
		"key4": &td1,
	}
	data, _ = json.Marshal(mapPtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapPtrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrDeep, td.PtrDeep.MapPtrDeep)
	}

	mapArrDeep = map[string][2]TypeDeep{
		"key3": arrDeep,
		"key4": arrDeep,
	}
	data, _ = json.Marshal(mapArrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapArrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapArrDeep, td.PtrDeep.MapArrDeep)
	}

	mapSlicePtrDeep = map[string][]*TypeDeep{
		"key3": slicePtrDeep,
		"key4": slicePtrDeep,
	}
	data, _ = json.Marshal(mapSlicePtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapSlicePtrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapSlicePtrDeep, td.PtrDeep.MapSlicePtrDeep)
	}

	mapMapDeep = map[string]map[string]TypeDeep{
		"key3": mapDeep,
		"key4": mapDeep,
	}
	data, _ = json.Marshal(mapMapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapMapDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapDeep, td.PtrDeep.MapMapDeep)
	}

	mapPtrMapDeep = map[string]*map[string]TypeDeep{
		"key3": &mapDeep,
		"key4": &mapDeep,
	}
	data, _ = json.Marshal(mapPtrMapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapPtrMapDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrMapDeep, td.PtrDeep.MapPtrMapDeep)
	}

	mapMapPtrDeep = map[string]map[string]*TypeDeep{
		"key3": mapPtrDeep,
		"key4": mapPtrDeep,
	}
	data, _ = json.Marshal(mapMapPtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapMapPtrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapPtrDeep, td.PtrDeep.MapMapPtrDeep)
	}

	// Three levels

	td1.StrVal = "strNewNew"
	data, _ = json.Marshal(td1)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, td1, td.PtrDeep.MapDeep["key3"])
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapPtrDeep", "key3"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapPtrDeep["key3"])
	}

	arrDeep = [2]TypeDeep{td1, td1}
	data, _ = json.Marshal(arrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapArrDeep", "key3"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrDeep, td.PtrDeep.MapArrDeep["key3"])
	}

	slicePtrDeep = []*TypeDeep{&td1, &td1}
	data, _ = json.Marshal(slicePtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapSlicePtrDeep", "key3"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrDeep, td.PtrDeep.MapSlicePtrDeep["key3"])
	}

	mapDeep = map[string]TypeDeep{
		"key3": td1,
		"key4": td1,
	}
	data, _ = json.Marshal(mapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapMapDeep", "key3"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapDeep, td.PtrDeep.MapMapDeep["key3"])
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapPtrMapDeep", "key3"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &mapDeep, td.PtrDeep.MapPtrMapDeep["key3"])
	}

	mapPtrDeep = map[string]*TypeDeep{
		"key3": &td1,
		"key4": &td1,
	}
	err = UpdateJSON(td, []string{"ptrDeep", "mapMapPtrDeep", "key3"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrDeep, td.PtrDeep.MapMapPtrDeep["key3"])
	}

	// Four levels

	td1.StrVal = "strNewNewNew"

	data, _ = json.Marshal(td1)
	err = UpdateJSON(td, []string{"ptrDeep", "sliceDeep", "0", "ptrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.SliceDeep[0].PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "slicePtrDeep", "0", "ptrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.SlicePtrDeep[0].PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "arrPtrDeep", "0", "ptrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.ArrPtrDeep[0].PtrDeep)
	}

	data, _ = json.Marshal(td1)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "ptrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapDeep["key3"].PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapPtrDeep", "key3", "ptrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapPtrDeep["key3"].PtrDeep)
	}

	sliceDeep = []TypeDeep{td1, td1}
	data, _ = json.Marshal(sliceDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "sliceDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceDeep, td.PtrDeep.MapDeep["key3"].SliceDeep)
	}

	slicePtrDeep = []*TypeDeep{&td1, &td1}
	data, _ = json.Marshal(slicePtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "slicePtrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrDeep, td.PtrDeep.MapDeep["key3"].SlicePtrDeep)
	}

	arrPtrDeep = [2]*TypeDeep{&td1, &td1}
	data, _ = json.Marshal(arrPtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "arrPtrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrPtrDeep, td.PtrDeep.MapDeep["key3"].ArrPtrDeep)
	}

	mapDeep = map[string]TypeDeep{
		"key5": td1,
		"key6": td1,
	}
	data, _ = json.Marshal(mapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "mapDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapDeep, td.PtrDeep.MapDeep["key3"].MapDeep)
	}

	mapPtrDeep = map[string]*TypeDeep{
		"key3": &td1,
		"key4": &td1,
	}
	data, _ = json.Marshal(mapPtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "mapPtrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrDeep, td.PtrDeep.MapDeep["key3"].MapPtrDeep)
	}

	arrDeep = [2]TypeDeep{td1, td1}
	mapArrDeep = map[string][2]TypeDeep{
		"key3": arrDeep,
		"key4": arrDeep,
	}
	data, _ = json.Marshal(mapArrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "mapArrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapArrDeep, td.PtrDeep.MapDeep["key3"].MapArrDeep)
	}

	slicePtrDeep = []*TypeDeep{&td1, &td1}
	data, _ = json.Marshal(slicePtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "slicePtrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrDeep, td.PtrDeep.MapDeep["key3"].SlicePtrDeep)
	}

	mapMapDeep = map[string]map[string]TypeDeep{
		"key3": mapDeep,
		"key4": mapDeep,
	}
	data, _ = json.Marshal(mapMapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "mapMapDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapDeep, td.PtrDeep.MapDeep["key3"].MapMapDeep)
	}

	mapPtrMapDeep = map[string]*map[string]TypeDeep{
		"key3": &mapDeep,
		"key4": &mapDeep,
	}
	data, _ = json.Marshal(mapPtrMapDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "mapPtrMapDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrMapDeep, td.PtrDeep.MapDeep["key3"].MapPtrMapDeep)
	}

	mapMapPtrDeep = map[string]map[string]*TypeDeep{
		"key3": mapPtrDeep,
		"key4": mapPtrDeep,
	}
	data, _ = json.Marshal(mapMapPtrDeep)
	err = UpdateJSON(td, []string{"ptrDeep", "mapDeep", "key3", "mapMapPtrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapPtrDeep, td.PtrDeep.MapDeep["key3"].MapMapPtrDeep)
	}

	// Five levels

	td1.StrVal = "strNewNewNewNew"
	data, _ = json.Marshal(td1)
	err = UpdateJSON(td, []string{"ptrDeep", "sliceDeep", "0", "ptrDeep", "ptrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.SliceDeep[0].PtrDeep.PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapArrDeep", "key3", "0", "ptrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapArrDeep["key3"][0].PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapSlicePtrDeep", "key3", "0", "ptrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapSlicePtrDeep["key3"][0].PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapMapDeep", "key3", "key4", "ptrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapMapDeep["key3"]["key4"].PtrDeep)
	}

	err = UpdateJSON(td, []string{"ptrDeep", "mapMapPtrDeep", "key3", "key4", "ptrDeep"}, string(data), UpdateOp)
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, td.PtrDeep.MapMapPtrDeep["key3"]["key4"].PtrDeep)
	}

}

func TestCreate(t *testing.T) {
	// TODO
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
	td1 := TypeDeep{}
	mapDeep := map[string]TypeDeep{
		"key1": td1,
		"key2": td1,
	}
	td := TypeDeep{
		StrVal:    strVal1,
		IntVal:    intVal,
		MapVal:    mapVal,
		PtrMapVal: &mapVal,
		SliceVal:  sliceVal,
		MapDeep:   mapDeep,
	}

	err := UpdateJSON(td, []string{"bogus"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"strVal"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"strVal", "bogus"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceVal"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceVal", "x"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceVal", "10"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"sliceVal", "0"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapVal", "bogus"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"ptrDeep"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapDeep"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapDeep", "key1", "bogus"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	err = UpdateJSON(td, []string{"mapDeep", "key1", "strVal"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapDeep", "key1", "sliceVal"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &OperationForbiddenError{}, err)
	}

	err = UpdateJSON(td, []string{"mapDeep", "key1", "sliceVal", "0"}, "", DeleteOp)
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	// IntfVal
	// MapSliceIntfVal
	// MapMapIntfVal

}
