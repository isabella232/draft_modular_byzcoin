package local

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/types"
)

var suite = pairing.NewSuiteBn256()

type localManager struct {
	instances map[string]*Onet
}

var manager = localManager{
	instances: make(map[string]*Onet),
}

// NewAddress creates an address compatible with local onet.
func NewAddress(id string) *types.Address {
	return &types.Address{
		Id: id,
	}
}

type sender struct {
	addr *types.Address
	in   chan *types.Envelope
}

func (s sender) Send(msg proto.Message, addrs ...*types.Address) error {
	a, err := ptypes.MarshalAny(msg)
	if err != nil {
		return err
	}

	go func() {
		s.in <- &types.Envelope{
			From:    s.addr,
			To:      addrs,
			Message: a,
		}
	}()

	return nil
}

type receiver struct {
	out  chan *types.Envelope
	errs chan error
}

func (r receiver) Recv(ctx context.Context) (*types.Address, proto.Message, error) {
	select {
	case env := <-r.out:
		var da ptypes.DynamicAny
		err := ptypes.UnmarshalAny(env.GetMessage(), &da)
		if err != nil {
			return nil, nil, err
		}

		return env.From, da.Message, nil
	case err := <-r.errs:
		return nil, nil, err
	case <-ctx.Done():
		return nil, nil, errors.New("timeout")
	}
}

// RPC is a registered handler that can send messages to other participants
// to the same handler type.
type RPC struct {
	path string
	h    onet.Handler
}

// Call sends the message to all participants and gather their reply.
func (rpc *RPC) Call(req proto.Message, addrs ...*types.Address) (<-chan proto.Message, error) {
	out := make(chan proto.Message, 1)

	go func() {
		for _, addr := range addrs {
			peer := manager.instances[addr.GetId()]
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

// Stream opens a stream. The caller is responsible for cancelling the context
// to close the stream.
func (rpc *RPC) Stream(ctx context.Context, addrs ...*types.Address) (onet.Sender, onet.Receiver) {
	in := make(chan *types.Envelope)
	out := make(chan *types.Envelope, 1)
	errs := make(chan error, 1)

	outs := make(map[string]receiver)

	for _, addr := range addrs {
		c := make(chan *types.Envelope, 1)
		outs[addr.GetId()] = receiver{out: c}

		peer := manager.instances[addr.GetId()]

		go func(r receiver) {
			s := sender{
				addr: peer.Address(),
				in:   in,
			}

			err := peer.rpcs[rpc.path].h.Stream(s, r)
			if err != nil {
				errs <- err
			}
		}(outs[addr.GetId()])
	}

	orchSender := sender{addr: &types.Address{}, in: in}
	orchRecv := receiver{out: out, errs: errs}

	go func() {
		for {
			select {
			case <-ctx.Done():
				// closes the orchestrator..
				close(out)
				// closes the participants..
				for _, r := range outs {
					close(r.out)
				}
				return
			case env := <-in:
				for _, to := range env.GetTo() {
					if to.GetId() == "" {
						orchRecv.out <- env
					} else {
						outs[to.GetId()].out <- env
					}
				}
			}
		}
	}()

	return orchSender, orchRecv
}

// Onet provides helpers to create handlers.
type Onet struct {
	addr *types.Address
	path string
	rpcs map[string]*RPC
}

// NewLocalOnet creates a new onet instance.
func NewLocalOnet(addr *types.Address) *Onet {
	o := &Onet{
		addr: addr,
		path: "",
		rpcs: make(map[string]*RPC),
	}
	manager.instances[addr.GetId()] = o
	return o
}

// Address returns the address.
func (o *Onet) Address() *types.Address {
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
