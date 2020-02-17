package local

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/phoenix/onet"
)

var suite = pairing.NewSuiteBn256()

// Address is the address of an onet instance.
type Address string

type localManager struct {
	instances map[Address]*Onet
}

var manager = localManager{
	instances: make(map[Address]*Onet),
}

// RPC is a registered handler that can send messages to other participants
// to the same handler type.
type RPC struct {
	path string
	h    onet.Handler
}

// Call sends the message to all participants and gather their reply.
func (rpc *RPC) Call(req proto.Message, addrs ...onet.Address) (<-chan proto.Message, error) {
	out := make(chan proto.Message, 1)

	go func() {
		for _, addr := range addrs {
			peer := manager.instances[addr.(Address)]
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

// Onet provides helpers to create handlers.
type Onet struct {
	addr Address
	path string
	rpcs map[string]*RPC
}

// NewLocalOnet creates a new onet instance.
func NewLocalOnet(addr Address) *Onet {
	o := &Onet{
		addr: addr,
		path: "",
		rpcs: make(map[string]*RPC),
	}
	manager.instances[addr] = o
	return o
}

// Address returns the address.
func (o *Onet) Address() onet.Address {
	return o.addr
}

// MakeNamespace creates a new namespace for the overlay.
func (o *Onet) MakeNamespace(name string) onet.Onet {
	return &Onet{
		path: fmt.Sprintf("%s/%s", o.path, name),
		rpcs: o.rpcs,
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
