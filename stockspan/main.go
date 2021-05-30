package main

import "fmt"

func main() {
	arr := []int{100, 80, 60, 70, 60, 75, 85}
	var result [7]int
	result[0] = 1
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			result[i] = 1
		} else {
			result[i] = 1
			temp := i - 1
			for temp >= 0 {
				if arr[temp] < arr[i] {
					result[i] = result[i] + result[temp]
				}
				temp = temp - result[temp]
			}
		}
	}
	fmt.Println(result)
}
