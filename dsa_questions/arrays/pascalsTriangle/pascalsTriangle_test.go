package pascalsTriangle

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPascalsTriangle(t *testing.T) {
	numRows := 5
	expected := [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}, []int{1, 3, 3, 1}, []int{1, 4, 6, 4, 1}}
	output := PascalsTriangle(numRows)

	require.Equal(t, expected, output)
}
