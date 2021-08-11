package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	array := []int{3,5,-4,8,11,1,-1,6}
	target := 10

	expected := []int{11,-1}
	output := TwoSum(array, target)
	require.Equal(t, expected, output) 
}