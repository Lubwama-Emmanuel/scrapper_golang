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

	for _, tc := range tests {
		tc := tc // created a local variable and assign the loop variable to it
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			got, err := scrapper.ReadFromFile(tc.args.fileName)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Error not expected but got one:\n"+"error: %q", err))
				return
			}
			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
				return
			}
			assert.Equal(t, tc.want, got)
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
		{
			testName: "500 status code",
			response: `<html><bodyhref="https://www.mukwano.com">Example</a></body></html>`,
			args: args{
				companyName: "mukwano",
			},
			want:    "",
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc // created a local variable and assign the loop variable to it
		// Run subtests
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			// Format the query parameter
			param := fmt.Sprintf("q=%v", tc.args.companyName)

			// Create test server for http mocking
			testServer := getTestServer(tc.response, param, tc.wantErr)
			defer testServer.Close()

			searchURL := testServer.URL + "/search?q=%s"

			got, err := scrapper.GoogleScrapper(searchURL, tc.args.companyName)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Error not expected but got one:\n"+"error: %q", err))
				return
			}

			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
				return
			}

			assert.Equal(t, tc.want, got)
		})
	}
}

func getTestServer(response, param string, err error) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if r.URL.Path == "/search" && r.URL.RawQuery == param {
			fmt.Fprintln(w, response)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
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

	for _, tc := range tests {
		tc := tc // created a local variable and assign the loop variable to it
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, tc.response)
			}))
			defer testServer.Close()
			got1, got2, err := scrapper.ScrapeCompanyWebsite(testServer.URL, tc.args.companyName)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Error not expected but got one:\n"+"error: %q", err))
				return
			}

			if tc.wantErr != nil {
				assert.EqualError(t, err, tc.wantErr.Error())
				return
			}

			assert.Equal(t, tc.want1, got1)
			assert.Equal(t, tc.want2, got2)
		})
	}
}
