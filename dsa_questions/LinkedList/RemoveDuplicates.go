package LinkedList

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	currentNode := linkedList

	for currentNode != nil {
		for currentNode.Next != nil && currentNode.Next.Value == currentNode.Value {
			currentNode.Next = currentNode.Next.Next
		}

		currentNode.Next = currentNode.Next
		currentNode = currentNode.Next
	}

	return linkedList
}