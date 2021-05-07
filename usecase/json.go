package main

import (
	"encoding/json"
	"os"
)

func readJSON(file_name string, filter func(map[string]interface{}) bool) []map[string]interface{} {
	file, _ := os.Open(file_name)
	defer file.Close()

	decoder := json.NewDecoder(file)
	filtered_data := []map[string]interface{}{}
	decoder.Token()

	data := map[string]interface{}{}
	for decoder.More() {
		decoder.Decode(&data)
		if filter(data) {
			filtered_data = append(filtered_data, data)
		}
	}
	return filtered_data
}
