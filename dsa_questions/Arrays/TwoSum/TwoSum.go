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
	hashMap := make(map[int]int)
	
	for _, value := range array{
		key := target - value
		_, ok := hashMap[key];
		if ok && key != value {
			return []int{key, value}
		} else {
			hashMap[value] = key
		}
	}

	return []int{}
}