package optikon

import (
	"fmt"
	"reflect"
)

// KeyNotFoundError is returned when trying to access a missing key.
type KeyNotFoundError struct {
	key string
}

// Error implements error interface, stating that key was not found.
func (e *KeyNotFoundError) Error() string {
	return "key not found: " + e.key
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

// KeyNotTraversableError is returned when trying to access a missing key.
type KeyNotTraversableError struct {
	key string
}

// Error implements error interface, stating that key was not found.
func (e *KeyNotTraversableError) Error() string {
	return "key not traversable: " + e.key
}

// Key returns the key that was missing.
func (e *KeyNotTraversableError) Key() string {
	return e.key
}

// OperationForbiddenError is returned when trying to apply a forbidden operation.
type OperationForbiddenError struct {
	key       string
	keyType   reflect.Type
	operation OpType
}

// Error implements error interface, stating that key was not found.
func (e *OperationForbiddenError) Error() string {
	return fmt.Sprintf("forbidden operation %s on key %s of type %s", OpNames[e.operation], e.key, e.keyType)
}

// Key returns the key that was operated on.
func (e *OperationForbiddenError) Key() string {
	return e.key
}

// KeyType returns the key type that was operated on.
func (e *OperationForbiddenError) KeyType() reflect.Type {
	return e.keyType
}

// Operation returns the operation that was forbidden.
func (e *OperationForbiddenError) Operation() OpType {
	return e.operation
}
