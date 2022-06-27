package SearchInsertPosition

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchInsertPosition(t *testing.T) {
	input := []int{1, 3, 5, 6}
	target := 2
	expected := 1
	output := SearchInsertPosition(input, target)

	require.Equal(t, expected, output)
}
