package main

import (
	"strings"

	scrapper "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/scrappers"
	"github.com/sirupsen/logrus"
)

func main() {
	collectionMap := make(map[string]string)
	companies := scrapper.ReadFromFile()

	for _, company := range companies {
		if strings.Contains(company, " ") {
			company = strings.ReplaceAll(company, " ", "")
		}
		companyLink := scrapper.GoogleScrapper(company)
		email, name := scrapper.CompanyScrapper(companyLink, company)
		collectionMap[name] = email
	}

	logrus.Info(collectionMap)
}
