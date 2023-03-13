//nolint:goerr113
package scrapper_test

import (
	"errors"
	"fmt"
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
			wantErr: errors.New("an error occurred trying to open the file open uploadedFiles/company_list-409695122.txt: no such file or directory"),
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
		args     args
		want     string
		wantErr  error
	}{
		// {
		// 	testName: "Url with www",
		// 	args: args{
		// 		companyName: "mukwano",
		// 	},
		// 	want:    "https://www.mukwano.com",
		// 	wantErr: nil,
		// },
		{
			testName: "Url without www",
			args: args{
				companyName: "picfare",
			},
			want:    "https://picfare.com",
			wantErr: nil,
		},
	}

	for i := range tests {
		i := i // created a local variable and assign the loop variable to it
		t.Run(tests[i].testName, func(t *testing.T) {
			t.Parallel()
			got, err := scrapper.GoogleScrapper(tests[i].args.companyName)
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
		args     args
		want1    string
		want2    string
		wantErr  error
	}{
		{
			testName: "Test email with .com",
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
			got1, got2, err := scrapper.ScrapeCompanyWebsite(tests[i].args.companyLink, tests[i].args.companyName)
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
