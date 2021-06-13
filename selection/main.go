package main

import "fmt"

func FindLargest(a []int) int {
	largest := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > largest {
			largest = a[i]
		}
	}
	return largest
}

func FindLargestAndSmallest(a []int) (int, int) {
	largest, smallest := a[0], a[0]
	for i := 0; i < len(a); {
		if a[i] < a[i+1] {
			if a[i] < smallest {
				smallest = a[i]
			}
			if a[i+1] > largest {
				largest = a[i+1]
			}
		} else {
			if a[i] > largest {
				largest = a[i]
			}
			if a[i+1] < smallest {
				smallest = a[i+1]
			}
		}
		i = i + 2
	}
	return smallest, largest
}

func patrition(low, high int, a []int) int {
	pivotElement := a[low]
	j := low
	for i := low + 1; i <= high; i++ {
		if a[i] < pivotElement {
			j++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[low], a[j] = a[j], a[low]
	return j
}

func median(a []int, start, end int) int {
	n := end - start + 1
	if n%2 == 1 {
		return a[start+(n/2)]
	} else {
		return (a[start+(n/2)] + a[(start+(n/2-1))]) / 2
	}
}

func main() {
	a := []int{3, 5, 2, 5, 7, 8}
	fmt.Printf("Array: %v, Largest: %v\n", a, FindLargest(a))
	smallest, largest := FindLargestAndSmallest(a)
	fmt.Printf("Array: %v, Smallest: %v, Largest: %v\n", a, smallest, largest)
	fmt.Println(patrition(0, len(a)-1, a))
	fmt.Println(a)
	fmt.Println(median(a, 0, len(a)-1))
}
