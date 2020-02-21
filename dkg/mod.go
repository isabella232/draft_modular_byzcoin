package dkg

import (
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/phoenix/types"
)

// SharedSecret is the result of a successful DKG.
type SharedSecret struct{}

// Interface provides the primitives to perform a DKG.
type Interface interface {
	Create([]*types.Address) (kyber.Point, error)
}
