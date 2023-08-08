package models

import "testing"

func TestLength(t *testing.T){
	n := DoublyLinkedList{}

	x := n.Length()

	if x != 0{
		t.Fatalf("Expecting 0 but got %d", x)
	}

	L := []int{1}

	n.AddNodes(L)

	y := n.Length()
	if y != 1{
		t.Fatalf("Expecting 1 but got %d", y)
	}

	A := []int{1,2,3}
	n.AddNodes(A)

	z:= n.Length()

	if z != 4{
		t.Fatalf("Expecting 3 but got %d",z)
	}
}

// func TestAddNodes(t *testing.T){
// 	n := DoublyLinkedList{}

// 	l := []int{1, 2, 3}
// 	n.AddNodes(l)
// }