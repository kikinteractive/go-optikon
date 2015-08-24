package optikon

// OpType defines possible operation types.
type OpType int

// OpType enumerates available operation types.
const (
	CreateOp OpType = iota + 1
	UpdateOp
	DeleteOp
)

// OpNames provides readable names for operation types.
var OpNames = map[OpType]string{
	CreateOp: "Create",
	UpdateOp: "Update",
	DeleteOp: "Delete",
}
