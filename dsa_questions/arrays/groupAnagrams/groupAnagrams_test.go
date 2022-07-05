package groupAnagrams

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	output := GroupAnagrams(input)

	require.Equal(t, expected, output)
}
