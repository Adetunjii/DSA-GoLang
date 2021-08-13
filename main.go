package main

import (
	"fmt"
	"math"
	"sort"
)

const HOME_TEAM_WON = 1
const AWAY_TEAM_WON = 0

func main() {

	fmt.Print(math.MinInt32)

}



func ThreeNumberSum(array []int, targetSum int){

	sort.Ints(array)
	
	startIndex := 0
	currentOutputIndex := 0
	outputArray := make([][]int, 3)


	for idx := len(array) - 1; idx > startIndex; idx-- {

		if currentOutputIndex == 3 {
			break;
		}

		currentSum := array[startIndex] + array[idx]
		lastNumber := targetSum - currentSum
		lastNumberIndex := isExists(array[startIndex : idx], lastNumber)

		foundArray := []int{}

		if(lastNumberIndex != -1) {	
			foundArray = []int{array[startIndex], array[idx], array[lastNumberIndex]}
		}
		sort.Ints(foundArray)
		outputArray[currentOutputIndex] =  foundArray


		if(len(foundArray) == 3) {
			currentOutputIndex += 1
			continue;
		} 

		fmt.Println("startIndex", startIndex)
		fmt.Println("currentIndex", idx)
		
	}

	fmt.Println(outputArray)
}

func isExists(array []int, number int) int {
	for index := 0; index < len(array); index++ {
		if(array[index] == number) {
			return index;
		}
	}
	return -1
}