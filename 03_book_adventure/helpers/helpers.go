package helpers

import (
	"encoding/json"
	"fmt"
	"os"
)

type chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Story map[string]chapter

// Parses JSON file to a struct
func ParseJSON(file *os.File) Story {
	d := json.NewDecoder(file)

	var jsonData Story
	err := d.Decode(&jsonData)
	if err != nil {
		fmt.Printf("Error parsing json data: %s", err)
	}

	return jsonData
}
