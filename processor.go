package main

import (
	"fmt"
)

// processCharacter is responsible for scraping and saving character data
func processCharacter(lsid string) {
	// Define the filename for the character's JSON file
	filename := fmt.Sprintf("characters/%s.json", lsid)

	// Check if the file exists first
	if isFileExist(filename) {
		// If the file exists, check if it has been updated in over 6 hours
		if isFileOld(filename) {
			fmt.Printf("%s has not been updated in over 6 hours. Proceeding with scrape...\n", lsid)
		} else {
			// If the file is within 6 hours old, skip the scrape
			fmt.Printf("%s is up-to-date (within the last 6 hours). Skipping scrape.\n\n", lsid)
			return
		}
	} else {
		// If the file doesn't exist, proceed with the scrape
		fmt.Printf("A JSON file for %s does not exist. Creating one now...\n", lsid)
	}
	// Load the base mounts (wishlist)
	baseMounts := LoadURLs() // This is defined in data.go

	// Scrape the character's data and mounts
	charData, charaMounts := ScrapeCharacter(lsid)

	// Compare mounts
	remainingMounts := mountCompare(baseMounts, charaMounts)

	// Update the remaining mounts in the struct
	charData.RemainingMounts = remainingMounts

	// Save character data to JSON
	if err := saveCharacterData(lsid, charData); err != nil {
		fmt.Println("Error saving character data:", err)
		return
	}

	// Print the number of remaining mounts
	if len(remainingMounts) > 0 {
		fmt.Printf("%d mounts remaining\n\n", len(remainingMounts))
	} else {
		fmt.Println("All mounts acquired!")
	}
}
