package rateAndSort

import (
	"fmt"
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

func init() {
	rand.Seed(time.Now().UnixNano())
}

func StartRating() {
	_ = godotenv.Load(envFileName)
	finFolder = os.Getenv("FINFOLDER")
	filePath = path.Join(finFolder, environment, "config", inputFileName)

	playTheGame()
}

func getStocksByType() map[StockType][]*stock {
	stocksByType := make(map[StockType][]*stock)

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

	return stocksByType
}

func playTheGame() {
	for {
		stocksByType := getStocksByType()

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
				_, _ = fmt.Scanln(&text)
				// time.Sleep(1 * time.Second)
				// if time.Second%2 == 0 {
				// 	text = "a"
				// } else {
				// 	text = "b"
				// }
				// text = "x"

				switch text {
				case "aa":
					makeEvaluation(stock1, stock2, 2)
				case "a":
					makeEvaluation(stock1, stock2, 1)
				case "bb":
					makeEvaluation(stock2, stock1, 2)
				case "b":
					makeEvaluation(stock2, stock1, 1)
				case "x", "q", "quit", "exit":
					writeData(stocksByType)
					return
				default:
					fmt.Println("bad input - no rating made")
				}
			}
			fmt.Println()
		}
		fmt.Println("writing results to file")
		_ = writeData(stocksByType)
	}
}

func makeEvaluation(i *stock, d *stock, v int) {
	i.increaseBy(v)
	d.decreaseBy(v)
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
