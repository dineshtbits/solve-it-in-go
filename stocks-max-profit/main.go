package main

import (
	"fmt"
)

// Problem:
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// Solution 1:
// Time Complexity O(n)
// Space Complexity O(1)
func maxProfit(a []int) int {
	min := a[0]
	profit := 0
	for _, element := range a	{
		if element < min {
			min = element
		} 
		currentProfit := element - min 
		if currentProfit > profit {
			profit = currentProfit
		}
	}
	return profit
}

func main() {
	arr := []int{7,1,5,3,6,4}
	fmt.Printf("Maximum profit is %v\n", maxProfit(arr))
}