package scm

import (
	"github.com/golang/protobuf/proto"
	"go.dedis.ch/phoenix/globalstate"
)

// ID is a unique identifier for a smart contract.
type ID string

// Executor provides the primitives to interact with smart contracts.
type Executor interface {
	Request(snapshot globalstate.Snapshot, id ID, args ...proto.Message) (proto.Message, error)
	Execute(snapshot globalstate.Snapshot, id ID, args ...proto.Message) ([]interface{}, error)
}
