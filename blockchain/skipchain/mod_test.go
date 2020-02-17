package skipchain

import (
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
	"go.dedis.ch/phoenix/blockchain"
)

func validate(b Block) error {
	return nil
}

func TestSkipchain_SimpleScenario(t *testing.T) {
	ro := blockchain.Roster{"a", "b", "c"}
	sc1 := NewSkipchain(ro[0], validate)
	NewSkipchain(ro[1], validate)
	NewSkipchain(ro[1], validate)

	ts := ptypes.TimestampNow()

	err := sc1.Store(ro, ts)
	require.NoError(t, err)

	err = sc1.Store(ro, ts)
	require.NoError(t, err)

	err = sc1.Store(ro, ts)
	require.NoError(t, err)

	proof, err := sc1.GetProof()
	require.NoError(t, err)
	require.NoError(t, proof.Verify())
}
