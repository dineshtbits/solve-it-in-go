package main

import (
	"fmt"
)

func linearSearch(a []int, target int) (int, bool) {
	var result int
	for i := 0; i < len(a); i++ {
		if a[i] == target {
			return i, true
		}
	}
	return result, false
}

func printResult(index int, result bool) {
	if result {
		fmt.Printf("element found at index %v\n", index)
	} else {
		fmt.Printf("element not found\n")
	}
}

func main() {
	arr := []int{4, 5, 62, 25, 2, 4}
	printResult(linearSearch(arr, 2))
}
