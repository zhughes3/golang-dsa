package llist

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	ll := LinkedList{}
	if ll.Size() != 0 {
		t.Errorf("Expected %d, got %d", 0, ll.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	list1 := LinkedList{}
	list2 := LinkedList{}
	list2.Insert(1)
	if !list1.IsEmpty() {
		t.Errorf("Expected %t, got %t", true, list1.IsEmpty())
	}
	if list2.IsEmpty() {
		t.Errorf("Expected %t, got %t", false, list2.IsEmpty())
	}
}

func TestInsertHead(t *testing.T) {
	ll := LinkedList{}
	ll.Insert(3)
	if ll.Size() != 1 {
		t.Errorf("Expected %d, got %d", 1, ll.Size())
	}
	head := ll.Head
	if head.Value != 3 {
		t.Errorf("Expected %d, got %d", 3, head.Value)
	}
}

func TestInsert(t *testing.T) {
	ll := LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	if ll.Size() != 2 {
		t.Errorf("Expected %d, got %d", 2, ll.Size())
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	if current.Value != 2 {
		t.Errorf("Expected %d, got %d", 2, current.Value)
	}
}

func TestDeleteHead(t *testing.T) {
	ll := LinkedList{}
	ll.Insert(1)
	ll.Delete(1)
	if ll.Size() != 0 {
		t.Errorf("Expected %d, got %d", 0, ll.Size())
	}
	if ll.Head != nil {
		t.Errorf("Expected %v, got %v", nil, ll.Head)
	}
}

func TestDelete(t *testing.T) {
	ll := LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Delete(2)
	if ll.Size() != 2 {
		t.Errorf("Expected %d, got %d", 2, ll.Size())
	}
	if ll.Head.Next.Value != 3 {
		t.Errorf("Expected %v, got %v", 3, ll.Head.Next.Value)
	}
}

func TestSize(t *testing.T) {
	ll := LinkedList{}
	if ll.Size() != 0 {
		t.Errorf("Expected %d, got %d", 0, ll.Size())
	}
	ll.Insert(1)
	if ll.Size() != 1 {
		t.Errorf("Expected %d, got %d", 1, ll.Size())
	}
}

func TestRetrieveNode(t *testing.T) {
	ll := LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	retrievedNode := ll.RetrieveNode(2)
	if retrievedNode.Value != 2 {
		t.Errorf("Expected %d, got %d", 2, retrievedNode.Value)
	}
	nodeNotFound := ll.RetrieveNode(4)
	if nodeNotFound != nil {
		t.Errorf("Expected %v, got %v", nil, nodeNotFound)
	}
}

func TestString(t *testing.T) {
	expected := "1 --> 2 --> "
	ll := LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	var b bytes.Buffer
	b.WriteString(ll.String())
	if b.String() != expected {
		t.Errorf("Expected %s, got %s", expected, b.String())
	}
}
