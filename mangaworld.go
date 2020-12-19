package mangaworld

import (
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
	"net/http"
)

/*type Mangaworld struct {
	Url string
}*/

const UrlSite = "https://www.mangaworld.cc/"
const UrlSearch = "https://www.mangaworld.cc/archive?"

/*func New(url string) *Mangaworld {
	return &Mangaworld{Url: url}
}*/

func (q *Query) Search(url string) (mangas []Manga, err error) {

	url = q.createQuery()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	htmlNode, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	divs, err := htmlutils.QuerySelector(htmlNode, "div", "class", "entry")
	if err != nil {
		return nil, err
	}

	for _, div := range divs {
		tagA, err := htmlutils.GetGeneralTags(div, "a")
		if err != nil {
			return nil, err
		}

		url, err := htmlutils.GetValueAttr(tagA[0], "a", "href")
		if err != nil {
			return nil, err
		}

		mangas = append(mangas, Manga{Url: string(url[0])})
	}
	return mangas, nil
}
