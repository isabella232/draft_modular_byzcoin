package skipchain

import (
	proto "github.com/golang/protobuf/proto"
	"go.dedis.ch/phoenix/blockchain/skipchain/cosi"
	"go.dedis.ch/phoenix/types"
)

// Proof is a data structure that contains the shortest chain to a block
// and the integrity can be verified.
type Proof struct {
	verifier cosi.Verifier
	block    Block
}

// NewProof creates a new proof.
func NewProof(block Block, v cosi.Verifier) Proof {
	return Proof{
		verifier: v,
		block:    block,
	}
}

// Payload returns the data of block.
func (p Proof) Payload() proto.Message {
	return p.block.Data
}

// Verify insures the integrity of the proof.
func (p Proof) Verify() error {
	hash, err := p.block.hash()
	if err != nil {
		return err
	}

	err = p.verifier(nil, hash, p.block.Signature)
	if err != nil {
		return err
	}

	return nil
}

// Pack creates a proof message that can be sent over the network.
func (p Proof) Pack() proto.Message {
	return &types.Proof{
		Block: p.block.Pack().(*types.Block),
	}
}
