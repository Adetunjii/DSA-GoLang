package SearchInsertPosition

func SearchInsertPosition(array []int, target int) int {

	if len(array) == 0 {
		return 0
	}

	left := 0
	right := len(array) - 1

	for left <= right {
		mid := (left + right) / 2

		if array[mid] == target {
			return mid
		} else if target < array[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
