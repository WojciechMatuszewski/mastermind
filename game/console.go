package game

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Console struct {
	input io.Reader
}

func NewConsole(input io.Reader) Console {
	return Console{input}
}

func (c Console) NextScore(letters []string) ([]string, error) {
	reader := bufio.NewReader(c.input)

	fmt.Printf("Enter score for: %v", strings.Join(letters, ""))
	fmt.Print("\nScore: ")

	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return strings.Split(text, ""), nil
}
