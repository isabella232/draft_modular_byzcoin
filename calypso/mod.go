package calypso

import (
	"context"
	"log"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/suites"
	"go.dedis.ch/kyber/v3/util/key"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/dkg"
	"go.dedis.ch/phoenix/dkg/pedersen"
	"go.dedis.ch/phoenix/ledger"
	"go.dedis.ch/phoenix/ledger/byzcoin"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/scm/ssc"
	"go.dedis.ch/phoenix/types"
)

// Suite is the cryptographic suite of the package.
var Suite = suites.MustFind("Ed25519")

// Interface is a private storage module.
type Interface interface {
	// AddWrite adds a write instance.
	AddWrite() error

	// AddRead adds a read instance.
	AddRead()

	DecryptKey()
}

// Calypso is the implementation of a private storage.
type Calypso struct {
	dkg    dkg.Interface
	ledger ledger.Ledger
}

// NewCalypso returns a new instance of the calypso module.
func NewCalypso(o onet.Onet, ro blockchain.Roster, kp *key.Pair, pubkeys []kyber.Point) *Calypso {
	sce := ssc.NewExecutor()
	sce.Register(ContractID, SmartContract{})

	return &Calypso{
		dkg:    pedersen.New(o, kp, pubkeys),
		ledger: byzcoin.NewByzcoin(o, ro, sce),
	}
}

// AddWrite creates a write instance.
func (c *Calypso) AddWrite() error {
	tx, err := c.ledger.GetTransactionFactory().Create(ContractID, ActionWrite, &types.CalypsoWrite{})
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := c.ledger.Watch(ctx)

	err = c.ledger.AddTransaction(tx)
	if err != nil {
		return err
	}

	tr := <-ch
	log.Printf("Transaction Result: %v\n", tr)

	return nil
}

func (c *Calypso) createLTS() error {
	_, err := c.dkg.Create(nil)
	if err != nil {
		return err
	}

	return nil
}
