package stack

import "testing"

func TestSize(t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	if s.Size() != 2 {
		t.Errorf("Expected %d, got %d", 2, s.Size())
	}
}

func TestPush(t *testing.T) {
	s := Stack{}
	s.Push(1)

	if s.Size() != 1 {
		t.Errorf("Expected %d, got %d", 1, s.Size())
	}
}

func TestPop(t *testing.T) {
	s := Stack{}
	s.Push(1)
	poppedItem := s.Pop()
	if poppedItem != 1 {
		t.Errorf("Expected %d, got %d", 1, poppedItem)
	}
	if s.Size() != 0 {
		t.Errorf("Expected %d, got %d", 0, s.Size())
	}
}

func TestPeek(t *testing.T) {
	s := Stack{}
	s.Push(1)
	elt := s.Peek()
	if elt != 1 {
		t.Errorf("Expected %d, got %d", 1, elt)
	}
	if s.Size() != 1 {
		t.Errorf("Expected %d, got %d", 1, s.Size())
	}
}
