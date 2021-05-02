package main

import (
	"fmt"
)

func indexOrders(orders []int) map[interface{}]interface{} {
	indexedOrders := make(map[interface{}]interface{})
	for index, value := range orders {
		indexedOrders[value] = index
	}
	return indexedOrders
}

func checkOrders(servedOrders []int, takeoutOrders []int, driveInOrders []int) {
	takeoutOrdersIndex := indexOrders(takeoutOrders)
	driveInOrdersIndex := indexOrders(driveInOrders)

	last_dinein_served_order_created_at := 0
	last_togo_served_order_created_at := 0
	for i := 0; i < len(servedOrders); i++ {
		v, ok := driveInOrdersIndex[servedOrders[i]].(int)
		if ok {
			if last_dinein_served_order_created_at > v {
				fmt.Println("Order is not servered correctly")
				break
			} else {
				last_dinein_served_order_created_at = v
			}
		} else {
			v, ok := takeoutOrdersIndex[servedOrders[i]].(int)
			if ok {
				if last_togo_served_order_created_at > v {
					fmt.Printf("Order is not servered correctly %v\n", servedOrders[i])
					break
				} else {
					last_togo_served_order_created_at = v
				}
			} else {
				fmt.Printf("This order is neither dine in nor take out\n")
			}
		}
	}
}

func main() {
	takeoutOrders := []int{1, 3, 5}
	driveInOrders := []int{2, 4, 6}
	servedOrders := []int{1, 2, 4, 6, 5, 3}
	checkOrders(servedOrders, takeoutOrders, driveInOrders)
}
