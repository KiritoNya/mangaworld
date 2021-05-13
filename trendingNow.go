package mangaworld

import (
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
	"strings"
)

//Trending is a object that contains all the information of the manga in trending.
type Trending struct {
	Manga   Manga
	Chapter Chapter
	node    *html.Node
}

//NewTrendingManga is a construct of Trending manga object.
func NewTrendingManga(n *html.Node) (*Trending, error) {

	_, err := htmlutils.QuerySelector(n, "a", "class", "thumb position-relative")
	if err != nil {
		return nil, err
	}

	return &Trending{node: n}, nil
}

//Add object Manga (only url field value) to the object.
func (t *Trending) GetManga() error {
	tagsA, err := htmlutils.GetGeneralTags(t.node, "a")
	if err != nil {
		return err
	}

	url, err := htmlutils.GetValueAttr(tagsA[0], "a", "href")
	if err != nil {
		return err
	}

	t.Manga.Url = string(url[0])
	return nil
}

//Add object Chapter (only url field value) to the object.
func (t *Trending) GetChapter() error {
	divs, err := htmlutils.QuerySelector(t.node, "div", "class", "chapter")
	if err != nil {
		return err
	}

	numString := string(htmlutils.GetNodeText(divs[0], "a"))
	t.Chapter.Number = strings.Replace(numString, "Capitolo ", "", -1)

	return nil
}
