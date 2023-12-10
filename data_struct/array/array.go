package main

import "fmt"

type Array struct {
	length int
	data   map[int]any
}

func NewArray() *Array {
	return &Array{
		length: 0,
		data:   map[int]any{},
	}
}

func (a *Array) Get(index int) any {
	return a.data[index]
}

func (a *Array) Push(item any) int {
	a.data[a.length] = item
	a.length++
	return a.length
}

func (a *Array) Pop() any {
	item := a.data[a.length-1]
	delete(a.data, a.length-1)
	a.length--
	return item
}

func (a *Array) Delete(index int) any {
	item := a.data[index]
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	delete(a.data, a.length-1)
	a.length--
	return item
}

func main() {
	fmt.Println("Implementing array")
	arr := NewArray()
	arr.Push(22)
	arr.Push(11)
	arr.Push(14)
	arr.Push(8)
	arr.Push(25)
	fmt.Printf("Initial array is %v\n", arr)
	fmt.Printf("Length of an array is %d\n", arr.length)
	getVal := arr.Get(0)
	fmt.Printf("Value at index 0 is %d\n", getVal)
	popVal := arr.Pop()
	fmt.Printf("Remove last value which is %d\n", popVal)
	fmt.Printf("Array after pop is %v\n", arr)
	fmt.Printf("Length of an array is %d\n", arr.length)
	delVal := arr.Delete(2)
	fmt.Printf("Delete item at index 2 which is %d\n", delVal)
	fmt.Printf("Array after delete is %v\n", arr)
	fmt.Printf("Length of an array is %d\n", arr.length)
}
