package parser

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mspraggs/quiz/models"
	"github.com/mspraggs/quiz/parser/mocks"
	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
	controller    *gomock.Controller
	mockCSVReader *mocks.MockCSVReader
}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}

func (s *ParserTestSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())
	s.mockCSVReader = mocks.NewMockCSVReader(s.controller)
}

func (s *ParserTestSuite) TestParse() {
	s.Run("parses rows and returns Quiz instance", func() {
		csvContents := [][]string{
			{"one", "foo"},
			{"two", "bar"},
		}
		s.mockCSVReader.EXPECT().
			ReadAll().Return(csvContents, nil)

		expectedQuiz := models.Quiz([]*models.Question{
			{Question: "one", Answer: "foo"},
			{Question: "two", Answer: "bar"},
		})

		parser := NewCSVParser(s.mockCSVReader)

		quiz, err := parser.Parse()

		s.Require().Nil(err)
		s.Require().Equal(expectedQuiz, quiz)
	})

	s.Run("returns error", func() {
		s.Run("when provided incorrect CSV layout", func() {
			csvContents := [][]string{
				{"one", "foo"},
				{"two"},
			}
			s.mockCSVReader.EXPECT().
				ReadAll().Return(csvContents, nil)

			parser := NewCSVParser(s.mockCSVReader)

			quiz, err := parser.Parse()

			s.Require().Nil(quiz)
			s.Require().EqualError(err, "expected CSV row with two entries")
		})
		s.Run("from CSV reader", func() {
			s.mockCSVReader.EXPECT().
				ReadAll().Return(nil, fmt.Errorf("oh no"))

			parser := NewCSVParser(s.mockCSVReader)

			quiz, err := parser.Parse()

			s.Require().Nil(quiz)
			s.Require().EqualError(err, "oh no")
		})
	})
}
