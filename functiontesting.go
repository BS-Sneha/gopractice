package main

import "fmt"

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func main() {
	totalPrice := calculateBill(10, 20)
	fmt.Println("Total Price", totalPrice)
}
