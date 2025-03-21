package naive

import (
	"go.dedis.ch/phoenix/executor"
	"go.dedis.ch/phoenix/perm"
)

// accessControl is a naive implementation of the access control that
// always gives access.
type accessControl struct{}

func (ac accessControl) CanUpdate(contractID executor.ContractID, action executor.Action) bool {
	return false
}

func (ac accessControl) CanSpawn(contractID executor.ContractID, action executor.Action) bool {
	return true
}

// Store is an implementation of the access control store for naive ones.
type Store struct{}

// Get returns a naive access control.
func (s Store) Get(id []byte) (perm.AccessControl, error) {
	return accessControl{}, nil
}
