package blockchain

import (
	"context"

	"github.com/golang/protobuf/proto"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/utils"
)

//go:generate protoc -I ./ --go_out=./ ./messages.proto

// Roster is a set of addresses.
type Roster []*onet.Address

// VerifiableBlock is the interface that provides the primitives to verify that a
// block is valid w.r.t. the genesis block.
type VerifiableBlock interface {
	// Payload returns the data of the latest block.
	Payload() proto.Message

	// Verify makes sure that the integrity of the block from the genesis block
	// is correct.
	Verify(publicKeys []kyber.Point) error

	Pack() proto.Message
}

// Blockchain is the interface that provides the primitives to interact with the
// blockchain.
type Blockchain interface {
	// Store stores any representation of a data structure into a new block.
	// The implementation is responsible for any validations required.
	Store(ro Roster, data proto.Message) error

	// GetVerifiableBlock returns a valid proof of the latest block.
	GetVerifiableBlock() (VerifiableBlock, error)

	// Watch takes an observer that will be notified for each new block
	// definitely appended to the chain.
	Watch(ctx context.Context, obs utils.Observer)
}
