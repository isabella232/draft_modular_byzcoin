package byzcoin

import (
	"context"
	"errors"
	"log"

	"github.com/golang/protobuf/ptypes"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/blockchain/skipchain"
	"go.dedis.ch/phoenix/globalstate/mem"
	"go.dedis.ch/phoenix/ledger"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/scm"
)

// Byzcoin is a ledger implementation.
type Byzcoin struct {
	roster    blockchain.Roster
	bc        blockchain.Blockchain
	validator validator
	factory   factory
}

// NewByzcoin creates a new byzcoin.
func NewByzcoin(o onet.Onet, ro blockchain.Roster, executor scm.Executor) *Byzcoin {
	store := mem.NewStore()
	validator := newValidator(store, executor)

	return &Byzcoin{
		roster:    ro,
		validator: validator,
		bc:        skipchain.NewSkipchain(o, validator),
	}
}

// GetTransactionFactory returns the factory.
func (b *Byzcoin) GetTransactionFactory() ledger.TransactionFactory {
	return b.factory
}

// AddTransaction adds a transaction to the ledger.
func (b *Byzcoin) AddTransaction(in ledger.Transaction) error {
	tx, ok := in.(Transaction)
	if !ok {
		return errors.New("wrong type of transaction")
	}

	instances, err := b.validator.execute(tx)
	if err != nil {
		return err
	}

	instanceIDs := make([][]byte, len(instances))
	for i, inst := range instances {
		instanceIDs[i] = inst.GetKey()
	}

	ptx, err := tx.Pack()
	if err != nil {
		return err
	}

	// The validator will take care of updating the global state and verifying
	// the access control.
	err = b.bc.Store(b.roster, &ledger.TransactionResult{
		Transaction: ptx.(*ledger.TransactionInput),
		Accepted:    true,
		Instances:   instanceIDs,
	})
	if err != nil {
		return err
	}

	return nil
}

type observer struct {
	ch chan *ledger.TransactionResult
}

func (o observer) NotifyCallback(event interface{}) {
	evt := event.(*blockchain.Event)

	var txr ledger.TransactionResult
	err := ptypes.UnmarshalAny(evt.Block.GetPayload(), &txr)

	if err == nil {
		log.Printf("Block [%d] added to the chain", evt.GetBlock().GetIndex())
		o.ch <- &txr
	} else {
		log.Printf("Error: %v", err)
	}
}

// Watch observes the ledger and notifies the new transactions.
func (b *Byzcoin) Watch(ctx context.Context) <-chan *ledger.TransactionResult {
	c := make(chan *ledger.TransactionResult, 100)
	b.bc.Watch(ctx, observer{ch: c})

	return c
}
