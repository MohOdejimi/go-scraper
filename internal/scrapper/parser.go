package scrapper

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

 
func ParseHTML(doc *goquery.Document, selectors string) ([]string, error) {
	var extractedData []string

	selectorList := strings.Split(selectors, ",")
 
	doc.Find(strings.Join(selectorList, ", ")).Each(func(i int, s *goquery.Selection) {
		extractedData = append(extractedData, s.Text())

		s.Find("img").Each(func(j int, img *goquery.Selection) {
            src, exists := img.Attr("src")
            if exists {
                extractedData = append(extractedData, "Image: "+src)
            }
        })
	})

	return ProcessData(extractedData), nil
}
