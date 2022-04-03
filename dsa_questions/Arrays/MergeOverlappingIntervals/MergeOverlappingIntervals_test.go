package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeOverlappingIntervals(t *testing.T) {
	input := [][]int{{1, 2}, {3, 5}, {4, 7}, {6, 8}, {9, 10}}
	expected := [][]int{{1, 2}, {3, 8}, {9, 10}}
	output := MergeOverlappingIntervals(input)
	require.Equal(t, expected, output)
}
