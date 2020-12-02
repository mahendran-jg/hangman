package datastore

import "hangman/hmgame"

// DataStore to store the game state and information
type DataStore interface {
	CreateHmGame(game hangman.HmGame) error
	FetchHmGame(id string) (hangman.HmGame, error)
	UpdateHmGame(game hangman.HmGame) error
	DeleteHmGame(id string) (bool, error)
}
