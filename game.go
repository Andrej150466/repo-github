package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	count := 0
	guess := -1
	number := rand.Intn(10)

	for count < 3 && guess != number {
		fmt.Print("Guess the number (0...9): ")
		fmt.Scanln(&guess)
		if guess != number {
			if guess < number {
				fmt.Println("Your number is less")
			} else {
				fmt.Println("Your number is bigger")
			}
			count++
		}
	}

	if guess == number {
		fmt.Println("Your WON")
	} else {
		fmt.Println("Your lose:", number)
	}
}
