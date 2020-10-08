package main

// LinkedList struct
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// ReverseLinkedList Time complexity : T: O(n) S: O(1)
func ReverseLinkedList(head *LinkedList) *LinkedList {
	var currentNode, previousNode *LinkedList
	currentNode = head
	previousNode = nil

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	return previousNode
}
