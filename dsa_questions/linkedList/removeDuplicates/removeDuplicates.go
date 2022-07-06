package removeDuplicates

import "github.com/Adetunjii/data_structures/dsa_questions/linkedList"

func RemoveDuplicatesFromLinkedList(linkedList *linkedList.LinkedList) *linkedList.LinkedList {
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
