package dsaquestions

func RemoveDuplicatesFromSortedArray(nums []int) int {
	leftPointer := 1

	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] != nums[i-1] {
			nums[leftPointer] = nums[i]
			leftPointer++
		}
	}

	return leftPointer
}
