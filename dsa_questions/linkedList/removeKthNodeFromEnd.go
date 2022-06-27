package linkedList

func RemoveKthNode(head *LinkedList, k int) *LinkedList {
	counter := 1
	firstPtr := head
	secondPtr := head

	for counter <= k {
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
