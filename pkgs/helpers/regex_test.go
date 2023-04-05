package regexhandler_test

import (
	"errors"
	"fmt"
	"testing"

	regexhandler "github.com/Lubwama-Emmanuel/scrapper_golang/pkgs/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGetSubString(t *testing.T) {
	t.Parallel()
	type args struct {
		s     string
		index []int
	}
	tests := []struct {
		testName string
		args     args
		want     string
		wantErr  error
	}{
		{
			testName: "Test1",
			args: args{
				s:     "/url?q=http://codebits.io/&sa=U&ved=2ahU",
				index: []int{7, 25},
			},
			want:    "http://codebits.io",
			wantErr: nil,
		},
		{
			testName: "Test2",
			args: args{
				s:     "/url?q=http://codebits.io/&sa=U&ved=2ahU",
				index: []int{},
			},
			want:    "",
			wantErr: errors.New("no index returned"), //nolint:goerr113
		},
		{
			testName: "Test2",
			args: args{
				s:     "/url?q=http://codebits.io/&sa=U&ved=2ahU",
				index: []int{4, 6, 7},
			},
			want:    "",
			wantErr: errors.New("index is out of Range"), //nolint:goerr113
		},
	}

	for _, tc := range tests {
		tc := tc // created a local variable and assign the loop variable to it
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			got, err := regexhandler.GetSubString(tc.args.s, tc.args.index)
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

func TestMatchCompanyLink(t *testing.T) {
	t.Parallel()
	type args struct {
		link        string
		companyName string
	}
	tests := []struct {
		testName string
		args     args
		want     string
		wantErr  error
	}{
		{
			testName: "Test link with www",
			args: args{
				link:        "/url?q=https://www.mukwano.com/&sa=U&ved=2ah",
				companyName: "mukwano",
			},
			want:    "https://www.mukwano.com",
			wantErr: nil,
		},
		{
			testName: "Test link without www",
			args: args{
				link:        "/url?q=http://codebits.io/&sa=U&ved=2ahU",
				companyName: "codebits",
			},
			want:    "http://codebits.io",
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc // created a local variable and assign the loop variable to it
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			got, err := regexhandler.MatchCompanyLink(tc.args.link, tc.args.companyName)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Error not expected but got one:\n"+"error: %q", err))
				return
			}

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestMatchEmail(t *testing.T) {
	t.Parallel()
	type args struct {
		link string
	}
	tests := []struct {
		testName string
		args     args
		want     string
		wantErr  error
	}{
		{
			testName: "Test email with .com",
			args: args{
				link: "mailto: customercare@mukwano.com",
			},
			want:    "customercare@mukwano.com",
			wantErr: nil,
		},
		{
			testName: "Test email without .com",
			args: args{
				link: "mailto:hello@codebits.io",
			},
			want:    "hello@codebits.io",
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc // created a local variable and assign the loop variable to it
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			got, err := regexhandler.MatchEmail(tc.args.link)
			if err != nil && tc.wantErr == nil {
				assert.Fail(t, fmt.Sprintf("Error not expected but got one:\n"+"error: %q", err))
				return
			}

			assert.Equal(t, tc.want, got)
		})
	}
}
