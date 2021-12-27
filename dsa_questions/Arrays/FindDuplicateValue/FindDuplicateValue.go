package dsaquestions

// Brute Force Solution O(N ^ 2) Time and O(1) Space Time
// func FindDuplicateValue(array []int) int {
// 	minDuplicateIndex := 7
// 	for i := 0; i < len(array); i++ {
// 		for j := i + 1; j < len(array); j++ {
// 			if(array[i] == array[j]) {
// 				minDuplicateIndex = min(j, minDuplicateIndex)
// 			}
// 		}
// 	}

// 	return array[minDuplicateIndex]
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

// func FindDuplicateValue(array []int) int {
// 	counts := make(map[int]int, len(array));
// 	for i := 0; i < len(array); i++ {
// 		 if counts[array[i]] + 1 == 2 {
// 			return array[i]
// 		}
// 		counts[array[i]] += 1
// 	}
// 	return -1
// }

//Optimal Solution
func FindDuplicateValue(array []int) int {
	for i := 0; i < len(array); i++ {
		currentIdx := abs(array[i]) - 1
		if array[currentIdx] < 0 {
			return abs(array[i])
		}

		array[currentIdx] *= -1
	}

	return -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
