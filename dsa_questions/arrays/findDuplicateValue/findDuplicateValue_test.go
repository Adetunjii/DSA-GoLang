package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDuplicateValue(t *testing.T) {
	array := []int{2, 1, 5, 3, 3, 2, 4}
	expected := 3
	output := FindDuplicateValue(array)
	require.Equal(t, expected, output)
}
