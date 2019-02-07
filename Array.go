package main

import "fmt"

func main() {
	var arrValue = [5]int32{1, 2, 3, 4, 5}

	for _, value := range arrValue {
		fmt.Printf("%d\n", value)
	}
}
