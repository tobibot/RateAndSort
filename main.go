package main

import (
	"log"
	"math/rand"
	"time"
)

var (
	inputFileName = "input.json" // todo args / flags
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	entries, err := readInputFile(inputFileName)
	if err != nil {
		log.Fatal()
	}

	playTheGame(entries)

}

func readInputFile(fileName string) ([]interface{}, error) {
	return nil, nil
}

func playTheGame(entries []interface{}) {

	for {
		for _, entry := range entries {
			// save order after each evaluation

			_ = entry
		}

		// todo if input == x, quit, exit => exit

	}

}
