package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	err := convertJsonToCsv()
	if err != nil {
		fmt.Println("Failed with error: %s", err)
		os.Exit(1)
	}

}

func convertJsonToCsv() error {
	baseFolder := `C:\git\data1.1`
	targetBaseFolder := "/data"
	numericFile := "sentences_ch_de_numerics.json"
	fullTranscribedFile := "sentences_ch_de_transcribed.json"
	dialects := []string{
		"ag",
		"be",
		"bs",
		"gr",
		"lu",
		"sg",
		"vs",
		"zh",
	}

	jsonNumericFile, err := os.Open(filepath.Join(baseFolder, numericFile))
	if err != nil {
		return err
	}
	defer jsonNumericFile.Close()
	jsonFullFile, err := os.Open(filepath.Join(baseFolder, fullTranscribedFile))
	if err != nil {
		return err
	}
	defer jsonFullFile.Close()

	bytesNumeric, _ := ioutil.ReadAll(jsonNumericFile)
	bytesFull, _ := ioutil.ReadAll(jsonFullFile)

	var resultNumeric []map[string]interface{}
	json.Unmarshal([]byte(bytesNumeric), &resultNumeric)
	var resultFull []map[string]interface{}
	json.Unmarshal([]byte(bytesFull), &resultFull)

	for _, dialect := range dialects {
		dialectFull := "ch_" + dialect
		csvFile, err := os.Create(dialectFull + ".csv")
		defer csvFile.Close()
		csvwriter := csv.NewWriter(csvFile)
		csvwriter.Comma = '|'
		if err != nil {
			return err
		}
		for index, entry := range resultNumeric {
			if entry[dialectFull] == nil {
				continue
			}
			text := entry[dialectFull].(string)
			text2 := resultFull[index][dialectFull].(string)
			id := int(entry["id"].(float64))
			fileName := fmt.Sprintf("%s/%s/%s_%04d.wav", targetBaseFolder, dialect, dialectFull, id)
			if id > 9999 {
				fileName = fmt.Sprintf("%s/%s/%s_%05d.wav", targetBaseFolder, dialect, dialectFull, id)
			}
			record := []string{
				fileName,
				text,
				text2,
			}
			err := csvwriter.Write(record)
			if err != nil {
				return err
			}
		}
		csvwriter.Flush()
	}
	return nil
}
