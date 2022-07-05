package dsaquestions

//func removeDuplicatesFromSortedArray(nums []int) int {
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
	counter := 1

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		counter += 1
	}
	return counter
}
