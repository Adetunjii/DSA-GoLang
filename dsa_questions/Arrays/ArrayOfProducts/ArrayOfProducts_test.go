package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayOfProducts(t *testing.T) {
	array := []int{5,1,4,2}
	expected := []int{8, 40, 10, 20}
	output := ArrayOfProducts(array)
	require.Equal(t, expected, output)
}