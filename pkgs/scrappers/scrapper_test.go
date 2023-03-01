package scrapper

import (
	"reflect"
	"testing"
)

func TestGoogleScrapper(t *testing.T) {
	got := GoogleScrapper("mukwano")
	want := "https://www.mukwano.com/"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected: %s got: %s", want, got)
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
		{"https://www.mukwano.com/", "mukwano", "customercare@mukwano.com"},
	}
	for _, tc := range tests {
		email, name := CompanyScrapper(tc.companyLink, tc.companyName)
		if email != tc.companyEmail || name != tc.companyName {
			t.Fatalf("Expected: Email: %v Name: %v, Got: Email: %v Name: %v", tc.companyEmail, tc.companyName, email, name)
		}
	}
}
