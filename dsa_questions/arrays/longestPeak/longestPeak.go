package dsaquestions

func LongestPeak(array []int) int {

	longestPeak := 0
	i := 1

	for i < len(array)-1 {
		isPeak := array[i-1] < array[i] && array[i+1] < array[i]

		if !isPeak {
			i += 1
			continue
		}
		leftIndex := i - 2
		for leftIndex >= 0 && array[leftIndex] < array[leftIndex+1] {
			leftIndex -= 1
		}
		rightIndex := i + 2
		for rightIndex <= len(array) && array[rightIndex] < array[rightIndex-1] {
			rightIndex += 1
		}
		currentLongestPeak := rightIndex - leftIndex - 1
		if currentLongestPeak > longestPeak {
			longestPeak = currentLongestPeak
		}
		i = rightIndex
	}
	return longestPeak
}
