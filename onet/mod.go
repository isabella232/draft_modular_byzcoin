// Package onet is the public API for the overlay network.
//
// Note: encoding must be self-describing.
package onet

import (
	"context"
	"errors"

	"github.com/golang/protobuf/proto"
	"go.dedis.ch/phoenix/types"
)

// Sender is an interface to provide primitives to send messages to recipients.
type Sender interface {
	Send(msg proto.Message, addrs ...*types.Address) error
}

// Receiver is an interface to provide primitives to receive messages from
// recipients.
type Receiver interface {
	Recv(context.Context) (*types.Address, proto.Message, error)
}

// RPC is a representation of a remote procedure call that can call a single
// distant procedure or multiple.
type RPC interface {
	// Call is a basic request to one or multiple distant peers.
	Call(req proto.Message, addrs ...*types.Address) (<-chan proto.Message, error)

	// Stream is a persistent request that will be closed only when the
	// orchestrator is done or an error occured.
	Stream(ctx context.Context, addrs ...*types.Address) (in Sender, out Receiver)
}

// Handler is the interface to implement to create a public endpoint.
type Handler interface {
	// Process handles a single request by producing the response according to
	// the request message.
	Process(req proto.Message) (resp proto.Message, err error)

	// Combine gives a chance to reduce the network load by combining multiple
	// messages for a collect call on the intermediate nodes.
	Combine(req []proto.Message) (resp []proto.Message, err error)

	// Stream is a handler for a stream request. It will open a stream with the
	// participants.
	Stream(in Sender, out Receiver) error
}

// DefaultHandler implements the Handler interface with default behaviour so
// that an implementation can focus on its needs.
type DefaultHandler struct{}

// Process is the default implementation for a handler. It will return an error.
func (h DefaultHandler) Process(req proto.Message) (proto.Message, error) {
	return nil, errors.New("rpc is not supported")
}

// Combine returns the messages without combining them.
func (h DefaultHandler) Combine(req []proto.Message) ([]proto.Message, error) {
	return req, nil
}

// Stream is the default implementation for a handler. It will return an error.
func (h DefaultHandler) Stream(in Sender, out Receiver) error {
	return errors.New("stream is not supported")
}

// Onet is a representation of a overlay network that allows the creation
// of namespaces for internal protocols and associate handlers to it.
type Onet interface {
	Address() *types.Address
	MakeNamespace(ns string) Onet
	MakeRPC(name string, h Handler) RPC
}
