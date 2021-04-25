package main

import "fmt"

func binarySearch(a []int, target int) (int, bool) {
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == target {
			return mid, true
		} else if a[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0, false
}

func printResult(index int, result bool) {
	if result {
		fmt.Printf("element found at index %v\n", index)
	} else {
		fmt.Printf("element not found\n")
	}
}

func main() {
	arr := []int{2, 4, 5, 25, 62, 74}
	printResult(binarySearch(arr, 62))
}
