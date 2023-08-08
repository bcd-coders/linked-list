package models

import "fmt"

// Node represents a single node in the doubly linked list
type Node struct {
	data int   // The data stored in the node, can be of any data type
	next *Node // A pointer to the next node in the list
	prev *Node // A pointer to the previous node in the list
}

// DoublyLinkedList represents a doubly linked list
type DoublyLinkedList struct {
	head *Node // A pointer to the first node (head) of the linked list.
}

func (list *DoublyLinkedList) Length() int {
	if list == nil {
		return 0
	}
	current := list.head
	ctr := 0
	for current != nil {
		current = current.next
		ctr++
	}
	return ctr
}

// AddNode adds a new node to the end of the linked list
func (list *DoublyLinkedList) AddNodes(data []int) {
	for _, value := range data {
		newNode := &Node{data: value}

		if list.head == nil {
			list.head = newNode
			continue
		}

		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
		newNode.prev = current
	}
}

// Forwards prints the elements of the linked list forwards
func (list *DoublyLinkedList) Forwards() {
	current := list.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

// Backwards prints the elements of the linked list backwards
func (list *DoublyLinkedList) Backwards() {
	current := list.head
	for current.next != nil {
		current = current.next
	}

	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.prev
	}
	fmt.Println()
}
