package utils

import (
	"encoding/csv"
	"os"
)

func ExportCsv(data [][]string) (string, error) {
	fileName := "public/sample.csv"
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	write := csv.NewWriter(file)
	err = write.WriteAll(data)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
