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
    <-timer.C

	numberOfCorrectAnswers := 0
	numberOfProblems := len(lines)
	for i, problem := range problems {
		fmt.Printf("Problem #%v: %v = \n", i+1, problem.q)
		var userResponse string
		fmt.Scanf("%v\n", &userResponse)
		if userResponse == problem.a {
			fmt.Println("Correct!")
			numberOfCorrectAnswers++
		} else {
			fmt.Println("Nope.")
		}
	}
	fmt.Printf("Game Over. You correctly answered %v out of %v problems.\n", numberOfCorrectAnswers, numberOfProblems)
}


func parseLines(lines [][]string) []problemSolutionPair {
	result := make([]problemSolutionPair, len(lines))
	for i, line := range lines {
		result[i] = problemSolutionPair{
			q: line[0],
			a: line[1],
		}
	}
	return result
}

type problemSolutionPair struct {
	q string
	a string
}

func exit(message string){
	fmt.Print(message)
}