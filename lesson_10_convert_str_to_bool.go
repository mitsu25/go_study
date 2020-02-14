package main

import (
	"fmt"
)

func main() {

	var result bool
	noErrors := true
	input := "yes"

	switch input {
	case "true", "yes", "1":
		result = true
	case "false", "no", "0":
		result = false
	default:
		noErrors = false
	}
	if noErrors {
		fmt.Println("result:", result)
	} else {
		fmt.Println("エラーです")
	}
}
