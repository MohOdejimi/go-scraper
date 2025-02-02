package scrapper

import (
	"html"
	"regexp"
	"strings"
)

func ProcessData(data []string) []string {
	var processedData []string
	seen := make(map[string]bool)

	re := regexp.MustCompile(`[^\w\s,.!?]`)

	for _, text := range data {
		text = html.UnescapeString(text)
		text = strings.TrimSpace(text)
		text = strings.ToLower(text)
		text = re.ReplaceAllString(text, "")

		if _, exists := seen[text]; !exists && text != "" {
			processedData = append(processedData, text)
			seen[text] = true
		}
	}

	return processedData
}