package game

import (
	"fmt"
	"os"
)

type LetterData struct {
	Letter string
	Sign   string
}

type GameData []LetterData

func Play() {
	console := NewConsole(os.Stdin)
	strategy := RandomStrategy{}

	guess := []string{"A", "B", "C", "D"}
	for {
		score, err := console.NextScore(guess)
		if err != nil {
			panic(err)
		}

		gameData := parseScore(guess, score)
		if isOver(gameData) {
			fmt.Println("Game finished!")
			os.Exit(0)
		}
		guess = strategy.Guess(gameData)
	}
}

func isOver(data GameData) bool {
	var isOver = true
	for _, d := range data {
		if d.Sign != "+" {
			isOver = false
		}
	}

	return isOver
}

func parseScore(guess []string, score []string) GameData {
	gameData := make(GameData, 4)

	for i := range gameData {
		var letterScore = ""
		if len(score) > i {
			letterScore = score[i]
		}

		gameData[i] = LetterData{
			Letter: guess[i],
			Sign:   letterScore,
		}
	}

	return gameData
}
