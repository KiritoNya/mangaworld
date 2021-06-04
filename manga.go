package mangaworld

import (
	"errors"
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

//Manga is a structure containing all information relating to the manga and its chapters.
type Manga struct {
	Url              string
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
	Chapters         []*Chapter
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

//NewManga is the constructor of the manga object.
func NewManga(urlManga string) (*Manga, error) {

	resp, err := http.Get(urlManga)
	if err != nil {
		return &Manga{}, err
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		return &Manga{}, err
	}

	return &Manga{Url: urlManga, resp: node}, nil
}

//Add the title to the object.
func (m *Manga) GetTitle() error {
	h1, err := htmlutils.QuerySelector(m.resp, "h1", "class", "name bigger")
	if err != nil {
		return err
	}

	m.Title = string(htmlutils.GetNodeText(h1[0], "h1"))

	return nil
}

//Add the alternative title to the object.
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

//Add the cover url to the object.
func (m *Manga) GetCoverUrl() error {
	images, err := htmlutils.QuerySelector(m.resp, "img", "class", "rounded")
	if err != nil {
		return err
	}

	link, err := htmlutils.GetValueAttr(images[0], "img", "src")
	m.CoverUrl = strings.Split(string(link[0]), "?")[0]
	return nil
}

//Add the genres to the object.
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

//Add authors to the object.
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

//Add artist to the object.
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

//Add type to the object.
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

//Add state to the object.
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

//Add visual to the object.
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

//Add start years to the object.
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

//Add fansub to the object.
//TODO: Mettere apposto il fansub e altri poichè non è sempre in settima posizione
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

//Add animeworld url to the object.
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

//Add My Anime List url to the object.
func (m *Manga) GetMalUrl() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12 col-md-6 p-0 mt-1")
	if err != nil {
		return err
	}

	tagA, err := htmlutils.GetGeneralTags(divs[1], "a")
	if err != nil {
		return err
	}

	url, err := htmlutils.GetValueAttr(tagA[0], "a", "href")
	if err != nil {
		return err
	}

	m.MALUrl = string(url[0])

	return nil

}

//Add anilist url to the object.
func (m *Manga) GetAnilistUrl() error {

	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12 col-md-6 p-0 mt-1")
	if err != nil {
		return err
	}

	for _, div := range divs {

		tagA, err := htmlutils.GetGeneralTags(div, "a")
		if err != nil {
			return err
		}

		span, err := htmlutils.GetGeneralTags(tagA[0], "span")
		if err != nil {
			return err
		}

		node := htmlutils.GetNodeText(span[0], "span")
		if string(node) == "AniList" {
			url, err := htmlutils.GetValueAttr(tagA[0], "a", "href")
			if err != nil {
				return err
			}

			m.AnilistUrl = string(url[0])
			return nil
		}
	}

	return errors.New("IdAnilist not found")

}

//Add manga updates url to the object.
func (m *Manga) GetMangaUpdatesUrl() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "col-12 col-md-6 p-0 mt-1")
	if err != nil {
		return err
	}

	tagA, err := htmlutils.GetGeneralTags(divs[3], "a")
	if err != nil {
		return err
	}

	url, err := htmlutils.GetValueAttr(tagA[0], "a", "href")
	if err != nil {
		return err
	}

	m.MangaUpdatesUrl = string(url[0])

	return nil

}

//Add plot to the object.
func (m *Manga) GetPlot() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "id", "noidungm")
	if err != nil {
		return err
	}

	m.Plot = string(htmlutils.GetNodeText(divs[0], "div"))

	return nil
}

//Add number of volumes to the object.
func (m *Manga) GetVolumsNum() error {
	tagsP, err := htmlutils.QuerySelector(m.resp, "p", "class", "volume-name d-inline")
	if err != nil {
		return err
	}

	volumString := string(htmlutils.GetNodeText(tagsP[0], "p"))
	volumString = strings.Replace(volumString, "Volume ", "", -1)
	m.VolumsNum, err = strconv.Atoi(volumString)

	return nil
}

