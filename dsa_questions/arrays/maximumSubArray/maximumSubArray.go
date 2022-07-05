package maximumSubArray

// maximumSubArray
/* Suboptimal approach
 * Time Complexity O(N ^ 3), Space Complexity O(1)
 * Literally loop through the array and get the sum of each sub array
 */
//func maximumSubArray(array []int) int {
//
//	maxSum := math.MinInt32
//
//	for i := 0; i < len(array); i++ {
//		for j := i + 1; j < len(array); j++ {
//			windowSum := 0
//			for k := i; k < j; k++ {
//				windowSum += array[k]
//			}
//			maxSum = max(maxSum, windowSum)
//		}
//	}
//
//	return maxSum
//}

// maximumSubArray
/* Suboptimal approach
 * Time Complexity O(N ^ 2), Space Complexity O(1)
 */
//func maximumSubArray(array []int) int {
//	maxSum := math.MinInt32
//
//	for i := 0; i < len(array)-1; i++ {
//		currentSum := 0
//
//		for j := i; j < len(array); j++ {
//
//			currentSum += array[j]
//			maxSum = max(maxSum, currentSum)
//		}
//	}
//
//	return maxSum
//}

// MaximumSubArray
/* Optimal Approach
 * Time Complexity O(N), Space Complexity O(1)
 */
func MaximumSubArray(array []int) int {
	maxSum := array[0]
	currentSum := array[0]

	for i := 1; i < len(array); i++ {
		currentSum = max(array[i], currentSum+array[i])
		maxSum = max(currentSum, maxSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
