package main

import "fmt"

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}
func paintNeeded(with float64, height float64) {
	area := with * height
	fmt.Printf("Total area: %.2f\n", area/10.0)
}
func double(number float64) float64 {
	return number * number
}
func main() {
	repeatLine("Hola", 2)
	paintNeeded(2, 4)
	fmt.Println("double: ", double(2))
}
