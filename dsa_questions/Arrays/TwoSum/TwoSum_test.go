package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	array := []int{2, 7, 11, 15}
	target := 9

	expected := []int{0, 1}
	output := TwoSum(array, target)
	require.Equal(t, expected, output)
}
