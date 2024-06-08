package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomStrategy(t *testing.T) {
	t.Run("Preserves the letters with +, changes letters with - and ''", func(t *testing.T) {
		gameData := GameData{
			{
				Letter: "A",
				Sign:   "+",
			},
			{
				Letter: "B",
				Sign:   "",
			},
			{
				Letter: "D",
				Sign:   "+",
			},
			{
				Letter: "C",
				Sign:   "-",
			},
		}

		randomStrategy := RandomStrategy{}

		letters := randomStrategy.Guess(gameData)

		assert.Equal(t, letters[0], gameData[0].Letter)
		assert.Equal(t, letters[2], gameData[2].Letter)

		assert.NotEqual(t, letters[1], gameData[1].Letter)
		assert.NotEqual(t, letters[3], gameData[3].Letter)
	})

}
