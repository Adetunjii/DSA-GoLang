package largestrange

// import (
// 	"sort"
// )

// ONLOGN - SubOptimal solution
// func LargestRange(array []int) []int {k
// 	sort.Ints(array)
// 	hashmap := make(map[int]int)
// 	smallestIdx := 0
// 	largestDifferenceKey := 0

// 	if len(array) < 2 {
// 		return []int{array[0], array[0]}
// 	}

// 	for i := 0; i < len(array)-1; i++ {
// 		if array[i+1]-array[i] > 1 {
// 			hashmap[array[smallestIdx]] = array[i]
// 			smallestIdx = i + 1
// 		}
// 		hashmap[array[smallestIdx]] = array[i+1]
// 	}

// 	currentDifference := 0
// 	largestDifference := 0
// 	for key, value := range hashmap {
// 		currentDifference = value - key
// 		if currentDifference > largestDifference {
// 			largestDifference = currentDifference
// 			largestDifferenceKey = key
// 		}
// 	}
// 	return []int{largestDifferenceKey, hashmap[largestDifferenceKey]}
// }

// O(N) Time --- O(N) Space => Optimal Solution
func LargestRange(array []int) []int {
	hashmap := make(map[int]bool)
	largestRange := 0
	result := []int{}

	if len(array) < 2 {
		return ([]int{array[0], array[0]})
	}

	for _, num := range array {
		hashmap[num] = true
	}

	for _, num := range array {
		if !hashmap[num] {
			continue
		}

		currentLength, left, right := 0, num-1, num+1

		for hashmap[left] {
			hashmap[left] = false
			currentLength += 1
			left -= 1
		}

		for hashmap[right] {
			hashmap[right] = false
			currentLength += 1
			right += 1
		}

		if currentLength > largestRange {
			largestRange = currentLength
			result = []int{left + 1, right - 1}
		}
	}

	return result
}
