package main

import (
	"fmt"
)

func main() {
	var quantity int
	var length, width float64
	var customerName string

	quantity = 5
	length, width = 1.2, 2.4
	customerName = "Javier"

	fmt.Println(customerName)
	fmt.Println("has order", quantity, "units")
	fmt.Println("total area is ", length*width)

}
