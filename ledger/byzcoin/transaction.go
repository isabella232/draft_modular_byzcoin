package byzcoin

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go.dedis.ch/phoenix/ledger"
	"go.dedis.ch/phoenix/scm"
)

// Transaction is the data structure for a transaction specific to Byzcoin.
type Transaction struct {
	ContractID scm.ID
	Action     scm.Action
	Arg        proto.Message
}

// FromProto returns a transaction from a protobuf message.
func FromProto(msg proto.Message) (Transaction, error) {
	tx := msg.(*ledger.TransactionInput)

	var da ptypes.DynamicAny
	err := ptypes.UnmarshalAny(tx.GetBody(), &da)
	if err != nil {
		return Transaction{}, err
	}

	return Transaction{
		ContractID: scm.ID(tx.ContractID),
		Action:     scm.Action(tx.Action),
		Arg:        da.Message,
	}, nil
}

// Pack returns the protobuf message of the transaction.
func (t Transaction) Pack() (proto.Message, error) {
	body, err := ptypes.MarshalAny(t.Arg)
	if err != nil {
		return nil, err
	}

	return &ledger.TransactionInput{
		ContractID: string(t.ContractID),
		Action:     string(t.Action),
		Body:       body,
	}, nil
}

type txFactory struct{}

func (f txFactory) Create(contractID scm.ID, action scm.Action, in proto.Message) (ledger.Transaction, error) {
	// TODO: sign tx

	return Transaction{ContractID: contractID, Action: action, Arg: in}, nil
}
