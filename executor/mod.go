package executor

import (
	"github.com/golang/protobuf/proto"
	"go.dedis.ch/phoenix/state"
)

// ContractID is a unique identifier for a smart contract.
type ContractID string

// Action is an action that the contract must perform.
type Action string

// Key is a combination of a contract identifier and an action name.
type Key struct {
	ContractID ContractID
	Action     Action
}

// Registry provides the primitives to interact with smart contracts.
type Registry interface {
	Execute(snapshot state.Snapshot, key Key, in proto.Message) ([]*state.Instance, error)
}
