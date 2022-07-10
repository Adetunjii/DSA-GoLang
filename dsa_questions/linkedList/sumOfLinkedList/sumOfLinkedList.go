package sumOfLinkedList

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func SumOfLinkedList(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	carry := 0
	result := &LinkedList{}
	currentNode := result

	for linkedListOne != nil || linkedListTwo != nil || carry != 0 {
		var valueOne, valueTwo int

		if linkedListOne != nil {
			valueOne = linkedListOne.Value
		}

		if linkedListTwo != nil {
			valueTwo = linkedListTwo.Value
		}

		sum := valueOne + valueTwo + carry

		newValue := sum % 10
		currentNode.Next = &LinkedList{Value: newValue}
		currentNode = currentNode.Next

		carry = sum / 10

		if linkedListOne != nil {
			linkedListOne = linkedListOne.Next
		}

		if linkedListTwo != nil {
			linkedListTwo = linkedListTwo.Next
		}

	}

	return result
}
