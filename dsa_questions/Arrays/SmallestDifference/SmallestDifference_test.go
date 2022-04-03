package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSmallestDifference(t *testing.T) {

	array1 := []int{-1, 5, 10, 28, 20, 3}
	array2 := []int{26, 134, 135, 15, 17}
	expected := []int{28, 26}
	output := SmallestDifference(array1, array2)
	require.Equal(t, expected, output)

}