//Add number of chapters to the object.
func (m *Manga) GetChaptersNum() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "has-shadow zing-section bg-white newest-chapters mt-4")
	if err != nil {
		return err
	}
	a, err := htmlutils.QuerySelector(divs[0], "a", "class", "chap")
	if err != nil {
		return err
	}

	m.ChaptersNum = len(a)

	return nil
}

//Add object Chapters (only url field value) to the object manga.
func (m *Manga) GetChapters(start int, end int) error {

	//Check parametres
	if start < 1 {
		return errors.New("<error> The \"start\" parameter must be greater than 0")
	}
	if end < 1 {
		return errors.New("<error> The \"end\" parameter must be greater than 0")
	}

	divs, err := htmlutils.QuerySelector(m.resp, "div", "id", "chapterList")

	tagsA, err := htmlutils.QuerySelector(divs[0], "a", "class", "chap")
	if err != nil {
		return err
	}

	//Check parametres
	if m.ChaptersNum == 0 {
		m.GetChaptersNum()
	}

	if end > m.ChaptersNum {
		return errors.New("<error> The \"end\" parameter is greater than the number of chapters.")
	}
	if start > m.ChaptersNum {
		return errors.New("<error> The \"start\" parameter is greater than the number of available chapters")
	}

	//Reverse
	a := make([]*html.Node, len(tagsA))
	copy(a, tagsA)
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	tagsA = a

	for i := start - 1; i < end; i++ {

		url, err := htmlutils.GetValueAttr(tagsA[i], "a", "href")
		if err != nil {
			return err
		}

		c, err := NewChapter(string(url[0]))
		if err != nil {
			return err
		}

		m.Chapters = append(m.Chapters, c)
	}

	return nil
}

//Add relation to the object.
func (m *Manga) GetRelations() error {
	var mangas []Manga

	//Extract manga relation list
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "entry vertical")
	if err != nil {
		return errors.New(htmlutils.RenderNode(m.resp))
	}

	for _, div := range divs {

		//Get url of manga relation
		tagA, err := htmlutils.GetGeneralTags(div, "a")
		if err != nil {
			return err
		}

		url, err := htmlutils.GetValueAttr(tagA[0], "a", "href")
		if err != nil {
			return err
		}

		//Get title of manga relation
		title, err := htmlutils.GetValueAttr(tagA[0], "a", "title")
		if err != nil {
			return err
		}

		//Get image of manga relation
		img, err := htmlutils.GetGeneralTags(div, "img")
		if err != nil {
			return err
		}

		imgUrl, err := htmlutils.GetValueAttr(img[0], "img", "src")
		if err != nil {
			return err
		}

		//Get Section with year+type
		divs2, err := htmlutils.GetGeneralTags(div, "div")
		if err != nil {
			return err
		}

		year := string(htmlutils.GetNodeText(divs2[0], "div"))
		typeMedia := string(htmlutils.GetNodeText(divs2[1], "div"))

		//Create manga object
		manga, err := NewManga(string(url[0]))
		if err != nil {
			return err
		}

		manga.Title = string(title[0])
		manga.CoverUrl = string(imgUrl[0])
		manga.YearsStart = year
		manga.Type = Type(typeMedia)

		mangas = append(mangas, *manga)
	}
	return nil
}

//Add keywords to the object.
func (m *Manga) GetKeywords() error {
	divs, err := htmlutils.QuerySelector(m.resp, "div", "class", "has-shadow top-wrapper p-3 mt-4 mb-3")
	if err != nil {
		return err
	}

	stripped := strip.StripTags(htmlutils.RenderNode(divs[0]))
	stripped = strings.Replace(stripped, "Keywords:", "", -1)
	stripped = strings.Replace(stripped, "  ", "", -1)
	//If there is a initial space
	if stripped[0] == ' ' {
		stripped = stripped[1:]
	}
	m.Keywords = strings.Split(stripped, " - ")

	return nil

}

//Download all chapters in the object in a folder defined by the dest parameter.
func (m *Manga) Download(dest string) error {
	for _, chapter := range m.Chapters {
		err := chapter.Download(dest)
		if err != nil {
			return err
		}
	}
	return nil
}
