package scrapper_test

import (
	"testing"

	scrapper "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/scrappers"
)

func TestGoogleScrapper(t *testing.T) {
	t.Parallel()
	type test struct {
		companyName string
		companyLink string
	}
	tests := []test{
		{"mukwano", "http://www.mukwano.com"},
	}

	for i := range tests {
		link := scrapper.GoogleScrapper(tests[i].companyName)

		if link != tests[i].companyLink {
			t.Fatalf("Expected: %v But Got: %v", tests[i].companyLink, link)
		}
	}
}

func TestCompanyScrapper(t *testing.T) {
	t.Parallel()
	type test struct {
		companyLink  string
		companyName  string
		companyEmail string
	}
	tests := []test{
		{"https://www.mukwano.com/", "mukwano", "customercare@mukwano.com"},
		{"https://www.mukwano.com/", "mukwano", "customercare@mukwano.com"},
	}

	for _, tc := range tests {
		email, name := scrapper.CompanyScrapper(tc.companyLink, tc.companyName)

		if email != tc.companyEmail || name != tc.companyName {
			t.Fatalf("Expected: Email: %v Name: %v, Got: Email: %v Name: %v", tc.companyEmail, tc.companyName, email, name)
		}
	}
}
