package subarraysort

import (
	"math"
)

func SubArraySort(array []int) []int {

	if len(array) < 2 {
		return nil
	}

	minOutOfOrder, maxOutOfOrder := math.MaxInt32, math.MinInt32

	for i := 0; i < len(array); i++ {
		if isOutOfOrder(i, array[i], array) {
			minOutOfOrder = min(minOutOfOrder, array[i])
			maxOutOfOrder = max(maxOutOfOrder, array[i])
		}
	}

	if minOutOfOrder == math.MaxInt32 {
		return []int{-1, -1}
	}

	leftIdx := 0
	for array[leftIdx] <= minOutOfOrder {
		leftIdx += 1
	}

	rightIdx := len(array) - 1
	for array[rightIdx] >= maxOutOfOrder {
		rightIdx -= 1
	}

	return []int{leftIdx, rightIdx}
}

func isOutOfOrder(i, num int, array []int) bool {
	if i == 0 {
		return num > array[i+1]
	}

	if i == len(array)-1 {
		return num < array[i-1]
	}

	return num < array[i-1] || num > array[i+1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
