package parser

import (
	"fmt"

	"github.com/mspraggs/quiz/models"
)

// CSVReader defines the Reader interface required by the CSV parser.
type CSVReader interface {
	ReadAll() ([][]string, error)
}

type csvParser struct {
	reader CSVReader
}

// NewCSVParser instantiates an CSV parser instance with the provided CSV reader
// instance.
func NewCSVParser(reader CSVReader) *csvParser {
	return &csvParser{
		reader: reader,
	}
}

// Parse extracts questions and answers from the CSV reader object and uses them
// to construct a Quiz model instance.
func (p *csvParser) Parse() (models.Quiz, error) {
	data, err := p.reader.ReadAll()
	if err != nil {
		return nil, err
	}

	questions := make([]*models.Question, len(data))

	for i, row := range data {
		question, err := p.parseRow(row)
		if err != nil {
			return nil, err
		}

		questions[i] = question
	}

	return models.Quiz(questions), nil
}

func (p *csvParser) parseRow(row []string) (*models.Question, error) {
	if len(row) != 2 {
		return nil, fmt.Errorf("expected CSV row with two entries")
	}

	return &models.Question{
		Question: row[0],
		Answer:   row[1],
	}, nil
}
