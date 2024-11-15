package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

// Scrape character data
func SearchCharacter(world, name string) {
	c := colly.NewCollector()

	// Error handler for scraping
	c.OnError(func(r *colly.Response, e error) {
		// Print the error message
		fmt.Println("Could not access Search page")
	})

	// Scrape character info
	c.OnHTML(".entry__link", func(e *colly.HTMLElement) {
		FullName := e.ChildText(".entry__name")
		World := e.ChildText(".entry__world")
		lsid := cleanID(e.Attr("href"))

		fmt.Println(lsid, "|", World, "|", FullName)
	})

	// Start scraping
	c.Visit(fmt.Sprintf("https://na.finalfantasyxiv.com/lodestone/character/?q=%s&worldname=%s&classjob=&race_tribe=&blog_lang=ja&blog_lang=en&blog_lang=de&blog_lang=fr&order=7", name, world,))


	return // lsid, FullName
}
