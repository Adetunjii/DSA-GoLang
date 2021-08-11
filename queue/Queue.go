package queue

import (
	"errors"

	LinkedList "github.com/Adetunjii/data_structures/linkedList"
)

type Queue struct{}

var list LinkedList.DoublyLinkedList = LinkedList.DoublyLinkedList{};

func (qu *Queue) Size() uint {
	return list.Size();
}

func (qu *Queue) isEmpty() bool {
	return qu.Size() == 0;
}

func (qu *Queue) Peek() interface{} {

	if(qu.isEmpty()) {
		panic(errors.New("queue empty"))
	}

	return list.PeekFirst();
}

func (qu *Queue) enqueue(element interface{}) {
	list.AddLast(element);
}

func (qu *Queue) dequeue(element interface{}) interface {} {
	if qu.isEmpty() {
		panic("queue empty")
	}
	return list.RemoveFirst()
}

