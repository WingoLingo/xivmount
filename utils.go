package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"
)

// Function to clean lsid links
func cleanID(id string) string {
	reNum := regexp.MustCompile("[^0-9]") // Matches non-numeric characters
	return reNum.ReplaceAllString(id, "") // Returns the cleaned ID (numeric only)
}

// Function to remove patch number from URLs
func removePatch(url string) string {
	rePng := regexp.MustCompile(`\.png.*`)
	return rePng.ReplaceAllString(url, ".png") // Replace everything after ".png" with an empty string
}

// Function to compare mounts
func mountCompare(baseMounts, charaMounts []string) []string {
	countMap := make(map[string]struct{})
	for _, mount := range charaMounts {
		countMap[mount] = struct{}{}
	}

	var uniqueMounts []string
	for _, mount := range baseMounts {
		if _, exists := countMap[mount]; !exists {
			uniqueMounts = append(uniqueMounts, mount)
		}
	}
	return uniqueMounts
}

// Function to save character data to JSON
func saveCharacterData(lsid string, data CharacterData) error {
	filename := fmt.Sprintf("characters/%s.json", lsid)

	// Marshal the character data into JSON format
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Create the characters directory if it doesn't exist
	if err := os.MkdirAll("characters", os.ModePerm); err != nil {
		return err
	}

	// Save the JSON file
	return os.WriteFile(filename, file, 0644)
}

// Function to save FC data to JSON
func saveFCData(fcid string, data FCData) error {
	filename := fmt.Sprintf("freecompany/%s.json", fcid)

	// Marshal the character data into JSON format
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Create the characters directory if it doesn't exist
	if err := os.MkdirAll("freecompany", os.ModePerm); err != nil {
		return err
	}

	// Save the JSON file
	return os.WriteFile(filename, file, 0644)
}

// Function to check if the JSON file exists
func isFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// Function to check if the JSON file hasn't been updated in over 6 hours
func isFileOld(filename string) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		// If the file doesn't exist, return false
		return false
	}

	// Check if the file hasn't been updated in over 6 hours
	modTime := fileInfo.ModTime()
	return time.Since(modTime) > 6*time.Hour
}
