package dsaquestions

func ArrayOfProducts(array []int) []int {

	products := make([]int, len(array))
	leftIdx := 0

	leftProducts := 0

	for i := range array {
		hasLeft := i > leftIdx

		for hasLeft {
			currentLeftIndex := i - 1
			leftProducts *= array[currentLeftIndex]
			currentLeftIndex -= 1
		}
	}

	return products
}
