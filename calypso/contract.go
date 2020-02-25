package calypso

import (
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go.dedis.ch/phoenix/executor"
	"go.dedis.ch/phoenix/state"
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

// Apply creates the write and read instances.
func (sc SmartContract) Apply(s state.Snapshot, action executor.Action, in proto.Message) ([]*state.Instance, error) {
	instances := []*state.Instance{}

	switch action {
	case ActionWrite:
		// do write
		write, err := ptypes.MarshalAny(in)
		if err != nil {
			return nil, err
		}

		instances = append(instances, &state.Instance{Key: []byte{1}, Value: write})
	case ActionRead:
		// do read
	default:
		return nil, errors.New("unknown action")
	}

	return instances, nil
}
