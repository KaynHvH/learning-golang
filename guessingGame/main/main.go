package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(100)

	fmt.Println("Guess a number between 1 and 100")
	fmt.Println(randNum)

	var guessStr string
	for {
		_, err := fmt.Scanln(&guessStr)
		if err != nil {
			fmt.Println("You have to enter a number")
			continue
		}

		guessNum, err := strconv.Atoi(guessStr)
		if err != nil {
			fmt.Println("You have to enter a number")
			continue
		}

		if guessNum < randNum {
			fmt.Println("Your number is too small")
		} else if guessNum > randNum {
			fmt.Println("Your number is too big")
		} else {
			fmt.Println("Congratulations! You won!")
			break
		}
	}
}
