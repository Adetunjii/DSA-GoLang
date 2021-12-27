package dsaquestions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeSortedArrays(t *testing.T) {
	nums1 := []int{4,5,6,0,0,0}
	nums2 := []int{1,2,3}
	m := 1
	n := 1
	expected := []int{1,2,3,4,5,6}
	output := MergeSortedArrays(nums1, nums2, m, n)
	require.Equal(t, expected, output)
}  