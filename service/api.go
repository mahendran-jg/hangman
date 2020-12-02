package service

import (
	hmgame "hangman/hmgame"
	"context"
	log "github.com/sirupsen/logrus"
	hangman_api "hangman/protos/hangman_api/v1"
)

// NewHangmanAPIServer constructs and returns a new HangmanAPIServer
func NewHangmanAPIServer() *HangmanAPIServer {
	return &HangmanAPIServer{}
}

func (h *HangmanAPIServer) CreateGame(ctx context.Context, req *hangman_api.CreateGameRequest) (*hangman_api.GameResponse, error) {
	words := hmgame.ReadFromFile(wordsFile)
	choosenWord := hmgame.PickRandomWord(words)
	game := hmgame.NewHmGame(choosenWord)
	store.CreateHmGame(game)
	return &hangman_api.GameResponse{
		Id: game.ID,
		GuessesLeft: int32(game.GuessesLeft),
		State: 		string(game.State),
	}, nil
}


func (h *HangmanAPIServer) GetGame(ctx context.Context, req *hangman_api.GameRequest) (*hangman_api.GameResponse, error) {
	game, err := store.FetchHmGame(req.GetId())
	if err != nil {
		return nil, err
	}
	gr := &hangman_api.GameResponse{
		Id:           game.ID,
		Word:         hmgame.RevealWordByGuesses(game.Letters, game.WrongGuesses),
		GuessesLeft:  int32(game.GuessesLeft),
		WrongGuesses: keysFromMap(game.WrongGuesses),
		State: 		string(game.State),
	}
	return gr, nil
}

func (h *HangmanAPIServer) UpdateGameGuesses(ctx context.Context, req *hangman_api.UpdateGameRequest) (*hangman_api.GameResponse, error) {
	game, err := store.FetchHmGame(req.GetId())
	if err != nil {
		return nil, err
	}

	game = hmgame.GuessLetter(game, req.GetGuess())
	store.UpdateHmGame(game)

	game, err = store.FetchHmGame(game.ID)

	gr := &hangman_api.GameResponse{
		Id:           game.ID,
		Word:         hmgame.RevealWordByGuesses(game.Letters, game.Guesses),
		GuessesLeft:  int32(game.GuessesLeft),
		WrongGuesses: keysFromMap(game.WrongGuesses),
		State: 		string(game.State),
	}
	return gr, nil
}

func (h *HangmanAPIServer) DeleteGame(ctx context.Context, req *hangman_api.DeleteGameRequest) (*hangman_api.DeleteGameResponse, error) {
	result, err := store.DeleteHmGame(req.GetId())
	if err != nil {
		log.Fatalf("Could not delete the game. Error: %s", err)
		return nil, err
	}
	return &hangman_api.DeleteGameResponse {
		Deleted: result,
	}, nil
}