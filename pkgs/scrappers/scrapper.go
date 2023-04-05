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
		return nil, fmt.Errorf("an error occurred trying to open the file %w", fileErr)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		companies = append(companies, scanner.Text())
	}

	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, fmt.Errorf("an error occurred during scanning %w", scannerErr)
	}

	return companies, nil
}

// Func takes in name of a company and makes a google search the returns link to company website.
func GoogleScrapper(uri, name string) (string, error) {
	url := fmt.Sprintf(uri, name)

	resp, httpErr := http.Get(url) //nolint
	if httpErr != nil {
		return "", fmt.Errorf("an error occurred trying to scrapper google for %s %w", name, httpErr)
	}

	// TODO: add check for status code

	defer resp.Body.Close()

	doc, queryErr := goquery.NewDocumentFromReader(resp.Body)
	if queryErr != nil {
		return "", fmt.Errorf("an error occurred loading goquery %w", queryErr)
	}

	var links []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		links = append(links, link)
	})

	var companyLink string

	for i := range links {
		answer, _ := regexhandler.MatchCompanyLink(links[i], name)
		if answer != "" {
			companyLink = answer
		}
	}

	return companyLink, nil
}

// Scraps the company website for their email.
func ScrapeCompanyWebsite(link, name string) (string, string, error) {
	resp, httpErr := http.Get(link) //nolint
	if httpErr != nil {
		return "", "nil", fmt.Errorf("an error occurred trying to scrapper company: %v website %w", name, httpErr)
	}

	defer resp.Body.Close()

	doc, queryErr := goquery.NewDocumentFromReader(resp.Body)
	// codecov: ignore-start
	if queryErr != nil {
		return "", "nil", fmt.Errorf("an error occurred loading goquery %w", queryErr)
	}
	// codecov: ignore-end

	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		links = append(links, link)
	})

	var email string

	for i := range links {
		answer, _ := regexhandler.MatchEmail(links[i])

		if answer != "" {
			email = answer
		}
	}

	return email, name, nil
}
