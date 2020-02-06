package skipchain

import (
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/blockchain/skipchain/cosi"
)

// Proof is a data structure that contains the shortest chain to a block
// and the integrity can be verified.
type Proof struct {
	verifier cosi.Verifier
	block    blockchain.Block
}

// NewProof creates a new proof.
func NewProof(block blockchain.Block, v cosi.Verifier) Proof {
	return Proof{
		verifier: v,
		block:    block,
	}
}

// LatestBlock returns the block.
func (p Proof) LatestBlock() blockchain.Block {
	return p.block
}

// Verify insures the integrity of the proof.
func (p Proof) Verify() error {
	payload := p.block.Payload.(Payload)

	err := p.verifier(p.block, payload.Signature)
	if err != nil {
		return err
	}

	return nil
}
