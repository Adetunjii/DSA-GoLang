package findNumberOfOccurrences

func FindNumberOfOccurrences(a, b []int) int {

	counter := 0
	startIdx := 0

	for i := 0; i < len(a); i++ {
		currentIdx := startIdx

		for currentIdx <= i {
			if a[i] == b[currentIdx] {
				counter++
				startIdx += 1
			}
			currentIdx++
		}
	}

	return counter
}
