package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/mspraggs/quiz/asker"
	"github.com/mspraggs/quiz/parser"
	"github.com/mspraggs/quiz/runner"
)

func main() {
	pathLoc := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	fileReader, err := os.Open(*pathLoc)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return
	}

	csvReader := csv.NewReader(fileReader)
	csvParser := parser.NewCSVParser(csvReader)
	questionAsker := asker.NewQuestionAsker(os.Stdout, os.Stdin)

	quizRunner := runner.NewQuizRunner(csvParser, questionAsker)
	score, err := quizRunner.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Final score: %d\n", score)
}
