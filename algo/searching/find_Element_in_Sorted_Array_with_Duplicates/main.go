package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Given a sorted array of integers and a target integer, find the first occurrence of the target and return its index.
// Return -1 if the target is not in the array.

// Input:

// arr = [1, 3, 3, 3, 3, 6, 10, 10, 10, 100]
// target = 3
// Output: 1

// Explanation: The first occurrence of 3 is at index 1.

// Input:

// arr = [2, 3, 5, 7, 11, 13, 17, 19]
// target = 6
// Output: -1

// Explanation: 6 does not exist in the array.

func findFirstOccurrence(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	result := -1

	for left <= right {
		mid := (left + right) / 2
		// Feasible function is arr[mid] >= target
		if arr[mid] == target {
			result = mid
			right = mid - 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func arrayAtoi(arr []string) []int {
	res := []int{}
	for _, x := range arr {
		v, _ := strconv.Atoi(x)
		res = append(res, v)
	}
	return res
}

func splitWords(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := arrayAtoi(splitWords(scanner.Text()))
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	res := findFirstOccurrence(arr, target)
	fmt.Println(res)
}
