package main

import "fmt"

type Queue struct {
	arr []int
}

func New(input []int) *Queue {
	q := new(Queue)
	for _, v := range input {
		q.Enqueue(v)
	}

	return q
}

func (q *Queue) Enqueue(v int) {
	fmt.Printf("Inserting %d into the queue\n", v)
	q.arr = append(q.arr, v)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.arr) == 0 {
		return -1, fmt.Errorf("empty queue")
	}

	item := q.arr[0]
	q.arr = q.arr[1:]
	return item, nil
}

func (q *Queue) Peek() (int, error) {
	if len(q.arr) == 0 {
		return -1, fmt.Errorf("empty queue")
	}

	return q.arr[0], nil
}

func main() {
	fmt.Println("Implementing queue using array")
	i := []int{3, 2, 4, 22, 55, 3311, 1}
	q := New(i)
	r, _ := q.Dequeue()
	fmt.Printf("Removing from the queue %d\n", r)
	p, _ := q.Peek()
	fmt.Printf("Peeking from the queue %d\n", p)
	fmt.Printf("Final state of queue %v\n", q.arr)
}
