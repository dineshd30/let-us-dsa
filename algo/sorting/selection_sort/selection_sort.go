package main

import "fmt"

// This function sorts given array in ascending order
// We call it selection sort because in every pass the smallest value index is selected and inserted into its correct position
func selectionSort(arr []int) []int {
	lastIndex := len(arr) - 1

	for i := 0; i <= lastIndex-1; i++ {
		lowIndex := i

		for j := i + 1; j <= lastIndex; j++ {
			if arr[j] < arr[lowIndex] {
				lowIndex = j
			}
		}

		if lowIndex != i {
			arr[lowIndex], arr[i] = arr[i], arr[lowIndex]
		}
	}

	return arr
}

func main() {
	fmt.Println("Implementing selection sort")
	arr := []int{33, 1, 44, 32, 7, 23}
	fmt.Printf("Sorted array - %v\n", selectionSort(arr))
}
