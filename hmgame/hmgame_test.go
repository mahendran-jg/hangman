package hangman

import (
	"testing"
)

func TestWinsWithoutSpl(t *testing.T) {
	g := NewHmGame("mellow")
	g = GuessLetter(g, "m")
	g = GuessLetter(g, "e")
	g = GuessLetter(g, "l")
	g = GuessLetter(g, "o")
	g = GuessLetter(g, "w")
	if g.State != Won {
		t.Errorf("Game state is wrong. Got %s, should be 'won'", g.State)
	}
}

func TestFindLetterInWord(t *testing.T) {
	word := []string{"h", "e", "l", "l", "0"}

	guess := "h"
	hasLetter := findLetterInWord(guess, word)
	if hasLetter != true {
		t.Errorf("Word %s does not contain letter %s", word, guess)
	}

	guess = "0"
	hasLetter = findLetterInWord(guess, word)
	if hasLetter != true {
		t.Errorf("Word %s does not contain letter %s", word, guess)
	}

	guess = "c"
	hasLetter = findLetterInWord(guess, word)
	if hasLetter == true {
		t.Errorf("Word %s should not contain letter %s", word, guess)
	}
}

func TestRevealWordByGuesses(t *testing.T) {
	letters := []string{"h", "e", "l", "l", "0"}
	goodGuesses := map[string]bool{"l": true}

	revealedWord := RevealWordByGuesses(letters, goodGuesses)
	if revealedWord != "__ll_" {
		t.Errorf("Revealed word is not correct. Got %s, should be '_oo'", revealedWord)
	}

	letters = []string{"b", "a", "r"}
	goodGuesses = map[string]bool{"o": true}

	revealedWord = RevealWordByGuesses(letters, goodGuesses)
	if revealedWord != "___" {
		t.Errorf("Revealed word is not correct. Got %s, should be '___'", revealedWord)
	}
}

func TestRevealWordSpecial(t *testing.T) {
	letters := []string{"h", "e", "l", "l", "o", "-", "w", "o", "r", "l", "d"}
	goodGuesses := map[string]bool{"l": true}

	revealedWord := RevealWordByGuesses(letters, goodGuesses)
	if revealedWord != "__ll_-___l_" {
		t.Errorf("Revealed word is not correct. Got %s, should be '__ll_-___l_'", revealedWord)
	}

	letters = []string{"b", "o", "b", "'", "s", " ", "p", "e", "n"}
	goodGuesses = map[string]bool{"b": true}

	revealedWord = RevealWordByGuesses(letters, goodGuesses)
	if revealedWord != "b_b'_ ___" {
		t.Errorf("Revealed word is not correct. Got %s, should be 'b_b'_ ___", revealedWord)
	}
}

func TestHmGameLose(t *testing.T) {
	g := NewHmGame("apple")
	// Misses 10 letters in a row so we lose the game
	g = GuessLetter(g, "a") // valid case
	g = GuessLetter(g, "b")
	g = GuessLetter(g, "c")
	g = GuessLetter(g, "d")
	g = GuessLetter(g, "e") // valid case
	g = GuessLetter(g, "f")
	g = GuessLetter(g, "g")
	g = GuessLetter(g, "h")
	g = GuessLetter(g, "i")
	g = GuessLetter(g, "j")
	g = GuessLetter(g, "k")
	g = GuessLetter(g, "t")

	if g.State != Lost {
		t.Errorf("HmGame state should be 'Lost'. Got %s", g.State)
	}
	if g.GuessesLeft != 0 {
		t.Errorf("HmGame must have no guesses left. Got %d", g.GuessesLeft)
	}
}
