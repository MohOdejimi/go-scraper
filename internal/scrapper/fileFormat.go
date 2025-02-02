package scrapper

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strings"
)

 
func DetermineFileFormat(formatFlag, outputFlag string, parsedData []string) string {
	var filename string
	format := strings.ToLower(formatFlag)
 
	content := strings.Join(parsedData, "\n")

	switch format {
	case "txt":
		filename = outputFlag + ".txt"
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %s", err)
		}

	case "json":
		filename = outputFlag + ".json"
		jsonData, err := json.MarshalIndent(map[string]interface{}{"content": parsedData}, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling JSON: %s", err)
		}
		err = os.WriteFile(filename, jsonData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %s", err)
		}

	case "csv":
		filename = outputFlag + ".csv"
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Error creating file: %s", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
 
		err = writer.Write([]string{"Scraped Content"})
		if err != nil {
			log.Fatalf("Error writing to CSV: %s", err)
		}
 
		for _, line := range parsedData {
			err = writer.Write([]string{line})
			if err != nil {
				log.Fatalf("Error writing to CSV: %s", err)
			}
		}

	default:
		log.Fatalf("Unsupported format: %s", format)
	}

	return filename
}
