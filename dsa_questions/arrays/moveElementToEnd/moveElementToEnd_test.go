package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoveElementToEnd(t *testing.T) {
	array := []int{2, 1, 2, 2, 3, 4, 2}
	toMove := 2
	expected := []int{4, 1, 3, 2, 2, 2, 2}
	output := MoveElementToEnd(array, toMove)
	require.Equal(t, expected, output)
}
