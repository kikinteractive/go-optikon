package optikon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TypeDeep struct {
	StrVal   string `json:"strVal"`
	IntVal   int    `json:"intVal"`
	NoTagVal string

	MapVal         map[string]string             `json:"mapVal"`
	MapIntVal      map[string]int                `json:"mapIntVal"`
	PtrMapVal      *map[string]string            `json:"ptrMapVal"`
	MapPtrVal      map[string]*string            `json:"mapPtrVal"`
	MapMapVal      map[string]map[string]string  `json:"mapMapVal"`
	MapPtrMapVal   map[string]*map[string]string `json:"mapPtrMapVal"`
	MapSliceVal    map[string][]string           `json:"mapSliceVal"`
	MapArrVal      map[string][2]string          `json:"mapArrVal"`
	MapPtrSliceVal map[string]*[]string          `json:"mapPtrSliceVal"`

	SliceVal         []string             `json:"sliceVal"`
	ArrVal           [2]string            `json:"arrVal"`
	SlicePtrVal      []*string            `json:"slicePtrVal"`
	SliceMapVal      []map[string]string  `json:"sliceMapVal"`
	SlicePtrMapVal   []*map[string]string `json:"slicePtrMapVal"`
	SliceSliceVal    [][]string           `json:"sliceSliceVal"`
	SlicePtrSliceVal []*[]string          `json:"slicePtrSliceVal"`

	IntfVal            interface{}                        `json:"intfVal"`
	ArrIntfVal         [2]interface{}                     `json:"arrIntfVal"`
	SliceIntfVal       []interface{}                      `json:"sliceIntfVal"`
	MapIntfVal         map[string]interface{}             `json:"mapIntfVal"`
	PtrMapIntfVal      *map[string]interface{}            `json:"ptrMapIntfVal"`
	MapArrIntfVal      map[string][2]interface{}          `json:"mapArrIntfVal"`
	MapSliceIntfVal    map[string][]interface{}           `json:"mapSliceIntfVal"`
	MapPtrSliceIntfVal map[string]*[]interface{}          `json:"mapPtrSliceIntfVal"`
	MapMapIntfVal      map[string]map[string]interface{}  `json:"mapMapIntfVal"`
	MapPtrMapIntfVal   map[string]*map[string]interface{} `json:"mapPtrMapIntfVal"`

	PtrDeep         *TypeDeep                       `json:"ptrDeep"`
	SliceDeep       []TypeDeep                      `json:"sliceDeep"`
	SlicePtrDeep    []*TypeDeep                     `json:"slicePtrDeep"`
	ArrPtrDeep      [2]*TypeDeep                    `json:"arrPtrDeep"`
	MapDeep         map[string]TypeDeep             `json:"mapDeep"`
	MapPtrDeep      map[string]*TypeDeep            `json:"mapPtrDeep"`
	MapArrDeep      map[string][2]TypeDeep          `json:"mapArrDeep"`
	MapSlicePtrDeep map[string][]*TypeDeep          `json:"mapSlicePtrDeep"`
	MapMapDeep      map[string]map[string]TypeDeep  `json:"mapMapDeep"`
	MapPtrMapDeep   map[string]*map[string]TypeDeep `json:"mapPtrMapDeep"`
	MapMapPtrDeep   map[string]map[string]*TypeDeep `json:"mapMapPtrDeep"`
}

func TestSelectPrimitive(t *testing.T) {
	td := &TypeDeep{
		StrVal: "strVal",
		IntVal: 5,
	}

	partOut, err := Select(td, []string{"strVal"})
	if assert.NoError(t, err) {
		assert.Equal(t, td.StrVal, partOut)
	}

	partOut, err = Select(td, []string{"intVal"})
	if assert.NoError(t, err) {
		assert.Equal(t, td.IntVal, partOut)
	}

	partOut, err = Select(td, []string{"NoTagVal"})
	if assert.NoError(t, err) {
		assert.Equal(t, td.NoTagVal, partOut)
	}
}

