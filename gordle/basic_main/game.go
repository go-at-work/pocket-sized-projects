package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const solutionLength = 5

type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}
	return g
}

func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := []rune(string(playerInput))

		if len(guess) != solutionLength {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution! Expected %d characters, got %d.\n", solutionLength, len(guess))
		} else {
			return guess
		}
	}
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	guess := g.ask()
	fmt.Printf("Your guess is: %s\n",
		string(guess))
}
