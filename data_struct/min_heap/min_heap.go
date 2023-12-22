package main

import "fmt"

// Min heap data structure a.k.a. priority queue
// With array index starting with 1
// Get parent index of any node 			- index/2
// Get left child index of any node 	- n*index
// Get right child index of any node 	- n*index + 1
type MinHeap struct {
	Size  int
	Array []int
}

// Create new min heap
func NewMinHeap(len int) *MinHeap {
	minHeap := &MinHeap{
		Size:  0,
		Array: make([]int, len),
	}
	minHeap.Array[0] = -1
	return minHeap
}

// Insert into min heap
func (m *MinHeap) Insert(v int) {
	fmt.Printf("Inserting %d into min heap\n", v)

	m.Size += 1
	m.Array[m.Size] = v

	smallestIndex := m.Size
	for smallestIndex > 0 {
		parentIndex := smallestIndex / 2                   // Get parent index of smallest
		if m.Array[smallestIndex] < m.Array[parentIndex] { // Check if smallest is smaller than parent
			m.SwapValue(parentIndex, smallestIndex) // Swap smallest with parent
			smallestIndex = parentIndex             // Parent is new smallest
			continue
		}
		break
	}
}

// Swap values in min heap array
func (m *MinHeap) SwapValue(i, j int) {
	m.Array[i], m.Array[j] = m.Array[j], m.Array[i]
}

// Delete from min heap - It always deletes the root node
func (m *MinHeap) Delete() {
	fmt.Println("Deleting from the min heap")
	m.Array[1] = m.Array[m.Size]
	m.Array[m.Size] = -1
	m.Size -= 1

	for index := 1; index < m.Size; index++ {
		smallestIndex := index
		leftChildIndex := index * 2
		rightChildIndex := index*2 + 1
		fmt.Printf("leftChildIndex %d rightChildIndex %d\n", leftChildIndex, rightChildIndex)

		if leftChildIndex <= m.Size && m.Array[leftChildIndex] < m.Array[smallestIndex] {
			smallestIndex = leftChildIndex
		}

		if rightChildIndex <= m.Size && m.Array[rightChildIndex] < m.Array[smallestIndex] {
			smallestIndex = rightChildIndex
		}

		fmt.Printf("index %d smallestIndex %d\n", index, smallestIndex)
		if index != smallestIndex {
			m.SwapValue(index, smallestIndex)
		} else {
			return
		}
	}
}

// Heapify min heap
func (m *MinHeap) Heapify(index int) {
	fmt.Printf("Heapifying min heap for index %d\n", index)
	smallest := index
	left := index * 2
	right := index*2 + 1

	if left <= m.Size && m.Array[left] < m.Array[smallest] {
		smallest = left
	}

	if right <= m.Size && m.Array[right] < m.Array[smallest] {
		smallest = right
	}

	if smallest != index {
		m.SwapValue(smallest, index)
		m.Heapify(smallest)
	}
}

// Run heapify on min heap
func (m *MinHeap) RunHeapify() {
	for i := m.Size / 2; i > 0; i-- {
		m.Heapify(i)
	}
}

// Print min heap
func (m *MinHeap) Print() {

	for i := 1; i <= m.Size; i++ {
		fmt.Printf("%d  ", m.Array[i])
	}
	fmt.Println()
}

func main() {
	fmt.Println("Implementing min heap")
	minHeap := NewMinHeap(10)
	minHeap.Insert(22)
	minHeap.Print()
	minHeap.Insert(44)
	minHeap.Print()
	minHeap.Insert(12)
	minHeap.Print()
	minHeap.Insert(67)
	minHeap.Print()
	minHeap.Insert(24)
	minHeap.Print()
	minHeap.Delete()
	minHeap.Print()
	fmt.Printf("Min heap size - %d\n", minHeap.Size)

	fmt.Println()
	fmt.Println("Demo of Heapify Algorithm")
	minHeap2 := NewMinHeap(10)
	minHeap2.Array = []int{-1, 22, 44, 12, 67, 24, 1}
	minHeap2.Size = 6
	minHeap2.Print()
	minHeap2.RunHeapify()
	minHeap2.Print()
}