func TestSelectMap(t *testing.T) {
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
	td := &TypeDeep{
		MapVal:         mapVal,
		PtrMapVal:      &mapVal,
		MapPtrVal:      mapPtrVal,
		MapMapVal:      mapMapVal,
		MapPtrMapVal:   mapPtrMapVal,
		MapSliceVal:    mapSliceVal,
		MapArrVal:      mapArrVal,
		MapPtrSliceVal: mapPtrSliceVal,
	}

	partOut, err := Select(td, []string{"mapVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapVal, partOut)
	}

	partOut, err = Select(td, []string{"ptrMapVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, &mapVal, partOut)
	}

	partOut, err = Select(td, []string{"mapPtrVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrVal, partOut)
	}

	partOut, err = Select(td, []string{"mapMapVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapVal, partOut)
	}

	partOut, err = Select(td, []string{"mapPtrMapVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrMapVal, partOut)
	}

	partOut, err = Select(td, []string{"mapSliceVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapSliceVal, partOut)
	}

	partOut, err = Select(td, []string{"mapArrVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapArrVal, partOut)
	}

	partOut, err = Select(td, []string{"mapPtrSliceVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrSliceVal, partOut)
	}
}

func TestSelectSlice(t *testing.T) {
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

	td := &TypeDeep{
		SliceVal:         sliceVal,
		ArrVal:           arrVal,
		SlicePtrVal:      slicePtrVal,
		SliceMapVal:      sliceMapVal,
		SlicePtrMapVal:   slicePtrMapVal,
		SliceSliceVal:    sliceSliceVal,
		SlicePtrSliceVal: slicePtrSliceVal,
	}

	partOut, err := Select(td, []string{"sliceVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceVal, partOut)
	}

	partOut, err = Select(td, []string{"arrVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrVal, partOut)
	}

	partOut, err = Select(td, []string{"slicePtrVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrVal, partOut)
	}

	partOut, err = Select(td, []string{"sliceMapVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceMapVal, partOut)
	}

	partOut, err = Select(td, []string{"slicePtrMapVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrMapVal, partOut)
	}

	partOut, err = Select(td, []string{"sliceSliceVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceSliceVal, partOut)
	}

	partOut, err = Select(td, []string{"slicePtrSliceVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrSliceVal, partOut)
	}
}

func TestSelectIntf(t *testing.T) {
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

	td := &TypeDeep{
		IntfVal:            strVal,
		ArrIntfVal:         arrIntfVal,
		SliceIntfVal:       sliceIntfVal,
		MapIntfVal:         mapIntfVal,
		PtrMapIntfVal:      &mapIntfVal,
		MapArrIntfVal:      mapArrIntfVal,
		MapSliceIntfVal:    mapSliceIntfVal,
		MapPtrSliceIntfVal: mapPtrSliceIntfVal,
		MapMapIntfVal:      mapMapIntfVal,
		MapPtrMapIntfVal:   mapPtrMapIntfVal,
	}

	partOut, err := Select(td, []string{"intfVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, strVal, partOut)
	}

	partOut, err = Select(td, []string{"arrIntfVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrIntfVal, partOut)
	}

	partOut, err = Select(td, []string{"mapIntfVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapIntfVal, partOut)
	}

	partOut, err = Select(td, []string{"ptrMapIntfVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, &mapIntfVal, partOut)
	}

	partOut, err = Select(td, []string{"mapSliceIntfVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapSliceIntfVal, partOut)
	}

	partOut, err = Select(td, []string{"mapPtrSliceIntfVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrSliceIntfVal, partOut)
	}

	partOut, err = Select(td, []string{"mapMapIntfVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapIntfVal, partOut)
	}

	partOut, err = Select(td, []string{"mapPtrMapIntfVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrMapIntfVal, partOut)
	}
}

func TestSelectDeep(t *testing.T) {
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

	partOut, err := Select(td, []string{"ptrDeep"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, partOut)
	}

	partOut, err = Select(td, []string{"sliceDeep"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, sliceDeep, partOut)
	}

	partOut, err = Select(td, []string{"slicePtrDeep"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrDeep, partOut)
	}

	partOut, err = Select(td, []string{"arrPtrDeep"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrPtrDeep, partOut)
	}

	partOut, err = Select(td, []string{"mapDeep"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapDeep, partOut)
	}

	partOut, err = Select(td, []string{"mapPtrDeep"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrDeep, partOut)
	}

	partOut, err = Select(td, []string{"mapArrDeep"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapArrDeep, partOut)
	}

	partOut, err = Select(td, []string{"mapSlicePtrDeep"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapSlicePtrDeep, partOut)
	}

	partOut, err = Select(td, []string{"mapMapDeep"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapDeep, partOut)
	}

	partOut, err = Select(td, []string{"mapMapPtrDeep"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapMapPtrDeep, partOut)
	}

	// Now traverse one level.
	partOut, err = Select(td, []string{"ptrDeep", "strVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, td.PtrDeep.StrVal, partOut)
	}

	partOut, err = Select(td, []string{"sliceDeep", "0"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, td1, partOut)
	}

	partOut, err = Select(td, []string{"slicePtrDeep", "1"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, partOut)
	}

	partOut, err = Select(td, []string{"arrPtrDeep", "0"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, partOut)
	}

	partOut, err = Select(td, []string{"mapDeep", "key1"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, td1, partOut)
	}

	partOut, err = Select(td, []string{"mapPtrDeep", "key1"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, partOut)
	}

	partOut, err = Select(td, []string{"mapArrDeep", "key1"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, arrDeep, partOut)
	}

	partOut, err = Select(td, []string{"mapSlicePtrDeep", "key1"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, slicePtrDeep, partOut)
	}

	partOut, err = Select(td, []string{"mapMapDeep", "key1"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapDeep, partOut)
	}

	partOut, err = Select(td, []string{"mapMapPtrDeep", "key1"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, mapPtrDeep, partOut)
	}

	// Now traverse two levels.
	partOut, err = Select(td, []string{"sliceDeep", "0", "strVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, td1.StrVal, partOut)
	}

	partOut, err = Select(td, []string{"slicePtrDeep", "1", "intVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, td1.IntVal, partOut)
	}

	partOut, err = Select(td, []string{"arrPtrDeep", "0", "intVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, td1.IntVal, partOut)
	}

	partOut, err = Select(td, []string{"mapDeep", "key1", "intVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, td1.IntVal, partOut)
	}

	partOut, err = Select(td, []string{"mapPtrDeep", "key1", "intVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, td1.IntVal, partOut)
	}

	partOut, err = Select(td, []string{"mapArrDeep", "key1", "0"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, td1, partOut)
	}

	partOut, err = Select(td, []string{"mapSlicePtrDeep", "key1", "0"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, partOut)
	}

	partOut, err = Select(td, []string{"mapMapDeep", "key1", "key2"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, td1, partOut)
	}

	partOut, err = Select(td, []string{"mapMapPtrDeep", "key1", "key2"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, &td1, partOut)
	}

	// Now traverse three levels.
	partOut, err = Select(td, []string{"sliceDeep", "0", "mapVal", "key1"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, strVal1, partOut)
	}

	partOut, err = Select(td, []string{"slicePtrDeep", "1", "mapVal", "key1"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, strVal1, partOut)
	}

	partOut, err = Select(td, []string{"arrPtrDeep", "0", "mapVal", "key1"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, strVal1, partOut)
	}

	partOut, err = Select(td, []string{"mapDeep", "key1", "mapVal", "key2"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, strVal2, partOut)
	}

	partOut, err = Select(td, []string{"mapPtrDeep", "key1", "mapVal", "key1"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, strVal1, partOut)
	}

	partOut, err = Select(td, []string{"mapArrDeep", "key1", "0", "strVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, strVal1, partOut)
	}

	partOut, err = Select(td, []string{"mapSlicePtrDeep", "key1", "0", "intVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, intVal, partOut)
	}

	partOut, err = Select(td, []string{"mapMapDeep", "key1", "key2", "strVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, strVal1, partOut)
	}

	partOut, err = Select(td, []string{"mapMapPtrDeep", "key1", "key2", "strVal"})
	if assert.NoError(t, err) {
		assert.EqualValues(t, strVal1, partOut)
	}
}

func TestSelectFails(t *testing.T) {
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
	slicePtrVal := []*string{&strVal1, &strVal2}
	sliceMapVal := []map[string]string{mapVal, mapVal}
	slicePtrMapVal := []*map[string]string{&mapVal, &mapVal}
	sliceSliceVal := [][]string{sliceVal, sliceVal}
	slicePtrSliceVal := []*[]string{&sliceVal, &sliceVal}

	td := &TypeDeep{
		StrVal:           strVal1,
		IntVal:           5,
		SliceVal:         sliceVal,
		ArrVal:           arrVal,
		SlicePtrVal:      slicePtrVal,
		SliceMapVal:      sliceMapVal,
		SlicePtrMapVal:   slicePtrMapVal,
		SliceSliceVal:    sliceSliceVal,
		SlicePtrSliceVal: slicePtrSliceVal,
		MapVal:           mapVal,
		PtrMapVal:        &mapVal,
		MapPtrVal:        mapPtrVal,
		MapMapVal:        mapMapVal,
		MapPtrMapVal:     mapPtrMapVal,
		MapSliceVal:      mapSliceVal,
		MapArrVal:        mapArrVal,
		MapPtrSliceVal:   mapPtrSliceVal,
	}

	_, err := Select(td, []string{"bogus"})
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
		assert.EqualError(t, err, "key not found: bogus")
		assert.Equal(t, "bogus", err.(*KeyNotFoundError).Key())
	}

	_, err = Select(td, []string{"strVal", "dummy"})
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotTraversableError{}, err)
		assert.EqualError(t, err, "key not traversable: dummy")
		assert.Equal(t, "dummy", err.(*KeyNotTraversableError).Key())
	}

	_, err = Select(td, []string{"sliceVal", "x"})
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	_, err = Select(td, []string{"sliceVal", "10"})
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	_, err = Select(td, []string{"sliceMapVal", "0", "dummy"})
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	_, err = Select(td, []string{"sliceSliceVal", "0", "x"})
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	_, err = Select(td, []string{"sliceSliceVal", "0", "10"})
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	_, err = Select(td, []string{"mapMapVal", "key1", "x"})
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	_, err = Select(td, []string{"mapMapVal", "key1", "key1", "strVal", "dummy"})
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotTraversableError{}, err)
	}

	_, err = Select(td, []string{"mapSliceVal", "key1", "x"})
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	_, err = Select(td, []string{"mapSliceVal", "key1", "10"})
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

	_, err = Select(td, []string{"ptrDeep", "strVal"})
	if assert.Error(t, err) {
		assert.IsType(t, &KeyNotFoundError{}, err)
	}

}
