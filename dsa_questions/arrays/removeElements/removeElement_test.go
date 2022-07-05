package dsaquestions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	array := []int{3, 2, 2, 3}
	val := 3
	expected := 2
	output := RemoveElements(array, val)
	require.Equal(t, expected, output)
}
