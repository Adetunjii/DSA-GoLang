package dsaquestions

import "sort"

func MergeOverlappingIntervals(array [][]int) [][]int {
	sort.Slice(array, func(i, j int) bool {
		return array[i][0] < array[j][0]
	})

	mergedIntervals := make([][]int, 0)
	currentInterval := array[0]
	mergedIntervals = append(mergedIntervals, currentInterval)

	for _, nextInterval := range array {
		currentIntervalEnd := currentInterval[1]
		nextIntervalStart, nextIntervalEnd := nextInterval[0], nextInterval[1]

		if currentIntervalEnd >= nextIntervalStart {
			currentInterval[1] = max(currentIntervalEnd, nextIntervalEnd)
		} else {
			currentInterval = nextInterval
			mergedIntervals = append(mergedIntervals, currentInterval)
		}
	}
	return mergedIntervals
}

func max(a, b int) int {
	if a > b{
		return a
 	}
	 return b
}