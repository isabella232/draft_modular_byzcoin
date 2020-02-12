package ledger

import (
	"context"

	"github.com/gogo/protobuf/proto"
)

// TransactionResult is the data structure sent back when a transaction is
// stored in the chain.
type TransactionResult struct {
	ID       []byte
	Accepted bool
}

// State is a verifiable value stored in the chain.
type State interface {
	Value() []byte
	Verify() error
	Pack() proto.Message
}

// Ledger is the interface that provides primitives to update a public ledger
// through transactions.
type Ledger interface {
	AddTransaction(tx proto.Message) error
	GetState(key string) (State, error)
	Watch(ctx context.Context) <-chan TransactionResult
}
