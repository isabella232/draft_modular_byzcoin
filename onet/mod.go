// Package onet is the public API for the overlay network.
//
// Note: encoding must be self-describing.
package onet

import "github.com/golang/protobuf/proto"

// Address represents a unique member of a network.
type Address interface{}

// RPC is a representation of a remote procedure call that can call a single
// distant procedure or multiple.
type RPC interface {
	Call(req proto.Message, addrs ...Address) (<-chan proto.Message, error)
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

// DefaultHandler implements optional functions.
type DefaultHandler struct{}

// Combine returns the messages without combining them.
func (h DefaultHandler) Combine(req []proto.Message) ([]proto.Message, error) {
	return req, nil
}

// Onet is a representation of a overlay network that allows the creation
// of namespaces for internal protocols and associate handlers to it.
type Onet interface {
	Address() Address
	MakeNamespace(ns string) Onet
	MakeRPC(name string, h Handler) RPC
}
