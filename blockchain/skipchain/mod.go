package skipchain

import (
	"context"
	"crypto/sha256"
	"encoding"
	"errors"

	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/blockchain/skipchain/cosi"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/onet/local"
	"go.dedis.ch/phoenix/utils"
)

// Roster is the representation of a membership.
type Roster []onet.Identity

// Block is the representation of the data structures that will be linked
// together.
type Block struct {
	Index     int64
	Roster    Roster
	Signature cosi.Signature
	Data      blockchain.Payload
}

// Hash returns a unique hash for each block.
func (b Block) Hash() ([]byte, error) {
	h := sha256.New()

	databuf, err := b.Data.MarshalBinary()
	if err != nil {
		return nil, err
	}

	_, err = h.Write(databuf)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Data is the representation of the stored data.
type Data interface {
	encoding.BinaryMarshaler
}

type blockValidator struct {
	db Database
}

func (b blockValidator) Validate(msg cosi.Hashable) error {
	block, ok := msg.(Block)
	if !ok {
		return errors.New("unknown type of message")
	}

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
	last, err := s.db.ReadLast()
	if err != nil {
		return err
	}

	block := Block{
		Index:  last.Index + 1,
		Roster: s.onet.Membership(),
		Data:   data,
	}

	sig, err := s.cosi.Sign(block)
	if err != nil {
		return err
	}

	block.Signature = sig

	err = s.db.Write(block)
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
