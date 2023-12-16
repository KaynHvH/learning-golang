package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var lower = "abcdefghjklmnoupqrstuvwxyz"
var upper = strings.ToUpper(lower)
var numbers = "1234567890"
var specialChars = "!?@#$%^&*()[]<>-=_+/;',.~"

var all = lower + upper + numbers + specialChars

var passLen int
var pass string

func main() {
	for {
		fmt.Println("How long should your password be?")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred:", err)
			continue
		}

		input = strings.TrimSpace(input)
		passLen, err = strconv.Atoi(input)
		if err != nil || passLen <= 0 {
			fmt.Println("You have to enter a positive number")
			continue
		}

		break
	}

	var timeNow = time.Now().UnixNano()
	random := rand.New(rand.NewSource(timeNow))

	for i := 0; i < passLen; i++ {
		pass += string(all[random.Intn(len(all))])
	}

	fmt.Printf("Your password is %v", pass)
}
