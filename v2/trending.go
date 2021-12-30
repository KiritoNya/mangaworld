package v2

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

//Trending is a object that contains all the information of the manga in trending.
type Trending struct {
	Manga   *Manga
	ChapterNum int
	doc *goquery.Document
}

//NewTrendingManga is a construct of Trending manga object.
func NewTrendingManga(node *html.Node) (*Trending, error) {
	var t Trending
	t.doc = goquery.NewDocumentFromNode(node)

	return &t, nil
}

//GetManga is a function that adds the relative manga to the Trending object.
func (t *Trending) GetManga() error {
	//Get manga url from manga trending
	url, found := t.doc.Find("a").Attr("href")
	if !found {
		return errors.New("Manga not found")
	}

	//Create manga object
	manga, err := NewManga(url)
	if err != nil {
		return err
	}

	t.Manga = manga
	return nil
}

//GetChapterNum is a function that adds the relative chapter number to the Trending object.
func (t *Trending) GetChapterNum() error {
	chapNumString := t.doc.Find("div.chapter").Text()
	chapNumString = strings.Replace(chapNumString, "Capitolo ", "", -1)
	chapNum, err := strconv.Atoi(chapNumString)
	if err != nil {
		return err
	}

	t.ChapterNum = chapNum
	return nil
}