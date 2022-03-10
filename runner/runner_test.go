package runner

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mspraggs/quiz/models"
	"github.com/mspraggs/quiz/runner/mocks"
	"github.com/stretchr/testify/suite"
)

type RunnerTestSuite struct {
	suite.Suite
	controller *gomock.Controller
	mockAsker  *mocks.MockQuestionAsker
	mockParser *mocks.MockParser
}

func TestRunnerTestSuite(t *testing.T) {
	suite.Run(t, new(RunnerTestSuite))
}

func (s *RunnerTestSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())
	s.mockAsker = mocks.NewMockQuestionAsker(s.controller)
	s.mockParser = mocks.NewMockParser(s.controller)
}

func (s *RunnerTestSuite) TestRun() {
	quiz := models.Quiz([]*models.Question{
		{
			Question: "one",
			Answer:   "yes",
		},
		{
			Question: "two",
			Answer:   "no",
		},
		{
			Question: "three",
			Answer:   "maybe",
		},
	})

	s.Run("runs quiz", func() {
		s.Run("with no questions and returns zero score", func() {
			s.mockParser.EXPECT().Parse().Return(models.Quiz([]*models.Question{}), nil)

			quizRunner := NewQuizRunner(s.mockParser, s.mockAsker)

			score, err := quizRunner.Run()

			s.Require().NoError(err)
			s.Require().Equal(models.Score(0), score)
		})
		s.Run("with some questions and returns score", func() {
			gomock.InOrder(
				s.mockParser.EXPECT().Parse().Return(quiz, nil),
				s.mockAsker.EXPECT().
					AskQuestion(quiz[0]).Return(true, nil),
				s.mockAsker.EXPECT().
					AskQuestion(quiz[1]).Return(false, nil),
				s.mockAsker.EXPECT().
					AskQuestion(quiz[2]).Return(true, nil),
			)

			quizRunner := NewQuizRunner(s.mockParser, s.mockAsker)

			score, err := quizRunner.Run()

			s.Require().NoError(err)
			s.Require().Equal(models.Score(2), score)
		})
	})

	s.Run("handles error", func() {
		s.Run("from parser and forwards to caller", func() {
			expectedErr := fmt.Errorf("oh no")
			s.mockParser.EXPECT().Parse().Return(nil, expectedErr)

			quizRunner := NewQuizRunner(s.mockParser, s.mockAsker)

			score, err := quizRunner.Run()

			s.Require().Equal(expectedErr, err)
			s.Require().Equal(models.Score(0), score)
		})
		s.Run("from asker and forwards to caller", func() {
			expectedErr := fmt.Errorf("oh no")
			s.mockParser.EXPECT().Parse().Return(quiz, nil)
			s.mockAsker.EXPECT().AskQuestion(quiz[0]).Return(false, expectedErr)

			quizRunner := NewQuizRunner(s.mockParser, s.mockAsker)

			score, err := quizRunner.Run()

			s.Require().Equal(expectedErr, err)
			s.Require().Equal(models.Score(0), score)
		})
	})
}
