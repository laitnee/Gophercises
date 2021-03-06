package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "QuestionsAndAnswers.csv", "a csv filenamein the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s \n", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse provided csv file")
	}
	var correct int = 0
	problems := parseLines(lines)
	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, prob.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == prob.a {
			correct++
		}
	}
	fmt.Printf("You scored %d of %d \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
