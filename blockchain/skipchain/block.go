package skipchain

import (
	"crypto/sha256"

	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/crypto"
)

// Block is the representation of the data structures that will be linked
// together.
type Block struct {
	Index     uint64
	Roster    blockchain.Roster
	Signature crypto.Signature
	Data      proto.Message
}

func (b Block) hash() ([]byte, error) {
	h := sha256.New()

	buffer, err := proto.Marshal(b.Data)
	if err != nil {
		return nil, err
	}

	_, err = h.Write(buffer)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Pack returns a network message.
func (b Block) Pack() (proto.Message, error) {
	payload, err := ptypes.MarshalAny(b.Data)
	if err != nil {
		return nil, err
	}

	var metadata *any.Any

	if b.Signature != nil {
		sig, err := b.Signature.Pack()
		if err != nil {
			return nil, err
		}

		sigany, err := ptypes.MarshalAny(sig)
		if err != nil {
			return nil, err
		}

		metadata, err = ptypes.MarshalAny(&BlockMetaData{
			Signature: sigany,
		})
		if err != nil {
			return nil, err
		}
	}

	return &blockchain.Block{
		Index:    b.Index,
		Payload:  payload,
		Metadata: metadata,
	}, nil
}

type blockFactory struct {
	verifier crypto.Verifier
}

func (f blockFactory) FromVerifiable(src *blockchain.VerifiableBlock, pubkeys []crypto.PublicKey) (interface{}, error) {
	var da ptypes.DynamicAny
	err := ptypes.UnmarshalAny(src.Block.GetPayload(), &da)
	if err != nil {
		return Block{}, err
	}

	var metadata BlockMetaData
	err = ptypes.UnmarshalAny(src.Block.GetMetadata(), &metadata)
	if err != nil {
		return Block{}, err
	}

	sig, err := f.verifier.GetSignatureFactory().FromAny(metadata.GetSignature())
	if err != nil {
		return Block{}, err
	}

	block := Block{
		Index:     src.Block.GetIndex(),
		Data:      da.Message,
		Signature: sig,
	}

	hash, err := block.hash()
	if err != nil {
		return block, err
	}

	err = f.verifier.Verify(pubkeys, hash, block.Signature)
	if err != nil {
		return block, err
	}

	return block, nil
}
