package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadJsonFile(filename string) []map[string]interface{} {

	plan, err := ioutil.ReadFile("./src/" + filename)
	if err != nil {
		log.Fatal("Unable to read input file "+filename, err)
	}
	var data []map[string]interface{}
	json.Unmarshal(plan, &data)

	file, err := os.Open("./src/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return data
}

func ReadJsonFile2(filename string, iProduct IProduct) *Product {
	NewProducts := new(Product)
	plan, err := ioutil.ReadFile("./src/" + filename)
	if err != nil {
		log.Fatal("Unable to read input file "+filename, err)
	}

	var f []interface{}
	err = json.Unmarshal(plan, &f)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range f {
		z := v.(map[string]interface{})
		NewProducts = iProduct.CompareJSON(z)
	}
	return NewProducts
}
