package llist

import "fmt"

// Node is a struct that holds an integer value and a pointer to the next node.
type Node struct {
	Value int
	Next  *Node
}

func (n Node) String() string {
	return fmt.Sprintf("%v --> ", n.Value)
}
