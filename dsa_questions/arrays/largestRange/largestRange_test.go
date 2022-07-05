package largestrange

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLargestRange(t *testing.T) {
	array := []int{19, -1, 18, 17, 2, 10, 3, 12, 5, 16, 4, 11, 8, 7, 6, 15, 12, 12, 2, 1, 6, 13, 14}
	expected := []int{10, 19}
	output := LargestRange(array)

	require.Equal(t, expected, output)
}
