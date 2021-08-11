package main

import (
	"fmt"
	"sort"
)

func main() {

	array := []int{10,-5,0,5,10}

	sort.Ints(array)

	fmt.Println(array)
}