package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	secretNumber := rand.Intn(100) + 1
	fmt.Println("I've chosen a number between 1 and 100")
	fmt.Println("Try guess it")

	reader := bufio.NewReader(os.Stdin)
	sucess := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Print("Guesses ", guesses+1, " your number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		//clear
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < secretNumber {
			fmt.Println("Your number is lower")
		} else if guess > secretNumber {
			fmt.Println("Your number is higher")
		} else {
			sucess = true
			break
		}
	}

	if !sucess {
		fmt.Println("Sorry!! the guess number was: ", secretNumber)
	}
}
