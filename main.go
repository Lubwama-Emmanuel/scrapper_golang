package main

import (
	"strings"

	errorhandler "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/errorHandler"
	scrapper "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/scrappers"
	log "github.com/sirupsen/logrus"
)

func main() {
	collectionMap := make(map[string]string)
	companies, err := scrapper.ReadFromFile("uploadedFiles/company_list-4096951222.txt")
	errorhandler.HanderError(err)

	for i := range companies {
		if strings.Contains(companies[i], " ") {
			companies[i] = strings.ReplaceAll(companies[i], " ", "")
		}
		companyLink, err := scrapper.GoogleScrapper(companies[i])
		errorhandler.HanderError(err)
		email, name, err := scrapper.ScrapeCompanyWebsite(companyLink, companies[i])
		errorhandler.HanderError(err)
		collectionMap[name] = email
	}

	log.Info(collectionMap)
}
