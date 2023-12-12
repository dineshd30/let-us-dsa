package main

import "fmt"

const SIZE int = 7

// Hash table structure
type HashTable struct {
	array [SIZE]*Bucket
}

// Initialize the hash table
func NewHashTable() *HashTable {
	table := &HashTable{}

	for i := 0; i < SIZE; i++ {
		table.array[i] = &Bucket{}
	}

	return table
}

// Insert into the hash table
func (h *HashTable) Insert(key string, value any) {
	hash := h.Hash(key)
	h.array[hash].Prepend(key, value)
}

// Delete from the hash table
func (h *HashTable) Delete(key string) {
	hash := h.Hash(key)
	h.array[hash].Remove(key)
}

// Search from the hash table
func (h *HashTable) Search(key string) (bool, string, any) {
	hash := h.Hash(key)
	return h.array[hash].Search(key)
}

// Bucket structure which implements singly linked list
type Bucket struct {
	head *Node
}

// Node structure
type Node struct {
	key   string
	value any
	next  *Node
}

// Insert into the Bucket
func (b *Bucket) Prepend(key string, value any) {
	node := &Node{
		key:   key,
		value: value,
	}
	currentHead := b.head
	b.head = node
	b.head.next = currentHead
}

// Delete from the Bucket
func (b *Bucket) Remove(key string) {
	// Validate if bucket is empty
	if b.head == nil {
		return
	}

	// First handle when deleting head node
	if b.head.key == key {
		b.head = b.head.next
		return
	}

	// Then handle when deleting nodes after head node
	previousNode := b.head
	currentNode := b.head.next
	for currentNode != nil {
		if currentNode.key == key {
			previousNode.next = currentNode.next
			return
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
}

// Search from the Bucket
func (b *Bucket) Search(key string) (bool, string, any) {
	if b.head == nil {
		return false, "", ""
	}

	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true, currentNode.key, currentNode.value
		}
		currentNode = currentNode.next
	}

	return false, "", ""
}

// Hash function
func (h *HashTable) Hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % SIZE
}

func main() {
	fmt.Println("Implementing hash table")
	table := NewHashTable()
	table.Insert("Randy", 22)
	table.Insert("Kyle", 11)
	table.Insert("Same", 4)
	table.Insert("Joe", 4)
	table.Insert("Tom", 2)
	table.Insert("Sam", 222)
	fmt.Println(table.Search("Joe"))
	table.Delete("Joe")
	fmt.Println(table.Search("Joe"))
	fmt.Println(table.Search("Randy"))
	table.Delete("Randy")
	fmt.Println(table.Search("Randy"))
}
