package minimumWindowSubString

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinimumWindowSubString(t *testing.T) {
	input := "a"
	sub := "a"
	expected := "a"
	output := MinimumWindowSubString(input, sub)

	require.Equal(t, expected, output)
}
