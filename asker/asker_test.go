package asker

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mspraggs/quiz/models"
	"github.com/stretchr/testify/suite"
)

type AskerTestSuite struct {
	suite.Suite
}

func TestAskerTestSuite(t *testing.T) {
	suite.Run(t, new(AskerTestSuite))
}

func (s *AskerTestSuite) TestAskQuestion() {
	question := &models.Question{
		Question: "question",
		Answer:   "answer",
	}

	s.Run("correct answer", func() {
		output := &strings.Builder{}

		asker := NewQuestionAsker(output, strings.NewReader("answer\n"))

		correct, err := asker.AskQuestion(question)

		s.Require().Equal("question? ", output.String())
		s.Require().True(correct, "Expected correct answer to question.")
		s.Require().NoError(err)
	})

	s.Run("incorrect answer", func() {
		output := &strings.Builder{}

		asker := NewQuestionAsker(output, strings.NewReader("foo\n"))

		correct, err := asker.AskQuestion(question)

		s.Require().Equal("question? ", output.String())
		s.Require().False(correct, "Expected incorrect answer to question.")
		s.Require().NoError(err)
	})

	s.Run("read error", func() {
		output := &strings.Builder{}
		expectedErr := fmt.Errorf("EOF")

		asker := NewQuestionAsker(output, strings.NewReader(""))

		correct, err := asker.AskQuestion(question)

		s.Require().Equal("question? ", output.String())
		s.Require().Equal(expectedErr, err)
		s.Require().False(correct, "Expected incorrect answer to question.")
	})
}
