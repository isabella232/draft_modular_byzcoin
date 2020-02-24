package scm

import (
	"github.com/golang/protobuf/proto"
	"go.dedis.ch/phoenix/globalstate"
)

// ID is a unique identifier for a smart contract.
type ID string

// Action is an action that the contract must perform.
type Action string

// Executor provides the primitives to interact with smart contracts.
type Executor interface {
	Request(snapshot globalstate.Snapshot, id ID, in proto.Message) (proto.Message, error)
	Execute(snapshot globalstate.Snapshot, id ID, action Action, in proto.Message) ([]*globalstate.Instance, error)
}
