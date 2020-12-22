package mangaworld

import (
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
)

//ChapterNew is an object of new chapters just released.
type ChapterNew struct {
	MangaNew Manga
	Chapters []Chapter
	node     *html.Node
}

//NewChapterNew is a construct of ChapterNew object.
func NewChapterNew(n *html.Node) (*ChapterNew, error) {

	_, err := htmlutils.QuerySelector(n, "a", "class", "thumb position-relative")
	if err != nil {
		return nil, err
	}

	_, err = htmlutils.QuerySelector(n, "div", "class", "content")
	if err != nil {
		return nil, err
	}

	return &ChapterNew{node: n}, nil
}

//Add object Manga (only url field value) to the object.
func (cn *ChapterNew) GetManga() error {

	tagsA, err := htmlutils.GetGeneralTags(cn.node, "a")
	if err != nil {
		return err
	}

	url, err := htmlutils.GetValueAttr(tagsA[0], "a", "href")
	if err != nil {
		return err
	}

	cn.MangaNew.Url = string(url[0])

	return nil
}

//Add object Chapter (only url field value) to the object.
func (cn *ChapterNew) GetChapter() error {

	divs, err := htmlutils.QuerySelector(cn.node, "div", "class", "content")
	if err != nil {
		return err
	}

	divsCh, err := htmlutils.QuerySelector(divs[0], "div", "class", "d-flex flex-wrap flex-row")
	if err != nil {
		return err
	}

	if _, err := htmlutils.QuerySelector(divsCh[0], "img", "alt", "nuovo"); err == nil {
		tagsA, err := htmlutils.GetGeneralTags(divsCh[0], "a")
		if err != nil {
			return err
		}

		url, err := htmlutils.GetValueAttr(tagsA[0], "a", "href")
		if err != nil {
			return err
		}

		cn.Chapters = append(cn.Chapters, Chapter{Url: string(url[0])})
	}
	return nil
}
