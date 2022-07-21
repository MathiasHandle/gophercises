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

func ParseJSON(filePath string) map[string]chapter {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %s", err)
	}

	var jsonData map[string]chapter
	err = json.Unmarshal(file, &jsonData)
	if err != nil {
		fmt.Printf("Error parsing json data: %s", err)
	}

	return jsonData
}
