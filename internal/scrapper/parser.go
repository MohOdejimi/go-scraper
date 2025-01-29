package scrapper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseHTML(html string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var titles []string
	doc.Find("h2.title").Each(func(i int, s *goquery.Selection) {
		titles = append(titles, s.Text())
	})
	return titles, nil
}
 