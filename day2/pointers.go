package main

import (
	"fmt"
)

type Order struct {
	salesOrderId int
	totalAmount  float64
}

func main() {

	customer := []string{"Alice", "Bob", "Charlie"}
	b := &customer[0]
	fmt.Printf(" %p\n", b) // memory address
	fmt.Println("Original value:", *b)
	changeName(b, "Lalisa")
	fmt.Println("New value:", *b)

	fmt.Println("====================")
	// Pointer
	newOrder := Order{salesOrderId: 900001234, totalAmount: 18900.50}
	fmt.Println(&newOrder.totalAmount)
	fmt.Println(newOrder.totalAmount)
	applyDiscountPointer(&newOrder, 0.03)
	fmt.Println(&newOrder.totalAmount)
	fmt.Println(newOrder.totalAmount)

	fmt.Println("====================")
	//Copy data not use the pointer
	newOrder2 := Order{salesOrderId: 900001234, totalAmount: 17900.50}
	fmt.Println(&newOrder2.totalAmount)
	fmt.Println(newOrder2)
	newOrder2.totalAmount = applyDiscount(newOrder2, 0.03)
	fmt.Println(&newOrder2.totalAmount)
	fmt.Println(newOrder2)

}

// pointer
func changeName(oldName *string, newName string) {
	fmt.Printf("%p\n", oldName) //memory address
	*oldName = newName
}

// pointer
func applyDiscountPointer(order *Order, discount float64) {
	fmt.Println("applyDiscountPointer")
	fmt.Println(&order.totalAmount) //memory address
	order.totalAmount = order.totalAmount - (order.totalAmount * discount)
}

// copy (not use pointer
func applyDiscount(order Order, discount float64) float64 {
	fmt.Println("applyDiscount")
	fmt.Println(&order.totalAmount) //memory address
	return order.totalAmount - (order.totalAmount * discount)
}
