package LinkedList

import (
	"errors"
	"fmt"

	"github.com/Adetunjii/data_structures/Exceptions"
)

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	size uint
	head *Node
	tail *Node
}

func (dl *DoublyLinkedList) Size() uint {
	return dl.size
}

func (dl *DoublyLinkedList) isEmpty() bool {
	return dl.Size() == 0
}

func (dl *DoublyLinkedList) AddFirst(T interface{}) {
	if dl.isEmpty() {
		dl.head = &Node{T, nil, nil}
		dl.tail = &Node{T, nil, nil}
	} else {
		dl.head.prev = &Node{T, nil, dl.head}
		dl.head = dl.head.prev
	}

	dl.size++
}

func (dl *DoublyLinkedList) AddLast(T interface{}) {
	if dl.isEmpty() {
		dl.head = &Node{T, nil, nil}
		dl.tail = &Node{T, nil, nil}
	} else {
		dl.tail.next = &Node{data: T, prev: dl.tail, next: nil}
		dl.tail = dl.tail.next
	}
	dl.size++
}

func (dl *DoublyLinkedList) PeekFirst() interface{} {
	if dl.isEmpty() {
		panic(errors.New("empty list"))
	}

	return dl.head
}

func (dl *DoublyLinkedList) PeekLast() interface{} {
	if dl.isEmpty() {
		panic(errors.New("empty list"))
	}

	return dl.tail
}

func (dl *DoublyLinkedList) RemoveFirst() interface{} {
	if dl.isEmpty() {
		panic(errors.New("empty list"))
	}

	data := dl.head.data
	dl.head = dl.head.next
	dl.size--

	if dl.isEmpty() {
		dl.tail = nil
	} else {
		dl.head.prev = nil
	}

	return data
}

func (dl *DoublyLinkedList) RemoveLast() interface{} {
	if dl.isEmpty() {
		panic(errors.New("empty list"))
	}

	data := dl.tail.data
	dl.tail = dl.tail.prev
	dl.size--

	if dl.isEmpty() {
		dl.head = nil
	} else {
		dl.tail.next = nil
	}

	return data
}

func (dl *DoublyLinkedList) Remove(node *Node) interface{} {
	if node.prev == nil {
		return dl.RemoveFirst()
	}
	if node.next == nil {
		return dl.RemoveLast()
	}

	node.next.prev = node.prev
	node.prev.next = node.next

	data := node.data
	node.data = nil

	node = nil
	node.prev = nil
	node.next = nil

	dl.size--
	return data
}

func (dl *DoublyLinkedList) RemoveAt(index uint) interface{} {
	if index < 0 || index >= dl.size {
		panic(Exceptions.IndexOutOfBoundsException())
	}

	var i uint
	var trav *Node

	if index < dl.size/2 {
		for i, trav = 0, dl.head; i != index; i++ {
			trav = trav.next
		}
	} else {
		for i, trav = dl.size-1, dl.tail; i != index; i-- {
			trav = trav.prev
		}
	}

	return dl.Remove(trav)
}

func (dl *DoublyLinkedList) RemoveObject(obj interface{}) bool {
	trav := dl.head

	if obj == nil {
		for trav = dl.head; trav != nil; trav = trav.next {
			if trav.data == nil {
				dl.Remove(trav)
				return true
			}
		}
	} else {
		for trav = dl.head; trav != nil; trav = trav.next {
			if obj == trav.data {
				dl.Remove(trav)
				return true
			}
		}
	}

	return false
}

func (dl *DoublyLinkedList) IndexOf(obj interface{}) int {
	index := 0
	trav := dl.head

	if obj == nil {
		for trav = dl.head; trav != nil; trav, index = trav.next, index+1 {
			if trav.data == nil {
				return index
			}
		}
	} else {
		for trav = dl.head; trav != nil; trav, index = trav.next, index+1 {
			if obj == trav.data {
				return index
			}
		}
	}
	return -1
}

func (dl *DoublyLinkedList) GetAllItems() {
	for trav := dl.head; trav != nil; trav = trav.next {
		fmt.Println(trav.data)
	}
}
