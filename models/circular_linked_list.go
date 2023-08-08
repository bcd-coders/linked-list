package models

import "fmt"

type CircularLinkedList struct {
	head *Node
}

func (list *CircularLinkedList) MakeCircular(data []int) {
	for _, value := range data {
		newNode := &Node{data: value}

		if list.head == nil {
			list.head = newNode
			newNode.next = newNode
			return
		}

		//  Loops through the list until current.next returns to the head
		current := list.head
		for current.next != list.head {
			current = current.next
		}

		current.next = newNode

		// For a circular linked list, point the new node back to the head
		newNode.next = list.head
	}
}

func (list *CircularLinkedList) IsCircular() bool{
	if list == nil || list.head == nil {
		// An empty list or a list with only one node is not circular
		fmt.Println("The linkedlist is not circular")
		return false
	}

	tortoise := list.head
	hare := list.head

	for hare != nil && hare.next != nil {
		tortoise = tortoise.next // Move tortoise one step
		hare = hare.next.next    // Move hare two steps

		// If they are equal, it means they met each other within the loop
		if tortoise == hare {
			fmt.Println("The linkedlist is circular")
			return true
		}
	}

	// If the hare reaches the end of the list (nil), the list is not circular
	fmt.Println("The linkedlist is not circular")
	return false
}
