//package main
//
//import (
//	"fmt"
//	"github.com/Adetunjii/data_structures/dsa_questions/LinkedList"
//)
//
//func main() {
//	println("hello world")
//
//	node := &LinkedList.Node{
//		Value: 1,
//		Prev: nil,
//		Next: nil,
//	}
//
//	newnode := &LinkedList.Node{
//		Value: 0,
//		Prev: node,
//		Next: nil,
//	}
//
//	lastnode := &LinkedList.Node{
//		Value: 5,
//		Prev: node,
//		Next: nil,
//	}
//	positionNode := &LinkedList.Node {
//		Value: 9,
//		Prev: newnode,
//		Next: lastnode,
//	}
//
//
//	linkedlist := LinkedList.NewDoublyLinkedList()
//	linkedlist.SetHead(node)
//	linkedlist.InsertBefore(node, newnode)
//	linkedlist.InsertAfter(linkedlist.Tail, lastnode)
//	linkedlist.InsertAtPosition(2, positionNode)
//	linkedlist.RemoveNodesWithValue(0)
//	fmt.Println(linkedlist.ContainsNodeWithValue(1))
//
//		currentNode := linkedlist.Head
//		for currentNode != nil {
//			fmt.Println(currentNode.Value)
//			currentNode = currentNode.Next
//		}
//
//}

package main

import "fmt"

func main() {
	nums := []int{0}

	slice := nums[0:1]

	fmt.Println(slice)
}
