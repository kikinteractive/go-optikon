package optikon

// OpType defines possible operation types.
type OpType int

// OpType enumeration.
const (
	CreateOp OpType = iota
	UpdateOp
	DeleteOp
)

// KeyNotFoundError is returned when trying to access a missing key.
type KeyNotFoundError struct {
	key string
}

// Error implements error interface, stating that key was not found.
func (e *KeyNotFoundError) Error() string {
	return "key not found or not traversable: " + e.key
}

// Key returns the key that was missing.
func (e *KeyNotFoundError) Key() string {
	return e.key
}

// KeyExistsError is returned when creating fails due to existing key.
type KeyExistsError struct {
	key string
}

// Error implements error interface, stating that key already exists.
func (e *KeyExistsError) Error() string {
	return "key exists: " + e.key
}

// Key returns the key that was duplicate.
func (e *KeyExistsError) Key() string {
	return e.key
}
