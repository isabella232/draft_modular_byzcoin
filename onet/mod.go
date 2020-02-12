// Package onet is the public API for the overlay network.
//
// Note: encoding must be self-describing.
package onet

import "github.com/golang/protobuf/proto"

type Identity interface{}

// RPC is a representation of a remote procedure call that can call a single
// distant procedure or multiple.
type RPC interface {
	Collect(req proto.Message) (<-chan proto.Message, error)
	Call(addr string, req proto.Message) (proto.Message, error)
}

// Handler is the interface to implement to create a public endpoint.
type Handler interface {
	// Process handles a single request by producing the response according to
	// the request message.
	Process(req proto.Message) (resp proto.Message, err error)

	// Combine gives a chance to reduce the network load by combining multiple
	// messages for a collect call on the intermediate nodes.
	Combine(req []proto.Message) (resp []proto.Message, err error)
}

// Onet is a representation of a overlay network that allows the creation
// of namespaces for internal protocols and associate handlers to it.
type Onet interface {
	Identity() Identity
	Membership() []Identity
	MakeNamespace(ns string) Onet
	MakeRPC(name string, h Handler) RPC
}
