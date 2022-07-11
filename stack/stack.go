package stack

import (
	"errors"

	LinkedList "github.com/Adetunjii/data_structures/LinkedList"
)

var list LinkedList.DoublyLinkedList = LinkedList.DoublyLinkedList{}

type Stack struct {
}

func (st *Stack) Size() uint {
	return list.Size()
}

func (st *Stack) isEmpty() bool {
	return st.Size() == 0
}

func (st *Stack) Push(element interface{}) {
	list.AddLast(element)
}

func (st *Stack) Pop() interface{} {
	if st.isEmpty() {
		panic(errors.New("empty stack"))
	}

	return list.RemoveLast()
}

func (st *Stack) Peek() interface{} {
	if st.isEmpty() {
		panic(errors.New("empty stack"))
	}

	return list.PeekLast()
}

func (st *Stack) Search(obj interface{}) int {
	return list.IndexOf(obj)
}

func (st *Stack) GetAllItems() {
	list.GetAllItems()
}
