package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

// Loads csv file and returns records from it
func getRecords(csvFlag *string, timerFlag *int) [][]string {

	file, err := os.Open(*csvFlag)
	if err != nil {
		exit(fmt.Sprintf("Error reading a csv file: %s\n", *csvFlag))
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
func askQuestions(problems []problem, timer *time.Timer) {
	var correctAnswers int

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %v = \n", i+1, problem.question)

		answerChan := make(chan string)
		go func() {
			var userInput string
			fmt.Scanln(&userInput)
			answerChan <- userInput
		}()

		select {
		case <-timer.C:
			fmt.Printf("%v correct answers out of %v total", correctAnswers, len(problems))
			return

		case answer := <-answerChan:
			if answer == problem.answer {
				correctAnswers++
			}
		}
	}

	fmt.Printf("%v correct answers out of %v total", correctAnswers, len(problems))
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "this is time limit for quiz in seconds")
	flag.Parse()

	records := getRecords(csvFilename, timeLimit)

	problems := getProblems(records)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	askQuestions(problems, timer)
}
