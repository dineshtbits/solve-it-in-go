package main

import (
	"fmt"
)

func twoSum(arr []int, target int) []int {
	myMap := make(map[interface{}]interface{})
	result := []int{}
	for index := 0; index < len(arr); index++ {
		complement := target - arr[index]
		v, ok := myMap[arr[index]]
		if ok {
			result = []int{v.(int), index}
			break
		} else {
			myMap[complement] = index
		}
	}
	return result
}

func main() {
	arr := []int{1,2,3,4}
	target := 4
	fmt.Printf("values matching target %v are %v\n", target, twoSum(arr, target))
}