package ledger

import (
	"context"

	"github.com/golang/protobuf/proto"
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

// Transaction is a set of instructions to be applied to the global state
// one after another.
type Transaction interface {
	proto.Message
}

// TransactionFactory is an interface to give an implementation of a ledger
// a chance to format the transactions with a specific format.
type TransactionFactory interface {
	Create(args ...proto.Message) Transaction
}

// Ledger is the interface that provides primitives to update a public ledger
// through transactions.
type Ledger interface {
	// The factory should be instantiated with stuff like the signer.
	GetTransactionFactory() TransactionFactory
	GetState(key string) (State, error)
	AddTransaction(tx Transaction) error
	Watch(ctx context.Context) <-chan TransactionResult
}
