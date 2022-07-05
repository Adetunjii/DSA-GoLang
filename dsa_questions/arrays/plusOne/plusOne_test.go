package plusOne

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlusOne(t *testing.T) {
	input := []int{9, 9, 9}
	expected := []int{1, 0, 0, 0}
	output := PlusOne(input)

	require.Equal(t, expected, output)
}
