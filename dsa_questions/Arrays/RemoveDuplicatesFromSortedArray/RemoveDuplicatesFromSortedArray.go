package dsaquestions

//func RemoveDuplicatesFromSortedArray(nums []int) int {
//	leftPointer := 1
//
//	for i := 1; i <= len(nums)-1; i++ {
//		if nums[i] != nums[i-1] {
//			nums[leftPointer] = nums[i]
//			leftPointer++
//		}
//	}
//
//	return leftPointer
//}

func RemoveDuplicatesFromSortedArray(nums []int) int {
	l := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[l] = nums[i]
			l++
		}
	}

	return l
}
