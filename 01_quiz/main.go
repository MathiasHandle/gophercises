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

type problem struct {
	question string
	answer   string
}

// Parses records and returns []problem
func getProblems(records [][]string) []problem {
	problems := make([]problem, len(records))

	for i, line := range records {
		problems[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}

	return problems
}

// Asks a question and validates user intput.
func askQuestions(problems []problem) {
	var correctAnswers int

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %v = \n", i+1, problem.question)

		var userInput string
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			fmt.Println("Error scanning answer from user: ", err)
		}

		if userInput == problem.answer {
			correctAnswers++
		}
	}

	fmt.Printf("%v correct answers out of %v total", correctAnswers, len(problems))
}

func main() {
	records := getRecords()

	problems := getProblems(records)

	askQuestions(problems)
}
