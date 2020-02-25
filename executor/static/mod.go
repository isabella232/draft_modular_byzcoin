package static

import (
	"errors"

	"github.com/golang/protobuf/proto"
	"go.dedis.ch/phoenix/executor"
	"go.dedis.ch/phoenix/state"
)

// SmartContract is the interface to implement to register a smart
// contract.
type SmartContract interface {
	Apply(snapshot state.Snapshot, action executor.Action, in proto.Message) ([]*state.Instance, error)
}

// Registry registers the smart contracts statically.
type Registry struct {
	contracts map[executor.ContractID]SmartContract
}

// NewExecutor returns a new static executor with an empty list of smart
// contracts.
func NewExecutor() *Registry {
	return &Registry{
		contracts: make(map[executor.ContractID]SmartContract),
	}
}

// Register makes a smart contract available to execution.
func (e *Registry) Register(id executor.ContractID, sc SmartContract) error {
	if _, ok := e.contracts[id]; ok {
		return errors.New("smart contract identifier already in used")
	}

	e.contracts[id] = sc

	return nil
}

// Execute runs a smart contract per its identifier.
func (e *Registry) Execute(snapshot state.Snapshot, key executor.Key, in proto.Message) ([]*state.Instance, error) {
	sc, ok := e.contracts[key.ContractID]
	if !ok {
		return nil, errors.New("no contract matching the identifier")
	}

	ret, err := sc.Apply(snapshot, key.Action, in)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
