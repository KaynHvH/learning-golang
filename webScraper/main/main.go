package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	var cNC = colly.NewCollector()
	cNC.OnHTML("p", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	if err := cNC.Visit("https://www.google.com"); err != nil {
		log.Fatal(err)
	}
}
