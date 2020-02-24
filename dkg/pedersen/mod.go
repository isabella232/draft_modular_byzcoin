package pedersen

import (
	"context"
	"log"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/util/key"
	"go.dedis.ch/phoenix/dkg"
	"go.dedis.ch/phoenix/onet"
)

// DKG is the implementation of the Pedersen DKG algorithm.
type DKG struct {
	rpc onet.RPC
}

// New instantiates a new module of the Pedersen DKG.
func New(o onet.Onet, kp *key.Pair, publicKeys []kyber.Point) *DKG {
	rpc := o.MakeRPC("pdkg", newHandler(kp, publicKeys))

	return &DKG{
		rpc: rpc,
	}
}

// Create starts a new distributed key generation. Each participant will store
// its own share and the caller can only orchestrate without participating.
func (p *DKG) Create(roster []*onet.Address) (kyber.Point, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sender, out := p.rpc.Stream(ctx, roster...)

	err := sender.Send(&dkg.Init{Addresses: roster}, roster...)
	if err != nil {
		return nil, err
	}

	var publicKey kyber.Point
	for i := 0; i < len(roster); i++ {
		_, msg, err := out.Recv(context.Background())
		if err != nil {
			return nil, err
		}

		resp := msg.(*dkg.Done)
		log.Printf("Public Key: %x\n", resp.GetPublicKey())

		publicKey = Suite.Point()
		publicKey.UnmarshalBinary(resp.GetPublicKey())
	}

	return publicKey, nil
}
