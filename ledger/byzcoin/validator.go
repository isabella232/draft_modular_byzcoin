package byzcoin

import (
	"errors"

	"go.dedis.ch/phoenix/ac"
	"go.dedis.ch/phoenix/ac/naive"
	"go.dedis.ch/phoenix/blockchain/skipchain"
	"go.dedis.ch/phoenix/globalstate"
	"go.dedis.ch/phoenix/ledger"
	"go.dedis.ch/phoenix/scm"
)

type validator struct {
	store    globalstate.Store
	ac       ac.AccessControlStore
	executor scm.Executor
}

func newValidator(store globalstate.Store, exec scm.Executor) validator {
	return validator{
		ac:       naive.Store{},
		store:    store,
		executor: exec,
	}
}

func (v validator) execute(tx Transaction) ([]*globalstate.Instance, error) {
	snapshot, err := v.store.GetCurrent()
	if err != nil {
		return nil, err
	}

	instances, err := v.executor.Execute(snapshot, tx.ContractID, tx.Action, tx.Arg)
	if err != nil {
		return nil, err
	}

	for _, instance := range instances {
		old, err := snapshot.Read(instance.GetKey())
		if err != nil {
			return nil, err
		}

		if old != nil {
			// A previous instance already exist thus it makes sure the previous
			// access control allows the update.
			ac, err := v.ac.Get(old.GetAccessControl())
			if err != nil {
				return nil, err
			}

			if !ac.CanUpdate(tx.ContractID, tx.Action) {
				return nil, errors.New("access control denied: forbidden to update")
			}
		} else {
			// New instance thus it makes sure the control access exists with
			// sufficient rights.
			ac, err := v.ac.Get(instance.GetAccessControl())
			if err != nil {
				return nil, err
			}

			if !ac.CanSpawn(tx.ContractID, tx.Action) {
				return nil, errors.New("access control denied: forbidden to spawn")
			}
		}
	}

	return instances, nil
}

func (v validator) Validate(block skipchain.Block) error {
	res := block.Data.(*ledger.TransactionResult)

	tx, err := FromProto(res.GetTransaction())
	if err != nil {
		return err
	}

	instances, err := v.execute(tx)
	if err != nil {
		return err
	}

	err = v.store.Update(func(io globalstate.IO) error {
		for _, inst := range instances {
			err := io.Write(inst)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
