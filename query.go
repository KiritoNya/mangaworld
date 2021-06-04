package mangaworld

import (
	"fmt"
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//Sort is a type that defines the type of sorting in query results.
type Sort string

const (
	AZ       Sort = "a-z"
	ZA       Sort = "z-a"
	MostRead Sort = "most_read"
	LessRead Sort = "less_read"
	Newest   Sort = "newest"
	Oldest   Sort = "oldest"
	minYears int  = 1968
)

//Query is an object query that allows you to build your own personalized search.
type Query struct {
	MangaName struct {
		Val    string
		Active bool
	}
	Genre struct {
		Val    []Genre
		Active bool
	}
	MangaType struct {
		Val    []Type
		Active bool
	}
	State struct {
		Val    []State
		Active bool
	}
	Author struct {
		Val    []string
		Active bool
	}
	Artist struct {
		Val    []string
		Active bool
	}
	Year struct {
		Val    []string
		Active bool
	}
	SortType Sort
}

//NewQuery is a constructor of object Query.
func NewQuery() *Query {
	return &Query{SortType: AZ}
}

//Set manga name value to the object.
func (q *Query) SetMangaName(name string) {
	q.MangaName.Active = true
	q.MangaName.Val = url.QueryEscape(name)
}

//Set manga genres value to the object.
func (q *Query) SetGenres(genres []Genre) {
	q.Genre.Active = true
	for _, genre := range genres {
		q.Genre.Val = append(q.Genre.Val, genre)
	}
}

//Set manga types value to the object.
func (q *Query) SetMangaTypes(mangaTypes []Type) {
	q.MangaType.Active = true
	for _, mangaType := range mangaTypes {
		q.MangaType.Val = append(q.MangaType.Val, mangaType)
	}
}

//Set manga status value to the object.
func (q *Query) SetStatus(states []State) {
	q.State.Active = true
	for _, state := range states {
		q.State.Val = append(q.State.Val, state)
	}
}

//Set manga authors value to the object.
func (q *Query) SetAuthors(authors []string) {
	q.Author.Active = true
	for _, auth := range authors {
		q.Author.Val = append(q.Author.Val, url.QueryEscape(auth))
	}
}

//Set manga artists value to the object.
func (q *Query) SetArtists(artists []string) {
	q.Artist.Active = true
	for _, artist := range artists {
		q.Artist.Val = append(q.Artist.Val, url.QueryEscape(artist))
	}
}

//Set manga years value to the object.
func (q *Query) SetYears(years []string) {

	for _, year := range years {
		match, _ := regexp.MatchString("[0-9][0-9][0-9][0-9]", year)
		if match == true {
			yearInt, _ := strconv.Atoi(year)
			if yearInt >= minYears && yearInt <= time.Now().Year() {
				q.Year.Val = append(q.Year.Val, year)
				q.Year.Active = true
			}
		}
	}
}

//Set sort result value to the object.
func (q *Query) SetSort(method Sort) {
	q.SortType = method
}

//Executes the query that was previously set.
func (q *Query) Do() (mangas []*Manga, err error) {

	if !serviceActive {
		err = NewDefaultService()
		if err != nil {
			return nil, err
		}
	}

	var htmlNode *html.Node

	query := q.createQuery()

	resp, err := doRequest(query)
	if err != nil {
		return nil, err
	}

	htmlNode, err = html.Parse(strings.NewReader(resp))
	if err != nil {
		return nil, err
	}

	li, err := htmlutils.QuerySelector(htmlNode, "li", "class", "page-item last")
	if err != nil {
		return nil, err
	}

	a, err := htmlutils.GetGeneralTags(li[0], "a")
	if err != nil {
		return nil, err
	}

	num := htmlutils.GetNodeText(a[0], "a")
	if err != nil {
		return nil, err
	}

	numInt, err := strconv.Atoi(string(num))
	if err != nil {
		return nil, err
	}

	for i:=1; i <= numInt;  i++ {

		query2 := query + "&page=" + strconv.Itoa(i)
		fmt.Println("QUERY: ", query2)

		if i != 1 {

			resp, err := http.Get(query2)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			htmlNode, err = html.Parse(resp.Body)
			if err != nil {
				return nil, err
			}

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

			manga, err := NewManga(string(url[0]))
			if err != nil {
				return nil, err
			}

			mangas = append(mangas, manga)
		}

	}

	return mangas, nil
}

//create the query string for mangaworld pre-custom search.
func (q *Query) createQuery() string {
	urlSearch := UrlSearch
	multiParams := false

	if q.MangaName.Active == true {
		urlSearch = urlSearch + "keyword=" + q.MangaName.Val
		multiParams = true
	}

	if q.Genre.Active == true {
		for _, gen := range q.Genre.Val {
			if multiParams == true {
				urlSearch = urlSearch + "&genre=" + strings.Replace(string(gen), " ", "-", -1)
			} else {
				urlSearch = urlSearch + "genre=" + strings.Replace(string(gen), " ", "-", -1)
				multiParams = true
			}
		}
	}

	if q.MangaType.Active == true {
		for _, typ := range q.MangaType.Val {
			if multiParams == true {
				urlSearch = urlSearch + "&type=" + string(typ)
			} else {
				urlSearch = urlSearch + "type=" + string(typ)
				multiParams = true
			}
		}
	}

	if q.State.Active == true {
		for _, stat := range q.State.Val {
			if multiParams == true {
				urlSearch = urlSearch + "&status=" + searchState(stat)
			} else {
				urlSearch = urlSearch + "status=" + searchState(stat)
				multiParams = true
			}
		}
	}

	if q.Author.Active == true {
		for _, auth := range q.Author.Val {
			if multiParams == true {
				urlSearch = urlSearch + "&author=" + auth
			} else {
				urlSearch = urlSearch + "author=" + auth
				multiParams = true
			}
		}
	}

	if q.Artist.Active == true {
		for _, art := range q.Artist.Val {
			if multiParams == true {
				urlSearch = urlSearch + "&artist=" + art
			} else {
				urlSearch = urlSearch + "artist=" + art
				multiParams = true
			}
		}
	}

	if q.Year.Active == true {
		for _, year := range q.Year.Val {
			if multiParams == true {
				urlSearch = urlSearch + "&year=" + year
			} else {
				urlSearch = urlSearch + "year=" + year
				multiParams = true
			}
		}
	}

	urlSearch = urlSearch + "&sort=" + string(q.SortType)

	return urlSearch
}
