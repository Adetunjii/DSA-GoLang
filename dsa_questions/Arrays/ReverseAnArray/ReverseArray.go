package dsaquestions

import "sort"

func ReverseArray(array []int) []int {
	sort.Ints(array);
	startIdx := 0
	endIdx := len(array) - 1

	for startIdx < endIdx {
			array[endIdx], array[startIdx] = array[startIdx], array[endIdx]
			startIdx++
			endIdx--
	}
	return array
}
