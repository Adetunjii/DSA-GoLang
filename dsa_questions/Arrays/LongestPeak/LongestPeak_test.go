package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestPeak(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 0, 10, 6, 5, -1, -3, -4, -12, -45, 2, 3}
	expected := 9
	output := LongestPeak(array)
	require.Equal(t, expected, output)
}
