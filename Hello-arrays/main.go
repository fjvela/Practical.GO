package main

import (
	"fmt"
	"log"

	"github.com/Practical.GO/Hello-arrays/read"
)

func arrays() {
	var notes [3]string = [3]string{"do", "re", "mi"}
	primes := [5]int{2, 3, 5, 7, 11}

	fmt.Println("notes: ", notes)
	fmt.Println("primes: ", primes)

	index := 1
	fmt.Println("index:", index, "note: ", notes[index])

	for i := 0; i < len(notes); i++ {
		fmt.Println("For: ", i, notes[i])
	}

	for i, item := range primes {
		fmt.Println("Range: ", i, item)
	}

	numbers := [3]float64{45.2, 34.44, 94.33}
	var total float64 = 0

	for _, number := range numbers {
		total += number
	}
	fmt.Println("Sum: ", total)

}
func main() {
	arrays()

	numbers, err := read.ReadFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var total float64 = 0
	for _, number := range numbers {
		total += number
	}
	count := float64(len(numbers))
	fmt.Println("Average: ", total/count)
}
