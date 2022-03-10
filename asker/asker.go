package asker

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/mspraggs/quiz/models"
)

type questionAsker struct {
	outputWriter io.Writer
	inputReader  io.Reader
}

func NewQuestionAsker(outputWriter io.Writer, inputReader io.Reader) *questionAsker {
	return &questionAsker{
		outputWriter: outputWriter,
		inputReader:  inputReader,
	}
}

func (a *questionAsker) AskQuestion(question *models.Question) (bool, error) {
	fmt.Fprintf(a.outputWriter, "%s? ", question.Question)

	reader := bufio.NewReader(a.inputReader)
	givenAnswer, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	givenAnswer = strings.Trim(givenAnswer, " \n")

	correctAnswer := false
	if givenAnswer == question.Answer {
		correctAnswer = true
	}

	return correctAnswer, nil
}
