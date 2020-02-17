package dkg

import "go.dedis.ch/phoenix/onet"

// SharedSecret is the result of a successful DKG.
type SharedSecret struct{}

// Interface provides the primitives to perform a DKG.
type Interface interface {
	Create([]onet.Address) (*SharedSecret, error)
}
