package scrapper

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	errorhandler "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/errorHandler"
	regexhandler "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/helpers"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

// Read company names from file.
func ReadFromFile(fileName string) []string {
	var companies []string

	f, err := os.Open(fileName)
	if err != nil {
		log.Error("Failed to open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		companies = append(companies, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Error(err)
	}

	return companies
}

// Func takes in name of a company and makes a google search the returns link to company website.
func GoogleScrapper(name string) string {
	url := fmt.Sprintf("https://www.google.com/search?q=%s", name)

	resp, err := http.Get(url) //nolint
	errorhandler.HanderError("Error hitting url", err)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	errorhandler.HanderError("Error reading from goquery", err)

	var links []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		links = append(links, link)

	})

	var companyLink string

	for i := range links {
		answer := regexhandler.MatchCompanyLink(links[i], name)
		if answer == "empty" {
			continue
		}
		companyLink = answer
	}

	return companyLink
}

// Scraps the company website for their email.
func ScrapeCompanyWebsite(link, name string) (string, string) {
	resp, err := http.Get(link) //nolint
	errorhandler.HanderError("Error getting hitting company link", err)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	errorhandler.HanderError("Error retrieving File", err)

	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		links = append(links, link)
	})

	var email string

	for i := range links {
		answer := regexhandler.MatchEmail(links[i])
		if answer == "empty" {
			continue
		}
		email = answer
	}

	return email, name
}
