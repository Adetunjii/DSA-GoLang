package dsaquestions

import "sort"

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)

	constructibleChange := 0
	if len(coins) == 0 {
		return 1
	}
	for _, coin := range coins {
		constructibleChange += coin
		if coin > constructibleChange+1 {
			return constructibleChange + 1
		}
	}

	return constructibleChange + 1
}
