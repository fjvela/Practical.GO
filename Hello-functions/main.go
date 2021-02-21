package main

import (
	"fmt"
	"log"
)

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}
func paintNeeded(width float64, height float64) (float64, error) {
	if height < 0 {
		return 0, fmt.Errorf("height %.2f is invalid", height)
	}
	if width < 0 {
		return 0, fmt.Errorf("width %.2f is invalid", width)
	}
	area := (width * height) / 10.0
	return area, nil
}
func double(number float64) float64 {
	return number * number
}
func doublePointer(number *float64) {
	*number *= 2
}
func main() {
	repeatLine("Hola", 2)
	area, err := paintNeeded(2, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total area: %.2f\n", area)

	amount := double(2)
	fmt.Println("double: ", amount)
	fmt.Println("double: ", &amount) //print addr

	amount = 6
	doublePointer(&amount)
	fmt.Println(amount)

	myInt := 4
	myIntPoint := &myInt
	fmt.Println("myInt:", myInt)
	fmt.Println(*myIntPoint)

	area, err = paintNeeded(-2, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total area: %.2f\n", area)
}
