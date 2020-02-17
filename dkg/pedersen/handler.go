package pedersen

import (
	"errors"

	"github.com/golang/protobuf/proto"
	"go.dedis.ch/phoenix/onet"
)

type handler struct {
	onet.DefaultHandler
}

func newHandler(o onet.Onet) handler {
	return handler{}
}

func (h handler) Process(msg proto.Message) (proto.Message, error) {
	switch msg.(type) {
	case *Deal:
	}

	return nil, errors.New("unknown type of message")
}
