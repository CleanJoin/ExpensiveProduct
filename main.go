package main

import (
	"flag"
	"fmt"
	"product/internal"
)

var flagvar string

func init() {
	const (
		defaultPath = "db.csv"
		usage       = "Флаг выбран"
	)
	flag.StringVar(&flagvar, "input", defaultPath, usage)
}
func main() {

	flag.Parse()
	if flagvar[len(flagvar)-2:] == "sv" {
		listProduct := internal.ReadCsvFile(flagvar)
		popularProduct := internal.CompareCSV(listProduct)
		fmt.Printf("%+v\n", popularProduct)
	} else {

		listProduct := internal.ReadJsonFile(flagvar)
		popularProduct := internal.CompareJSON(listProduct)
		fmt.Printf("%+v\n", popularProduct)

	}

}
