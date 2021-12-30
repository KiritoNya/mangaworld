package v2

import (
	"errors"
	"github.com/KiritoNya/gotaku"
	"github.com/KiritoNya/gotaku/image"
	"github.com/KiritoNya/gotaku/manga"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//Fansub is a struct that contains all information relating the fansub.
type Fansub struct {
	Name string
	Url  string
}

//Manga is a structure containing all information relating to the manga and its chapters.
type Manga struct {
	manga.Manga
	Volumes []*Volume
	Relations []*Manga
	doc *goquery.Document
}

//NewManga is the constructor of the manga object.
func NewManga(urlManga string) (*Manga, error) {
	var m Manga
	var err error

	m.doc, err = doRequest(urlManga)
	if err != nil {
		return nil, err
	}
	m.Url = urlManga

	return &m, nil
}

//NewMangaFromHtml is the constructor of the manga object from html node.
func NewMangaFromHtml(html *html.Node) (*Manga, error) {
	var m Manga

	//Check error
	if html == nil {
		return nil, errors.New("Html is empty")
	}
	m.doc = goquery.NewDocumentFromNode(html)
	return &m, nil
}

//NewMangaFromFile is the constructor of the manga object from html file.
func NewMangaFromFile(path string) (*Manga, error) {
	var m Manga

	//Read file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	//Parse html
	htmlNode, err := html.Parse(f)
	if err != nil {
		return nil, err
	}

	//Check error
	if htmlNode == nil {
		return nil, errors.New("Html is empty")
	}

	//Create document
	m.doc = goquery.NewDocumentFromNode(htmlNode)
	return &m, nil
}

//GetTitle - Add the title to the object.
func (m *Manga) GetTitle() error {
	//Get title
	title := m.doc.Find(".name.bigger").Text()

	//Check title
	if title == "" {
		return errors.New("Title not found")
	}

	m.Title.Pretty = title
	return nil
}

//GetAlternativeTitle - Add the alternative title to the object.
func (m *Manga) GetAlternativeTitle() ([]string, error) {
	titlesDiv := m.doc.Find(".meta-data.row").Find(".col-12").Get(0)
	rowTitles := goquery.NewDocumentFromNode(titlesDiv).Text()
	if rowTitles == "" {
		return nil, errors.New("Alternative titles not found")
	}
	rowTitles = strings.Replace(rowTitles, "Titoli alternativi:", "", -1)
	rowTitles = strings.Replace(rowTitles, " ", "", -1)
	titles := strings.Split(rowTitles, ",")
	return titles, nil
}

//GetCoverUrl - Add the cover url to the object.
func (m *Manga) GetCoverUrl() error {
	//Get image url
	imgUrl, ok := m.doc.Find("img.rounded").Attr("src")

	//Check image url
	if !ok {
		return errors.New("Image url not found")
	}

	//Get image info
	imgUrlMatrix := strings.Split(imgUrl, "?")
	imgName := filepath.Base(imgUrlMatrix[0])
	imageNameMatrix := strings.Split(imgName, ".")

	//Assign cover object to manga object
	m.CoverImage = image.Cover{
		Id:         "",
		ExtraLarge: image.Image{
			Id:     imgUrlMatrix[1],
			Url:    imgUrlMatrix[0],
			Name:   imgName,
			Data:   nil,
			Format: imageNameMatrix[1],
		},
		Large:      image.Image{},
		Medium:     image.Image{},
		Color:      "",
	}

	return nil
}

//GetGenres is a function that add the genres to the object.
func (m *Manga) GetGenres() error {
	//Get genres
	section := m.doc.Find("div.col-12").Get(1)
	goquery.NewDocumentFromNode(section).Find("a").Each(func (index int, element *goquery.Selection) {
		m.Genres = append(m.Genres, gotaku.Genre{Name: element.Text()})
	})

	//Check genres
	if m.Genres == nil {
		return errors.New("Genres not found")
	}

	return nil
}

//GetAuthors is a function that add authors to the object.
func (m *Manga) GetAuthors() error {
	m.doc.Find("div.col-12.col-md-6").EachWithBreak(func (index int, element *goquery.Selection) bool {
		if checkCategoryLabel(element, "Autore: ") {
			textNode := element.Find("a").Text()
			textNode = strings.Replace(textNode, "Autore:", "", -1)
			authors := strings.Split(textNode, ",")
			for _, author := range authors {
				m.getAuthor(author)
			}
			return false
		}
		return true
	})
	return nil
}

//GetArtists is a function that add artists to the object.
func (m *Manga) GetArtists() error {
	m.doc.Find("div.col-12.col-md-6").EachWithBreak(func (index int, element *goquery.Selection) bool {
		if checkCategoryLabel(element, "Artista: ") {
			textNode := element.Find("a").Text()
			textNode = strings.Replace(textNode, "Artista:", "", -1)
			artists := strings.Split(textNode, ",")
			for _, artist := range artists {
				m.getArtist(artist)
			}
			return false
		}
		return true
	})
	return nil
}

//GetType is a function that add type to the object.
func (m *Manga) GetType() error {
	m.doc.Find("div.col-12.col-md-6").EachWithBreak(func (index int, element *goquery.Selection) bool {
		if checkCategoryLabel(element, "Tipo: ") {
			m.Type = gotaku.MediaType(element.Find("a").Text())
			return false
		}
		return true
	})

	//Check type
	if m.Type == "" {
		return errors.New("Manga type not found")
	}

	return nil
}

//GetStatus is a function that add status to the object
func (m *Manga) GetStatus() error {
	m.doc.Find("div.col-12.col-md-6").EachWithBreak(func (index int, element *goquery.Selection) bool {
		if checkCategoryLabel(element, "Stato: ") {
			m.Status = gotaku.Status(element.Find("a").Text())
			return false
		}
		return true
	})

	//Check type
	if m.Status == "" {
		return errors.New("Manga status not found")
	}

	return nil
}

//GetReleaseDate is a function that add status to the object
func (m *Manga) GetReleaseDate() error {
	m.doc.Find("div.col-12.col-md-6").EachWithBreak(func (index int, element *goquery.Selection) bool {
		if checkCategoryLabel(element, "Anno di uscita: ") {
			//Get year
			yearInt, err := strconv.Atoi(element.Find("a").Text())
			if err != nil {
				return false
			}
			m.StartDate = time.Date(yearInt, 1,1, 1, 1, 1,1,time.Local)
			return false
		}
		return true
	})

	//Check type
	if m.StartDate.IsZero() {
		return errors.New("Manga release date not found")
	}

	return nil
}

//GetFansub is a function that add status to the object
func (m *Manga) GetFansub() error {
	m.doc.Find("div.col-12.col-md-6").EachWithBreak(func (index int, element *goquery.Selection) bool {
		if checkCategoryLabel(element, "Fansub: ") {
			m.Fansub.Name = element.Find("a").Text()
			m.Fansub.Url, _ = element.Find("a").Attr("href")
			return false
		}
		return true
	})

	//Check type
	if m.Fansub.Name == "" {
		return errors.New("Manga fansub not found")
	}

	return nil
}

//GetAnilist is a function that adds external anilist links to the manga object
func (m *Manga) GetAnilist() error {
	//Get all external links
	links, err := m.getExternalLinks()
	if err != nil {
		return err
	}

	//Foreach link
	for index, link := range links {
		//Check anilist link
		if link.Site == "anilist" {
			m.ExternalLinks = append(m.ExternalLinks, link)
			break
		}

		if index == len(links) {
			return errors.New("Anilist link not found")
		}
	}

	return nil
}

//GetMyAnimeList is a function that adds external MyAnimeList link to the manga object
func (m *Manga) GetMyAnimeList() error {
	//Get all external links
	links, err := m.getExternalLinks()
	if err != nil {
		return err
	}

	//Foreach link
	for index, link := range links {
		//Check anilist link
		if link.Site == "myanimelist" {
			m.ExternalLinks = append(m.ExternalLinks, link)
			break
		}

		if index == len(links) {
			return errors.New("MyAnimeList link not found")
		}
	}

	return nil
}

//GetMangaUpdates is a function that adds external MangaUpdates link to the manga object
func (m *Manga) GetMangaUpdates() error {
	//Get all external links
	links, err := m.getExternalLinks()
	if err != nil {
		return err
	}

	//Foreach link
	for index, link := range links {
		//Check anilist link
		if link.Site == "mangaupdates" {
			m.ExternalLinks = append(m.ExternalLinks, link)
			break
		}

		if index == len(links) {
			return errors.New("MangaUpdates link not found")
		}
	}

	return nil
}

//GetExternalLinks is a function that adds all external links to the manga object
func (m *Manga) GetExternalLinks() error {
	links, err := m.getExternalLinks()
	if err != nil {
		return err
	}

	m.ExternalLinks = links
	return nil
}

//GetPlot is a function that adds plot to the manga object
func (m *Manga) GetPlot() error {
	plot := m.doc.Find("#noidungm").Text()
	if plot == "" {
		return errors.New("Manga plot not found")
	}

	m.Description = plot
	return nil
}

//GetVolumesNum is a function that adds the number of volumes to the object.
func (m *Manga) GetVolumesNum() error {
	m.doc.Find(".volume-name.d-inline").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		volNumString := selection.Text()
		volNumString = strings.Replace(volNumString, "Volume ", "", -1)
		volNum, _ := strconv.Atoi(volNumString)
		m.NumVolumes = volNum
		return false
	})

	// Check error
	if m.NumVolumes == 0 {
		return errors.New("Number of volumes not found")
	}

	return nil
}

