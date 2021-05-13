package animeworld

import (
	"errors"
	"fmt"
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

type Archive struct {
	Seasons []*Season
	TotalPages int
	url string
	htmlPage *html.Node
}

func NewArchive(method string) (*Archive, error) {

	var a Archive
	var err error

	a.url = composeURL(method)

	resp, err := doRequest(a.url)
	if err != nil {
		return nil, err
	}

	a.htmlPage, err = html.Parse(strings.NewReader(resp))
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *Archive) GetSeason() (err error) {

	if a.htmlPage == nil {
		resp, err := doRequest(composeURL("all"))
		if err != nil {
			return err
		}

		a.htmlPage, err = html.Parse(strings.NewReader(resp))
		if err != nil {
			return err
		}
	}

	err = a.GetTotalPages()
	if err != nil {
		a.TotalPages=1
	}

	for i:=1; i<=a.TotalPages ; i++ {

		htmlPage := a.htmlPage

		if i != 1 {

			link := a.url + "?page=" + strconv.Itoa(i)

			resp, err := doRequest(link)
			if err != nil {
				return err
			}

			htmlPage, err = html.Parse(strings.NewReader(resp))
			if err != nil {
				return err
			}

		}

		items, err := htmlutils.QuerySelector(htmlPage, "div", "class", "items")
		if err != nil {
			return errors.New("Archive error: " + err.Error())
		}

		//Get all seasons html sections
		seasonsNode, err := htmlutils.QuerySelector(items[0], "div", "class", "item")
		if err != nil {
			return errors.New("Error to obtain seasons section in the archive")
		}

		for i, seasonNode := range seasonsNode {

			if i == 0 {
				continue
			}

			tagA, err := htmlutils.QuerySelector(seasonNode, "a", "class", "name")
			if err != nil {
				return errors.New("Error to obtain tag \"a\" of seasons section")
			}

			href, err := htmlutils.GetValueAttr(tagA[0], "a", "href")
			if err != nil {
				return errors.New("Error to obtain link of seasons section")
			}

			fmt.Println(BaseUrl + string(href[0]))

			season, err := NewSeason( BaseUrl + string(href[0]))
			if err != nil {
				return errors.New("Error to create season object")
			}

			a.Seasons = append(a.Seasons, season)
		}

	}

	return nil
}

func (a *Archive) GetTotalPages() error {
	div, err := htmlutils.QuerySelector(a.htmlPage, "div", "class", "paging-wrapper")
	if err != nil {
		return errors.New("Error to get paging wrapper for calculate max pages")
	}

	totalNum, err := htmlutils.QuerySelector(div[0], "span", "class", "total")
	if err != nil {
		return errors.New("Error to get total number of page.")
	}

	numTotal := htmlutils.GetNodeText(totalNum[0], "span")

	a.TotalPages, err = strconv.Atoi(string(numTotal))
	if err != nil {
		return err
	}

	return nil
}

func composeURL(method string) string {
	var link string

	if method == "all" {
		link = ArchiveUrl
	} else {
		link= ArchiveUrl + method
	}

	return link
}

