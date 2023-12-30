package main

import "fmt"

type Stack struct {
	arr []int
}

func New(input []int) *Stack {
	q := new(Stack)
	for _, v := range input {
		q.Push(v)
	}

	return q
}

func (q *Stack) Push(v int) {
	fmt.Printf("Inserting %d into the stack\n", v)
	q.arr = append(q.arr, v)
}

func (q *Stack) Pop() (int, error) {
	if len(q.arr) == 0 {
		return -1, fmt.Errorf("empty stack")
	}

	item := q.arr[len(q.arr)-1]
	q.arr = q.arr[:len(q.arr)-1]
	return item, nil
}

func (q *Stack) Peek() (int, error) {
	if len(q.arr) == 0 {
		return -1, fmt.Errorf("empty stack")
	}

	return q.arr[len(q.arr)-1], nil
}

func main() {
	fmt.Println("Implementing stack using array")
	i := []int{3, 2, 4, 22, 55, 3311, 1}
	q := New(i)
	r, _ := q.Pop()
	fmt.Printf("Removing from the stack %d\n", r)
	p, _ := q.Peek()
	fmt.Printf("Peeking from the stack %d\n", p)
	fmt.Printf("Final state of stack %v\n", q.arr)
}
