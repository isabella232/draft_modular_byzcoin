package globalstate

//go:generate protoc -I ./ --go_out=./ ./messages.proto

// Snapshot is a read-only interface to the global state.
type Snapshot interface {
	Read(key []byte) (*Instance, error)
}

// IO is provided to the updater function to update the store.
type IO interface {
	Read(key []byte) (*Instance, error)
	Write(instance *Instance) error
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
