package service

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"hangman/datastore"
	hangman_api "hangman/protos/hangman_api/v1"
	"net/http"
	"os"
)

var (
	store     datastore.DataStore
	wordsFile = "words/words.txt"
)

func serveHealth(mux *runtime.ServeMux) {
	healthFunc := func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Up"))
	}
	err := mux.HandlePath("GET", "/health", healthFunc)
	if err != nil {
		log.Errorf("failed to get health status,", err)
	}
}

func runExternalAPIServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	err := hangman_api.RegisterHangmanAPIHandlerServer(ctx, mux, NewHangmanAPIServer())
	if err != nil {
		return err
	}
	serveHealth(mux)
	// Start HTTP server
	return http.ListenAndServe(":8003", mux)
}

// Retrieve all keys from a map
func keysFromMap(used map[string]bool) []string {
	keys := make([]string, 0, len(used))
	for k := range used {
		keys = append(keys, k)
	}
	return keys
}

func Start() {
	var err error
	// vaildate the word file and fail fast
	if _, err = os.Stat(wordsFile); err != nil {
		log.Fatalf("failed to find or validate the words file: %s\n", err)
	}
	// initialize the in memory datastore
	store, err = datastore.NewInMemoryStore()
	if err != nil {
		log.Fatalf("failed to initialize in-memory datastore: %s\n", err)
	}
	log.Println("Starting HTTP server...")
	runExternalAPIServer()
}
