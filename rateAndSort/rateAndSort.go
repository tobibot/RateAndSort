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

	stocksByType = make(map[StockType][]*stock)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func StartRating() {
	_ = godotenv.Load(envFileName)
	finFolder = os.Getenv("FINFOLDER")
	filePath = path.Join(finFolder, environment, "config", inputFileName)

	playTheGame()
}

func playTheGame() {
	for {
		stocks, err := readEvaluationFile()
		if err != nil {
			log.Fatal()
		}

		randomize(stocks)
		randomizeTypes(stockTypes)

		for _, t := range stockTypes {
			stocksByType[t] = make([]*stock, 0)
			for _, stock := range stocks {
				stock1 := stock
				if stock.Type == t {
					stocksByType[t] = append(stocksByType[t], &stock1)
				}
			}
		}

		for _, t := range stockTypes {
			fmt.Printf("\nStarting type '%v'\n", t)
			for i := 0; i < len(stocksByType[t]); i += 2 {
				if i+1 == len(stocksByType[t]) {
					continue
				}
				stock1 := stocksByType[t][i]
				stock2 := stocksByType[t][i+1]
				msg := fmt.Sprintf("\nWhich one do you prefer?\n"+
					"\t(a) %s\n"+
					"\t(b) %s", stock1.Name, stock2.Name)
				fmt.Println(msg)

				text := ""
				fmt.Scanln(&text)

				switch text {
				case "aa":
					makeEvaluation(stock1, stock2, 2)
					break
				case "a":
					makeEvaluation(stock1, stock2, 1)
					break
				case "bb":
					makeEvaluation(stock2, stock1, 2)
					break
				case "b":
					makeEvaluation(stock2, stock1, 1)
					break
				case "x":
					fallthrough
				case "quit":
					fallthrough
				case "exit":
					writeData()
					return
				default:
					fmt.Println("bad input - no rating made")
				}
			}
		}
		fmt.Println("writing results to file")
		writeData()
	}
}

func writeData() {
	newStocks := make([]stock, 0)
	for _, t := range stockTypes {
		sortStocks(stocksByType[t])
		for _, s := range stocksByType[t] {
			newStocks = append(newStocks, *s)
		}
	}
	err := writeEvaluationFile(newStocks)
	if err != nil {
		log.Printf("Error writing ev-File: %v\n", err)
	}
}

func makeEvaluation(i *stock, d *stock, v int) {
	//fmt.Printf("increased %s from %d to %d\n", i.Symbol, i.Value, i.Value + v)
	//fmt.Printf("decreased %s from %d to %d\n", d.Symbol, d.Value, d.Value - v)
	i.increaseBy(v)
	d.decreaseBy(v)
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

func randomizeTypes(data []StockType) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
}

func sortStocks(data []*stock) {
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Value > data[j].Value
	})
}
