package skipchain

import (
	fmt "fmt"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/onet/local"
	"go.dedis.ch/phoenix/types"
)

type testValidator struct{}

func (v testValidator) Validate(b Block) error {
	return nil
}

func TestSkipchain_SimpleScenario(t *testing.T) {
	onets, addrs := makeRoster(3)

	ro := blockchain.Roster(addrs)
	sc1 := NewSkipchain(onets[0], testValidator{})
	NewSkipchain(onets[1], testValidator{})
	NewSkipchain(onets[2], testValidator{})

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

func makeRoster(n int) ([]onet.Onet, []*types.Address) {
	onets := make([]onet.Onet, n)
	addrs := make([]*types.Address, n)

	for i := 0; i < n; i++ {
		addrs[i] = local.NewAddress(fmt.Sprintf("node%d", i))
		onets[i] = local.NewLocalOnet(addrs[i])
	}

	return onets, addrs
}
