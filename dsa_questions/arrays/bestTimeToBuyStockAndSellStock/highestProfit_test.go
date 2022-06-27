package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHighestProfit(t *testing.T) {
	prices := []int{1, 2}
	expected := 1
	output := HighestProfit(prices)
	require.Equal(t, expected, output)
}
