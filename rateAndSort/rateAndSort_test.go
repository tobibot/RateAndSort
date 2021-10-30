package rateAndSort

import (
	"fmt"
	"os"
	"testing"
)

func Test_SortAndThenWriteFile(t *testing.T) {
	evalFile := "evaluation.json"
	filePath = evalFile
	tData := getStocksByType()

	for _, t := range stockTypes {
		sortStocks(tData[t])
	}
	tests := []struct {
		name    string
		data    map[StockType][]*stock
		wantErr bool
	}{
		{name: "test should save 1000 entries", data: tData, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath = evalFile
			if err := writeData(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("writeEvaluationFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_writeEvaluationFile(t *testing.T) {
	testFileName := "evaluation.json"
	deleteFile := false
	fakeData := false
	nOfStockTypes := 4
	sn := 100
	runs := 200
	tData := make([]stock, 0)

	if fakeData {
		for i := 0; i <= sn; i++ {
			t := StockType(fmt.Sprintf("HyperTech%d", i%nOfStockTypes))
			tData = append(tData, stock{Symbol: fmt.Sprintf("sym%d.de", i), Name: fmt.Sprintf("symbol %d corp", i), Value: 50, Type: t})
		}
	} else {
		filePath = testFileName
		tData, _ = readEvaluationFile()
	}

	tests := []struct {
		name    string
		data    []stock
		wantErr bool
	}{
		{name: "test should save 1000 entries", data: tData, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath = testFileName
			_, _ = os.Create(filePath)
			for i := 0; i < runs; i++ {
				if err := writeEvaluationFile(tt.data); (err != nil) != tt.wantErr {
					t.Errorf("writeEvaluationFile() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			if deleteFile {
				_ = os.Remove(filePath)
			}
		})
	}
}
