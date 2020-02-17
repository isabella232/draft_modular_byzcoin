package globalstate

import "github.com/gogo/protobuf/proto"

type Snapshot interface {
	Read(key []byte) (proto.Message, error)
}
