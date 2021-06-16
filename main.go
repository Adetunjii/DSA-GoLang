package main

import (
	"fmt"
	Arrays "github.com/Adetunjii/data_structures/Arrays"
)

func main() {
	dynamicArray := Arrays.DynamicArray{}
	dynamicArray.Add(3);
	dynamicArray.Add(4);
	dynamicArray.Add(6);


	fmt.Println(dynamicArray.GetAll());
	fmt.Println(dynamicArray.Contains(3))


}