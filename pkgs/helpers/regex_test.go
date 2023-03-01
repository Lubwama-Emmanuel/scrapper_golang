package regexhandler

import "testing"

func TestMatchCompanyLink(t *testing.T) {
	type test struct {
		link        string
		companyLink string
		companyName string
	}
	tests := []test{
		{"/url?q=https://www.mukwano.com/&sa=U&ved=2ahUKEwiex4XXtbr9AhWZHLkGHZokBnAQFnoECAgQAg&usg=AOvVaw2tr-7pdNFoa_J9JVdxmLTX", "https://www.mukwano.com", "mukwano"},
		{"/url?q=http://www.mukwano.com/&sa=U&ved=2ahUKEwiex4XXtbr9AhWZHLkGHZokBnAQgU96BAgKEAQ&usg=AOvVaw29cym8cXd9e24lBYwu2G44", "http://www.mukwano.com", "mukwano"},
	}
	for i := range tests {
		got := MatchCompanyLink(tests[i].link, tests[i].companyName)

		if got != tests[i].companyLink {
			t.Fatalf("Expected: %v But Got: %v", tests[i].companyLink, got)
		}
	}
}

func TestMatchEmail(t *testing.T) {
	type test struct {
		link string
		want string
	}
	tests := []test{
		{"customercare@mukwano.com", "customercare@mukwano.com"},
	}
	for i := range tests {
		got := MatchEmail(tests[i].link)

		if got != tests[i].want {
			t.Fatalf("Expected: %v But Got: %v", tests[i].want, got)
		}
	}
}
