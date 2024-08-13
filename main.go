package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	c := colly.NewCollector()

	url := "https://books.toscrape.com/"

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Title:", e.Text)
	})

	c.OnHTML("article", func(e *colly.HTMLElement) {
		title := e.DOM.Find("a").Text()
		price := e.DOM.Find(".price_color").Text()
		fmt.Println("Title:", title)
		fmt.Println("Price:", price)
	})

	c.Visit(url)

}
