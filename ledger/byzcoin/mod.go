package byzcoin

import (
	"context"
	"errors"
	"log"

	"github.com/golang/protobuf/ptypes"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/blockchain/skipchain"
	"go.dedis.ch/phoenix/executor"
	"go.dedis.ch/phoenix/ledger"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/state/mem"
)

// Byzcoin is a ledger implementation.
type Byzcoin struct {
	roster    blockchain.Roster
	bc        blockchain.Blockchain
	validator validator
}

// NewByzcoin creates a new byzcoin.
func NewByzcoin(o onet.Onet, ro blockchain.Roster, registry executor.Registry) *Byzcoin {
	store := mem.NewStore()
	validator := newValidator(store, registry)

	return &Byzcoin{
		roster:    ro,
		validator: validator,
		bc:        skipchain.NewSkipchain(o, validator),
	}
}

// GetTransactionFactory returns the transaction factory.
func (b *Byzcoin) GetTransactionFactory() ledger.TransactionFactory {
	return txFactory{}
}

// GetInstanceFactory returns the instance factory.
func (b *Byzcoin) GetInstanceFactory() ledger.InstanceFactory {
	return instanceFactory{}
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

// GetVerifiableInstance returns the instance if it exists alongside a proof
// that it is included in the latest block.
func (b *Byzcoin) GetVerifiableInstance() (*ledger.VerifiableInstance, error) {
	block, err := b.bc.GetVerifiableBlock()
	if err != nil {
		return nil, err
	}

	return &ledger.VerifiableInstance{Block: block}, nil
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
