package skipchain

import (
	fmt "fmt"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
	"go.dedis.ch/phoenix/blockchain"
	"go.dedis.ch/phoenix/crypto"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/onet/local"
)

type testValidator struct{}

func (v testValidator) Validate(b Block) error {
	return nil
}

func TestSkipchain_SimpleScenario(t *testing.T) {
	onets, addrs := makeRoster(3)

	ro := blockchain.Roster(addrs)
	sc1 := NewSkipchain(onets[0], testValidator{})
	sc2 := NewSkipchain(onets[1], testValidator{})
	sc3 := NewSkipchain(onets[2], testValidator{})

	pubkeys := []crypto.PublicKey{sc1.PublicKey(), sc2.PublicKey(), sc3.PublicKey()}

	ts := ptypes.TimestampNow()

	err := sc1.Store(ro, ts)
	require.NoError(t, err)

	err = sc1.Store(ro, ts)
	require.NoError(t, err)

	err = sc1.Store(ro, ts)
	require.NoError(t, err)

	verifiableBlock, err := sc1.GetBlock()
	require.NoError(t, err)

	block, err := sc1.GetBlockFactory().Create(verifiableBlock, pubkeys)
	require.NoError(t, err)
	require.NotNil(t, block)
}

func makeRoster(n int) ([]onet.Onet, []*onet.Address) {
	onets := make([]onet.Onet, n)
	addrs := make([]*onet.Address, n)

	for i := 0; i < n; i++ {
		addrs[i] = local.NewAddress(fmt.Sprintf("node%d", i))
		onets[i] = local.NewLocalOnet(addrs[i])
	}

	return onets, addrs
}
