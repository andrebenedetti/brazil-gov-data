package aggregator

import (
	"encoding/csv"
	"fmt"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func ParseFile(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)
	r.Comma = ';'
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

// For some reason, gov-br data comes in Windows1252 encoding.
// This function takes a Windows1252 string and decodes it to UTF-8.
// If you provide a string with any other encoding, the result will be incorrect.
func win1252ToUtf8(str string) (string, error) {
	return charmap.Windows1252.NewDecoder().String(str)
}

func PrintAllRecords(records [][]string) {
	for _, record := range records {
		fmt.Println(record[1])
		fmt.Println(win1252ToUtf8(record[1]))
	}
}
