package byzcoin

import (
	"context"
	"errors"

	"github.com/gogo/protobuf/proto"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/blockchain/skipchain"
	"go.dedis.ch/phoenix/ledger"
)

//go:generate protoc -I ../../ -I ./ --go_out=./ ./messages.proto

type State struct {
	proof blockchain.Proof
	value []byte
}

func (s State) Verify() error {
	return s.proof.Verify()
}

func (s State) Value() []byte {
	return s.value
}

func (s State) Pack() proto.Message {
	return &StateMessage{
		Proof: s.proof.Pack().(*skipchain.ProofMessage),
		Value: s.value,
	}
}

type Byzcoin struct {
	bc blockchain.Blockchain
}

func (b *Byzcoin) AddTransaction(in proto.Message) error {
	tx, ok := in.(*Transaction)
	if !ok {
		return errors.New("wrong type of transaction")
	}

	err := b.bc.Store(tx)
	if err != nil {
		return err
	}

	return nil
}

func (b *Byzcoin) GetState(key string) (ledger.State, error) {
	proof, err := b.bc.GetProof()
	if err != nil {
		return nil, err
	}

	state := State{
		proof: proof,
		value: []byte{},
	}

	return state, nil
}

type observer struct{}

func (o observer) NotifyCallback(event interface{}) {

}

func (b *Byzcoin) Watch(ctx context.Context) <-chan ledger.TransactionResult {
	c := make(chan ledger.TransactionResult)
	b.bc.Watch(context.Background(), observer{})
	return c
}
