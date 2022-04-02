package Arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArray(t *testing.T) {

	input := []int{1, 2, 3, 4, 5, 6, 7}
	expected := []int{7, 6, 5, 4, 3, 2, 1}
	output := ReverseArray(input)
	require.Equal(t, output, expected)

}
