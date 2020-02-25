package skipchain

import (
	"context"
	"errors"

	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/blockchain/skipchain/cosi"
	"go.dedis.ch/phoenix/crypto"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/utils"
)

//go:generate protoc -I ./ --go_out=./ ./messages.proto

// Validator is the validator provided by the user of the skipchain module.
type Validator interface {
	Validate(b Block) error
}

type blockValidator struct {
	db Database
	v  Validator
}

func (b blockValidator) Validate(msg proto.Message) ([]byte, error) {
	bm, ok := msg.(*blockchain.Block)
	if !ok {
		return nil, errors.New("unknown type of message")
	}

	var da ptypes.DynamicAny
	err := ptypes.UnmarshalAny(bm.GetPayload(), &da)
	if err != nil {
		return nil, err
	}

	block := Block{
		Index: bm.GetIndex(),
		Data:  da.Message,
	}

	err = b.v.Validate(block)
	if err != nil {
		return nil, err
	}

	last, err := b.db.ReadLast()
	if err != nil {
		return nil, err
	}

	if block.Index <= last.Index {
		return nil, errors.New("wrong index")
	}

	return block.hash()
}

// Skipchain is an implementation of the Blockchain interface that is using
// collective signing to create links between the blocks.
type Skipchain struct {
	db      Database
	onet    onet.Onet
	cosi    cosi.CollectiveSigning
	watcher utils.Observable
	factory blockFactory
}

// NewSkipchain creates a new skipchain-powered blockchain.
func NewSkipchain(o onet.Onet, v Validator) *Skipchain {
	db := NewInMemoryDatabase()
	cosi := cosi.NewBlsCoSi(o, blockValidator{db: db, v: v})

	return &Skipchain{
		db:      db,
		onet:    o,
		cosi:    cosi,
		watcher: utils.NewWatcher(),
		factory: blockFactory{verifier: cosi.MakeVerifier()},
	}
}

// PublicKey returns the cosi public key.
func (s *Skipchain) PublicKey() crypto.PublicKey {
	return s.cosi.PublicKey()
}

// GetBlockFactory returns a factory to create blocks.
func (s *Skipchain) GetBlockFactory() blockchain.BlockFactory {
	return s.factory
}

// Store creates a new block with the data as the payload.
func (s *Skipchain) Store(ro blockchain.Roster, data proto.Message) error {
	last, err := s.db.ReadLast()
	if err != nil {
		return err
	}

	block := Block{
		Index:  last.Index + 1,
		Roster: ro,
		Data:   data,
	}

	packed, err := block.Pack()
	if err != nil {
		return err
	}

	blockproto := packed.(*blockchain.Block)

	sig, err := s.cosi.Sign(ro, blockproto)
	if err != nil {
		return err
	}

	block.Signature = sig

	err = s.db.Write(block)
	if err != nil {
		return err
	}

	go s.watcher.Notify(&blockchain.Event{Block: blockproto})

	return nil
}

// GetBlock reads the latest block of the chain.
func (s *Skipchain) GetBlock() (*blockchain.Block, error) {
	block, err := s.db.ReadLast()
	if err != nil {
		return nil, err
	}

	packed, err := block.Pack()
	if err != nil {
		return nil, err
	}

	return packed.(*blockchain.Block), nil
}

// GetVerifiableBlock reads the latest block of the chain and creates a verifiable
// proof of the shortest chain from the genesis to the block.
func (s *Skipchain) GetVerifiableBlock() (*blockchain.VerifiableBlock, error) {
	block, err := s.GetBlock()
	if err != nil {
		return nil, err
	}

	return &blockchain.VerifiableBlock{Block: block}, nil
}

// Watch registers the observer so that it will be notified of new blocks.
func (s *Skipchain) Watch(ctx context.Context, obs utils.Observer) {
	s.watcher.Add(obs)

	go func() {
		<-ctx.Done()

		s.watcher.Remove(obs)
	}()
}
