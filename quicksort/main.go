package main

import "fmt"

func quickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	pivot := a[0]
	var leftSubList, rightSubList, sortedList []int

	for i := 1; i < len(a); i++ {
		if a[i] > pivot {
			rightSubList = append(rightSubList, a[i])
		} else {
			leftSubList = append(leftSubList, a[i])
		}
	}
	sortedList = append(sortedList, quickSort(leftSubList)...)
	sortedList = append(sortedList, pivot)
	sortedList = append(sortedList, quickSort(rightSubList)...)
	return sortedList
}

func parition(low, high int, a []int) int {
	pivotElement := a[low]
	j := low
	for i := low + 1; i < len(a); i++ {
		if a[i] < pivotElement {
			j++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[j], a[low] = a[low], a[j]
	return j
}

func quickSortImproved(low, high int, a []int) []int {
	if high > low {
		p := parition(low, high, a)
		quickSortImproved(low, p-1, a)
		quickSortImproved(p+1, high, a)

	}
	return a
}

func main() {
	arr := []int{-317, -381, -14, -215, -590, -243, -412, 380, -312, 925, 158, -46, 177, 22, -482, 273, 217, 514, -392, 424}
	fmt.Printf("Original array %v\n", arr)
	fmt.Printf("Sorted array %v\n", quickSort(arr))
	fmt.Printf("Sorted array %v\n", quickSortImproved(0, len(arr)-1, arr))
}
