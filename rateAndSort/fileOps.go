package rateAndSort

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func writeData(stocksByType map[StockType][]*stock) error {
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
		return err
	}
	return nil
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
	if err != nil {
		return err
	}
	return nil
}

func readEvaluationFile() ([]stock, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, _ = os.Create(filePath)
		return make([]stock, 0), err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return make([]stock, 0), err
	}

	defer func() { _ = file.Close() }()
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
