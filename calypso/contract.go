package calypso

import (
	"github.com/golang/protobuf/proto"
	"go.dedis.ch/phoenix/globalstate"
)

type SmartContract struct{}

func (sc SmartContract) Get(s globalstate.Snapshot, args ...proto.Message) (proto.Message, error) {
	// Returns read and write instances depending on the request.

	return nil, nil
}

func (sc SmartContract) Post(s globalstate.Snapshot, args ...proto.Message) ([]interface{}, error) {
	// Create read or write instances..

	return nil, nil
}
