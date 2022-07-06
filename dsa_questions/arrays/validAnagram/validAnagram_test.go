package validAnagram

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	x := "cat"
	y := "rat"
	output := ValidAnagram(x, y)

	require.False(t, output)
}
