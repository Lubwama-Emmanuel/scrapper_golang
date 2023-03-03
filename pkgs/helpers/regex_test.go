package regexhandler_test

import (
	"testing"

	regexhandler "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/helpers"
)

func TestMatchCompanyLink(t *testing.T) {
	t.Parallel()
	type test struct {
		link        string
		companyLink string
		companyName string
	}
	tests := []test{
		{"/url?q=https://www.mukwano.com/&sa=U&ved=2ah", "https://www.mukwano.com", "mukwano"},
		{"/url?q=http://www.mukwano.com/&sa=U&ved=2ahUK", "http://www.mukwano.com", "mukwano"},
	}

	for i := range tests {
		got := regexhandler.MatchCompanyLink(tests[i].link, tests[i].companyName)

		if got != tests[i].companyLink {
			t.Fatalf("Expected: %v But Got: %v", tests[i].companyLink, got)
		}
	}
}

func TestMatchEmail(t *testing.T) {
	t.Parallel()
	type test struct {
		link string
		want string
	}
	tests := []test{
		{"customercare@mukwano.com", "customercare@mukwano.com"},
	}

	for i := range tests {
		got := regexhandler.MatchEmail(tests[i].link)

		if got != tests[i].want {
			t.Fatalf("Expected: %v But Got: %v", tests[i].want, got)
		}
	}
}