//GetChaptersNum is a function that adds the number of chapters to the object.
func (m *Manga) GetChaptersNum() error {
	m.doc.Find(".has-shadow.zing-section.bg-white.newest-chapters.mt-4").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		chapterNum := selection.Find(".chap").Length()
		m.NumChapters = chapterNum
		return false
	})

	// Check error
	if m.NumChapters == 0 {
		return errors.New("Chapters number not found")
	}

	return nil
}

//GetVolumes is a functiont that adds volumes object to the object manga.
func (m *Manga) GetVolumes() error {
	//Get volumes
	m.doc.Find(".volume-element.pl-2").Each(func(i int, selection *goquery.Selection) {
		v := NewVolume(selection)
		m.Volumes = append(m.Volumes, v)
	})

	//Check volumes
	if m.Volumes == nil {
		return errors.New("Volumes not found")
	}

	return nil
}

//GetRelations is a function that adds relations manga to the object.
func (m *Manga) GetRelations() error {
	var relErrors []string

	m.doc.Find(".entry.vertical").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		url, _ := selection.Find("a").Attr("href")

		//Create new manga object
		m2, err := NewManga(url)
		if err != nil {
			relErrors = append(relErrors, err.Error())
			return false
		}

		m.Relations = append(m.Relations, m2)
		return true
	})

	//Check errors
	if relErrors != nil {
		return errors.New(strings.Join(relErrors, "\n"))
	}
	return nil
}

