package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
	data int   //The data stored in the node, can be of any data type (e.g., integer, string, custom struct, etc.)
	next *Node //A pointer to the next node in the list. In Go, pointers are represented using the * symbol.
}

// LinkedList represents a singly linked list
type LinkedList struct {
	head *Node //A pointer to the first node (head) of the linked list.
}

type CircularLinkedList struct {
	head *Node
}

// AddNode adds a new node to the end of the linked list
func (list *LinkedList) AddNode(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// AddNode adds a new node to the end of the circular linked list
func (list *CircularLinkedList) AddNode(data int) {
	newNode := &Node{data: data}

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

func (list *CircularLinkedList) isCircular() {
	if list == nil || list.head == nil {
		// An empty list or a list with only one node is not circular
		fmt.Println("The linkedlist is not circular")
	}

	tortoise := list.head
	hare := list.head

	for hare != nil && hare.next != nil {
		tortoise = tortoise.next // Move tortoise one step
		hare = hare.next.next    // Move hare two steps

		if tortoise == hare {
			fmt.Println("The linkedlist is circular")
			return
		}
	}

	// If the hare reaches the end of the list (nil), the list is not circular
	fmt.Println("The linkedlist is not circular")
}

// Display prints the elements of the linked list
func (list *LinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	// Create a new linked list
	myList := LinkedList{}

	// Create a circular linked list
	myList2 := CircularLinkedList{}

	// Add elements to the linked list
	myList.AddNode(1)
	myList.AddNode(2)
	myList.AddNode(3)
	myList.AddNode(4)
	myList.AddNode(5)

	// Add elements to the circular linked list
	myList2.AddNode(1)
	myList2.AddNode(2)
	myList2.AddNode(3)
	myList2.AddNode(4)
	myList2.AddNode(5)

	// Display the linked list
	fmt.Println("Linked List:")
	myList.Display()

	myList2.isCircular()
}
