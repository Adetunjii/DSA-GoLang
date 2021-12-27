package dsaquestions

func TwoSum(array []int, target int)[]int {

	//O(N²) Time, O(1) Space Time
	// for i := 0; i < len(array); i++ {
	// 	for j := 0; j < len(array); j++ {
	// 		if(i == j) {
	// 			continue;
	// 		}
	// 		if(array[i] + array[j] == target) {
	// 			return []int{array[i], array[j]}
	// 		}
	// 	}
	// }

	// return []int{}

	//Optimal Solution O(N) time, O(N) space
	//hashMap := make(map[int]int)
	//
	//for _, value := range array{
	//	key := target - value
	//	_, ok := hashMap[key];
	//	if ok && key != value {
	//		return []int{key, value}
	//	} else {
	//		hashMap[value] = key
	//	}
	//}
	//
	//return []int{}


	startIdx := 0
	endIdx := len(array) -1

	for startIdx < endIdx {
		if array[startIdx] + array[endIdx] == target {
			return []int{startIdx, endIdx}
		} else if array[startIdx] + array[endIdx] < target {
			startIdx++
		} else {
			endIdx--
		}
	}
	return []int{}
}