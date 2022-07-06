package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := 1
	output := GetNthFib(2)
	require.Equal(t, expected, output)
}
