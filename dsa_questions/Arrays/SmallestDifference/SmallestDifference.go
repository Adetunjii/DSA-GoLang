package dsaquestions

import (
	"fmt"
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {

	sort.Ints(array1)
	sort.Ints(array2)


	current := math.MaxInt32
	result := []int{}

	leftIdx, rightIdx := 0, 0

	for leftIdx < len(array1) && rightIdx < len(array2) {

		fmt.Println(current)
		left := array1[leftIdx]
		right := array2[rightIdx]

		if left < right {
			current = right - left
			leftIdx += 1
		} else if left > right {
			current = left - right
			rightIdx += 1
		} else {
			return []int{left, right}
		}
	}
	return result
}

//	sort.Ints(array1)
//	sort.Ints(array2)
//
//	smallest, current := math.MaxInt32, math.MaxInt32
//	idx1, idx2 := 0,0
//	output := []int{}
//
//	for idx1 < len(array1) && idx2 < len(array2) {
//		left := array1[idx1]
//		right := array2[idx2]
//
//		if left < right {
//			current = right - left
//			idx1 += 1
//		} else if left > right {
//			current = left - right
//			idx2 += 1
//		} else {
//			return []int{left, right}
//		}
//
//		if smallest > current {
//			smallest = current
//			output = []int{left, right}
//		}
//	}
//	return output
//}
//
//
//
