package fournumbersum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFourNumberSum(t *testing.T) {
	array := []int{7, 6, 4, -1, 1, 2}
	target := 16
	expected := [][]int{{7, 6, 4, -1}, {7, 6, 1, 2}}
	output := FourNumberSum(array, target)
	require.Equal(t, output, expected)
}
