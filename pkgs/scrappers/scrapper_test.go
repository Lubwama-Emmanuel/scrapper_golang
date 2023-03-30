//nolint:goerr113
package scrapper_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	scrapper "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/scrappers"
	"github.com/stretchr/testify/assert"
)

func TestReadFromFile(t *testing.T) {
	t.Parallel()
	type args struct {
		fileName string
	}
	tests := []struct {
		testName string
		args     args
		want     []string
		wantErr  error
	}{
		{
			testName: "Wrong filename",
			args: args{
				fileName: "uploadedFiles/company_list-409695122.txt",
			},
			want:    []string{},
			wantErr: errors.New("an error occurred trying to open the file open uploadedFiles/company_list-409695122.txt: no such file or directory"), //nolint:lll
		},
		{
			testName: "Correct filename",
			args: args{
				fileName: "../../uploadedFiles/company_list-4096951222.txt",
			},
			want:    []string{"codebits", "sunbird", "picfare", "mukwano"},
			wantErr: nil,
		},
	}

	for i := range tests {
		i := i // created a local variable and assign the loop variable to it
		t.Run(tests[i].testName, func(t *testing.T) {
			t.Parallel()
			got, err := scrapper.ReadFromFile(tests[i].args.fileName)
			if err != nil && tests[i].wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Error not expected but got one:\n"+"error: %q", err))
				return
			}
			if tests[i].wantErr != nil {
				assert.EqualError(t, err, tests[i].wantErr.Error())
				return
			}
			assert.Equal(t, tests[i].want, got)
		})
	}
}

func TestGoogleScrapper(t *testing.T) {
	t.Parallel()
	type args struct {
		companyName string
	}
	tests := []struct {
		testName string
		response string
		args     args
		want     string
		wantErr  error
	}{
		{
			testName: "Url without www",
			response: `<html><body><a href="https://picfare.com">Example</a></body></html>`,
			args: args{
				companyName: "picfare",
			},
			want:    "https://picfare.com",
			wantErr: nil,
		},
		{
			testName: "Url with www",
			response: `<html><body><a href="https://www.mukwano.com">Example</a></body></html>`,
			args: args{
				companyName: "mukwano",
			},
			want:    "https://www.mukwano.com",
			wantErr: nil,
		},
	}

	for i := range tests {
		i := i // created a local variable and assign the loop variable to it
		// Run subtests
		t.Run(tests[i].testName, func(t *testing.T) {
			t.Parallel()
			// Format the query parameter
			param := fmt.Sprintf("q=%v", tests[i].args.companyName)

			// Create test server for http mocking
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/search" && r.URL.RawQuery == param {
					fmt.Fprintln(w, tests[i].response)
				} else {
					w.WriteHeader(http.StatusNotFound)
				}
			}))
			defer testServer.Close()

			searchURL := testServer.URL + "/search?q=%s"

			got, err := scrapper.GoogleScrapper(searchURL, tests[i].args.companyName)
			if err != nil && tests[i].wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Error not expected but got one:\n"+"error: %q", err))
				return
			}

			if tests[i].wantErr != nil {
				assert.EqualError(t, err, tests[i].wantErr.Error())
				return
			}

			assert.Equal(t, tests[i].want, got)
		})
	}
}

func TestScrapeCompanyWebsite(t *testing.T) {
	t.Parallel()

	type args struct {
		companyLink string
		companyName string
	}

	tests := []struct {
		testName string
		response string
		args     args
		want1    string
		want2    string
		wantErr  error
	}{
		{
			testName: "Test email with .com",
			response: `<html><body><a href="customercare@mukwano.com">Example</a></body></html>`,
			args: args{
				companyLink: "https://www.mukwano.com/",
				companyName: "mukwano",
			},
			want1:   "customercare@mukwano.com",
			want2:   "mukwano",
			wantErr: nil,
		},
		{
			testName: "Test email without .com",
			response: `<html><body><a href="hello@codebits.io">Example</a></body></html>`,
			args: args{
				companyLink: "http://codebits.io",
				companyName: "codebits",
			},
			want1:   "hello@codebits.io",
			want2:   "codebits",
			wantErr: nil,
		},
	}

	for i := range tests {
		i := i // created a local variable and assign the loop variable to it
		t.Run(tests[i].testName, func(t *testing.T) {
			t.Parallel()
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, tests[i].response)
			}))
			defer testServer.Close()
			got1, got2, err := scrapper.ScrapeCompanyWebsite(testServer.URL, tests[i].args.companyName)
			if err != nil && tests[i].wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Error not expected but got one:\n"+"error: %q", err))
				return
			}

			if tests[i].wantErr != nil {
				assert.EqualError(t, err, tests[i].wantErr.Error())
				return
			}

			assert.Equal(t, tests[i].want1, got1)
			assert.Equal(t, tests[i].want2, got2)
		})
	}
}

// func TestGoogleScrapper(t *testing.T) {
// 	// Create a test server that returns a mock response for the search query
// 	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.Path == "/search" && r.URL.RawQuery == "q=test" {
// 			w.Write([]byte(`<html><body><a href="https://test.com">Example</a></body></html>`))
// 		} else {
// 			w.WriteHeader(http.StatusNotFound)
// 		}
// 	}))
// 	defer testServer.Close()

// 	// Set the test server URL for the GoogleScrapper function
// 	// googleURL := "https://www.google.com/search?q=%s"

// 	// oldURL := googleURL
// 	// googleURL = testServer.URL + "/search"
// 	// defer func() { googleURL = oldURL }()

// 	// Call GoogleScrapper with the test query
// 	companyLink, err := scrapper.GoogleScrapper("test")
// 	if err != nil {
// 		t.Errorf("GoogleScrapper returned an error: %v", err)
// 	}

// 	// Verify that the company link matches the expected value
// 	expectedLink := "https://test.com"
// 	if companyLink != expectedLink {
// 		t.Errorf("GoogleScrapper returned an unexpected company link: got %v, expected %v", companyLink, expectedLink)
// 	}

// 	// Restore the original Google URL
// 	// googleURL = oldURL
// }
