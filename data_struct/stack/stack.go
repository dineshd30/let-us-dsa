package main

import "fmt"

// Stack structure
type Stack struct {
	top    *Node
	bottom *Node
	length int
}

// Stack Node
type Node struct {
	value any
	next  *Node
}

// Peek top node
func (s *Stack) Peek() any {
	if s.top == nil {
		return nil
	}

	return s.top.value
}

// Push into the stack
func (s *Stack) Push(v any) {
	node := &Node{value: v}
	if s.top == nil {
		s.top = node
		s.bottom = node
	} else {
		currentTop := s.top
		s.top = node
		s.top.next = currentTop
	}

	s.length++
}

// Pop from the stack
func (s *Stack) Pop() {
	if s.top == s.bottom {
		s.top = nil
	} else {
		s.top = s.top.next
	}
	s.length--
}

func main() {
	fmt.Println("Implementing stack")
	stack := Stack{}
	fmt.Printf("Stack Len - %d\n", stack.length)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Printf("Stack Len - %d\n", stack.length)
	fmt.Printf("Stack Peek - %d\n", stack.Peek())
	stack.Pop()
	stack.Pop()
	fmt.Printf("Stack Len - %d\n", stack.length)
	fmt.Printf("Stack Peek - %d\n", stack.Peek())
	stack.Pop()
	fmt.Printf("Stack Len - %d\n", stack.length)
	fmt.Printf("Stack Peek - %d\n", stack.Peek())
}
