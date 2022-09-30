package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz, in seconds.")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open CSV file: %v", *csvFilename)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Application terminated.")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	numberOfCorrectAnswers := 0
problemloop:
	for i, pair := range problems {
		fmt.Printf("Problem #%v: %v = \n", i+1, pair.q)
		answerChannel := make(chan string)
		go func() {
			var userResponse string
			fmt.Scanf("&v\n", &userResponse)
			answerChannel <- userResponse
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerChannel:
			if answer == pair.a {
				fmt.Println("Correct!")
				numberOfCorrectAnswers++
			}
		}
	}
	gameOver(numberOfCorrectAnswers, len(lines))
}

func gameOver(correct int, total int) {
	fmt.Printf("Game Over. You correctly answered %v out of %v problems.\n", correct, total)
}

func parseLines(lines [][]string) []questionAnswerPair {
	result := make([]questionAnswerPair, len(lines))
	for i, line := range lines {
		result[i] = questionAnswerPair{
			q: line[0],
			a: line[1],
		}
	}
	return result
}

type questionAnswerPair struct {
	q string
	a string
}

func exit(message string) {
	fmt.Print(message)
}
