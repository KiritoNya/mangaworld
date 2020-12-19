package mangaworld

import (
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
)

type ChapterNew struct {
	MangaNew  Manga
	ChapteNew Chapter
	node      *html.Node
}

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

func (cn *ChapterNew) GetManga() error {

	/*resp, err := http.Get(UrlSite)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	grid, err := htmlutils.QuerySelector(node, "div", "class", "comics-grid")
	if err != nil {
		return nil, err
	}

	chDivs, err := htmlutils.QuerySelector(grid[0], "li", "class", "entry")
	if err != nil {
		return nil, err
	}

	numPag, err := getNewestNumPage(node)
	if err != nil {
		return nil, err
	}

	if num > ((numPag-1)*16) - (len(chDivs)-1) {
		return nil, errors.New("Number inserted greater than the new chapters released")
	}*/

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

/*func getNewestNumPage(node *html.Node) (int, error) {
	divs, err := htmlutils.QuerySelector(node, "div", "class", "pagination-container d-flex justify-content-center")
	if err != nil {
		return 0, err
	}

	li, err := htmlutils.QuerySelector(divs[0], "li", "class", "page-link")
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(string(htmlutils.GetNodeText(li[0], "li")))
	if err != nil {
		return 0, err
	}

	return num, nil
}*/

func (cn *ChapterNew) GetChapter() error {

	divs, err := htmlutils.QuerySelector(cn.node, "div", "class", "content")
	if err != nil {
		return err
	}

	divsCh, err := htmlutils.QuerySelector(divs[0], "div", "class", "d-flex flex-wrap flex-row")
	if err != nil {
		return err
	}

	tagsA, err := htmlutils.GetGeneralTags(divsCh[0], "a")
	if err != nil {
		return err
	}

	url, err := htmlutils.GetValueAttr(tagsA[0], "a", "href")
	if err != nil {
		return err
	}

	cn.ChapteNew.Url = string(url[0])

	return nil
}