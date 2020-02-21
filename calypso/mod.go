package calypso

import (
	"go.dedis.ch/phoenix/dkg"
	"go.dedis.ch/phoenix/dkg/pedersen"
	"go.dedis.ch/phoenix/ledger"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/scm"
	"go.dedis.ch/phoenix/scm/ssc"
)

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
	sce    scm.Executor
	ledger ledger.Ledger
}

// NewCalypso returns a new instance of the calypso module.
func NewCalypso(o onet.Onet) *Calypso {
	sce := ssc.NewExecutor()
	sce.Register("calypso", SmartContract{})

	return &Calypso{
		dkg: pedersen.New(o, nil, nil),
		sce: sce,
	}
}

func (c *Calypso) AddWrite() error {
	tx := c.ledger.GetTransactionFactory().Create()
	err := c.ledger.AddTransaction(tx)
	if err != nil {
		return err
	}

	return nil
}

func (c *Calypso) createLTS() error {
	_, err := c.dkg.Create(nil)
	if err != nil {
		return err
	}

	return nil
}
