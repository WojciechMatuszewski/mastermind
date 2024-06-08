package game

import (
	"math/rand"
	"time"
)

type RandomStrategy struct{}

func (rs RandomStrategy) Guess(data GameData) []string {
	var newLetters []string

	for _, letterData := range data {
		if letterData.Sign == "+" {
			newLetters = append(newLetters, letterData.Letter)
		}

		if letterData.Sign == "" {
			newLetters = append(newLetters, rs.randomDifferentLetter(letterData.Letter))
		}

		if letterData.Sign == "-" {
			newLetters = append(newLetters, rs.randomDifferentLetter(letterData.Letter))
		}
	}

	return newLetters
}

func (rs RandomStrategy) randomLetter() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	letter := letters[r.Intn(len(letters))]
	return string(letter)
}

func (rs RandomStrategy) randomDifferentLetter(currentLetter string) string {
	for {
		newLetter := rs.randomLetter()
		if newLetter != currentLetter {
			return newLetter
		}

	}
}
