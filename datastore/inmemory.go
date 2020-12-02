package datastore

import (
	"errors"
	"fmt"
	"hangman/hmgame"
	"log"
)

// inMemoryStore : store gama data and state in memory database
type inMemoryStore struct {
	m map[string]hangman.HmGame
}

// NewInMemoryStore : Initialize an in memory datastore
func NewInMemoryStore() (DataStore, error) {
	return &inMemoryStore{m: make(map[string]hangman.HmGame)}, nil
}

// CreateHmGame : create new hangman game
func (imStore inMemoryStore) CreateHmGame(game hangman.HmGame) error {
	imStore.m[game.ID] = game
	log.Printf("HmGame ID %s inserted", game.ID)
	return nil
}

// FetchHmGame : get the hangman info
func (imStore inMemoryStore) FetchHmGame(id string) (hangman.HmGame, error) {
	game, ok := imStore.m[id]
	if !ok {
		errmsg := fmt.Sprintf("No game was found for ID: %s\n", id)
		log.Print(errmsg)
		return hangman.HmGame{}, errors.New(errmsg)
	}
	return game, nil
}

// UpdateHmGame : update hangman game info & state
func (imStore inMemoryStore) UpdateHmGame(game hangman.HmGame) error {
	imStore.m[game.ID] = game
	log.Printf("HmGame ID %s updated", game.ID)
	return nil
}


// DeleteHmGame : delete the hangman game by id
func (imStore inMemoryStore) DeleteHmGame(id string) (bool, error) {
	_, ok := imStore.m[id]
	if !ok {
		return false, nil
	}
	delete(imStore.m, id)
	return true, nil
}
