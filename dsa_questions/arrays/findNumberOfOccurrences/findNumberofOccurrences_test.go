package findNumberOfOccurrences

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindNumberOfOccurrences(t *testing.T) {
	a := []int{13, 27, 35, 40, 49, 55, 59}
	b := []int{17, 35, 39, 40, 55, 58, 60}

	expected := 3
	output := FindNumberOfOccurrences(a, b)

	require.Equal(t, expected, output)
}
