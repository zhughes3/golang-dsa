package llist

import (
	"strings"
)

// LinkedList is a singly linked list of integers.
type LinkedList struct {
	Head *Node
	size int
}

//IsEmpty checks if list is empty.
func (ll *LinkedList) IsEmpty() bool {
	return ll.Size() == 0
}

// Insert a new node into the list with supplied value.
func (ll *LinkedList) Insert(val int) {
	n := Node{Value: val}
	if ll.Head == nil {
		ll.Head = &n
	} else {
		current := ll.Head
		last := ll.Head
		for current != nil {
			last = current
			current = current.Next
		}
		current = &n
		last.Next = current
	}
	ll.size++
}

// Delete first node found with supplied value.
func (ll *LinkedList) Delete(val int) {
	current := ll.Head
	if current.Value == val {
		ll.Head = nil
		ll.size--
		return
	}
	last := ll.Head
	for current != nil {
		if current.Value == val {
			last.Next = current.Next
			ll.size--
			return
		}
		last = current
		current = current.Next
	}
}

// Size returns number of Nodes in LinkedList.
func (ll *LinkedList) Size() int {
	return ll.size
}

//RetrieveNode returns the first Node found with the given value.
func (ll *LinkedList) RetrieveNode(val int) *Node {
	current := ll.Head
	for current != nil {
		if current.Value == val {
			return current
		}
		current = current.Next
	}
	return nil
}

//String outputs the list in a readable format.
func (ll *LinkedList) String() string {
	var sb strings.Builder
	current := ll.Head
	for current != nil {
		sb.WriteString(current.String())
		current = current.Next
	}
	return sb.String()
}
