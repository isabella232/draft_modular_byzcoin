package blockchain

import (
	"context"

	"github.com/golang/protobuf/proto"
	"go.dedis.ch/phoenix/crypto"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/utils"
)

//go:generate protoc -I ./ --go_out=./ ./messages.proto

// Roster is a set of addresses.
type Roster []*onet.Address

// BlockFactory provides primitives to create blocks from a untrusted source.
type BlockFactory interface {
	Create(src *VerifiableBlock, originPublicKeys []crypto.PublicKey) (interface{}, error)
}

// Blockchain is the interface that provides the primitives to interact with the
// blockchain.
type Blockchain interface {
	GetBlockFactory() BlockFactory

	// Store stores any representation of a data structure into a new block.
	// The implementation is responsible for any validations required.
	Store(ro Roster, data proto.Message) error

	// GetBlock returns a valid proof of the latest block.
	GetBlock() (*VerifiableBlock, error)

	// Watch takes an observer that will be notified for each new block
	// definitely appended to the chain.
	Watch(ctx context.Context, obs utils.Observer)
}
