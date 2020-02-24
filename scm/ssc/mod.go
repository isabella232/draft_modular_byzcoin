package ssc

import (
	"errors"

	"github.com/golang/protobuf/proto"
	"go.dedis.ch/phoenix/globalstate"
	"go.dedis.ch/phoenix/scm"
)

// StaticSmartContract is the interface to implement to register a smart
// contract.
type StaticSmartContract interface {
	Get(snapshot globalstate.Snapshot, in proto.Message) (proto.Message, error)
	Post(snapshot globalstate.Snapshot, action scm.Action, in proto.Message) ([]*globalstate.Instance, error)
}

// StaticExecutor registers the smart contracts statically.
type StaticExecutor struct {
	contracts map[scm.ID]StaticSmartContract
}

// NewExecutor returns a new static executor with an empty list of smart
// contracts.
func NewExecutor() *StaticExecutor {
	return &StaticExecutor{
		contracts: make(map[scm.ID]StaticSmartContract),
	}
}

// Register makes a smart contract available to execution.
func (e *StaticExecutor) Register(id scm.ID, sc StaticSmartContract) error {
	if _, ok := e.contracts[id]; ok {
		return errors.New("smart contract identifier already in used")
	}

	e.contracts[id] = sc

	return nil
}

// Request executes the smart contract to read the current state.
func (e *StaticExecutor) Request(snapshot globalstate.Snapshot, id scm.ID, in proto.Message) (proto.Message, error) {
	sc, ok := e.contracts[id]
	if !ok {
		return nil, errors.New("no contract matching the identifier")
	}

	ret, err := sc.Get(snapshot, in)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// Execute runs a smart contract per its identifier.
func (e *StaticExecutor) Execute(snapshot globalstate.Snapshot, id scm.ID, action scm.Action, in proto.Message) ([]*globalstate.Instance, error) {
	sc, ok := e.contracts[id]
	if !ok {
		return nil, errors.New("no contract matching the identifier")
	}

	ret, err := sc.Post(snapshot, action, in)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
