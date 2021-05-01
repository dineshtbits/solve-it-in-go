package main

import (
	"fmt"
	"sort"
)

func twoSum(arr []int, target int) [][]int {
	myMap := make(map[interface{}]interface{})
	result := [][]int{}
	for index := 0; index < len(arr); index++ {
		complement := target - arr[index]
		v, ok := myMap[arr[index]]
		if ok {
			result = append(result, []int{v.(int), arr[index]})
		} else {
			myMap[complement] = arr[index]
		}
	}
	fmt.Printf("target: %v, input: %v, result: %v\n", target, arr, result)
	return result
}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			a := twoSum(nums[i+1:], 0-nums[i])
			if len(a) > 0 {
				for _, v := range a {
					v = append(v, nums[i])
					result = append(result, v)
				}
			}
		}
	}
	return result
}

func main() {
	arr := []int{0, 0, 0, 0}
	fmt.Printf("values matching target %v\n", threeSum(arr))
}
