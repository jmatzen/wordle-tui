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
	maxAttempts = 6
	wordLength  = 5
)

var (
	// ANSI color codes
	reset     = "\033[0m"
	greenBg   = "\033[42m"
	yellowBg  = "\033[43m"
	grayBg    = "\033[40;37m" // Dark gray background with white text
	boldText  = "\033[1m"
)

var wordList = []string{
	"APPLE", "HOUSE", "PLANT", "TRAIN", "WORLD", "MUSIC", "WATER", "EARTH",
	"BRAIN", "CLOCK", "SMILE", "PHONE", "LIGHT", "BEACH", "SPACE", "CHAIR",
	"TABLE", "RIVER", "PIZZA", "OCEAN", "TIGER", "SNAKE", "QUEEN", "GHOST",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Select a random word
	targetWord := wordList[rand.Intn(len(wordList))]
	
	fmt.Println("=== WORDLE TUI ===")
	fmt.Println("Guess the 5-letter word. You have 6 attempts.")
	fmt.Println("After each guess, the colors will show:")
	fmt.Printf("%s GREEN %s = correct letter in correct position\n", greenBg, reset)
	fmt.Printf("%s YELLOW %s = correct letter in wrong position\n", yellowBg, reset)
	fmt.Printf("%s GRAY %s = letter not in the word\n\n", grayBg, reset)

	scanner := bufio.NewScanner(os.Stdin)
	attempts := 0
	guessHistory := []string{}
	
	// Game loop
	for attempts < maxAttempts {
		fmt.Printf("Attempt %d/%d > ", attempts+1, maxAttempts)
		scanner.Scan()
		guess := strings.ToUpper(strings.TrimSpace(scanner.Text()))
		
		// Validate input
		if len(guess) != wordLength {
			fmt.Printf("Please enter a %d-letter word\n", wordLength)
			continue
		}
		
		attempts++
		result := checkGuess(targetWord, guess)
		guessHistory = append(guessHistory, result)
		
		// Display all guesses with colors
		fmt.Println("\nYour guesses so far:")
		for _, coloredGuess := range guessHistory {
			fmt.Println(coloredGuess)
		}
		fmt.Println()
		
		// Check if player won
		if guess == targetWord {
			fmt.Println("🎉 Congratulations! You guessed the word correctly!")
			return
		}
	}
	
	// Game over - player lost
	fmt.Printf("Game over! The word was: %s\n", targetWord)
}

// checkGuess compares the guess with the target word and returns a colored string
func checkGuess(target, guess string) string {
	// Create array to track which letters in target have been matched
	targetChars := []rune(target)
	matched := make([]bool, len(targetChars))
	
	// Create array to store the result for each position
	resultColors := make([]string, len(guess))
	
	// First pass: find exact matches (green)
	for i, char := range guess {
		if i < len(target) && string(char) == string(targetChars[i]) {
			resultColors[i] = greenBg
			matched[i] = true
		}
	}
	
	// Second pass: find partial matches (yellow) or misses (gray)
	for i, char := range guess {
		if resultColors[i] != "" {
			// Already matched in the first pass
			continue
		}
		
		found := false
		// Look for the character elsewhere in the target
		for j, targetChar := range targetChars {
			if !matched[j] && char == targetChar {
				resultColors[i] = yellowBg
				matched[j] = true
				found = true
				break
			}
		}
		
		if !found {
			resultColors[i] = grayBg
		}
	}
	
	var result strings.Builder
	for i, char := range guess {
		result.WriteString(fmt.Sprintf("%s %c %s", resultColors[i], char, reset))
	}
	
	return result.String()
}
