package dkg

import (
	"go.dedis.ch/kyber/v3"
	onet "go.dedis.ch/phoenix/onet"
)

//go:generate protoc -I ./ --proto_path=../ --go_out=Monet/messages.proto=go.dedis.ch/phoenix/onet:. ./messages.proto

// SharedSecret is the result of a successful DKG.
type SharedSecret struct{}

// Interface provides the primitives to perform a DKG.
type Interface interface {
	Create([]*onet.Address) (kyber.Point, error)
}
