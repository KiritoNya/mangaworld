package v2

import (
	"errors"
	"github.com/KiritoNya/gotaku/image"
	"github.com/KiritoNya/gotaku/manga"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"os"
	"path/filepath"
	"strconv"
)

//Page is struct with all chapter page info
type Page struct {
	manga.Page
	doc *goquery.Document
}

//NewPage is a constructor for the page object.
func NewPage(pageUrl string) (*Page, error) {
	var p Page

	doc, err := doRequest(pageUrl)
	if err != nil {
		return nil, err
	}

	p.doc = doc
	p.Url = pageUrl
	return &p, nil
}

//NewPageFromNode is a constructor for the page object that use the html node.
func NewPageFromNode(node *html.Node) (*Page, error){
	var p Page
	p.doc = goquery.NewDocumentFromNode(node)

	return &p, nil
}

//NewPageFromFile is a constructor for the page object that use the file.
func NewPageFromFile(filePath string) (*Page, error){
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	contentHtml, err := html.Parse(f)
	if err != nil {
		return nil, err
	}

	p, err := NewPageFromNode(contentHtml)
	if err != nil {
		return nil, err
	}

	return p, nil
}

//GetNumber is a function that adds the number of page to the page object.
func (p *Page) GetNumber() error {
	if p.Url == "" {
		return errors.New("Page url not setted")
	}

	num, err := strconv.Atoi(filepath.Base(p.Url))
	if err != nil {
		return err
	}

	p.Number = num
	return nil
}

//GetImage is a function that adds the object image to the object page.
func (p *Page) GetImage() error {
	var i image.Image

	pageUrl, ok := p.doc.Find("#page").Find("img.img-fluid").Attr("src")
	if !ok {
		return errors.New("Page image not found")
	}
	i.Url = pageUrl

	p.Image = &i
	return nil
}