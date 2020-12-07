package mangaworld

import (

	//"fmt"
	"github.com/KiritoNya/htmlutils"
	"github.com/grokify/html-strip-tags-go"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
	"strings"
)

type Fansub struct {
	Name string
	Url  string
}

type Manga struct {
	Title            string
	TitleAlternative []string
	CoverUrl         string
	Genres           []Genre
	Authors          []string
	Artists          []string
	Type             Type
	State            State
	Plot             string
	YearsStart       string
	VolumsNum        int
	ChaptersNum      int
	Chapters         []Chapter
	Relations        []Manga
	Visual           int
	Fansub           Fansub
	AnimeworldUrl    string
	AnilistUrl       string
	MALUrl           string
	MangaUpdatesUrl  string
	Keywords         []string
	resp             *html.Node
}

func NewManga(urlManga string) (*Manga, error) {
	var m Manga

	resp, err := http.Get(urlManga)
	if err != nil {
		return &Manga{}, err
	}
	defer resp.Body.Close()

	m.resp, err = html.Parse(resp.Body)
	if err != nil {
		return &Manga{}, err
	}

	return &m, nil
}

func (m *Manga) GetTitle() error {
	h1, err := htmlutils.QuerySelector(m.resp, "h1", "class", "name bigger")
	if err != nil {
		return err
	}

	m.Title = string(htmlutils.GetNodeText(h1[0], "h1"))

	return nil
}

func (m *Manga) GetAlternativeTitle() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12")
	if err != nil {
		return err
	}

	stripped := strip.StripTags(htmlutils.RenderNode(divs[0]))

	stripped = strings.Replace(stripped, "Titoli alternativi:  ", "", -1)
	m.TitleAlternative = strings.Split(stripped, ", ")
	return nil
}

func (m *Manga) GetCoverUrl() error {
	images, err := htmlutils.QuerySelector(m.resp, "img", "class", "rounded")
	if err != nil {
		return err
	}

	link, err := htmlutils.GetValueAttr(images[0], "img", "src")
	m.CoverUrl = strings.Split(string(link[0]), "?")[0]
	return nil
}

func (m *Manga) GetGenre() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12")
	if err != nil {
		return err
	}

	tagsA, err := htmlutils.GetGeneralTags(divs[1], "a")
	if err != nil {
		return err
	}

	for _, tagA := range tagsA {
		genre := htmlutils.GetNodeText(tagA, "a")
		m.Genres = append(m.Genres, Genre(genre))
	}

	return nil
}

func (m *Manga) GetAuthors() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12 col-md-6")
	if err != nil {
		return err
	}

	tagsA, err := htmlutils.GetGeneralTags(divs[0], "a")
	if err != nil {
		return err
	}

	for _, tagA := range tagsA {
		auth := htmlutils.GetNodeText(tagA, "a")
		m.Authors = append(m.Authors, string(auth))
	}

	return nil
}

func (m *Manga) GetArtists() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12 col-md-6")
	if err != nil {
		return err
	}

	tagsA, err := htmlutils.GetGeneralTags(divs[1], "a")
	if err != nil {
		return err
	}

	for _, tagA := range tagsA {
		artist := htmlutils.GetNodeText(tagA, "a")
		m.Artists = append(m.Artists, string(artist))
	}

	return nil
}

func (m *Manga) GetType() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12 col-md-6")
	if err != nil {
		return err
	}

	tagsA, err := htmlutils.GetGeneralTags(divs[2], "a")
	if err != nil {
		return err
	}

	artist := htmlutils.GetNodeText(tagsA[0], "a")
	m.Type = Type(artist)

	return nil
}

func (m *Manga) GetState() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12 col-md-6")
	if err != nil {
		return err
	}

	tagsA, err := htmlutils.GetGeneralTags(divs[3], "a")
	if err != nil {
		return err
	}

	state := htmlutils.GetNodeText(tagsA[0], "a")
	m.State = State(state)

	return nil
}

func (m *Manga) GetVisual() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12 col-md-6")
	if err != nil {
		return err
	}

	spans, err := htmlutils.GetGeneralTags(divs[4], "span")
	if err != nil {
		return err
	}

	visual := htmlutils.GetNodeText(spans[1], "span")
	m.Visual, err = strconv.Atoi(string(visual))
	if err != nil {
		return err
	}

	return nil
}

func (m *Manga) GetYearsStart() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12 col-md-6")
	if err != nil {
		return err
	}

	tagA, err := htmlutils.GetGeneralTags(divs[5], "a")
	if err != nil {
		return err
	}

	years := htmlutils.GetNodeText(tagA[0], "a")
	m.YearsStart = string(years)

	return nil
}

func (m *Manga) GetFansub() error {
	var f Fansub

	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12 col-md-6")
	if err != nil {
		return err
	}

	tagA, err := htmlutils.GetGeneralTags(divs[6], "a")
	if err != nil {
		return err
	}

	f.Name = string(htmlutils.GetNodeText(tagA[0], "a"))
	url, err := htmlutils.GetValueAttr(tagA[0], "a", "href")
	if err != nil {
		return err
	}
	f.Url = string(url[0])

	m.Fansub = f

	return nil
}

func (m *Manga) GetAnimeworldUrl() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12 col-md-6 p-0 mt-1")
	if err != nil {
		return err
	}

	tagA, err := htmlutils.GetGeneralTags(divs[0], "a")
	if err != nil {
		return err
	}

	url, err := htmlutils.GetValueAttr(tagA[0], "a", "href")
	if err != nil {
		return err
	}

	m.AnimeworldUrl = string(url[0])

	return nil

}
