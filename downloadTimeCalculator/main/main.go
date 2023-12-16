package main

import (
	"fmt"
	"strings"
)

func main() {
	var fileSize, internetSpeed float64
	var unit string

	for {
		fmt.Println("Is your file in MB or GB?")
		fmt.Scanln(&unit)

		unit = strings.ToLower(unit)

		if unit != "mb" && unit != "gb" {
			fmt.Println("Wrong format. Please enter 'MB' or 'GB'.")
			continue
		}

		fmt.Println("What's your file size?")
		fmt.Scanln(&fileSize)

		if unit == "gb" {
			fileSize *= 1024.0
		}

		fmt.Println("What's your internet speed in MB/s?")
		fmt.Scanln(&internetSpeed)

		var convertToTime = fileSize / internetSpeed / 60.0

		if convertToTime >= 60 {
			fmt.Printf("Time: %.2f hours\n", convertToTime/60)
		} else if convertToTime >= 1 {
			fmt.Printf("Time: %.2f minutes\n", convertToTime)
		} else {
			fmt.Printf("Time: %.2f seconds\n", convertToTime*60)
		}
		break
	}
}
