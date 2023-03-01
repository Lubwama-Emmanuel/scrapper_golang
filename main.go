package main

import (
	"fmt"
	"strings"

	scrapper "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/scrappers"
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
	fmt.Println("Here it is", collectionMap)
}
