package MaximumSubArray

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaximumSubArray(t *testing.T) {

	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6
	output := MaximumSubArray(input)

	require.Equal(t, expected, output)
}
