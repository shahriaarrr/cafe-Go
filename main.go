package main

import (
	"cafeManagement/cafe"
	"fmt"
)

func main() {
	orderItems := []string{"Coffee", "Tea", "mocha", "Latte", "Espresso", "Cappuccino", "Hot Chocolate"}

	for _, item := range orderItems {
		if err := cafe.AddOrder(item); err != nil {
			fmt.Printf("%sError adding order %s: %v%s\n", cafe.ColorRed, item, err, cafe.ColorReset)
		} else {
			fmt.Printf("Order %s added successfully\n", item)
		}
	}

	cafe.ListOrders()

	for i := 0; i < cafe.MaxOrders; i++ {
		if err := cafe.MarkOrderReady(i); err != nil {
			fmt.Printf("%sError marking order %d as ready: %v%s\n", cafe.ColorRed, i+1, err, cafe.ColorRed)
		} else {
			fmt.Printf("Order %d marked as ready\n", i+1)
		}
	}

	cafe.ListOrders()
}
