package cosi

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/sign/bls"
	"go.dedis.ch/kyber/v3/util/key"
	"go.dedis.ch/phoenix/onet"
)

//go:generate protoc -I ./ --go_out=./ ./types.proto

var suite = pairing.NewSuiteBn256()

// Verifier is the function used to make sure a signature matches the message
// with a specific list of identities.
type Verifier func(roster []onet.Identity, msg []byte, sig []byte) error

// Signature is the response type of a collective signing protocol.
type Signature []byte

// CollectiveSigning is the interface that provides the primitives to sign
// a message by members of a network.
type CollectiveSigning interface {
	Sign(msg proto.Message) (Signature, error)
	MakeVerifier() Verifier
}

// Validator is the interface that is used to validate a block.
type Validator interface {
	Validate(msg proto.Message) ([]byte, error)
}

// BlsCoSi is an implementation of the collective signing interface by
// using BLS signatures.
type BlsCoSi struct {
	onet onet.RPC
}

// NewBlsCoSi returns a new collective signing instance.
func NewBlsCoSi(o onet.Onet, v Validator) *BlsCoSi {
	return &BlsCoSi{
		onet: o.MakeRPC("cosi", newHandler(o, v)),
	}
}

// Sign returns the collective signature of the block.
func (cosi *BlsCoSi) Sign(msg proto.Message) (Signature, error) {
	data, err := ptypes.MarshalAny(msg)
	if err != nil {
		return nil, err
	}

	msgs, err := cosi.onet.Collect(&SignatureRequest{Message: data})
	if err != nil {
		return nil, err
	}

	var agg []byte
	ok := true
	var resp proto.Message
	for ok {
		resp, ok = <-msgs
		if ok {
			reply := resp.(*SignatureResponse)

			if agg == nil {
				agg = reply.GetSignature()
			} else {
				agg, err = bls.AggregateSignatures(suite, agg, reply.GetSignature())
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return agg, nil
}

// MakeVerifier returns a verifier that can be used to verify signatures
// from this collective signing.
func (cosi *BlsCoSi) MakeVerifier() Verifier {
	return blsVerifier
}

// BlsVerifier verifies that a signature matches the message for the roster public keys.
func blsVerifier(roster []onet.Identity, msg []byte, sig []byte) error {
	points := make([]kyber.Point, 0)

	for _, identity := range roster {
		points = append(points, identity.(*key.Pair).Public)
	}

	publicKey := bls.AggregatePublicKeys(suite, points...)

	err := bls.Verify(suite, publicKey, msg, sig)
	if err != nil {
		return err
	}

	return nil
}
