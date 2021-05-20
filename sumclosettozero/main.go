package main

import (
	"fmt"
	"sort"
)

func sumClosestToZeroUsingSort(arr []int) ([]int, int) {

	start := 0
	end := len(arr) - 1
	sort.Ints(arr)
	sum := arr[start] + arr[end]
	fmt.Println(arr)
	var result []int

	for start < end {
		sum = arr[start] + arr[end]
		if sum == 0 {
			result = []int{arr[start], arr[end]}
			break
		} else if sum > 0 {
			result = []int{arr[start], arr[end]}
			end--
		} else {
			result = []int{arr[start], arr[end]}
			start++
		}
	}
	return result, sum
}

func main() {
	arr := []int{4, -3, 5, 82, -3, 1, 8}
	result, sum := sumClosestToZeroUsingSort(arr)
	fmt.Printf("Sum closest to zero in %v is %v, sum: %v\n", arr, result, sum)
}
