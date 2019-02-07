package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	c := max(a, b)
	fmt.Println(c)
}

func max(a int, b int) int {
	var result int

	if a > b {
		result = a
	} else {
		result = b
	}
	return result
}
