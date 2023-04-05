package regexhandler

import (
	"errors"
	"fmt"
	"regexp"
)

// Matches received index and returns the string.
func GetSubString(s string, index []int) (string, error) {
	if len(index) == 0 {
		err := errors.New("no index returned") //nolint:goerr113
		return "", err
	} else if len(index) > 2 {
		err := errors.New("index is out of Range") //nolint:goerr113
		return "", err
	}

	return s[index[0]:index[1]], nil
}

// Matches valid company website link and returns the Link.
func MatchCompanyLink(link, name string) (string, error) {
	url := fmt.Sprintf(`(https?:\/\/)(www\.)?(%v\.)+[a-z]{2,}`, name)
	pattern := regexp.MustCompile(url)
	matchedIndex := pattern.FindStringIndex(link)

	return GetSubString(link, matchedIndex)
}

// Matches valid email address.
func MatchEmail(link string) (string, error) {
	pattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	matchedIndex := pattern.FindStringIndex(link)

	return GetSubString(link, matchedIndex)
}
