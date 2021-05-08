package rateAndSort

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/joho/godotenv"
)

var (
	inputFileName = "evaluation.json" // todo args / flags
	environment   = "prod"
	envFileName   = environment + ".env"
	finFolder     = ""
	filePath      = ""
)

type stock struct {
	Symbol string
	Name   string
	Value  int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func StartRating() {
	_ = godotenv.Load(envFileName)
	finFolder = os.Getenv("FINFOLDER")
	filePath = path.Join(finFolder, environment, "config", inputFileName)

	entries, err := readEvaluationFile()
	if err != nil {
		log.Fatal()
	}

	playTheGame(entries)
}

func playTheGame(stocks []stock) {
	for {
		randomize(stocks)

		for i := 0; i < len(stocks); i += 2 {
			if i+1 == len(stocks) {
				continue
			}
			stock1 := &stocks[i]
			stock2 := &stocks[i+1]
			msg := fmt.Sprintf("\nWhich one do you prefer?\n"+
				"\t(a) %s\n"+
				"\t(b) %s\n", stock1.Name, stock2.Name)
			fmt.Println(msg)

			text := ""
			fmt.Scanln(&text)

			if text == "a" {
				stock1.Value += 1
				stock2.Value -= 1
				writeEvaluationFile(stocks)
			} else if text == "b" {
				stock1.Value -= 1
				stock2.Value += 1
				writeEvaluationFile(stocks)
			} else if text == "x" || text == "quit" || text == "exit" {
				//sort(stocks) // Todo
				//writeEvaluationFile(stocks)
				return
			} else {
				fmt.Println("bad input")
			}
		}
	}

}

func writeEvaluationFile(data []stock) (err error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	err = encoder.Encode(data)
	return err
}

func readEvaluationFile() ([]stock, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Create(filePath)
		return make([]stock, 0), err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return make([]stock, 0), err
	}

	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		return make([]stock, 0), err
	}

	data := make([]stock, 0)

	if len(byteValue) == 0 {
		return make([]stock, 0), fmt.Errorf("file contains no data")
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return make([]stock, 0), err
	}
	return data, nil
}

func randomize(sliceIn []stock) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(sliceIn), func(i, j int) { sliceIn[i], sliceIn[j] = sliceIn[j], sliceIn[i] })
}
