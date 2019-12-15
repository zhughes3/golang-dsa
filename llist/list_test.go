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
