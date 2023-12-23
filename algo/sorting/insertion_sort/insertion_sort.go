package main

import "fmt"

// This function sorts given array in ascending order
// We call it insertion sort because in every pass the selected temp value is inserted into correct gap position index
func insertionSort(arr []int) []int {
	for i := 1; i <= len(arr)-1; i++ {
		temp := arr[i]    // Store value from gap position index into temp
		position := i - 1 // Set the initial position for this pass Note: gap position index is always position+1

		for position >= 0 {
			if temp < arr[position] { // If temp is less than current position value
				arr[position+1] = arr[position] // Then shift its value into the gap position index which is position+1
				position -= 1                   // Move the gap to the left
				continue                        // Continue with next iteration
			}
			break // Else break from the loop
		}

		arr[position+1] = temp // Put back the temp value into the arrived gap position index which is position+1
	}

	return arr
}

func main() {
	fmt.Println("Implementing insertion sort")
	arr := []int{33, 2, 1, 5, 44, 11, 66}
	fmt.Printf("Sorted array is %v\n", insertionSort(arr))
}
