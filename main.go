package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5}

	for _, arr := range array {
		fmt.Println(arr)
	}
}
