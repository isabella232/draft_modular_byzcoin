package cosi

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/util/key"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/crypto"
	"go.dedis.ch/phoenix/crypto/bls"
	"go.dedis.ch/phoenix/onet"
)

//go:generate protoc -I ./ --go_out=./ ./types.proto

var suite = pairing.NewSuiteBn256()

// CollectiveSigning is the interface that provides the primitives to sign
// a message by members of a network.
type CollectiveSigning interface {
	PublicKey() crypto.PublicKey
	Sign(ro blockchain.Roster, msg proto.Message) (crypto.Signature, error)
	MakeVerifier() crypto.Verifier
}

// Validator is the interface that is used to validate a block.
type Validator interface {
	Validate(msg proto.Message) ([]byte, error)
}

// BlsCoSi is an implementation of the collective signing interface by
// using BLS signatures.
type BlsCoSi struct {
	rpc    onet.RPC
	signer crypto.AggregateSigner
}

// NewBlsCoSi returns a new collective signing instance.
func NewBlsCoSi(o onet.Onet, v Validator) *BlsCoSi {
	kp := key.NewKeyPair(pairing.NewSuiteBn256())
	signer := bls.NewSigner(kp)

	return &BlsCoSi{
		rpc:    o.MakeRPC("cosi", newHandler(o, signer, v)),
		signer: signer,
	}
}

// PublicKey returns the public key for this instance.
func (cosi *BlsCoSi) PublicKey() crypto.PublicKey {
	return cosi.signer.PublicKey()
}

// Sign returns the collective signature of the block.
func (cosi *BlsCoSi) Sign(ro blockchain.Roster, msg proto.Message) (crypto.Signature, error) {
	data, err := ptypes.MarshalAny(msg)
	if err != nil {
		return nil, err
	}

	msgs, err := cosi.rpc.Call(&SignatureRequest{Message: data}, ro...)
	if err != nil {
		return nil, err
	}

	var agg crypto.Signature
	ok := true
	var resp proto.Message
	for ok {
		resp, ok = <-msgs
		if ok {
			reply := resp.(*SignatureResponse)
			sig, err := cosi.signer.GetSignatureFactory().FromAny(reply.GetSignature())
			if err != nil {
				return nil, err
			}

			if agg == nil {
				agg = sig
			} else {
				agg, err = cosi.signer.Aggregate(agg, sig)
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
func (cosi *BlsCoSi) MakeVerifier() crypto.Verifier {
	return bls.NewVerifier()
}
