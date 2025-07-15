package cafe

import (
	"fmt"
	"time"
)

const (
	MaxOrders  = 5
	ColorRed   = "\033[0;31m"
	ColorReset = "\033[0m"
)

type Order struct {
	Item      string
	IsReady   bool
	OrderTime time.Time
}

var orders []Order

func AddOrder(item string) error {
	if item == "" {
		return ErrInvalidItem
	}

	if len(orders) >= MaxOrders {
		return ErrOrderLimitReached
	}

	orders = append(orders, Order{
		Item:      item,
		IsReady:   false,
		OrderTime: time.Now(),
	})
	return nil
}

func ListOrders() {
	fmt.Println("\nCurrent Orders:")
	for i, order := range orders {
		fmt.Printf("%d. Item: %s\nReady: %t\nOrder Time: %s\n\n",
			i+1, order.Item, order.IsReady, order.OrderTime.Format("15:04:05"))
		fmt.Println("====================================")
	}
}

func MarkOrderReady(index int) error {
	if index < 0 || index >= MaxOrders {
		return ErrInvalidIndex
	}
	if orders[index].IsReady {
		return ErrAlreadyReady
	}
	orders[index].IsReady = true
	return nil
}
