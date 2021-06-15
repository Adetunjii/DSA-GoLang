package Arrays

import (
	"errors"
	"fmt"
)

type DynamicArray struct {
	capacity uint
	len uint
	arr []interface{}
}

var defaultCapacity uint = 16

func (d *DynamicArray) Size() uint {
	return d.len
}

func IndexOutOfBoundsException() error{
	return errors.New("Index out of bounds")
}

func (d *DynamicArray) Get(index uint) (interface{}, error) {

	if(index > d.len || index < 0) {
		return nil, IndexOutOfBoundsException();
	}
	
	return d.arr[index], nil
}

func (d *DynamicArray) Set(index int, elem interface{}) {
	d.arr[index] = elem
}

func (d *DynamicArray) Clear() {
	for i := range d.arr {
		d.arr[i] = nil
	}
	d.len = 0
}

func (d *DynamicArray) Add(element interface{}) {
	if((d.len + 1) >= d.capacity){
		if(d.capacity == 0 ) { 
			d.capacity = defaultCapacity
		} else {
			//bit-shift the capacity i.e (capacity x 2^1)
			d.capacity = d.capacity << 1
		}

		newArray := make([]interface{}, d.capacity)
		fmt.Println(newArray)
		copy(newArray, d.arr)
		
		d.arr = newArray
		d.arr[d.len] = element
		d.len++
	}
	d.arr[d.len] = element;  
	
}


func (d *DynamicArray) RemoveAt(index uint) error {
	if(index >= d.len) {
		return  IndexOutOfBoundsException()
	}

	//get the data
	// data := d.arr[index]



	// copy the element from the next index to the end 
	// to the index we're removing from 
	copy(d.arr[index: ], d.arr[index +1: d.len])

	fmt.Println(d.arr)

	return nil;
}

func (d *DynamicArray) Print() {
	fmt.Print(d.arr)
}
