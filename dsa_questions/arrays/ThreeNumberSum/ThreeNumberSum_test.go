package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThreeNumberSum(t *testing.T) {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	targetSum := 0

	expected := [][]int{
		{-8, 2, 6},
		{-8, 3, 5},
		{-6, 1, 5},
	}
	output := ThreeNumberSum(array, targetSum)
	require.Equal(t, expected, output)
}
