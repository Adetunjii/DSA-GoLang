package fournumbersum

func FourNumberSum(array []int, target int) [][]int {
	allPairs := make(map[int][][]int)
	result := [][]int{}

	for i := 1; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			currentSum := array[i] + array[j]
			difference := target - currentSum

			if pairs, ok := allPairs[difference]; ok {
				for _, pair := range pairs {
					newQuad := append(pair, array[i], array[j])
					result = append(result, newQuad)
				}
			}
		}

		for k := 0; k < i; k++ {
			currentSum := array[i] + array[k]
			allPairs[currentSum] = append(allPairs[currentSum], []int{array[k], array[i]})
		}
	}

	return result
}
