package linkedList

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveKthNode(t *testing.T) {

	input := &LinkedList{
		Value: 0,
		Next:  nil,
	}

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, elem := range arr {
		input.add(elem)
	}

	expected := &LinkedList{
		Value: 0,
		Next:  nil,
	}

	arr2 := []int{1, 2, 3, 4, 6, 7, 8}
	for _, elem := range arr2 {
		expected.add(elem)
	}

	k := 4

	output := RemoveKthNode(input, k)
	expected.printNodes()
	output.printNodes()
	require.Equal(t, expected, output)
}
