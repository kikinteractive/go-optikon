package optikon

// OpType defines possible operation types.
type OpType int

// OpType enumerates available operation types.
const (
	CreateOp OpType = 0x1
	UpdateOp OpType = 0x2
	SetOp    OpType = CreateOp | UpdateOp
	DeleteOp OpType = 0x4
)

// OpNames provides readable names for operation types.
var OpNames = map[OpType]string{
	CreateOp: "Create",
	UpdateOp: "Update",
	SetOp:    "Set",
	DeleteOp: "Delete",
}
