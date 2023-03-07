package regexhandler_test

import (
	"testing"

	regexhandler "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/helpers"
)

func TestMatchCompanyLink(t *testing.T) {
	t.Parallel()
	type test struct {
		testName    string
		link        string
		companyLink string
		companyName string
	}
	tests := []test{
		{"Test link with www", "/url?q=https://www.mukwano.com/&sa=U&ved=2ah", "https://www.mukwano.com", "mukwano"},
		{"Test link without www", "/url?q=http://codebits.io/&sa=U&ved=2ahU", "http://codebits.io", "codebits"},
	}

	for i := range tests {
		got := regexhandler.MatchCompanyLink(tests[i].link, tests[i].companyName)

		if got != tests[i].companyLink {
			t.Fatalf("%v: Expected: %v But Got: %v", tests[i].testName, tests[i].companyLink, got)
		}
	}
}

func TestMatchEmail(t *testing.T) {
	t.Parallel()
	type test struct {
		testName string
		link     string
		want     string
	}
	tests := []test{
		{"Test email with .com", "customercare@mukwano.com", "customercare@mukwano.com"},
		{"Test email without .com", "mailto:hello@codebits.io", "hello@codebits.io"},
	}

	for i := range tests {
		got := regexhandler.MatchEmail(tests[i].link)

		if got != tests[i].want {
			t.Fatalf("%v: Expected: %v But Got: %v", tests[i].testName, tests[i].want, got)
		}
	}
}
