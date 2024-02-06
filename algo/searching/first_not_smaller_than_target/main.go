package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// First Element Not Smaller Than Target
// Given an array of integers sorted in increasing order and a target, find the index of the first element
// in the array that is larger than or equal to the target. Assume that it is guaranteed to find a satisfying number.

// Input:

// arr = [1, 3, 3, 5, 8, 8, 10]
// target = 2
// Output: 1

// Explanation: the first element larger than 2 is 3 which has index 1.

// Input:

// arr = [2, 3, 5, 7, 11, 13, 17, 19]
// target = 6
// Output: 3

// Explanation: the first element larger than 6 is 7 which has index 3.

func firstNotSmallerThanTarget(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	result := -1
	for left <= right {
		mid := (left + right) / 2
		fmt.Println(mid)
		if arr[mid] >= target {
			fmt.Println(mid)
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func main() {
	fmt.Println("First Element Not Smaller Than Target")
	fmt.Println("Enter number of elements to insert")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numberOfElements, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Enter int values in increasing order")
	arr := []int{}
	for i := 0; i < numberOfElements; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, val)
	}
	fmt.Printf("Given array : %v\n", arr)
	fmt.Printf("Search result : %d\n", firstNotSmallerThanTarget(arr, 10))
}
