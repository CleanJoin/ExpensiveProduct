package internal

import (
	"encoding/csv"
	"log"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func ReadCsvFile(filename string) [][]string {
	f, err := os.Open("./src/" + filename)
	if err != nil {
		log.Fatal("Unable to read input file "+filename, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(charmap.Windows1251.NewDecoder().Reader(f))
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filename, err)
	}

	return records
}
