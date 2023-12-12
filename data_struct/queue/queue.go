package main

import "fmt"

// Queue structure
type Queue struct {
	first *Item
	last  *Item
	len   int
}

// Queue item
type Item struct {
	value any
	next  *Item
}

// Queue peek first item
func (q *Queue) Peek() any {
	// Handle if queue is empty
	if q.first == nil {
		return nil
	}

	return q.first.value
}

// Queue enqueue item
func (q *Queue) Enqueue(v any) {
	item := &Item{value: v}
	if q.first == nil {
		q.first = item
		q.last = item
	} else {
		currentLast := q.last
		currentLast.next = item
		q.last = item
	}
	q.len++
}

// Queue dequeue item
func (q *Queue) Dequeue() any {
	// Handle if queue is empty
	if q.first == nil {
		return nil
	}

	value := q.first.value
	if q.first == q.last {
		q.first = nil
	} else {
		q.first = q.first.next
	}
	q.len--
	return value
}

func main() {
	fmt.Println("Implementing queue")
	queue := &Queue{}
	queue.Enqueue(12)
	queue.Enqueue(2)
	queue.Enqueue(100)
	fmt.Printf("Queue Peek - %d\n", queue.Peek())
	fmt.Printf("Queue Len - %d\n", queue.len)
	fmt.Printf("Queue Dequeue - %d\n", queue.Dequeue())
	fmt.Printf("Queue Dequeue - %d\n", queue.Dequeue())
	fmt.Printf("Queue Peek - %d\n", queue.Peek())
	fmt.Printf("Queue Len - %d\n", queue.len)
	fmt.Printf("Queue Dequeue - %d\n", queue.Dequeue())
	fmt.Printf("Queue Len - %d\n", queue.len)
}
