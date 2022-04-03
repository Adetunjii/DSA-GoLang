package LinkedList


import "fmt"

func main() {
	fmt.Println(NewDoublyLinkedList())
}

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}

	ll.InsertBefore(ll.Head, node)
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	if ll.Tail == nil {
		ll.SetHead(node)
		return
	}

	ll.InsertAfter(ll.Tail, node)
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}
	ll.Remove(nodeToInsert)


	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node


	if node.Prev == nil {
		ll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}
	node.Prev = nodeToInsert

}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}
	ll.Remove(nodeToInsert)
	nodeToInsert.Next = node.Next
	nodeToInsert.Prev = node

	if node.Next == nil {
		ll.Tail = nodeToInsert
	} else {
		node.Next.Prev = nodeToInsert
	}
	node.Next = nodeToInsert

}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {

	if position == 1{
		ll.SetHead(nodeToInsert)
	}

	currentNode := ll.Head
	currentPosition := 1

	for currentNode != nil && currentPosition != position {
		currentNode = currentNode.Next
		currentPosition += 1
	}

	if currentNode != nil {
		ll.InsertBefore(currentNode, nodeToInsert)
	} else {
		ll.SetTail(nodeToInsert)
	}
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	if ll.Head == node {
		ll.Head = ll.Head.Next
	}

	if ll.Tail == node {
		ll.Tail = ll.Tail.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Next = nil
	node.Prev = nil
}

// Remove one or more methods with the same value

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	currentNode := ll.Head
	for currentNode != nil {
		nodeToRemove := currentNode
		currentNode = currentNode.Next

		if nodeToRemove.Value == value {
			ll.Remove(nodeToRemove)
			return
		}
	}

}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	currentNode := ll.Head
	for currentNode != nil {
		if currentNode.Value == value {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}
