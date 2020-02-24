package pedersen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/util/key"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/onet/local"
)

func TestPedersen_Basic(t *testing.T) {
	onets, addrs, kps, pubkeys := makeRoster(3)
	dkgs := make([]*DKG, 3)

	for i := range dkgs {
		dkgs[i] = New(onets[i], kps[i], pubkeys)
	}

	share, err := dkgs[0].Create(addrs)
	require.NoError(t, err)
	require.NotNil(t, share)
}

func makeRoster(n int) ([]onet.Onet, []*onet.Address, []*key.Pair, []kyber.Point) {
	onets := make([]onet.Onet, n)
	addrs := make([]*onet.Address, n)
	kps := make([]*key.Pair, n)
	pubkeys := make([]kyber.Point, n)

	for i := 0; i < n; i++ {
		addrs[i] = local.NewAddress(fmt.Sprintf("node%d", i))
		onets[i] = local.NewLocalOnet(addrs[i])
		kp := key.NewKeyPair(Suite)
		kps[i] = kp
		pubkeys[i] = kp.Public
	}

	return onets, addrs, kps, pubkeys
}
