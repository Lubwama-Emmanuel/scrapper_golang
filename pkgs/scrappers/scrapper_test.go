package scrapper

import (
	"testing"
)

func TestGoogleScrapper(t *testing.T) {
	type test struct {
		companyName string
		companyLink string
	}
	tests := []test{
		{"mukwano", "https://www.mukwano.com"},
		{"kanzucode", "https://kanzucode.com"},
	}
	for _, tc := range tests {
		link := GoogleScrapper(tc.companyName)
		if link != tc.companyLink {
			t.Fatalf("Expected: %v, Got: %v", tc.companyLink, link)
		}
	}
}

func TestCompanyScrapper(t *testing.T) {
	type test struct {
		companyLink  string
		companyName  string
		companyEmail string
	}
	tests := []test{
		{"https://www.mukwano.com/", "mukwano", "customercare@mukwano.com"},
	}
	for _, tc := range tests {
		email, name := CompanyScrapper(tc.companyLink, tc.companyName)
		if email != tc.companyEmail || name != tc.companyName {
			t.Fatalf("Expected: Email: %v Name: %v, Got: Email: %v Name: %v", tc.companyEmail, tc.companyName, email, name)
		}
	}

}
