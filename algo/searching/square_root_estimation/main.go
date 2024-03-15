package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// mid^2 > n

func squareRoot(n int) int {
	// WRITE YOUR BRILLIANT CODE HERE
	left := 0
	right := n
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		if (mid * mid) == n {
			ans = mid
			right = mid - 1
		} else if (mid * mid) > n {
			ans = mid - 1
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	res := squareRoot(n)
	fmt.Println(res)
}
