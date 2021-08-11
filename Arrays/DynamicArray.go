package arrays

import Exception "github.com/Adetunjii/data_structures/Exceptions"


type DynamicArray struct {
	capacity uint
	len uint
	arr []interface{}
}

var defaultCapacity uint = 16

func (d *DynamicArray) Size() uint {
	return d.len
}

func (d *DynamicArray) Get(index uint) (interface{}, error) {

	if(index > d.len) {
		return nil, Exception.IndexOutOfBoundsException();
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
		
		if d.len > 0 {
			copy(newArray, d.arr);
		}
		
		d.arr = newArray
		d.arr[d.len] = element
	}
	d.arr[d.len] = element;  
	d.len++
}

func (d *DynamicArray) RemoveAt(index uint) (interface{}, error) {
	if(index >= d.len) {
		return  nil, Exception.IndexOutOfBoundsException()
	}

	data := d.arr[index]

	copy(d.arr[index: ], d.arr[index + 1 : ])
	newArray := make([]interface{}, d.capacity - 1)
	copy(newArray, d.arr);

	
	return data, nil;
}

func (d * DynamicArray) IndexOf(object interface{}) int {
	for index, item := range d.arr {
		if item == object {
			return index;
		}
	}
	return -1
} 

func (d *DynamicArray) GetAll() []interface{} {
	return d.arr[ : d.len];
}

func (d *DynamicArray) Contains(object interface{}) bool {
	return d.IndexOf(object) != -1;
}