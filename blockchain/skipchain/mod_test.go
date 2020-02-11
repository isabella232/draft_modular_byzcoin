package skipchain

import (
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
)

func validate(b Block) error {
	return nil
}

func TestSkipchain_SimpleScenario(t *testing.T) {
	sc1 := NewSkipchain(validate)
	NewSkipchain(validate)
	NewSkipchain(validate)

	ts := ptypes.TimestampNow()

	err := sc1.Store(ts)
	require.NoError(t, err)

	err = sc1.Store(ts)
	require.NoError(t, err)

	err = sc1.Store(ts)
	require.NoError(t, err)

	proof, err := sc1.GetProof()
	require.NoError(t, err)
	require.NoError(t, proof.Verify())
}
