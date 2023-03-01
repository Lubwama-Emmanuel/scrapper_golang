package scrapper

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/errorHandler"
	regexHandler "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/helpers"
	"github.com/PuerkitoBio/goquery"
)

func ReadFromFile() []string {
	// content, err := os.ReadFile("uploadedFiles/company_list-4096951222.txt")
	// errorHandler.HanderError("Error reading file", err)

	// fmt.Println(string(content))
	var companies []string
	f, err := os.Open("uploadedFiles/company_list-4096951222.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		companies = append(companies, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(companies)
	return companies
}

// Func takes in name of a company and makes a google search the returns link to company website
func GoogleScrapper(name string) string {
	url := fmt.Sprintf("https://www.google.com/search?q=%s", name)

	resp, err := http.Get(url)
	errorHandler.HanderError("Error hitting url", err)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	errorHandler.HanderError("Error reading from goquery", err)

	var links []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		links = append(links, link)
	})

	var companyLink string
	for _, link := range links {
		answer := regexHandler.MatchCompanyLink(link, name)
		if answer == "empty" {
			continue
		}
		companyLink = answer
	}
	return companyLink
}

// Scraps the company website for their email or contact-us page
func CompanyScrapper(link, name string) (string, string) {
	fmt.Println("gotten link", link)
	resp, err := http.Get(link)
	errorHandler.HanderError("Error getting hitting company link", err)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	errorHandler.HanderError("Error retriving File", err)

	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		links = append(links, link)
		// fmt.Println(link)
	})

	var email string
	for _, link := range links {
		answer := regexHandler.MatchEmail(link)
		if answer == "empty" {
			continue
		}
		email = answer
	}
	fmt.Println("Here is the email", email)
	return email, name
}

func ContactUsScrapper(link string) {
	fmt.Println("Here is the contactus", link)
	resp, err := http.Get(link)
	errorHandler.HanderError("Error retriving File", err)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	errorHandler.HanderError("Error retriving File", err)

	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		links = append(links, link)
		fmt.Println(link)
	})

	var email string
	for _, link := range links {
		answer := regexHandler.MatchEmail(link)

		if answer != "nil" {
			continue
		}
		email = answer
		fmt.Println("here is the email", email)
	}
}
