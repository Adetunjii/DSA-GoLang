package linkedList

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (l *LinkedList) add(value int) {
	if l == nil {
		l.Value = value
		l.Next = nil
	}
	currentNode := l

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = &LinkedList{Value: value, Next: nil}
}

func (l *LinkedList) printNodes() {
	currentNode := l
	for currentNode.Next != nil {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}
}
