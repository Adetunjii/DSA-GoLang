package removeKthNodeFromEnd

import "github.com/Adetunjii/data_structures/dsa_questions/linkedList"

func RemoveKthNode(head *linkedList.LinkedList, k int) *linkedList.LinkedList {
	counter := 0
	firstPtr := head
	secondPtr := head

	for counter < k {
		secondPtr = secondPtr.Next
		counter += 1
	}

	if secondPtr == nil {
		head.Value = head.Next.Value
		head.Next = head.Next.Next
		return head
	}

	for secondPtr.Next != nil {
		firstPtr = firstPtr.Next
		secondPtr = secondPtr.Next
	}

	firstPtr.Next = firstPtr.Next.Next

	return head
}
