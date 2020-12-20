package mangaworld

/*type Mangaworld struct {
	Url string
}*/

const UrlSite = "https://www.mangaworld.cc/"
const UrlSearch = "https://www.mangaworld.cc/archive?"

func SearchByName(name string) (manga []Manga, err error) {
	q := NewQuery()
	q.SetMangaName(name)
	return q.Do()
}

func SearchByGenre(genres []Genre) (manga []Manga, err error) {
	q := NewQuery()
	q.SetGenres(genres)
	return q.Do()
}

func SearchByType(types []Type) (manga []Manga, err error) {
	q := NewQuery()
	q.SetMangaTypes(types)
	return q.Do()
}

func SearchByStatus(states []State) (manga []Manga, err error) {
	q := NewQuery()
	q.SetStatus(states)
	return q.Do()
}
