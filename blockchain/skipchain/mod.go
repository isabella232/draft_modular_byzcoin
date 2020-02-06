package skipchain

import (
	"context"
	"encoding"
	"errors"

	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/blockchain/skipchain/cosi"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/onet/local"
	"go.dedis.ch/phoenix/utils"
)

// Data is the representation of the stored data.
type Data interface {
	encoding.BinaryMarshaler
}

// Payload is the data structure stored in the chain.
type Payload struct {
	Data      Data
	Signature cosi.Signature
}

// MarshalBinary implements the BinaryMarshaler interface.
func (p Payload) MarshalBinary() ([]byte, error) {
	return p.Data.MarshalBinary()
}

type blockValidator struct {
	db Database
}

func (b blockValidator) Validate(block blockchain.Block) error {
	last, err := b.db.ReadLast()
	if err != nil {
		return err
	}

	if block.Index <= last.Index {
		return errors.New("wrong index")
	}

	return nil
}

// Skipchain is an implementation of the Blockchain interface that is using
// collective signing to create links between the blocks.
type Skipchain struct {
	db      Database
	onet    onet.Onet
	cosi    cosi.CollectiveSigning
	watcher utils.Observable
}

// NewSkipchain creates a new skipchain-powered blockchain.
func NewSkipchain() *Skipchain {
	db := NewInMemoryDatabase()
	onet := local.NewLocalOnet()

	return &Skipchain{
		db:      db,
		onet:    onet,
		cosi:    cosi.NewBlsCoSi(onet, blockValidator{db: db}),
		watcher: utils.NewWatcher(),
	}
}

// Store creates a new block with the data as the payload.
func (s *Skipchain) Store(data blockchain.Payload) error {
	payload := Payload{
		Data: data,
	}

	last, err := s.db.ReadLast()
	if err != nil {
		return err
	}

	roster := blockchain.Roster{}
	for _, m := range s.onet.Membership() {
		roster = append(roster, m)
	}

	block := blockchain.Block{
		Index:   last.Index + 1,
		Roster:  roster,
		Payload: payload,
	}

	sig, err := s.cosi.Sign(block)
	if err != nil {
		return err
	}

	payload.Signature = sig
	block.Payload = payload

	err = s.db.Write(block.Index, block)
	if err != nil {
		return err
	}

	s.watcher.Notify(block)

	return nil
}

// GetProof reads the latest block of the chain and creates a verifiable proof
// of the shortest chain from the genesis to the block.
func (s *Skipchain) GetProof() (blockchain.Proof, error) {
	block, err := s.db.ReadLast()
	if err != nil {
		return nil, err
	}

	return NewProof(block, s.cosi.MakeVerifier()), nil
}

// Watch registers the observer so that it will be notified of new blocks.
func (s *Skipchain) Watch(ctx context.Context, obs utils.Observer) {
	s.watcher.Add(obs)

	go func() {
		<-ctx.Done()

		s.watcher.Remove(obs)
	}()
}
