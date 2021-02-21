package main

import (
	"fmt"
	"log"

	"github.com/Practical.GO/Hello-packages/read"
)

func testGrade() {
	fmt.Print("Grade: ")
	grade, err := read.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "Passed"
	} else {
		status = "Failed"
	}
	fmt.Println("a grade of ", grade, " is ", status)
}

func convertTemperature() {
	fmt.Print("Farenheit: ")
	farenheit, err := read.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (farenheit - 32) * 5 / 9
	fmt.Print("%0.2f degrees Celsius\n", celsius)
}

func main() {
	testGrade()
	convertTemperature()

}
