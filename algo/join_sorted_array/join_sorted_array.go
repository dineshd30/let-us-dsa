package main

import (
	"errors"
	"fmt"
)

func join(arr1 []int, arr2 []int) ([]int, error) {
	fmt.Printf("arr1 - %v, arr2 - %v\n", arr1, arr2)
	if len(arr1) == 0 && len(arr2) == 0 {
		return nil, errors.New("failed to get valid array")
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
	for arr1Index < len(arr1) && arr2Index < len(arr2) {
		if arr1[arr1Index] < arr2[arr2Index] {
			mergedArr = append(mergedArr, arr1[arr1Index])
			arr1Index += 1
		} else {
			mergedArr = append(mergedArr, arr2[arr2Index])
			arr2Index += 1
		}
	}

	if arr1Index < len(arr1) {
		mergedArr = append(mergedArr, arr1[arr1Index:]...)
	}

	if arr2Index < len(arr2) {
		mergedArr = append(mergedArr, arr2[arr2Index:]...)
	}

	return mergedArr, nil
}

func main() {
	fmt.Println("Join sorted integer array")
	merged, _ := join([]int{0, 3, 4, 31, 32}, []int{-1, 1, 4, 6, 30})
	fmt.Printf("Merged sorted integer array - %v\n", merged)
}
