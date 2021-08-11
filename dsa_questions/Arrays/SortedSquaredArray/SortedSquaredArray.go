package dsaquestions

func SortedSquaredArray(array []int) []int {


	//Naive Solution O(NLOGN) Time, O(1) Space Time

	// for i, value := range array {
	// 	array[i] = value * value
	// }
	// sort.Ints(array)
	// return array

	//Optimal solution for already sorted input
	// O(N) Time, O(N) Space Time

	sortedArray := make([]int, len(array))
	startIndex := 0
	endIndex := len(array) - 1

	for index := len(array) - 1; index >= 0; index-- {
		start := array[startIndex]
		end := array[endIndex]
		if abs(start) > abs(end) {
			sortedArray[index] = array[startIndex] * array[startIndex]
			startIndex++
		} else {
			sortedArray[index] = array[endIndex] * array[endIndex]
			endIndex--
		}
	} 

	return sortedArray
}

func abs(a int) int {
	if a < 0{
		return -a // -(-ve value) 
	}
	return a
}