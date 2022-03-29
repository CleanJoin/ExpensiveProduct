package main

import (
	"flag"
	"fmt"
	"product/internal"
)

var flagvar string

func init() {
	const (
		defaultPath = "db.json"
		usage       = "Флаг выбран"
	)
	flag.StringVar(&flagvar, "input", defaultPath, usage)
}
func main() {

	flag.Parse()

	interProduct := internal.NewProduct()
	if flagvar[len(flagvar)-2:] == "sv" {
		product := internal.ReadCsvFile2(flagvar, interProduct)
		fmt.Printf("%+v", product)

	} else {
		product := internal.ReadJsonFile2(flagvar, interProduct)
		fmt.Printf("%+v", product)

	}

}
