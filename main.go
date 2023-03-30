package main

import (
	"strings"

	scrapper "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/scrappers"
	log "github.com/sirupsen/logrus"
)

func main() {
	collectionMap := make(map[string]string)
	const googleSearchUrl = "https://www.google.com/search?q=%s"

	companies, err := scrapper.ReadFromFile("uploadedFiles/company_list-4096951222.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i := range companies {
		if strings.Contains(companies[i], " ") {
			companies[i] = strings.ReplaceAll(companies[i], " ", "")
		}

		companyLink, err := scrapper.GoogleScrapper(googleSearchUrl, companies[i])
		if err != nil {
			log.Error(err)
		}

		email, name, err := scrapper.ScrapeCompanyWebsite(companyLink, companies[i])
		if err != nil {
			log.Error(err)
		}
		collectionMap[name] = email
	}

	log.Info(collectionMap)
}
