package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Ensure we have the correct number of command-line arguments
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run . <LodestoneID> , go run . fc <FreeCompanyID> or go run . search <World> <Name>")
		return
	}

	// Check if the first argument is "search" to look up a lsid
	if os.Args[1] == "search" {
		var name string
		world := os.Args[2]
		if len(os.Args) > 3 {
			name = strings.Join(os.Args[3:], "+")
		} else {
			name = os.Args[3]
		}


		// Run the SearchCharacter function to 
		SearchCharacter(world, name)

		return
		}

	// Check if the first argument is "fc" to scrape a list of members
	if os.Args[1] == "fc" {
		fcid := os.Args[2]
		// Run the ScrapeMember function to scrape multiple members (a list of lsid)
		fcData := ScrapeMember(fcid)

		// Save FC data to JSON
		if err := saveFCData(fcid, fcData); err != nil {
			fmt.Printf("Error saving character data:", err)
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

	fmt.Printf("Finished!")
}
