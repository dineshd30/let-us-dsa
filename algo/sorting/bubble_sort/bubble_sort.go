package main

import "fmt"

// This function sorts given array in ascending order
// We call it bubble sort because in every pass the largest value bubbled up to its correct position
// If we do not perform any swaps in pass then we will stop the sorting and consider the array is sorted
func bubbleSort(arr []int) []int {
	isSorted := false
	lastIndex := len(arr) - 1

	for !isSorted {
		isSorted = true

		for i := 0; i < lastIndex; i++ {
			// if first value is greater than next value then we swap and consider array needs another pass
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}

		lastIndex -= 1
	}

	return arr
}

func main() {
	fmt.Println("Implementing bubble sort")
	arr := []int{33, 1, 44, 32, 7, 23}
	fmt.Printf("Sorted array - %v\n", bubbleSort(arr))
}
