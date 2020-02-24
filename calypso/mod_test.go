package calypso

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/util/key"
	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/onet/local"
)

func TestCalypso_Basic(t *testing.T) {
	onets, addrs, kps, pubkeys := makeRoster(3)

	cs := make([]*Calypso, 3)
	for i := range onets {
		cs[i] = NewCalypso(onets[i], addrs, kps[i], pubkeys)
	}

	err := cs[0].AddWrite()
	require.NoError(t, err)

	err = cs[0].AddWrite()
	require.Error(t, err)
	require.Equal(t, "access control denied: forbidden to update", err.Error())
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
