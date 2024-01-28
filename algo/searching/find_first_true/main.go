package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// An array of boolean values is divided into two sections; the left section consists of all false
// and the right section consists of all true. Find the First True in a Sorted Boolean Array of the right section,
// i.e. the index of the first true element. If there is no true element, return -1.

// Input: arr = [false, false, true, true, true]

// Output: 2

// Explanation: first true's index is 2.

func findFirstTrue(arr []bool, left int, right int) int {
	recordedTrue := []int{}
	for left > -1 && right < len(arr) && left <= right {
		midIndex := (left + right) / 2
		if arr[midIndex] {
			recordedTrue = append(recordedTrue, midIndex)
			right = midIndex - 1
		} else {
			left = midIndex + 1
		}
	}

	if len(recordedTrue) > 0 {
		return recordedTrue[len(recordedTrue)-1]
	}

	return -1
}

func main() {
	fmt.Println("Find the First True in a Sorted Boolean Array")
	fmt.Println("Enter number of elements to insert")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numberOfElements, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Enter boolean values in increasing order")
	arr := []bool{}
	for i := 0; i < numberOfElements; i++ {
		scanner.Scan()
		val, _ := strconv.ParseBool(scanner.Text())
		arr = append(arr, val)
	}
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Search result : %d\n", findFirstTrue(arr, 0, len(arr)-1))
}
