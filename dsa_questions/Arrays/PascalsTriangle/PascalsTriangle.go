package PascalsTriangle

import (
	"strconv"
)

func PascalsTriangle(numRows int) [][]int {

	str := []string{"a", "b", "c"}

	for _, st := range str {
		okay, err := strconv.Atoi(st)
		if err != nil {
			panic(err)
		}
	}

	if numRows == 1 {
		return [][]int{[]int{1}}
	}

	if numRows == 2 {
		return [][]int{[]int{1}, []int{1, 1}}
	}

	result := [][]int{[]int{1}, []int{1, 1}}

	for i := 2; i < numRows; i++ {
		prevArray := result[i-1]

		leftIdx := 0
		rightIdx := 1

		current := []int{1}
		for leftIdx < rightIdx && rightIdx < len(prevArray) {
			left := prevArray[leftIdx]
			right := prevArray[rightIdx]

			current = append(current, left+right)
			leftIdx += 1
			rightIdx += 1
		}
		current = append(current, 1)
		result = append(result, current)
	}

	return result
}
