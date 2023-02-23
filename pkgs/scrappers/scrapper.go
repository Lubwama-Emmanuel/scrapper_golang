package scrapper

import (
	"fmt"
	"net/http"

	"github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/errorHandler"
	regexHandler "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/helpers"
	"github.com/PuerkitoBio/goquery"
)

func Scrapper() string {
	resp, err := http.Get("https://www.mukwano.com/")
	errorHandler.HanderError(err)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	errorHandler.HanderError(err)

	var links []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		links = append(links, link)
	})

	var email string
	for _, link := range links {
		answer := regexHandler.RegExp(link)
		if answer == "An empty string" {
			continue
		}
		email = answer
		fmt.Println(email)
	}

	return email
}

func GetEmails(link string) {
	resp, err := http.Get(link)
	errorHandler.HanderError(err)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	errorHandler.HanderError(err)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		email, _ := s.Attr("href")

		fmt.Println(email)
	})
}
