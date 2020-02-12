package skipchain

import (
	"context"
	"crypto/sha256"
	"errors"

	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/blockchain/skipchain/cosi"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/onet/local"
	"go.dedis.ch/phoenix/utils"
)

//go:generate protoc -I ./ --go_out=./ ./types.proto

// Roster is the representation of a membership.
type Roster []onet.Identity

// Block is the representation of the data structures that will be linked
// together.
type Block struct {
	Index     int64
	Roster    Roster
	Signature cosi.Signature
	Data      proto.Message
}

func (b Block) hash() ([]byte, error) {
	h := sha256.New()

	buffer, err := proto.Marshal(b.Data)
	if err != nil {
		return nil, err
	}

	_, err = h.Write(buffer)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Pack returns a network message.
func (b Block) Pack() proto.Message {
	any, _ := ptypes.MarshalAny(b.Data)

	return &BlockMessage{
		Index: b.Index,
		Data:  any,
	}
}

// PayloadValidator is the validator provided by the user of the skipchain module.
type PayloadValidator func(b Block) error

type blockValidator struct {
	db Database
	pv PayloadValidator
}

func (b blockValidator) Validate(msg proto.Message) ([]byte, error) {
	bm, ok := msg.(*BlockMessage)
	if !ok {
		return nil, errors.New("unknown type of message")
	}

	var da ptypes.DynamicAny
	err := ptypes.UnmarshalAny(bm.GetData(), &da)
	if err != nil {
		return nil, err
	}

	block := Block{
		Index: bm.GetIndex(),
		Data:  da.Message,
	}

	err = b.pv(block)
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
}

// NewSkipchain creates a new skipchain-powered blockchain.
func NewSkipchain(v PayloadValidator) *Skipchain {
	db := NewInMemoryDatabase()
	onet := local.NewLocalOnet()

	return &Skipchain{
		db:      db,
		onet:    onet,
		cosi:    cosi.NewBlsCoSi(onet, blockValidator{db: db, pv: v}),
		watcher: utils.NewWatcher(),
	}
}

// Store creates a new block with the data as the payload.
func (s *Skipchain) Store(data proto.Message) error {
	last, err := s.db.ReadLast()
	if err != nil {
		return err
	}

	block := Block{
		Index:  last.Index + 1,
		Roster: s.onet.Membership(),
		Data:   data,
	}

	sig, err := s.cosi.Sign(block.Pack())
	if err != nil {
		return err
	}

	block.Signature = sig

	err = s.db.Write(block)
	if err != nil {
		return err
	}

	s.watcher.Notify(blockchain.Event{})

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
