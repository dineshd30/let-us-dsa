package main

import (
	"errors"
	"fmt"
)

func reverse(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("failed to get valid string")
	}

	var reversed []byte
	for i := len(str) - 1; i >= 0; i-- {
		reversed = append(reversed, str[i])
	}
	return string(reversed), nil
}

func main() {
	fmt.Println("Reversing the string")
	reversed, _ := reverse("Hi my name is Dinesh")
	fmt.Println(reversed)
}
