package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/errorHandler"
	scrapper "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/scrappers"
)

var companyName string

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading a file endpoint")

	// 1. parse input, type multipart/form-data
	r.ParseMultipartForm(10 << 20)

	// 2. retrieve file from posted form-data
	file, handler, err := r.FormFile("myFile")
	errorHandler.HanderError("Error retrieving file form form-data", err)

	defer file.Close()

	fmt.Printf("Uploaded file: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// 3. Write temporary file on our server
	tempFile, err := ioutil.TempFile("uploadedFiles", "company_list-*.txt")
	errorHandler.HanderError("Error loading tempFile", err)

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	errorHandler.HanderError("Error reading fileBytes", err)

	tempFile.Write(fileBytes)
	// 4. Return whether or not this has been successfull
	fmt.Fprintf(w, "Successfully uploaded the file")
}

func routes() {
	http.HandleFunc("/uploadFile", uploadFile)
	http.ListenAndServe(":8000", nil)
}
func main() {
	collectionMap := make(map[string]string)
	// companyName = "mukwano"
	// companyLink, compName := scrapper.GoogleScrapper(companyName)
	// scrapper.CompanyScrapper(companyLink, compName)
	// scrapper.ContactUsScrapper(contactLink)

	// emails := []string{
	// 	"mukwano@email",
	// 	"creec@email",
	// }

	// companies := []string{
	// 	"mukwano",
	// 	"creec",
	// }

	// for _, company := range companies {
	// 	for _, email := range emails {
	// 		collectionMap[company] = email
	// 	}
	// }

	companies := scrapper.ReadFromFile()
	for _, company := range companies {
		companyLink := scrapper.GoogleScrapper(company)
		email, name := scrapper.CompanyScrapper(companyLink, company)
		collectionMap[name] = email
	}
	// fmt.Println("Here it is", collectionMap)
}
