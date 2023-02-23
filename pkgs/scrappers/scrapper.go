package scrapper

import (
	"fmt"
	"net/http"

	"github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/errorHandler"
	"github.com/PuerkitoBio/goquery"
)

func Scrapper() {
	resp, err := http.Get("https://www.mukwano.com/")
	errorHandler.HanderError(err)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	errorHandler.HanderError(err)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link := s.Text()
		fmt.Println(link)
	})
}
