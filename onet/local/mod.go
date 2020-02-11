package local

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/util/key"
	"go.dedis.ch/phoenix/onet"
)

var suite = pairing.NewSuiteBn256()

type localManager struct {
	instances []*Onet
}

var manager = localManager{
	instances: make([]*Onet, 0),
}

// RPC is a registered handler that can send messages to other participants
// to the same handler type.
type RPC struct {
	path string
	h    onet.Handler
}

// Collect sends the message to all participants and gather their reply.
func (rpc *RPC) Collect(req proto.Message) (<-chan proto.Message, error) {
	out := make(chan proto.Message, 1)

	go func() {
		for _, peer := range manager.instances {
			resp, err := peer.rpcs[rpc.path].h.Process(req)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				close(out)
				return
			}

			out <- resp
		}

		close(out)
	}()

	return out, nil
}

// Call performs a single rpc call.
func (rpc *RPC) Call(addr string, req proto.Message) (proto.Message, error) {
	return rpc.h.Process(req)
}

// Onet provides helpers to create handlers.
type Onet struct {
	path     string
	rpcs     map[string]*RPC
	identity onet.Identity
}

// NewLocalOnet creates a new onet instance.
func NewLocalOnet() *Onet {
	kp := key.NewKeyPair(suite)

	o := &Onet{
		path:     "v1",
		rpcs:     make(map[string]*RPC),
		identity: kp,
	}
	manager.instances = append(manager.instances, o)
	return o
}

// Identity returns the identity of the current onet.
func (o *Onet) Identity() onet.Identity {
	return o.identity
}

// Membership returns the participants.
func (o *Onet) Membership() []onet.Identity {
	m := make([]onet.Identity, 0)
	for _, p := range manager.instances {
		m = append(m, p.Identity())
	}

	return m
}

// MakeNamespace creates a new namespace for the overlay.
func (o *Onet) MakeNamespace(name string) onet.Onet {
	return &Onet{
		path:     fmt.Sprintf("%s/%s", o.path, name),
		rpcs:     o.rpcs,
		identity: o.identity,
	}
}

// MakeRPC creates a new rpc at the given endpoint.
func (o *Onet) MakeRPC(name string, h onet.Handler) onet.RPC {
	rpc := &RPC{
		path: fmt.Sprintf("%s/%s", o.path, name),
		h:    h,
	}

	o.rpcs[rpc.path] = rpc

	return rpc
}
