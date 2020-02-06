package skipchain

import (
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/blockchain/skipchain/cosi"
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
func (p Proof) Payload() blockchain.Payload {
	return p.block.Data
}

// Verify insures the integrity of the proof.
func (p Proof) Verify() error {
	err := p.verifier(p.block.Roster, p.block, p.block.Signature)
	if err != nil {
		return err
	}

	return nil
}
