package regexHandler

import (
	"fmt"
	"regexp"
)

func GetSubString(s string, index []int) string {
	if index == nil {
		return "empty"
	}
	return s[index[0]:index[1]]
}

func MatchCompanyLink(link, name string) string {
	url := fmt.Sprintf(`https:\/\/www\.%s\.(com|org)\/`, name)
	pattern := regexp.MustCompile(url)
	matchedIndex := pattern.FindStringIndex(link)
	return GetSubString(link, matchedIndex)
}

func MatchEmail(link string) string {
	pattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	matchedIndex := pattern.FindStringIndex(link)
	return GetSubString(link, matchedIndex)
}

func MatchContactUs(link string) string {
	pattern := regexp.MustCompile("https://www.creec.or.ug/contact-us-2/")
	matchedIndex := pattern.FindStringIndex(link)
	return GetSubString(link, matchedIndex)
}
