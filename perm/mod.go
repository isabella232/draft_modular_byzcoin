package perm

import "go.dedis.ch/phoenix/executor"

// AccessControl provides the primitive to control the access of an instance
// stored in the global state.
type AccessControl interface {
	CanUpdate(executor.ContractID, executor.Action) bool
	CanSpawn(executor.ContractID, executor.Action) bool
}

// AccessControlStore provides the primitives to store and get the access
// control implementations.
type AccessControlStore interface {
	Get(id []byte) (AccessControl, error)
}
