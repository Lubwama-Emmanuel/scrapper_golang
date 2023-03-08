package main

import (
	"strings"

	scrapper "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/scrappers"
	"github.com/sirupsen/logrus"
)

func main() {
	collectionMap := make(map[string]string)
	companies := scrapper.ReadFromFile("uploadedFiles/company_list-4096951222.txt")

	for i := range companies {
		if strings.Contains(companies[i], " ") {
			companies[i] = strings.ReplaceAll(companies[i], " ", "")
		}
		companyLink := scrapper.GoogleScrapper(companies[i])
		email, name := scrapper.ScrapeCompanyWebsite(companyLink, companies[i])
		collectionMap[name] = email
	}

	logrus.Info(collectionMap)
}
