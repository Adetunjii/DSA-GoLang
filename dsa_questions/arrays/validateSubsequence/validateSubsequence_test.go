package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateSubSequence(t *testing.T) {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	expected := true
	output := IsValidSubsequence(array, sequence)
	require.Equal(t, expected, output)
}
