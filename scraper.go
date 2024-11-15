package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

// Struct to hold character data
type CharacterData struct {
	LodestoneID     string   `json:"LodestoneID"`
	CharacterName   string   `json:"Character Name"`
	CharacterWorld  string   `json:"Character World"`
	Avatar          string   `json:"Avatar"`
	RemainingMounts []string `json:"Remaining Mounts"`
}

// Scrape character data
func ScrapeCharacter(lsid string) (CharacterData, []string) {
	var charData CharacterData
	var charaMounts []string

	c := colly.NewCollector()

	// Error handler for scraping
	c.OnError(func(r *colly.Response, e error) {
		// Print the error message
		fmt.Println("Could not access 'Mounts' page")
	})

	// Scrape character info
	c.OnHTML(".frame__chara__link.state_btn", func(e *colly.HTMLElement) {
		charData.CharacterName = e.ChildText(".frame__chara__name")
		charData.CharacterWorld = e.ChildText(".frame__chara__world")
		charData.Avatar = e.ChildAttr("img", "src")
	})

	// Scrape mount list
	c.OnHTML(".character__icon__list li", func(e *colly.HTMLElement) {
		cleanedMount := removePatch(e.ChildAttr("img", "src"))
		charaMounts = append(charaMounts, cleanedMount)
	})

	// Start scraping
	c.Visit(fmt.Sprintf("https://na.finalfantasyxiv.com/lodestone/character/%s/mount/", lsid))

	// Set the LodestoneID in the struct before returning
	charData.LodestoneID = lsid

	return charData, charaMounts
}
