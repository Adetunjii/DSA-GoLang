package subarraysort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubArraySort(t *testing.T) {
	array := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 20}
	expected := []int{3, 9}
	output := SubArraySort(array)
	require.Equal(t, expected, output)
}
