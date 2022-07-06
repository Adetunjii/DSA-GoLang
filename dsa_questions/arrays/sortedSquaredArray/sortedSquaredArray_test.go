package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedSquaredArray(t *testing.T) {
	array := []int{-10, -5, 0, 5, 10}
	expected := []int{0, 25, 25, 100, 100}
	output := SortedSquaredArray(array)
	require.Equal(t, expected, output)
}
