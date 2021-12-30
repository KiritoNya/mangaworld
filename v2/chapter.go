package v2

import (
	"errors"
	"fmt"
	"github.com/KiritoNya/gotaku/manga"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//Chapter is a struct with all chapter info
type Chapter struct {
	manga.Chapter
	Pages []*Page
	doc *goquery.Document
}

//NewChapter is a constructor of chapter object
func NewChapter(chapterUrl string) (*Chapter, error) {
	var c Chapter

	//Do request
	doc, err := doRequest(chapterUrl)
	if err != nil {
		return nil, err
	}

	c.Url = chapterUrl
	c.doc = doc

	return &c, nil
}

//NewChapterFromNode is a constructor of chapter object from node
func NewChapterFromNode(html *html.Node) (*Chapter, error) {
	var c Chapter

	if html == nil {
		return nil, errors.New("Html is empty")
	}

	doc := goquery.NewDocumentFromNode(html)
	c.doc = doc
	return &c, nil
}

//NewChapterFromFile is a constructor of chapter object from file
func NewChapterFromFile(path string) (*Chapter, error){
	//Open file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	//Parse html
	html, err := html.Parse(f)
	if err != nil {
		return nil, err
	}

	//Create chapter
	c, err := NewChapterFromNode(html)
	if err != nil {
		return nil, err
	}

	return c, nil
}

//GetId is a functiont that adds id to the chapter object
func (c *Chapter) GetId() error {
	if c.Url == "" {
		return errors.New("Url not setted")
	}

	c.Id = filepath.Base(c.Url)
	return nil
}

//GetNumber is a function that adds the number of chapter to the object.
func (c *Chapter) GetNumber() error {
	titleChap := c.doc.Find("head").Find("title").Text()
	fmt.Printf("\"%s\"", titleChap)

	if strings.Contains(titleChap, "Capitolo ") {
		matrix := strings.Split(titleChap, "Capitolo ")
		tmp := strings.Split(matrix[1], " ")

		//CASE ... Capitolo Extra ...
		if tmp[0] == "Extra" {
			num,  err := strconv.Atoi(tmp[1])
			if err != nil {
				return err
			}

			c.Number = num
			c.Type = tmp[0]
			return nil
		}

		//CASE ... Capitolo [0-9]+ ...
		num, err  := strconv.Atoi(tmp[0])
		if err != nil {
			return err
		}

		c.Number = num
		c.Type = "Standard"
		return nil
	}

	//CASE ... Capitolo ...
	if strings.Contains(titleChap, "Capitolo ") {

		matrix := strings.Split(titleChap, "Capitolo ")
		tmp := strings.Split(matrix[1], " ")

		//CASE ... Capitolo Extra ...
		if tmp[0] == "Extra" {
			num, err := strconv.Atoi(tmp[1])
			if err != nil {
				return err
			}

			c.Number = num
			c.Type = tmp[0]
			return nil
		}

		//CASE ... Capitolo [0-9]+ ...
		num, err  := strconv.Atoi(tmp[0])
		if err != nil {
			return err
		}

		c.Number = num
		c.Type = "Extra"
		return nil
	}

	if strings.Contains(string(titleChap), "CApitolo "){
		matrix := strings.Split(string(titleChap), "CApitolo ")
		tmp := strings.Split(matrix[1], " ")

		//CASE ... Capitolo Extra ...
		if tmp[0] == "Extra" {
			num, err := strconv.Atoi(tmp[1])
			if err != nil {
				return err
			}

			c.Number = num
			c.Type = tmp[0]
			return nil
		}

		//CASE ... Capitolo [0-9]+ ...
		num, err  := strconv.Atoi(tmp[0])
		if err != nil {
			return err
		}

		c.Number = num
		c.Type = "Standard"
		return nil
	}

	//CASE ... Oneshot ...
	if strings.Contains(titleChap, "Oneshot") {
		c.Number = 1
		c.Type = "Oneshot"
		return nil
	}

	return errors.New("Manga number not found")
}

//GetPages is a function that adds the pages to the chapter object.
func (c *Chapter) GetPages() error {
	numPages := c.GetPageNum()
	if numPages == -1 {
		return errors.New("Number of pages not found")
	}

	for i := 1; i <= numPages; i++ {
		p, err := NewPage(c.Url + "/" + strconv.Itoa(i))
		if err != nil {
			return err
		}

		c.Pages = append(c.Pages, p)
	}

	return nil
}

//GetReleaseDate is a function that adds the release chapter date to the chapter object.
func (c *Chapter) GetReleaseDate() error {
	dateString := c.doc.Find(".col-12.col-md-4.d-flex.justify-content-start.justify-content-md-center.align-items-center").Text()
	dateString = strings.Split(dateString, "aggiunta:\u00a0")[1]
	fmt.Println(dateString)
	dateStringParsed := strings.Split(dateString, " ")
	day, err := strconv.Atoi(dateStringParsed[0])
	if err != nil {
		return err
	}
	year, err := strconv.Atoi(dateStringParsed[2])
	if err != nil {
		return err
	}
	c.ReleaseDate = time.Date(year, time.Month(convertItalianMonth(dateStringParsed[1])), day, 0, 0,0, 0, time.Local)
	return nil
}

//GetKeywords is a function that returns the chapter search keywords.
func (c *Chapter) GetKeywords() ([]string, error) {
	keysContainer := c.doc.Find(".has-shadow.top-wrapper.p-3.mt-4.mb-3").Last().Find("h2").Text()
	if keysContainer == "" {
		return nil, errors.New("Search keys not found")
	}

	keys := strings.Split(keysContainer, " - ")
	return keys,nil
}

//GetPageNum is a function that returns the number of pages of chapter. Return -1 if errored.
func (c *Chapter) GetPageNum() int {
	str := c.doc.Find("select.page.custom-select").Find("option").First().Text()
	str = strings.Split(str, "/")[1]
	numPage, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}

	return numPage
}

//GetVisual is a function that returns the visual of chapter. Return -1 if errored.
func (c *Chapter) GetVisual() int {
	stripped := c.doc.Find(".col-12.col-md-4.d-flex.justify-content-start.align-items-center").Text()
	fmt.Printf("\"%s\"\n", stripped)
	stripped = strings.Replace(stripped, "\"", "", -1)
	stripped = strings.Split(stripped, "Data")[0]
	stripped = strings.Replace(stripped, "Visualizzazioni:\u00a0", "", -1)
	visual, err := strconv.Atoi(stripped)
	if err != nil {
		return -1
	}

	return visual
}

//GetVisualToday is a function that returns the today visual of chapter. Return -1 if errored.
func (c *Chapter) GetVisualToday() int {
	stripped := c.doc.Find(".col-12.col-md-4.d-flex.justify-content-start.align-items-center").Text()
	fmt.Printf("\"%s\"\n", stripped)
	stripped = strings.Replace(stripped, "\"", "", -1)
	stripped = strings.Split(stripped, "oggi:\u00a0")[1]
	visual, err := strconv.Atoi(stripped)
	if err != nil {
		return -1
	}

	return visual
}