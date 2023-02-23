package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("AN ERROR OCCURED", err)
	}
}

func getSubString(s string, index []int) string {
	if index == nil {
		return ("An empty string")
	}
	return s[index[0]:index[1]]
}

func regExp(link string) string {
	pattern := regexp.MustCompile(`https://www.facebook.com/[a-zA-Z]*/`)
	matchedIndex := pattern.FindStringIndex(link)
	fmt.Println(matchedIndex)
	return getSubString(link, matchedIndex)
}

func main() {
	resp, err := http.Get("https://www.google.com/search?q=makerere+university&sxsrf=AJOqlzXjBg4Obq2pWl8rt4NEKPnj7MoX-g%3A1677060062375&ei=3uf1Y4jKFsDLkdUP_ciM8AU&oq=ma&gs_lcp=Cgxnd3Mtd2l6LXNlcnAQARgAMgQIIxAnMgQIIxAnMgQIIxAnMgQILhBDMgQILhBDMgQILhBDMgQILhBDMgQILhBDMgQIABBDMgQILhBDOgoIABBHENYEELADOgoILhDUAhCwAxBDOgcILhCwAxBDOgcIABCwAxBDOg0IABDkAhDWBBCwAxgBOgwILhDIAxCwAxBDGAI6EgguEK8BEMcBEMgDELADEEMYAjoSCC4QxwEQrwEQyAMQsAMQQxgCOg0ILhCABBAUEIcCENQCOg0IABCABBAUEIcCELEDOgUIABCABDoKCC4QxwEQrwEQQzoHCCMQ6gIQJzoPCC4Q1AIQ6gIQtAIQQxgDOgwILhDqAhC0AhBDGAM6EgguEK8BEMcBEOoCELQCEEMYAzoMCAAQ6gIQtAIQQxgDOgcILhDUAhBDSgQIQRgAUPgFWIE4YLpLaAJwAXgEgAGbAogBnw2SAQMyLTeYAQCgAQGwARTIARLAAQHaAQYIARABGAnaAQYIAhABGAjaAQYIAxABGAE&sclient=gws-wiz-serp")
	checkError(err)

	defer resp.Body.Close()
	url := "/url?q=https://www.facebook.com/mukwanogroup&sa=U&ved=2ahUKEwiRlbri86j9AhW9BbkGHfkCD28QFnoECAUQAg&usg=AOvVaw2SWzjgJfbHwyxcvUYT_XoR"

	answer := regExp(url)
	fmt.Println("Here is the result:", answer)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("AN ERROR OCCURED", err)
	}

	// // Find all the article titles and print them
	count := doc.Find(".GyAeWb").Size()
	fmt.Println(count)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Attr("href")

		fmt.Println("Here is the result:", title)
		// fmt.Printf("A: %d %s\n", i, title)
	})

}
