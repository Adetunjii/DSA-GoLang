//package main
//
//import (
//	"fmt"
//	"github.com/Adetunjii/data_structures/dsa_questions/linkedList"
//)
//
//func main() {
//	println("hello world")
//
//	node := &linkedList.Node{
//		Value: 1,
//		Prev: nil,
//		Next: nil,
//	}
//
//	newnode := &linkedList.Node{
//		Value: 0,
//		Prev: node,
//		Next: nil,
//	}
//
//	lastnode := &linkedList.Node{
//		Value: 5,
//		Prev: node,
//		Next: nil,
//	}
//	positionNode := &linkedList.Node {
//		Value: 9,
//		Prev: newnode,
//		Next: lastnode,
//	}
//
//
//	linkedlist := linkedList.NewDoublyLinkedList()
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
