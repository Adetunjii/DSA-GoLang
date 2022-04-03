package dsaquestions

func RemoveElements(nums []int, val int) int {
	left := 0

	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] != val {
			nums[left] = nums[i]
			left++
		}
	}

	return left
}
