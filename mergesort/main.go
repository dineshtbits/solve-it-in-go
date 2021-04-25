package main

import "fmt"

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	mid := len(a) / 2

	leftSubList := mergeSort(a[:mid])
	rightSubList := mergeSort(a[mid:])

	var sortedList []int
	var leftPointer, rightPointer int = 0, 0

	fmt.Printf("Now working on %v    %v\n", leftSubList, rightSubList)
	for leftPointer < len(leftSubList) && rightPointer < len(rightSubList) {
		if leftSubList[leftPointer] > rightSubList[rightPointer] {
			sortedList = append(sortedList, rightSubList[rightPointer])
			rightPointer += 1
		} else {
			sortedList = append(sortedList, leftSubList[leftPointer])
			leftPointer += 1
		}
	}
	sortedList = append(sortedList, leftSubList[leftPointer:]...)
	sortedList = append(sortedList, rightSubList[rightPointer:]...)
	return sortedList
}

func main() {
	arr := []int{-317, -381, -14, -215, -590, -243, -412, 380, -312, 925, 158, -46, 177, 22, -482, 273, 217, 514, -392, 424}
	fmt.Printf("Original array %v\n", arr)
	fmt.Printf("Sorted array %v\n", mergeSort(arr))
}
