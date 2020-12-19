package mangaworld

import (
	"golang.org/x/net/html"
)

type Sort string

const (
	AZ       Sort = "a-z"
	ZA       Sort = "z-a"
	MostRead Sort = "most_read"
	LessRead Sort = "less_read"
	Newest   Sort = "newest"
	Oldest   Sort = "oldest"
)

type Query struct {
	MangaName struct {
		Value  string
		Active bool
	}
	Genre []struct {
		Value  Genre
		Active bool
	}
	MangaType []struct {
		Value  Type
		Active bool
	}
	State []struct {
		Value  State
		Active bool
	}
	Author []struct {
		Value  string
		Active bool
	}
	Artist []struct {
		Value  string
		Active bool
	}
	Years []struct {
		Value  string
		Active bool
	}
	SortType Sort
}

func NewQuery() *Query {
	return &Query{SortType: AZ}
}

func (q *Query) createQuery() string {
	url := UrlSearch
	multiParams := false

	if q.MangaName.Active == true {
		url = url + "keyword=" + html.EscapeString(q.MangaName.Value)
		multiParams = true
	}

	for _, gen := range q.Genre {
		if gen.Active == true {
			if multiParams == true {
				url = url + "&genre=" + html.EscapeString(string(gen.Value))
			} else {
				url = url + "genre=" + html.EscapeString(string(gen.Value))
				multiParams = true
			}
		}
	}
	for _, typ := range q.MangaType {
		if typ.Active == true {
			if multiParams == true {
				url = url + "&type=" + html.EscapeString(string(typ.Value))
			} else {
				url = url + "type=" + html.EscapeString(string(typ.Value))
				multiParams = true
			}
		}
	}

	for _, stat := range q.State {
		if stat.Active == true {
			if multiParams == true {
				url = url + "&state=" + html.EscapeString(string(stat.Value))
			} else {
				url = url + "state=" + html.EscapeString(string(stat.Value))
				multiParams = true
			}
		}
	}

	for _, auth := range q.Author {
		if auth.Active == true {
			if multiParams == true {
				url = url + "&author=" + html.EscapeString(auth.Value)
			} else {
				url = url + "author=" + html.EscapeString(auth.Value)
				multiParams = true
			}
		}
	}

	for _, art := range q.Artist {
		if art.Active == true {
			if multiParams == true {
				url = url + "&artist=" + html.EscapeString(art.Value)
			} else {
				url = url + "artist=" + html.EscapeString(art.Value)
				multiParams = true
			}
		}
	}

	for _, year := range q.Genre {
		if year.Active == true {
			if multiParams == true {
				url = url + "&year=" + html.EscapeString(string(year.Value))
			} else {
				url = url + "year=" + html.EscapeString(string(year.Value))
				multiParams = true
			}
		}
	}

	url = url + "&sort=" + html.EscapeString(string(q.SortType))

	return url
}
