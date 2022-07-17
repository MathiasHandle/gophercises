package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// Loads csv file and returns records from it
func getRecords() [][]string {
	csvFilename := flag.String("csv", "problems.csv", "csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Error reading a csv file: %s\n", *csvFilename))
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Error reading from reader %s", err))
	}

	return records
}

// Prints out error message and quits with os.Exit(1)
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Asks questions and returns total questions count and correct answers.
func askQuestions(records [][]string) (totalQ, correctA int) {
	totalQuestions := len(records)
	var correctAnswers int

	for i, record := range records {
		question := record[0]
		answer := record[1]

		fmt.Printf("\n\nQuestion %v/%v\n", i+1, totalQuestions)

		var userInput string
		fmt.Printf("What is %s ?\n", question)
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			fmt.Println("Error scanning answer from user: ", err)
		}

		if userInput == answer {
			correctAnswers += 1
		}
	}

	return totalQuestions, correctAnswers
}

func main() {
	records := getRecords()

	totalQuestions, correctAnswers := askQuestions(records)

	fmt.Printf("%v correct answers out of %v total questions asked", correctAnswers, totalQuestions)
}
