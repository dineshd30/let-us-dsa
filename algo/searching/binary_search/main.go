package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func binarySearch(arr []int, item int) bool {
	if len(arr) == 0 {
		return false
	}

	midIndex := len(arr) / 2
	if arr[midIndex] == item { // Item found
		return true
	} else if item > arr[midIndex] {
		return binarySearch(arr[midIndex+1:], item) // Search in right part of the array
	} else {
		return binarySearch(arr[:midIndex], item) // Search in left part of the array
	}
}

func main() {
	fmt.Println("Implementing binary search")
	fmt.Println("Enter number of elements to insert")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numberOfElements, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Enter number in increasing order")
	arr := []int{}
	for i := 0; i < numberOfElements; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, num)
	}
	fmt.Printf("Array: %v\n", arr)
	fmt.Println("Enter number to be searched in the array")
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	if binarySearch(arr, x) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
