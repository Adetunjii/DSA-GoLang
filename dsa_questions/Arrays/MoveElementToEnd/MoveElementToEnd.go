package dsaquestions

func MoveElementToEnd(array []int, toMove int) []int {

	startIdx, endIdx := 0, len(array) -1 

	for startIdx < endIdx {
		if array[startIdx] != toMove {
			startIdx += 1
		} else if array[endIdx] == toMove {
			endIdx -= 1
		 } else {
			array[startIdx], array[endIdx] = array[endIdx], array[startIdx]
			startIdx += 1
		 }
	}
	return array
}