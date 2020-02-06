package onet

import "context"

type Identity interface{}

// Message is an interface that represents the messages passed from a peer
// to another.
type Message interface{}

// Handler receives a message as input and produces a response, or an error
// if something goes wrong.
type Handler func(context.Context, Message) (Message, error)

// Onet is an interface that provides primitives to communicate with
// members of a network.
type Onet interface {
	Identity() Identity
	Membership() []Identity
	Collect(msg Message) (<-chan Message, error)
	MakeHandler(id string, h Handler) Onet
}
