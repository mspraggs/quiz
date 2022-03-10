package runner

import (
	"github.com/mspraggs/quiz/models"
)

// Parser defines the quiz parsing interface required by the QuizRunner.
type Parser interface {
	Parse() (models.Quiz, error)
}

// QuestionAnswer defines the interface used by the QuizRunner to ask for
// answers to quiz questions.
type QuestionAsker interface {
	AskQuestion(question *models.Question) (bool, error)
}

// QuizRunner encapsulates the logic required to run a quiz.
type QuizRunner struct {
	parser Parser
	asker  QuestionAsker
}

// NewQuizRunner instantiates a new QuizRunner instance with the provided Parser
// and QuestionAsker.
func NewQuizRunner(parser Parser, asker QuestionAsker) *QuizRunner {
	return &QuizRunner{
		parser: parser,
		asker:  asker,
	}
}

// Run runs the quiz, returning the total score.
func (r *QuizRunner) Run() (models.Score, error) {
	score := models.Score(0)

	quiz, err := r.parser.Parse()
	if err != nil {
		return 0, err
	}

	for _, question := range quiz {
		result, err := r.asker.AskQuestion(question)
		if err != nil {
			return 0, err
		}

		if result {
			score.Inc()
		}
	}

	return score, nil
}
