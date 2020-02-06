package skipchain

import (
	"errors"

	"go.dedis.ch/phoenix/blockchain"
)

// Database is an interface that provides the primitives to read and write
// blocks to a storage.
type Database interface {
	Write(index int64, block blockchain.Block) error
	Read(index int64) (blockchain.Block, error)
	ReadLast() (blockchain.Block, error)
}

// InMemoryDatabase is an implementation of the database interface that is
// an in-memory storage.
type InMemoryDatabase struct {
	blocks []blockchain.Block
}

// NewInMemoryDatabase creates a new in-memory storage for blocks.
func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		blocks: make([]blockchain.Block, 1),
	}
}

func (db *InMemoryDatabase) Write(index int64, block blockchain.Block) error {
	if int64(len(db.blocks)) == index {
		db.blocks = append(db.blocks, block)
	} else if int64(len(db.blocks)) > index {
		db.blocks[index] = block
	} else {
		return errors.New("missing intermediate blocks")
	}

	return nil
}

func (db *InMemoryDatabase) Read(index int64) (blockchain.Block, error) {
	if index < int64(len(db.blocks)) {
		return db.blocks[index], nil
	}

	return blockchain.Block{}, errors.New("block not found")
}

// ReadLast reads the last known block of the chain.
func (db *InMemoryDatabase) ReadLast() (blockchain.Block, error) {
	if len(db.blocks) == 0 {
		return blockchain.Block{}, errors.New("missing genesis block")
	}

	return db.blocks[len(db.blocks)-1], nil
}
