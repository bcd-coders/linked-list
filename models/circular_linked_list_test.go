package models

import "testing"

func TestMakeCircular(t *testing.T) {

	testCases := []struct {
		Data       []int
		IsCircular bool
	}{
		{Data: []int{}, IsCircular: false},
		{Data: []int{1}, IsCircular: true},
		{Data: []int{1, 2, 3}, IsCircular: true},
	}

	for _, c := range testCases {
		n := CircularLinkedList{}
		n.MakeCircular(c.Data)

		if c.IsCircular != n.IsCircular() {
			t.Fatalf("expected circular: %v, but got: %v", c.IsCircular, n.IsCircular())
		}
	}
}

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
