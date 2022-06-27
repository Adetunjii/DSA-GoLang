package dsaquestions

import "sort"

func ThreeNumberSum(array []int, targetSum int) [][]int {

	sort.Ints(array)
	output := make([][]int, 0)

	for index := 0; index < len(array)-2; index++ {
		left := index + 1
		right := len(array) - 1

		for left < right {

			currentSum := array[index] + array[left] + array[right]

			if currentSum == targetSum {
				output = append(output, []int{array[index], array[left], array[right]})
				left += 1
				right -= 1
			} else if currentSum < targetSum {
				left += 1
			} else if currentSum > targetSum {
				right -= 1
			}
		}
	}
	return output
}
