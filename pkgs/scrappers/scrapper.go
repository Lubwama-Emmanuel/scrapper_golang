package scrapper

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	regexhandler "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/helpers"
	"github.com/PuerkitoBio/goquery"
)

// Read company names from file.
func ReadFromFile(fileName string) ([]string, error) {
	var companies []string

	f, fileErr := os.Open(fileName)
	if fileErr != nil {
		err := fmt.Errorf("an error occurred trying to open the file %w", fileErr)
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		companies = append(companies, scanner.Text())
	}

	if scannerErr := scanner.Err(); scannerErr != nil {
		err := fmt.Errorf("an error occurred during scanning %w", scannerErr)
		return nil, err
	}

	return companies, nil
}

// Func takes in name of a company and makes a google search the returns link to company website.
func GoogleScrapper(name string) (string, error) {
	url := fmt.Sprintf("https://www.google.com/search?q=%s", name)

	resp, httpErr := http.Get(url) //nolint
	if httpErr != nil {
		err := fmt.Errorf("an error occurred trying to scrapper google for %s %w", name, httpErr)
		return "", err
	}

	defer resp.Body.Close()

	doc, queryErr := goquery.NewDocumentFromReader(resp.Body)
	if queryErr != nil {
		err := fmt.Errorf("an error occurred loading goquery %w", queryErr)
		return "", err
	}

	var links []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		links = append(links, link)
	})

	var companyLink string

	for i := range links {
		answer, _ := regexhandler.MatchCompanyLink(links[i], name)
		if answer == "" {
			continue
		}
		companyLink = answer
	}

	return companyLink, nil
}

// Scraps the company website for their email.
func ScrapeCompanyWebsite(link, name string) (string, string, error) {
	resp, httpErr := http.Get(link) //nolint
	if httpErr != nil {
		err := fmt.Errorf("an error occurred trying to scrapper company: %v website %w", name, httpErr)
		return "", "nil", err
	}

	defer resp.Body.Close()

	doc, queryErr := goquery.NewDocumentFromReader(resp.Body)
	if queryErr != nil {
		err := fmt.Errorf("an error occurred loading goquery %w", queryErr)
		return "", "nil", err
	}

	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		links = append(links, link)
	})

	var email string

	for i := range links {
		answer, _ := regexhandler.MatchEmail(links[i])

		if answer == "" {
			continue
		}
		email = answer
	}

	return email, name, nil
}
