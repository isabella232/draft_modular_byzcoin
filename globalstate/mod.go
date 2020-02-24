package globalstate

import "go.dedis.ch/phoenix/types"

// Snapshot is a read-only interface to the global state.
type Snapshot interface {
	Read(key []byte) (*types.Instance, error)
}

// IO is provided to the updater function to update the store.
type IO interface {
	Read(key []byte) (*types.Instance, error)
	Write(instance *types.Instance) error
}

// Updater is the function that is called to create a new version of the store.
type Updater func(io IO) error

// Store provides the primitives to read and write the global state.
type Store interface {
	Update(fn Updater) error
	Remove(version uint64) error
	Snapshot(version uint64) (Snapshot, error)
	GetCurrent() (Snapshot, error)
}
