package mem

import (
	"errors"
	"sync"

	"go.dedis.ch/phoenix/globalstate"
)

// Key is the type of the instance keys.
type Key [32]byte

// InMemorySnapshot is an immutable state of the store.
type InMemorySnapshot map[Key]*globalstate.Instance

func (s InMemorySnapshot) Read(key []byte) (*globalstate.Instance, error) {
	k := Key{}
	copy(k[:], key)

	inst := s[k]

	return inst, nil
}

type io struct {
	snapshot InMemorySnapshot
}

func (u io) Read(key []byte) (*globalstate.Instance, error) {
	return u.snapshot.Read(key)
}

func (u io) Write(instance *globalstate.Instance) error {
	k := Key{}
	copy(k[:], instance.GetKey())

	u.snapshot[k] = instance
	return nil
}

// InMemoryStore is a implementation of the store interface that can be used
// as a global state.
type InMemoryStore struct {
	sync.Mutex
	snapshots []InMemorySnapshot
}

// NewStore creates a new in-memory store.
func NewStore() *InMemoryStore {
	return &InMemoryStore{
		snapshots: []InMemorySnapshot{{}},
	}
}

// Update adds a snapshot to the store.
func (s *InMemoryStore) Update(fn globalstate.Updater) error {
	s.Lock()
	defer s.Unlock()

	latest := s.snapshots[len(s.snapshots)-1]

	newest := make(InMemorySnapshot)
	for key, value := range latest {
		newest[key] = value
	}

	err := fn(io{snapshot: newest})
	if err != nil {
		return err
	}

	s.snapshots = append(s.snapshots, newest)

	return nil
}

// Remove deletes a snaphost of the store.
func (s *InMemoryStore) Remove(version uint64) error {
	s.Lock()
	defer s.Unlock()

	return nil
}

// Snapshot returns the snapshot for the given version.
func (s *InMemoryStore) Snapshot(version uint64) (globalstate.Snapshot, error) {
	s.Lock()
	defer s.Unlock()

	if version >= uint64(len(s.snapshots)) {
		return nil, errors.New("unknown version")
	}

	return s.snapshots[version], nil
}

// GetCurrent returns the latest version of the store.
func (s *InMemoryStore) GetCurrent() (globalstate.Snapshot, error) {
	s.Lock()
	defer s.Unlock()

	return s.snapshots[len(s.snapshots)-1], nil
}
