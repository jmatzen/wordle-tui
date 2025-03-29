# Wordle TUI

A simple, terminal-based implementation of the popular word-guessing game Wordle.

## Description

Wordle TUI is a text-based version of the Wordle game that runs in your terminal. Players have 6 attempts to guess a random 5-letter word. After each guess, the game provides color-coded feedback to help you narrow down the correct word.

## Features

- Terminal user interface with color feedback
- 24 built-in words
- 6 attempts to guess the correct word
- Visual feedback using colors:
  - Green: Correct letter in correct position
  - Yellow: Correct letter in wrong position
  - Gray: Letter not in the word

## Installation

1. Make sure you have Go installed on your system. If not, download and install it from [go.dev](https://go.dev/dl/).

2. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/wordle-tui.git
   cd wordle-tui
   ```

3. Build the program:
   ```bash
   go build -o wordle
   ```

## Usage

Run the game with the following command:

```bash
./wordle
```

## How to Play

1. The game will select a random 5-letter word.
2. Enter your guess when prompted (must be a 5-letter word).
3. After each guess, you'll receive color-coded feedback:
   - 🟩 GREEN: The letter is correct and in the right position
   - 🟨 YELLOW: The letter is in the word but in the wrong position
   - ⬛ GRAY: The letter is not in the word
4. Use the feedback to inform your next guess.
5. Try to guess the word within 6 attempts!

## Example Gameplay

```
=== WORDLE TUI ===
Guess the 5-letter word. You have 6 attempts.
After each guess, the colors will show:
GREEN  = correct letter in correct position
YELLOW  = correct letter in wrong position
GRAY  = letter not in the word

Attempt 1/6 > TRAIN

Your guesses so far:
GRAY T GRAY R YELLOW A GRAY I GRAY N 

Attempt 2/6 > SPACE

Your guesses so far:
GRAY T GRAY R YELLOW A GRAY I GRAY N 
GRAY S GRAY P GREEN A GRAY C GRAY E 

...
```

## Customization

You can easily modify the word list by editing the `wordList` variable in the source code.

## Dependencies

This game uses only standard Go libraries:
- `bufio` - For reading user input
- `fmt` - For formatting output
- `math/rand` - For selecting a random word
- `os` - For accessing standard input/output
- `strings` - For string manipulation
- `time` - For random seed initialization

## License

[MIT License](LICENSE)

## Future Improvements

- Add a dictionary check to validate guesses
- Expand the word list
- Add difficulty levels
- Add statistics tracking
- Implement keyboard display showing used letters

---

Feel free to contribute to this project by submitting issues or pull requests!
