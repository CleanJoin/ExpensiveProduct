package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadJsonFile(filename string) []map[string]interface{} {

	plan, err := ioutil.ReadFile("./src/" + filename)
	if err != nil {
		log.Fatal("Unable to read input file "+filename, err)
	}
	var data []map[string]interface{}
	json.Unmarshal(plan, &data)
	return data
}
