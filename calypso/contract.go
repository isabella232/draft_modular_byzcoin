package calypso

import (
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go.dedis.ch/phoenix/globalstate"
	"go.dedis.ch/phoenix/scm"
	"go.dedis.ch/phoenix/types"
)

const (
	// ContractID is the unique identifier of the contract.
	ContractID = "go.dedis.ch/phoenix/calypso"

	// ActionWrite is an identifier to a write request that must be allowed by
	// the access control.
	ActionWrite = "write"

	// ActionRead is an identifier to a read request that must be allowed by
	// the access control.
	ActionRead = "read"
)

// SmartContract is the implemenation of the static smart contract interface
// for Calypso.
type SmartContract struct{}

// Get reads the global state.
func (sc SmartContract) Get(s globalstate.Snapshot, in proto.Message) (proto.Message, error) {
	// Returns read and write instances depending on the request.

	return nil, nil
}

// Post creates the write and read instances.
func (sc SmartContract) Post(s globalstate.Snapshot, action scm.Action, in proto.Message) ([]*types.Instance, error) {
	instances := []*types.Instance{}

	switch action {
	case ActionWrite:
		// do write
		write, err := ptypes.MarshalAny(in)
		if err != nil {
			return nil, err
		}

		instances = append(instances, &types.Instance{Key: []byte{1}, Value: write})
	case ActionRead:
		// do read
	default:
		return nil, errors.New("unknown action")
	}

	return instances, nil
}
