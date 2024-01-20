package main

import (
	"fmt"
)

func partition(arr []int) ([]int, int, []int) {
	pivot := len(arr) - 1
	i, j := 0, len(arr)-1
	for i != j {
		if arr[i] < arr[pivot] {
			i += 1
			continue
		}

		if arr[j] >= arr[pivot] {
			j -= 1
			continue
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	if i == j && pivot != j {
		arr[pivot], arr[i] = arr[i], arr[pivot]
	}

	return arr[:i], arr[i], arr[i+1:]
}

func quickSort(arr []int) []int {
	// base case
	if len(arr) <= 1 {
		return arr
	}

	left, middle, right := partition(arr)
	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)
	return append(append(sortedLeft, middle), sortedRight...)
}

func main() {
	fmt.Println("Implementing quick sort")
	arr := []int{10, 5, 11, 33, 7, 2, 1, 22, 1}
	fmt.Printf("%d\n", arr)
	// Built in function to sort slice in go
	// sort.SliceStable(arr, func(i, j int) bool {
	// 	return arr[i] > arr[j]
	// })
	// fmt.Println(arr);
	fmt.Printf("%d\n", quickSort(arr))
}
