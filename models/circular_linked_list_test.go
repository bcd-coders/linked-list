package models

import "testing"

func TestIsCircular(t *testing.T) {
	n := CircularLinkedList{}

	if n.IsCircular() {
		t.Fatalf("Empty lists are not circular")
	}

	n.head = &Node{data: 1}
	if n.IsCircular() {
		t.Fatalf("A single node is not circular")
	}

	newNode := &Node{data: 2}
	n.head.next = newNode
	newNode.next = n.head

	if !n.IsCircular() {
		t.Fatalf("Should have detected circular list")
	}
}
func TestAddNode(t *testing.T) {

	nonCircular := CircularLinkedList{}

	l := []int{1, 2, 3}
	nonCircular.MakeCircular(l)

}
