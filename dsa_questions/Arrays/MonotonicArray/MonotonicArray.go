package dsaquestions

func IsMonotonicArray(array []int) bool {

	isDecreasing := true
	isIncreasing := true

	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			isDecreasing = false
		}
		if array[i] < array[i+1] {
			isIncreasing = false
		}
	}
	return isIncreasing || isDecreasing
}
