package byzcoin

import (
	"errors"

	"go.dedis.ch/phoenix/blockchain/skipchain"
	"go.dedis.ch/phoenix/executor"
	"go.dedis.ch/phoenix/ledger"
	"go.dedis.ch/phoenix/perm"
	"go.dedis.ch/phoenix/perm/naive"
	"go.dedis.ch/phoenix/state"
)

type validator struct {
	store     state.Store
	permStore perm.AccessControlStore
	registry  executor.Registry
}

func newValidator(store state.Store, reg executor.Registry) validator {
	return validator{
		permStore: naive.Store{},
		store:     store,
		registry:  reg,
	}
}

func (v validator) execute(tx Transaction) ([]*state.Instance, error) {
	snapshot, err := v.store.GetCurrent()
	if err != nil {
		return nil, err
	}

	key := executor.Key{ContractID: tx.ContractID, Action: tx.Action}

	instances, err := v.registry.Execute(snapshot, key, tx.Arg)
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
			ac, err := v.permStore.Get(old.GetAccessControl())
			if err != nil {
				return nil, err
			}

			if !ac.CanUpdate(tx.ContractID, tx.Action) {
				return nil, errors.New("access control denied: forbidden to update")
			}
		} else {
			// New instance thus it makes sure the control access exists with
			// sufficient rights.
			ac, err := v.permStore.Get(instance.GetAccessControl())
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

	err = v.store.Update(func(io state.IO) error {
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
