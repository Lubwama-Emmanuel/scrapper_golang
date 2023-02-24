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
	fmt.Println("Uploading file from here")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	errorHandler.HanderError("Error retriving File", err)

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func routes() {
	http.HandleFunc("/uploadFile", uploadFile)
	http.ListenAndServe(":8000", nil)
}
func main() {
	companyName = "mukwano"
	routes()
	companyLink, compName := scrapper.GoogleScrapper(companyName)
	scrapper.CompanyScrapper(companyLink, compName)
	// scrapper.ContactUsScrapper(contactLink)
}
