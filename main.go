package main

import (
	"fmt"

	"github.com/tanggol69/linked-list/tree/doubly-linked-list/models"
)

func main() {
	// Create a new linked list
	myList := models.DoublyLinkedList{}

	// elements to the linked list
	l := []int{1, 2, 3, 4, 5}
	myList.AddNodes(l)

	// Display the linked list forwards and backwards
	fmt.Println("Linked List (Forwards):")
	myList.Forwards()

	fmt.Println("Linked List (Backwards):")
	myList.Backwards()
}
