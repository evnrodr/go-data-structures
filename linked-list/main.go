package main

import "fmt"

type node struct {
	next *node
	data int
}

type linkedlist struct {
	head *node
	size int
}

func createNode(value int) *node {
	return &node{data: value, next: nil}
}

func insertAtBeginning(linkedlist *linkedlist, value int) {
	newNode := createNode(value)
	newNode.next = linkedlist.head
	linkedlist.head = newNode
	linkedlist.size++
}

func insertAtEnd(linkedlist *linkedlist, value int) {
	newNode := createNode(value)

	if linkedlist.head == nil {
		linkedlist.head = newNode
		return
	}

	currentNode := linkedlist.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = newNode
	linkedlist.size++
}

func insertAtPosition(linkedlist *linkedlist, position int, value int) {
	if linkedlist.size == 0 {
		fmt.Println("Linked list is currently empty. Inserting at beginning.")
		insertAtBeginning(linkedlist, value)
		return
	}

	if position <= 0 {
		fmt.Println("Given position is equal or below 0. Inserting at beginning.")
		insertAtBeginning(linkedlist, value)
		return
	}

	if position > linkedlist.size {
		fmt.Println("Given position is bigger than linked list size. Inserting at end.")
		insertAtEnd(linkedlist, value)
		return
	}

	newNode := createNode(value)
	currentPosition := 0
	currentNode := linkedlist.head

	for currentPosition < position-1 {
		currentNode = currentNode.next
		currentPosition++
	}

	remainingNodes := currentNode.next
	currentNode.next = newNode
	newNode.next = remainingNodes
	linkedlist.size++
}

func removeFromStart(linkedlist *linkedlist) {
	if linkedlist.size == 0 {
		fmt.Println("Linked list is currently empty, there is nothing to remove.")
		return
	}

	linkedlist.head = linkedlist.head.next
	linkedlist.size--
}

func removeFromEnd(linkedlist *linkedlist) {
	if linkedlist.size == 0 {
		fmt.Println("Linked list is currently empty, there is nothing to remove.")
		return
	}

	currentNode := linkedlist.head

	for currentNode.next.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = nil
	linkedlist.size--
}

func removeFromPosition(linkedlist *linkedlist, position int) {
	if linkedlist.size == 0 {
		fmt.Println("Linked list is currently empty. Removing from beginning.")
		removeFromStart(linkedlist)
		return
	}

	if position <= 0 {
		fmt.Println("Given position is equal or below 0. Remove from beginning.")
		removeFromStart(linkedlist)
		return
	}

	if position > linkedlist.size {
		fmt.Println("Given position is bigger than linked list size. Removing from end.")
		removeFromEnd(linkedlist)
		return
	}

	currentPosition := 0
	currentNode := linkedlist.head

	for currentPosition < position-1 {
		currentNode = currentNode.next
		currentPosition++
	}

	currentNode.next = currentNode.next.next
	linkedlist.size--
}

func findNode(linkedlist *linkedlist, value int) {
	if linkedlist.size == 0 {
		fmt.Println("Linked list is currently empty.")
	}

	currentPosition := 0
	currentNode := linkedlist.head

	for currentNode != nil {
		if currentNode.data == value {
			fmt.Println("Element found at position:", currentPosition+1)
			return
		}
		currentNode = currentNode.next
		currentPosition++
	}

	fmt.Println("Element not found.")
}

func display(linkedlist *linkedlist) {
	if linkedlist.head == nil {
		fmt.Println("Linked list is curently empty.")
		return
	}

	currentNode := linkedlist.head

	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {
	var linkedlist linkedlist = linkedlist{head: nil, size: 0}

	fmt.Println("Adding in the end")
	insertAtEnd(&linkedlist, 10)
	insertAtEnd(&linkedlist, 20)
	insertAtEnd(&linkedlist, 30)

	fmt.Println("Adding in the beggining")
	insertAtBeginning(&linkedlist, 40)
	insertAtBeginning(&linkedlist, 50)
	insertAtBeginning(&linkedlist, 60)

	fmt.Println("Adding at position")
	insertAtPosition(&linkedlist, 3, 100)
	insertAtPosition(&linkedlist, 4, 200)
	insertAtPosition(&linkedlist, 5, 300)

	fmt.Println("\nCurrent linked list:")
	display(&linkedlist)

	fmt.Println("\nRemoving from start")
	removeFromStart(&linkedlist)

	fmt.Println("\nCurrent linked list:")
	display(&linkedlist)

	fmt.Println("\nRemoving from end")
	removeFromEnd(&linkedlist)

	fmt.Println("\nCurrent linked list:")
	display(&linkedlist)

	fmt.Println("\nRemoving from position")
	removeFromPosition(&linkedlist, 3)

	fmt.Println("\nCurrent linked list:")
	display(&linkedlist)

	fmt.Println("\nFinding element: 100")
	findNode(&linkedlist, 100)

	fmt.Println("\nCurrent linked list:")
	display(&linkedlist)
}
