package regexHandler

import "regexp"

func GetSubString(s string, index []int) string {
	if index == nil {
		return ("An empty string")
	}
	return s[index[0]:index[1]]
}

func RegExp(link string) string {
	pattern := regexp.MustCompile(`[a-z]*@[a-z]*\.(com|org)`)
	matchedIndex := pattern.FindStringIndex(link)
	return GetSubString(link, matchedIndex)
}
