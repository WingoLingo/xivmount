package main

import (
	"fmt"
	"os"
)

func main() {
	// Ensure we have the correct number of command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <LodestoneID> or go run . fc <FreeCompanyID>")
		return
	}

	// Check if the first argument is "fc" to scrape a list of members
	if os.Args[1] == "fc" {
		fcid := os.Args[2]
		// Run the ScrapeMember function to scrape multiple members (a list of lsid)
		fcData := ScrapeMember(fcid)

		// Save FC data to JSON
		if err := saveFCData(fcid, fcData); err != nil {
			fmt.Println("Error saving character data:", err)
			return
		}

		// Process each member in the list
		for _, lsid := range fcData.FreeCompanyMembers {
			processCharacter(lsid)
		}
	} else {
		// Process a single Lodestone ID
		lsid := os.Args[1]
		processCharacter(lsid)
	}

	fmt.Println("Finished!")
}

// processCharacter is responsible for scraping and saving character data
func processCharacter(lsid string) {
	// Define the filename for the character's JSON file
	filename := fmt.Sprintf("characters/%s.json", lsid)

	// Check if the file exists first
	if isFileExist(filename) {
		// If the file exists, check if it has been updated in over 6 hours
		if isFileOld(filename) {
			fmt.Println("%s has not been updated in over 6 hours. Proceeding with scrape...", lsid)
		} else {
			// If the file is within 6 hours old, skip the scrape
			fmt.Println("%s is up-to-date (within the last 6 hours). Skipping scrape.", lsid)
			return
		}
	} else {
		// If the file doesn't exist, proceed with the scrape
		fmt.Println("A JSON file for %s does not exist. Creating one now...", lsid)
	}
	// Load the base mounts (wishlist)
	baseMounts := LoadURLs() // This should be defined somewhere else, e.g., loading from a file

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
