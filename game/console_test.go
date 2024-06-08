package game

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsole(t *testing.T) {
	t.Run("Parses the input", func(t *testing.T) {
		userScore := "+-\n"
		userInput := strings.NewReader(userScore)

		console := NewConsole(userInput)
		receivedScore, err := console.NextScore([]string{"DCED"})

		assert.Nil(t, err)
		assert.Equal(t, strings.TrimSuffix(userScore, "\n"), strings.Join(receivedScore, ""))
	})

}
