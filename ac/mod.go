package ac

import "go.dedis.ch/phoenix/scm"

// AccessControl provides the primitive to control the access of an instance
// stored in the global state.
type AccessControl interface {
	CanUpdate(scm.ID, scm.Action) bool
	CanSpawn(scm.ID, scm.Action) bool
}

// AccessControlStore provides the primitives to store and get the access
// control implementations.
type AccessControlStore interface {
	Get(id []byte) (AccessControl, error)
}
