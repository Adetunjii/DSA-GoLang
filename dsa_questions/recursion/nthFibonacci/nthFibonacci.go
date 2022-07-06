package dsaquestions

//NAIVE APPROACH | O(2ⁿ) Time, O(n)Space Time
func GetNthFib(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return GetNthFib(n-1) + GetNthFib(n-2)
}

//MEMOIZATION | O(N) Time, O(n) Space Time
// func GetNthFib(n int) int {
// 	memoize := map[int]int {
// 		1: 0,
// 		2: 1,
// 	}

// 	if _, ok := memoize[n]; ok {
// 		return memoize[n]
// 	}
// 	memoize[n] = GetNthFib(n - 1) + GetNthFib(n - 2)
// 	return memoize[n]
// }

//ITERATIVE SOLUTION O(n) Time, O(1) Space Time
// func GetNthFib(n int)int {
// 	lastTwo := []int{0,1}

// 	if(n == 1) {return lastTwo[0]}

// 	for counter := 3; counter <= n; counter++ {
// 		nextFib := lastTwo[0] + lastTwo[1]
// 		lastTwo[0] = lastTwo[1]
// 		lastTwo[1] = nextFib
// 	}

// 	return lastTwo[1]
// }
