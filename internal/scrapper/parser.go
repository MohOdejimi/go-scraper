package scrapper

import (
	"github.com/PuerkitoBio/goquery"
)

 
func ParseHTML(doc *goquery.Document) ([]string, error) {
	var extractedData []string
 
	doc.Find("h1, h2, p").Each(func(i int, s *goquery.Selection) {
		extractedData = append(extractedData, s.Text())
	})

	return extractedData, nil
}