//GetKeyboards is a function that adds keyboards to the manga object
func (m *Manga) GetKeyboards() error {
	m.Keywords = m.doc.Find(".has-shadow.top-wrapper.p-3.mt-4.mb-3").Last().Text()
	m.Keywords = strings.Replace(m.Keywords, "Keywords:", "", -1)
	m.Keywords = strings.Replace(m.Keywords, "  ", "", -1)
	if m.Keywords[0] == ' ' {
		m.Keywords = m.Keywords[1:]
	}

	// Check error
	if m.Keywords == "" {
		return errors.New("Manga keyboards not found")
	}

	return nil
}

/*
==================================
			PRIVATE METHODS
==================================
 */

//getAuthor is a function that creates the author staff object and assign it to the manga object
func (m *Manga) getAuthor(author string) {
		var s gotaku.Staff
		author = strings.ToLower(author)
		author = strings.Title(author)
		authorName := strings.Split(author, " ")

		s.Name = gotaku.StaffName{
			First:       authorName[0],
			Middle:      "",
			Last:        authorName[1],
			Full:        "",
			Native:      "",
			Alternative: nil,
		}
		s.Role = "Author"

		m.Staff = append(m.Staff, &s)
}

//getArtist is a function that creates the artist staff object and assign it to the manga object
func (m *Manga) getArtist(artist string) {
	var s gotaku.Staff
	artist = strings.ToLower(artist)
	artist = strings.Title(artist)
	artistName := strings.Split(artist, " ")

	s.Name = gotaku.StaffName{
		First:       artistName[0],
		Middle:      "",
		Last:        artistName[1],
		Full:        "",
		Native:      "",
		Alternative: nil,
	}
	s.Role = "Artist"

	m.Staff = append(m.Staff, &s)
}

//getExternalLinks is a function that get all manga references
func (m *Manga) getExternalLinks() ([]*gotaku.ExternalLinks, error) {
	var links []*gotaku.ExternalLinks

	m.doc.Find(".references.mt-1.row").Find("a").Each(func(i int, selection *goquery.Selection) {
		var e gotaku.ExternalLinks
		siteName := strings.Replace(selection.Find("span.fix-weird-shit").Text(), " ", "", -1)
		siteName = strings.ToLower(siteName)
		e.Site = siteName
		e.Url, _ = selection.Attr("href")
		links = append(links, &e)
	})

	if links == nil {
		return nil, errors.New("Manga external links not found")
	}

	return links, nil
}

//checkCategoryLabel is a function that checks if the row of manga info table is the target category
func checkCategoryLabel(row *goquery.Selection, target string) bool {
	find := strings.Replace(row.Find("span").Text(), "  ", "", -1)
	if strings.ToLower(find) == strings.ToLower(target) {
		return true
	}
	return false
}