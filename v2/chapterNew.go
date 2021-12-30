package v2

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"strings"
)

//ChapterNew is an object of new chapters just released.
type ChapterNew struct {
	Manga *Manga
	Chapters []*Chapter
	doc *goquery.Document
}

//NewChapterNew is a construct of ChapterNew object.
func NewChapterNew(node *html.Node) (*ChapterNew, error) {
	var cn ChapterNew

	if node == nil {
		return nil , errors.New("node is empty")
	}

	doc := goquery.NewDocumentFromNode(node)
	cn.doc = doc
	return &cn, nil
}

//GetManga is a function that adds the object Manga to the object.
func (cn *ChapterNew) GetManga() error {
	//Get url of manga
	mangaUrl, found := cn.doc.Find("a").Attr("href")
	if !found {
		return errors.New("Manga of chapter new not found")
	}

	//Create manga object
	m, err := NewManga(mangaUrl)
	if err != nil {
		return err
	}

	cn.Manga = m
	return nil
}

//GetChapters is a function that adds the newest chapters objects to the ChapterNew object.
func (cn *ChapterNew) GetChapters() error {
	content := cn.doc.Find(".content").Find(".d-flex.flex-wrap.flex-row")
	alt, ok := content.Find("img").Attr("alt")
	if !ok {
		return errors.New("img alt for the newest manga not found")
	}

	if strings.ToLower(alt) == "nuovo" {
		//Get url
		url, found := content.Find("a").Attr("href")
		if !found {
			return errors.New("Chapter new url not found")
		}

		//Create chapter
		ch, err := NewChapter(url)
		if err != nil {
			return err
		}

		cn.Chapters = append(cn.Chapters, ch)
	}

	return nil
}