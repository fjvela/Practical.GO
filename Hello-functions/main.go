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
func main() {
	repeatLine("Hola", 2)
	area, err := paintNeeded(2, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total area: %.2f\n", area)

	area, err = paintNeeded(-2, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total area: %.2f\n", area)

	fmt.Println("double: ", double(2))
}
