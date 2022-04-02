package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShuffleTheArray(t *testing.T) {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	expected := []int{2, 3, 5, 4, 1, 7}
	output := ShuffleTheArray(nums, n)
	require.Equal(t, expected, output)
}
