package skipchain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type value struct {
	str string
}

func (v value) MarshalBinary() ([]byte, error) {
	return []byte(v.str), nil
}

func TestSkipchain_SimpleScenario(t *testing.T) {
	sc1 := NewSkipchain()
	NewSkipchain()
	NewSkipchain()

	err := sc1.Store(value{str: "this is a simple text"})
	require.NoError(t, err)

	err = sc1.Store(value{str: "this is a simple text"})
	require.NoError(t, err)

	err = sc1.Store(value{str: "this is a simple text"})
	require.NoError(t, err)

	proof, err := sc1.GetProof()
	require.NoError(t, err)
	require.NoError(t, proof.Verify())
}
