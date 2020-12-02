# hangman
Go code for hangman game

To compile: make all

To run: make run

API design and documentation: protos/hangman_api/v1/doc/index.html

Also atatched compliled binary from MacOs v 10.14 in bin/main


Examples:
Create Game:
curl --location --request POST 'http://localhost:8003/api/v1/games' \
--header 'Content-Type: application/json' \
--data-raw ''

{
    "id": "a11dfa00-3478-4fee-84dc-5f2393ac9d77",
    "word": "",
    "guessesLeft": 10,
    "wrongGuesses": [],
    "state": "Started"
}

Get Game:
curl --location --request GET 'http://localhost:8003/api/v1/game/a11dfa00-3478-4fee-84dc-5f2393ac9d77' \
--header 'Content-Type: application/json' \
--data-raw ''

{
    "id": "a11dfa00-3478-4fee-84dc-5f2393ac9d77",
    "word": "",
    "guessesLeft": 10,
    "wrongGuesses": [],
    "state": "Started"
}

Guess Game
curl --location --request PATCH 'http://localhost:8003/api/v1/games/a11dfa00-3478-4fee-84dc-5f2393ac9d77/guesses' \
--header 'Content-Type: application/json' \
--data-raw '{
  "guess": "a"
}'

{
    "id": "4acec45c-3c23-4a1c-badb-c433fb8d664d",
    "word": "___ a__ ___",
    "guessesLeft": 9,
    "wrongGuesses": [e],
    "state": "GoodGuess"
}

Delete Game
curl --location --request DELETE 'http://localhost:8003/api/v1/games/a11dfa00-3478-4fee-84dc-5f2393ac9d77'
{
    "deleted": true
}
