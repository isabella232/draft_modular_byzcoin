package pedersen

import (
	"go.dedis.ch/phoenix/dkg"
	"go.dedis.ch/phoenix/onet"
)

// TODO: think about a stream based interaction with onet setting up
// a mesh so that the DKG can be done in one shot and then clean all
// the resources. (grpc duplex feature...).

//go:generate protoc -I ./ --go_out=./ ./messages.proto

// DKG is the implementation of the Pedersen DKG algorithm.
type DKG struct {
	rpc onet.RPC
}

// New instantiates a new module of the Pedersen DKG.
func New(o onet.Onet) *DKG {
	rpc := o.MakeRPC("pdkg", newHandler(o))

	return &DKG{
		rpc: rpc,
	}
}

// Create starts a new distributed key generation. Each participant will store
// its own share and the caller can only orchestrate without participating.
func (p *DKG) Create(roster []onet.Address) (*dkg.SharedSecret, error) {
	return &dkg.SharedSecret{}, nil
}
