package internal

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

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
	fmt.Println(records)
	return records
}
func ReadCsvFile2(filename string, iProduct IProduct) *Product {
	NewProducts := new(Product)

	f, err := os.Open("./src/" + filename)
	if err != nil {
		log.Fatal("Unable to read input file "+filename, err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(charmap.Windows1251.NewDecoder().Reader(f))

	for scanner.Scan() {
		descProduct := strings.Split(scanner.Text(), ";")
		NewProducts = iProduct.CompareCSV(descProduct)
	}

	return NewProducts
}
