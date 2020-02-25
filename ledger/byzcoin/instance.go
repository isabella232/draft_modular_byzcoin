package byzcoin

import (
	"go.dedis.ch/phoenix/globalstate"
	"go.dedis.ch/phoenix/ledger"
)

type instanceFactory struct{}

func (f instanceFactory) FromVerifiable(src *ledger.VerifiableInstance) (*globalstate.Instance, error) {
	// 1. Get the block from the verifiable block.
	// 2. Check the Merkle trie root against the instance hash.

	return src.GetInstance(), nil
}
