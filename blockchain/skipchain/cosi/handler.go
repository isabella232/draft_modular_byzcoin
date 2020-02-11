package cosi

import (
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go.dedis.ch/kyber/v3/sign/bls"
	"go.dedis.ch/kyber/v3/util/key"
	"go.dedis.ch/phoenix/onet"
)

type handler struct {
	onet      onet.Onet
	validator Validator
}

func newHandler(o onet.Onet, v Validator) handler {
	return handler{
		onet:      o,
		validator: v,
	}
}

func (h handler) Process(msg proto.Message) (proto.Message, error) {
	switch req := msg.(type) {
	case *SignatureRequest:
		var da ptypes.DynamicAny
		err := ptypes.UnmarshalAny(req.Message, &da)
		if err != nil {
			return nil, err
		}

		buf, err := h.validator.Validate(da.Message)
		if err != nil {
			return nil, err
		}

		id := h.onet.Identity().(*key.Pair)

		sig, err := bls.Sign(suite, id.Private, buf)
		if err != nil {
			return nil, err
		}

		return &SignatureResponse{Signature: sig}, nil
	}

	return nil, errors.New("unknown type of message")
}
