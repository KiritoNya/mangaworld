package v2

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

//doRequest is a function that do the http request and create goquery document from response.
func doRequest(url string) (*goquery.Document, error) {
	//Do request
	req, err := ClientHttp.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	//Create html document
	return goquery.NewDocumentFromReader(req.Body)
}

//convertItalianMonth is a convert function for the month (EX: dicembre -> 12)
func convertItalianMonth(month string) int {
	monthLow := strings.ToLower(month)
	switch monthLow {
	case "gennaio":
		return 1
	case "febbraio":
		return 2
	case "marzo":
		return 3
	case "aprile":
		return 4
	case "maggio":
		return 5
	case "giugno":
		return 6
	case "luglio":
		return 7
	case "agosto":
		return 8
	case "settembre":
		return 9
	case "ottobre":
		return 10
	case "novembre":
		return 11
	case "dicembre":
		return 12
	}

	return 0
}

