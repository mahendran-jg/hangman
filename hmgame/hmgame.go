package hangman

import (
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type GameState string

const(
	Lost GameState = "Lost"
	Won = "Won"
	BadGuess = "BadGuess"
	GoodGuess = "GoodGuess"
	AlreadyGuessed = "AlreadyGuessed"
	Started = "Started"
	TotalGueses = 10
)
// HmGame : Hangman game struct
type HmGame struct {
	ID             string          // identifier
	State          GameState          // HmGame state
	GuessesLeft    int             // Remaining attempts
	Letters        []string        // Letters in the word
	Guesses   	   map[string]bool // Correct guesses
	WrongGuesses   map[string]bool // Wrong guesses
}

// NewHmGame : Start a new hangman game
func NewHmGame(word string) HmGame {
	letters := strings.Split(word, "")
	return HmGame{ID: uuid.New().String(),
		State:          Started,
		GuessesLeft:    TotalGueses, // starting with total no. of guesses
		Letters:        letters,
		Guesses:   		make(map[string]bool),
		WrongGuesses:   make(map[string]bool),
	}
}

func isSplLetter(wordLetter string) bool {
	return strings.ContainsAny("-' ", wordLetter)
}

// RevealWordByGuesses : reveal the word by the guesses
func RevealWordByGuesses(letters []string, used map[string]bool) string {
	revealedWord := ""

	for _, wordLetter := range letters {
		if used[wordLetter] {
			revealedWord += wordLetter
		} else if isSplLetter(wordLetter) {
			revealedWord += wordLetter

		} else {
			revealedWord += "_"
		}
	}

	return revealedWord
}

func hasWon(letters []string, used map[string]bool) bool {
	foundCnt := 0
	for _, letter := range letters {
		if used[letter] || isSplLetter(letter) {
			foundCnt++
		}
	}
	return foundCnt >= len(letters)
}

func findLetterInWord(guess string, letters []string) bool {
	for _, letter := range letters {
		if guess == letter {
			return true
		}
	}
	return false
}

// GuessLetter : Validate and update the guess
func GuessLetter(game HmGame, guess string) HmGame {
	// Return if the game is reached end state
	if game.State == Lost || game.State == Won {
		return game
	}

	// If already guessed this letter...
	if game.Guesses[guess] || game.WrongGuesses[guess]{
		game.State = AlreadyGuessed
	} else if findLetterInWord(guess, game.Letters) {
		game.Guesses[guess] = true
		game.State = GoodGuess
		if hasWon(game.Letters, game.Guesses) {
			game.State = Won
		}
	} else {
		game.GuessesLeft--
		game.State = BadGuess
		game.WrongGuesses[guess] = true
		if game.GuessesLeft == 0 {
			game.State = Lost
		}
	}

	return game
}

// PickRandomWord : Get a random word based on random index
func PickRandomWord(words []string) string {
	rand.Seed(time.Now().Unix())
	return words[rand.Intn(len(words))]
}

// ReadFromFile : Read words from a text file
func ReadFromFile(filePath string) []string {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	words := strings.Split(string(b), "\n")
	return words
}
