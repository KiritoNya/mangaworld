package mangaworld

type Mangaworld struct {
	Url string
}

var UrlSite = "https://www.mangaworld.cc/"

func New(url string) *Mangaworld {
	return &Mangaworld{Url: url}
}
