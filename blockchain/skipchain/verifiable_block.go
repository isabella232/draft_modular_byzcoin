package skipchain

import (
	proto "github.com/golang/protobuf/proto"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/blockchain/skipchain/cosi"
)

// VerifiableBlock is a data structure that contains the shortest chain to a block
// and the integrity can be verified.
type VerifiableBlock struct {
	verifier cosi.Verifier
	block    Block
}

// NewVerifiableBlock creates a new proof.
func NewVerifiableBlock(block Block, v cosi.Verifier) VerifiableBlock {
	return VerifiableBlock{
		verifier: v,
		block:    block,
	}
}

// Payload returns the data of block.
func (p VerifiableBlock) Payload() proto.Message {
	return p.block.Data
}

// Verify insures the integrity of the proof.
func (p VerifiableBlock) Verify(publicKeys []kyber.Point) error {
	hash, err := p.block.hash()
	if err != nil {
		return err
	}

	err = p.verifier(publicKeys, hash, p.block.Signature)
	if err != nil {
		return err
	}

	return nil
}

// Pack creates a proof message that can be sent over the network.
func (p VerifiableBlock) Pack() proto.Message {
	return &blockchain.Chain{
		Block: p.block.Pack().(*blockchain.Block),
	}
}
