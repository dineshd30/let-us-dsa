package main

import (
	"errors"
	"fmt"
)

func join(arr1 []int, arr2 []int) ([]int, error) {
	fmt.Printf("arr1 - %v, arr2 - %v\n", arr1, arr2)
	if len(arr1) == 0 && len(arr2) == 0 {
		return nil, errors.New("failed to get valid string")
	}

	if len(arr1) > 0 && len(arr2) == 0 {
		return arr1, nil
	}

	if len(arr2) > 0 && len(arr1) == 0 {
		return arr2, nil
	}

	arr1Index := 0
	arr2Index := 0

	var mergedArr []int
	for arr1Index < len(arr1) || arr2Index < len(arr2) {
		value1, value1Found := getValueSafely(arr1Index, arr1)
		value2, value2Found := getValueSafely(arr2Index, arr2)
		if value1Found && value2Found && value1 < value2 {
			mergedArr = append(mergedArr, value1)
			arr1Index++
		} else if value1Found && value2Found && value1 > value2 {
			mergedArr = append(mergedArr, value2)
			arr2Index++
		} else if value1Found {
			mergedArr = append(mergedArr, value1)
			arr1Index++
		} else if value2Found {
			mergedArr = append(mergedArr, value2)
			arr2Index++
		} else {
			arr1Index++
			arr2Index++
		}
	}

	return mergedArr, nil
}

func getValueSafely(index int, arr []int) (int, bool) {
	if index >= 0 && index < len(arr) {
		return arr[index], true
	}

	return 0, false
}

func main() {
	fmt.Println("Join sorted integer array")
	merged, _ := join([]int{0, 3, 4, 31}, []int{-1, 1, 4, 6, 30})
	fmt.Printf("Merged sorted integer array - %v\n", merged)
}
