package ledger

import (
	"context"

	"github.com/golang/protobuf/proto"
	"go.dedis.ch/phoenix/encoding"
	"go.dedis.ch/phoenix/scm"
	"go.dedis.ch/phoenix/state"
)

//go:generate protoc -I ./ --proto_path=../ --go_out=Mstate/messages.proto=go.dedis.ch/phoenix/state,Mblockchain/messages.proto=go.dedis.ch/phoenix/blockchain:. ./messages.proto

// Transaction is a set of instructions to be applied to the global state
// one after another.
type Transaction interface {
	encoding.Packable
}

// TransactionFactory is an interface to give an implementation of a ledger
// a chance to format the transactions with a specific format.
type TransactionFactory interface {
	Create(contractID scm.ID, action scm.Action, in proto.Message) (Transaction, error)
}

// InstanceFactory is an interface to create instances from verifiables ones.
type InstanceFactory interface {
	FromVerifiable(src *VerifiableInstance) (*state.Instance, error)
}

// Ledger is the interface that provides primitives to update a public ledger
// through transactions.
type Ledger interface {
	// The factory should be instantiated with stuff like the signer.
	GetTransactionFactory() TransactionFactory

	// GetInstanceFactory returns a factory to create instances from different
	// sources.
	GetInstanceFactory() InstanceFactory

	// AddTransaction gossips the transaction to add it to the ledger.
	AddTransaction(tx Transaction) error

	// GetVerifiableInstance returns an instance with a proof of existence.
	GetVerifiableInstance() (*VerifiableInstance, error)

	// Watch notifies the channel for every new transaction.
	Watch(ctx context.Context) <-chan *TransactionResult
}
