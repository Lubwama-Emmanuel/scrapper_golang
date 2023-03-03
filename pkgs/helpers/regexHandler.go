package regexhandler

import (
	"fmt"
	"regexp"
)

// Matches received index and returns the string.
func GetSubString(s string, index []int) string {
	if index == nil {
		return "empty"
	}

	return s[index[0]:index[1]]
}

// Matches valid company website link and returns the Link.
func MatchCompanyLink(link, name string) string {
	url := fmt.Sprintf(`(https?:\/\/)(www\.)?(%v\.)+[a-z]{2,}`, name)
	pattern := regexp.MustCompile(url)
	matchedIndex := pattern.FindStringIndex(link)

	return GetSubString(link, matchedIndex)
}

// Matches valid email address.
func MatchEmail(link string) string {
	pattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	matchedIndex := pattern.FindStringIndex(link)

	return GetSubString(link, matchedIndex)
}
