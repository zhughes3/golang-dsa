package llist

import (
	"fmt"
	"strings"
)

// LinkedList is a linkedlist
type LinkedList struct {
	Head *Node
	size int
}

// Insert does
func (ll *LinkedList) Insert(val int) {
	n := Node{Value: val}
	if ll.Head == nil {
		fmt.Println("nil, inserting head")
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

// Size returns number of Nodes in LinkedList
func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) String() string {
	var sb strings.Builder
	current := ll.Head
	for current != nil {
		sb.WriteString(current.String())
		current = current.Next
	}
	return sb.String()
}
