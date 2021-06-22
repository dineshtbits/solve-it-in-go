package main

import "fmt"

func span(input []int) []int {
	result := make([]int, len(input))

	if len(input) <= 0 {
		return input
	}

	fmt.Println(result)

	result[0] = 1
	for i := 1; i < len(input); i++ {
		if input[i-1] > input[i] {
			result[i] = 1
		} else {
			result[i] = result[i-1] + 1
		}
	}

	return result
}

func main() {
	a := []int{6, 3, 4, 5, 2}
	fmt.Printf("span for array %v is %v\n", a, span(a))
}
