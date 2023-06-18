package shared

import (
	"encoding/json"
	"log"
	"os"
)

// ParseFile parses a file and apply its values to a struct
func ParseFile(path string, dataType interface{}) {
	jsonFile, err := os.ReadFile(path)
	if err != nil {
		log.Println("Failed to read file")
		log.Fatal(err)
	}

	if err = json.Unmarshal(jsonFile, dataType); err != nil {
		log.Println("Failed to unmarshal file")
		log.Fatal(err)
	}
}
