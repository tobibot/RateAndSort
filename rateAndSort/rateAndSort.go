package rateAndSort

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"sort"
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

			switch text {
			case "aa":
				makeEvaluation(stocks, stock1, stock2, 2)
				break
			case "a":
				makeEvaluation(stocks, stock1, stock2, 1)
				break
			case "bb":
				makeEvaluation(stocks, stock2, stock1, 2)
				break
			case "b":
				makeEvaluation(stocks, stock2, stock1, 1)
				break
			case "x":
				fallthrough
			case "quit":
				fallthrough
			case "exit":
				sortStocks(stocks)
				writeEvaluationFile(stocks)
				return
			default:
				fmt.Println("bad input - no rating made")
			}
		}
	}

}

func makeEvaluation(stocks []stock, i *stock, d *stock, v int) {
	i.increaseBy(v)
	d.decreaseBy(v)
	writeEvaluationFile(stocks)
}

func (s *stock) decreaseBy(b int) {
	s.Value -= b
	if s.Value < 0 {
		s.Value = 0
	}
}

func (s *stock) increaseBy(b int) {
	s.Value += b
	if s.Value > 100 {
		s.Value = 100
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

func randomize(data []stock) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
}

func sortStocks(data []stock) {
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Value > data[j].Value
	})
}

func sortThenQuit(data []stock) {
	sortStocks(data)
	writeEvaluationFile(data)
}
