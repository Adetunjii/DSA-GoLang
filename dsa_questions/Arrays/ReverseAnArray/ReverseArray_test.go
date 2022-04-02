package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseArray(t *testing.T) {
	array := []int{1, 4, 5, 0, 9, 11}
	expected := []int{11, 9, 5, 4, 1, 0}
	output := ReverseArray(array)
	require.Equal(t, expected, output)
}
