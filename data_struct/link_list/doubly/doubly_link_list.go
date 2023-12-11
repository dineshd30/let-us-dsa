package main

import "fmt"

type LinkList struct {
	length int
	head   *LinkListNode
	tail   *LinkListNode
}

func NewLinkList() *LinkList {
	return &LinkList{
		length: 0,
	}
}

func (l *LinkList) Append(value any) {
	newNode := NewLinkListNode(value)
	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
		l.length++
		return
	}

	tailNode := l.tail
	tailNode.next = newNode
	newNode.prev = tailNode
	l.tail = newNode
	l.length++
}

func (l *LinkList) Prepend(value any) {
	newNode := NewLinkListNode(value)
	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
		l.length++
		return
	}

	headNode := l.head
	newNode.next = headNode
	headNode.prev = newNode
	l.head = newNode
	l.length++
}

func (l *LinkList) Insert(index int, value any) {
	if index <= 0 {
		l.Prepend(value)
		return
	}

	if index >= l.length {
		l.Append(value)
		return
	}

	newNode := NewLinkListNode(value)
	leftNode := l.traverseToIndex(index)
	rightNode := leftNode.next
	leftNode.next = newNode
	newNode.prev = leftNode
	newNode.next = rightNode
	rightNode.prev = newNode
	l.length++
}

func (l *LinkList) Delete(index int) {
	if index <= 0 {
		currentHead := l.head
		l.head = currentHead.next
		l.length--
		return
	}

	leftNode := l.traverseToIndex(index)
	nodeToBeDeleted := leftNode.next
	if nodeToBeDeleted == l.tail {
		l.tail = leftNode
	}
	rightNode := nodeToBeDeleted.next
	leftNode.next = rightNode
	rightNode.prev = leftNode
	l.length--
}

func (l *LinkList) traverseToIndex(index int) *LinkListNode {
	if index <= 0 {
		return nil
	}

	if index >= l.length {
		index = l.length - 1
	}

	counter := 0
	currentNode := l.head
	for counter != index-1 && currentNode != nil {
		currentNode = currentNode.next
		counter++
	}
	return currentNode
}

func (l *LinkList) ToArray() []any {
	if l.length == 0 {
		return []any{}
	}

	var arr []any
	currentNode := l.head
	for currentNode != nil {
		arr = append(arr, currentNode.value)
		currentNode = currentNode.next
	}
	return arr
}

type LinkListNode struct {
	value any
	next  *LinkListNode
	prev  *LinkListNode
}

func NewLinkListNode(v any) *LinkListNode {
	return &LinkListNode{
		value: v,
	}
}

func main() {
	fmt.Println("Implementing Singly Link List")
	ll := NewLinkList()
	ll.Append(12)
	ll.Append(44)
	ll.Append(33)
	fmt.Printf("Singly Link List is %v\n", ll.ToArray())
	fmt.Printf("Head value is %d and Tail value is %d and length is %d\n", ll.head.value, ll.tail.value, ll.length)
	ll.Prepend(99)
	ll.Prepend(211)
	fmt.Printf("Singly Link List is %v\n", ll.ToArray())
	fmt.Printf("Head value is %d and Tail value is %d and length is %d\n", ll.head.value, ll.tail.value, ll.length)
	ll.Insert(4, 888)
	fmt.Printf("Singly Link List is %v\n", ll.ToArray())
	fmt.Printf("Head value is %d and Tail value is %d and length is %d\n", ll.head.value, ll.tail.value, ll.length)
	ll.Delete(0)
	ll.Delete(3)
	fmt.Printf("Singly Link List is %v\n", ll.ToArray())
	fmt.Printf("Head value is %d and Tail value is %d and length is %d\n", ll.head.value, ll.tail.value, ll.length)
}
