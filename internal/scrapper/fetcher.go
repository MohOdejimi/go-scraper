package scrapper

import (
	"net/http" 
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func Fetch(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)

    if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("non-200 response code: %d", resp.StatusCode)
	}

    doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to parse the HTML document: %w", err)
	}

	return doc, nil
}