package main

import (
	"fmt"

	Arrays "github.com/Adetunjii/data_structures/Arrays"
)

func main() {
	dynamicArray := Arrays.DynamicArray{}
	dynamicArray.Add(3);
	dynamicArray.Add(4);
	dynamicArray.Add(5);

	fmt.Println(dynamicArray.Contains(3))


}