package stack

//Stack is a LIFO collection of integers.
type Stack struct {
	elements []int
}

//Pop removes the elt at the top of stack and returns it.
func (s *Stack) Pop() int {
	elt := s.Peek()
	s.elements = s.elements[s.Size():]
	return elt
}

//Push adds an integer to the stack.
func (s *Stack) Push(val int) {
	s.elements = append(s.elements, val)
}

//Peek returns the elt at the top of the stack.
func (s *Stack) Peek() int {
	return s.elements[s.Size()-1]
}

//Size returns length of Stack collection.
func (s *Stack) Size() int {
	return len(s.elements)
}
