package dsaquestions

func ShuffleTheArray(nums []int, n int) []int {

	for i := n; i < len(nums); i++ {
		nums[i] = (nums[i] * 1024) + nums[i-n]
	}

	index := 0

	for i := n; i < len(nums); i++ {
		nums[index] = nums[i] % 1024
		nums[index+1] = nums[i] / 1024
		index += 2
	}

	return nums
}
