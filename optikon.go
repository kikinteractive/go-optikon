package optikon

import "reflect"

// OpType defines possible operation types.
type OpType int

// OpType enumeration.
const (
	CreateOp OpType = iota
	UpdateOp
	DeleteOp
)

// This helper function determines whether this fieldKind is traversable, i.e.
// can be drilled down into.
func isTraversable(fieldKind reflect.Kind) bool {
	return fieldKind == reflect.Map || fieldKind == reflect.Struct ||
		fieldKind == reflect.Array || fieldKind == reflect.Slice ||
		fieldKind == reflect.Interface || fieldKind == reflect.Ptr
}

func IndirectType(t reflect.Type) reflect.Type {
	if isTraversable(t.Kind()) {
		return t.Elem()
	}
	return t
}
