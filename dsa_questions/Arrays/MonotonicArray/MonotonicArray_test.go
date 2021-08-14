package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestMonotonicArray(t *testing.T) {
	array := []int{1, 1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 8, 9, 10, 11}
	expected := true
	output := IsMonotonicArray(array)
	require.Equal(t, expected, output)
}