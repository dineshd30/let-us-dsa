package main

import "fmt"

func firstRecurringChar(arr []int) (int, bool) {
	mapRecurChar := make(map[int]bool)
	for _, value := range arr {
		if _, ok := mapRecurChar[value]; ok {
			return value, true
		}
		mapRecurChar[value] = true
	}
	return -1, false
}

func main() {
	fmt.Println("Finding first recurring integer")
	char, found := firstRecurringChar([]int{1, 2, 3, 3, 4})
	fmt.Printf("First first recurring char is - %d and is found - %t\n", char, found)
}
