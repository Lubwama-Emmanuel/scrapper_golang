package scrapper_test

import (
	"testing"

	scrapper "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/scrappers"
)

func TestGoogleScrapper(t *testing.T) {
	t.Parallel()
	type test struct {
		testName    string
		companyName string
		companyLink string
	}
	tests := []test{
		{"Url with www", "mukwano", "http://www.mukwano.com"},
		{"Url without www", "picfare", "https://picfare.com"},
	}

	for i := range tests {
		link := scrapper.GoogleScrapper(tests[i].companyName)

		if link != tests[i].companyLink {
			t.Fatalf("%v: Expected: %v But Got: %v", tests[i].testName, tests[i].companyLink, link)
		}
	}
}

func TestScrapeCompanyWebsite(t *testing.T) {
	t.Parallel()
	type test struct {
		testName     string
		companyLink  string
		companyName  string
		companyEmail string
	}
	tests := []test{
		{"Test email with .com", "https://www.mukwano.com/", "mukwano", "customercare@mukwano.com"},
		{"Test email without .com", "http://codebits.io", "codebits", "hello@codebits.io"},
	}

	for i := range tests {
		email, name := scrapper.ScrapeCompanyWebsite(tests[i].companyLink, tests[i].companyName)

		if email != tests[i].companyEmail || name != tests[i].companyName {
			t.Fatalf("%v: Expected: Email: %v Name: %v, Got: Email: %v Name: %v", tests[i].testName, tests[i].companyEmail, tests[i].companyName, email, name) //nolint
		}
	}
}
