package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
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
	for i, problem := range problems {
		fmt.Printf("Problem #%v: %v = \n", i+1, problem.q)
	}
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