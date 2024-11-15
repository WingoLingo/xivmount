package main

import (
	"fmt"
	"os"
	"regexp"
	"github.com/gocolly/colly/v2"
	"strconv"
)

// Struct to hold character data
type FCData struct {
	FreeCompanyID      string   `json:"Free Company ID"`
	FreeCompanyName    string   `json:"Free Company Name"`
//	FreeCompanyWorld   string   `json:"Free Company World"`
//	FreeCompanyCrest   []string `json:"Free Company Crest"`
	FreeCompanyMembers []string `json:"Free Company Members"`
}

// Scrape FC Members page for Character Lodestone IDs
func ScrapeMember(fcid string) (FCData) {
	var freeData FCData
	page := 1
	fmt.Println("Checking FC", fcid)

	c := colly.NewCollector()

	// Error handler for scraping
	c.OnError(func(r *colly.Response, e error) {
		// Print the error message
		fmt.Println("Error occurred:", e)

		// Exit the program if there is an error during scraping
		os.Exit(1)
	})

	// Scrape FC info
	c.OnHTML(".entry__freecompany__box", func(e *colly.HTMLElement) {
		if page <2 {
			freeData.FreeCompanyName = e.ChildText(".entry__freecompany__name")
//			freeData.FreeCompanyWorld = e.ChildText(".entry__freecompany__gc")
//			fmt.Println(freeData.FreeCompanyName, freeData.FreeCompanyWorld)
		}
	})



	// Scrape total page count
	var totalPages int
	c.OnHTML(".btn__pager", func(e *colly.HTMLElement) {
		// Extract the total number of pages from the pagination info
		// This is assuming that the page info is in the format of "Page 1 of 5"
		pageInfo := e.Text
		re := regexp.MustCompile(`Page (\d+) of (\d+)`)
		matches := re.FindStringSubmatch(pageInfo)
		if len(matches) > 2 {
			totalPages, _ = strconv.Atoi(matches[2]) // Second match is the total page number
		}
	})

	// Scrape member list
	c.OnHTML("a.entry__bg", func(e *colly.HTMLElement) {
		member := e.Attr("href")
		member = cleanID(member)
		freeData.FreeCompanyMembers = append(freeData.FreeCompanyMembers, member)
	})

	// Start scraping the first page
	fmt.Printf("Visiting page 1\n")
	c.Visit(fmt.Sprintf("https://na.finalfantasyxiv.com/lodestone/freecompany/%s/member/?page=1", fcid))

	// Now that we have the total page count, visit all pages
	for page = 2; page <= totalPages; page++ {
		pageURL := fmt.Sprintf("https://na.finalfantasyxiv.com/lodestone/freecompany/%s/member/?page=%d", fcid, page)
		fmt.Printf("Visiting page %d\n", page)
		c.Visit(pageURL)
	}

	// Set the LodestoneID in the struct before returning
	freeData.FreeCompanyID = fcid

	return freeData
}
