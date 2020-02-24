package naive

import (
	"go.dedis.ch/phoenix/ac"
	"go.dedis.ch/phoenix/scm"
)

// accessControl is a naive implementation of the access control that
// always gives access.
type accessControl struct{}

func (ac accessControl) CanUpdate(contractID scm.ID, action scm.Action) bool {
	return false
}

func (ac accessControl) CanSpawn(contractID scm.ID, action scm.Action) bool {
	return true
}

// Store is an implementation of the access control store for naive ones.
type Store struct{}

// Get returns a naive access control.
func (s Store) Get(id []byte) (ac.AccessControl, error) {
	return accessControl{}, nil
}
