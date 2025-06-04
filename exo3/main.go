package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	CorrectPosition = iota
	WrongPosition
	Absent
)

type LetterStatus int

type GuessResult []LetterStatus

type SecretGenerator interface {
	Generate(wordLength int) string
}

type DictionaryGenerator struct {
	Dictionary []string
}

func (d DictionaryGenerator) Generate(wordLength int) string {
	candidates := []string{}
	for _, w := range d.Dictionary {
		if len(w) == wordLength {
			candidates = append(candidates, w)
		}
	}
	if len(candidates) == 0 {
		return ""
	}
	return candidates[rand.Intn(len(candidates))]
}

func RandomLettersGenerator(wordLength int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	var sb strings.Builder
	for i := 0; i < wordLength; i++ {
		sb.WriteRune(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

type RandomGenerator struct{}

func (r RandomGenerator) Generate(wordLength int) string {
	return RandomLettersGenerator(wordLength)
}

type Guesser interface {
	Guess(wordLength int) string
}

type HumanGuesser struct{}

func (h HumanGuesser) Guess(wordLength int) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter a word of %d letters: ", wordLength)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) == wordLength && len(input) > 0 {
			return input
		}
		fmt.Println("Size incorrect. Please try again.")
	}
}

type IAGuesser struct {
	Dictionary []string
}

func (ia IAGuesser) Guess(wordLength int) string {
	candidates := []string{}
	for _, w := range ia.Dictionary {
		if len(w) == wordLength {
			candidates = append(candidates, w)
		}
	}
	if len(candidates) == 0 {
		return RandomLettersGenerator(wordLength)
	}
	return candidates[rand.Intn(len(candidates))]
}

func EvaluateGuess(secret, guess string) GuessResult {
	secretRunes := []rune(secret)
	guessRunes := []rune(guess)
	result := make(GuessResult, len(guessRunes))
	secretUsed := make([]bool, len(secretRunes))

	for i := range result {
		result[i] = Absent
	}

	for i := 0; i < len(guessRunes) && i < len(secretRunes); i++ {
		if guessRunes[i] == secretRunes[i] {
			result[i] = CorrectPosition
			secretUsed[i] = true
		}
	}

	for i := 0; i < len(guessRunes); i++ {
		if result[i] == CorrectPosition {
			continue
		}
		
		for j := 0; j < len(secretRunes); j++ {
			if !secretUsed[j] && guessRunes[i] == secretRunes[j] {
				result[i] = WrongPosition
				secretUsed[j] = true
				break
			}
		}
	}

	return result
}

func PrintResult(guess string, result GuessResult) {
	for i, c := range guess {
		switch result[i] {
		case CorrectPosition:
			fmt.Printf("[%c]", c)
		case WrongPosition:
			fmt.Printf("(%c)", c)
		case Absent:
			fmt.Printf(" %c ", c)
		}
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wordLength := 5
	maxTries := 6
	dictionary := []string{"apple", "grape", "peach", "melon", "berry", "lemon", "mango", "olive", "plumb", "guava"}

	reader := bufio.NewReader(os.Stdin)

	var generator SecretGenerator
	fmt.Println("Choose the secret generation method:")
	fmt.Println("1. Dictionary")
	fmt.Println("2. Random")
	fmt.Print("Your choice: ")

	var genChoice int
	fmt.Scan(&genChoice)
	reader.ReadString('\n')

	if genChoice == 1 {
		generator = DictionaryGenerator{Dictionary: dictionary}
	} else {
		generator = RandomGenerator{}
	}

	secret := generator.Generate(wordLength)
	if secret == "" {
		fmt.Println("No possible word for this length.")
		return
	}

	var guesser Guesser
	fmt.Println("\nChoose the game method:")
	fmt.Println("1. Human")
	fmt.Println("2. AI")
	fmt.Print("Your choice: ")

	var guessChoice int
	fmt.Scan(&guessChoice)
	reader.ReadString('\n')

	if guessChoice == 1 {
		guesser = HumanGuesser{}
	} else {
		guesser = IAGuesser{Dictionary: dictionary}
	}

	fmt.Printf("\nSecret word generated (length %d)\n", wordLength)
	fmt.Println("Legend: [letter] = correct position, (letter) = wrong position, letter = absent")
	fmt.Println()

	for tries := 1; tries <= maxTries; tries++ {
		guess := guesser.Guess(wordLength)
		result := EvaluateGuess(secret, guess)
		fmt.Printf("Try %d/%d: ", tries, maxTries)
		PrintResult(guess, result)

		if guess == secret {
			fmt.Println("Congratulations, you found the word!")
			return
		}
	}

	fmt.Printf("You lost! The word was: %s\n", secret)
}