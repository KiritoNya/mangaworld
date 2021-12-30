package v2

import (
	"errors"
	"fmt"
	"github.com/KiritoNya/gotaku/manga"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

//Volume is a struct that contains all volume info
type Volume struct {
	manga.Volume
	Chapters []*Chapter
	doc *goquery.Selection
}

//NewVolume is a constructor for the Volume object
func NewVolume(volumeHtml *goquery.Selection) *Volume {
	return &Volume{doc: volumeHtml}
}

//GetNumber is a function that adds volume number to the object
func (v *Volume) GetNumber() error {
	volNumString := strings.Replace(v.doc.Find("p").Text(), "Volume ", "", -1)
	volNum, err := strconv.Atoi(volNumString)
	if err != nil {
		return err
	}

	v.Number= volNum
	return nil
}

//GetChapters is a function that add chapters to the object
func (v *Volume) GetChapters() error {
	var chapterErrors []string

	v.doc.Find(".chapter").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		chapterUrl, find := selection.Find("a").Attr("href")
		if !find {
			chapterErrors = append(chapterErrors, fmt.Sprintf("Chapter %d not found", i+1))
			return false
		}

		c, err := NewChapter(chapterUrl)
		if err != nil {
			chapterErrors = append(chapterErrors, err.Error())
			return false
		}

		v.Chapters = append(v.Chapters, c)
		return true
	})

	// Check errors
	if chapterErrors != nil {
		return errors.New(strings.Join(chapterErrors, "\n"))
	}

	return nil
}

