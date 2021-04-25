package main

import "fmt"

func binarySearch(a []int, target int) bool {
	if len(a) == 0 {
		return false
	}
	mid := len(a) / 2
	if a[mid] == target {
		return true
	} else if a[mid] > target {
		return binarySearch(a[:mid], target)
	} else {
		return binarySearch(a[mid+1:], target)
	}
}

func printResult(result bool) {
	if result {
		fmt.Printf("element found\n")
	} else {
		fmt.Printf("element not found\n")
	}
}

func main() {
	arr := []int{2, 4, 5, 25, 62, 74}
	printResult(binarySearch(arr, 74))
}
