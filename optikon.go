// Package optikon provides a way to manipulate deep Go structures using
// simple relative path selectors.
package optikon

import "fmt"

// OpType defines possible operation types.
type OpType int

// OpType enumerates available operation types.
const (
	CreateOp OpType = 0x1
	UpdateOp OpType = 0x2
	SetOp    OpType = CreateOp | UpdateOp
	DeleteOp OpType = 0x4
)

func (t OpType) String() string {
	switch t {
	case CreateOp:
		return "Create"
	case UpdateOp:
		return "Update"
	case SetOp:
		return "Set"
	case DeleteOp:
		return "Delete"
	default:
		panic(fmt.Sprintf("operation type not known: %d", t))
	}
}
