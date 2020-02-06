package local

import (
	"context"
	"errors"
	"fmt"

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

type Onet struct {
	identity *key.Pair
	path     string
	handlers map[string]onet.Handler
}

func NewLocalOnet() *Onet {
	kp := key.NewKeyPair(suite)

	o := &Onet{
		identity: kp,
		path:     "",
		handlers: make(map[string]onet.Handler),
	}
	manager.instances = append(manager.instances, o)
	return o
}

func (o *Onet) handle(path string, msg onet.Message) (onet.Message, error) {
	h := o.handlers[path]
	if h == nil {
		return nil, errors.New("unknown path")
	}

	ctx := context.Background()

	return h(ctx, msg)
}

func (o *Onet) Identity() onet.Identity {
	return o.identity
}

func (o *Onet) Membership() []onet.Identity {
	roster := make([]onet.Identity, len(manager.instances))
	for i, inst := range manager.instances {
		roster[i] = inst.identity
	}

	return roster
}

func (o *Onet) Collect(msg onet.Message) (<-chan onet.Message, error) {
	out := make(chan onet.Message)

	go func() {
		for _, peer := range manager.instances {
			resp, err := peer.handle(o.path, msg)
			if err != nil {
				close(out)
				return
			}

			out <- resp
		}

		close(out)
	}()

	return out, nil
}

func (o *Onet) MakeHandler(id string, h onet.Handler) onet.Onet {
	path := fmt.Sprintf("%s/%s", o.path, id)
	o.handlers[path] = h

	return &Onet{
		path:     path,
		handlers: o.handlers,
	}
}
