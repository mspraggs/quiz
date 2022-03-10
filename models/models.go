package models

// Question wraps a quiz question and the associated answer.
type Question struct {
	Question string
	Answer   string
}

// Quiz wraps a series of quiz questions.
type Quiz []*Question

// Score represents a quiz score.
type Score int

// Inc increments the score by 1.
func (s *Score) Inc() {
	*s = *s + 1
}
