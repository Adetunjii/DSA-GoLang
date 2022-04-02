package Arrays

func ReverseArray(array []int) []int {
	// split the array into two
	halfIdx := len(array) / 2
	leftIdx := 0
	rightIdx := len(array) - 1

	for leftIdx < halfIdx && rightIdx > halfIdx {
		array[leftIdx], array[rightIdx] = array[rightIdx], array[leftIdx]

		leftIdx += 1
		rightIdx -= 1
	}

	return array
}
