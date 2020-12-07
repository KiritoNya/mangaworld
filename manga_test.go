package mangaworld

import "testing"

const (
	link = "https://www.mangaworld.cc/manga/1876/citrus-1"
)

var testManga = Manga{
	Title: "Citrus+",
	TitleAlternative: []string{
		"Citrus Plus",
		"シトラスプラス",
	},
	CoverUrl: "https://cdn.mangaworld.cc/mangas/5fa4d1817a7a701817de97e1.jpg",
	Genres: []Genre{
		"Drammatico",
		"Romantico",
		"Scolastico",
		"Slice of Life",
		"Yuri",
	},
	Authors: []string{
		"Saburouta",
	},
	Artists: []string{
		"Saburouta",
	},
	Type:       Manga_type,
	State:      Releasing,
	Visual:     4215,
	YearsStart: "2018",
	Fansub: Fansub{
		Name: "Phoenix Scans",
		Url:  "https://www.phantomreader.com/",
	},
	AnimeworldUrl:   "https://www.animeworld.tv/play/citrus.vci_L/",
	MALUrl:          "https://myanimelist.net/manga/117083",
	AnilistUrl:      "https://anilist.co/manga/103884",
	MangaUpdatesUrl: "https://www.mangaupdates.com/series.html?id=151981",
	Plot:            "Sequel di Citrus che racconta le nuove vicende di Mei e Yuzu.",
	VolumsNum:       3,
	ChaptersNum:     16,
	Relations: []Manga{
		{
			Title:      "Citrus",
			CoverUrl:   "https://cdn.mangaworld.cc/mangas/5f73f50ac360051e41af919e.jpg?1607352479165",
			YearsStart: "2012",
			Type:       Manga_type,
		},
	},
}

func TestManga_GetTitle(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetTitle()
	if err != nil {
		t.Error("Error to get title")
	}

	if m.Title != testManga.Title {
		t.Error("Error not obtain", testManga.Title, "but obtain", m.Title)
	} else {
		t.Log("Title [OK]")
	}
}

func TestManga_GetAlternativeTitle(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetAlternativeTitle()
	if err != nil {
		t.Error("Error to get alternative title")
	}

	for i, title := range m.TitleAlternative {
		if title != testManga.TitleAlternative[i] {
			t.Error("Error not obtain", title, "but obtain", testManga.TitleAlternative[i])
		} else {
			t.Log("Alternative title", i, "[OK]")
		}
	}
}

func TestManga_CoverUrl(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetCoverUrl()
	if err != nil {
		t.Error("Error to get cover")
	}

	if m.CoverUrl != testManga.CoverUrl {
		t.Error("Error not obtain", testManga.CoverUrl, "but obtain", m.CoverUrl)
	} else {
		t.Log("Title [OK]")
	}
}

func TestManga_GetGenre(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetAlternativeTitle()
	if err != nil {
		t.Error("Error to get genres")
	}

	for i, genre := range m.Genres {
		if genre != testManga.Genres[i] {
			t.Error("Error not obtain", genre, "but obtain", testManga.Genres[i])
		} else {
			t.Log("Alternative title", i, "[OK]")
		}
	}
}

func TestManga_GetAuthor(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetAuthors()
	if err != nil {
		t.Error("Error to get authors")
	}

	for i, auth := range m.Authors {
		if auth != testManga.Authors[i] {
			t.Error("Error not obtain", auth, "but obtain", testManga.Authors[i])
		} else {
			t.Log("Author", i, "[OK]")
		}
	}
}

func TestManga_GetArtists(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetAuthors()
	if err != nil {
		t.Error("Error to get artist")
	}

	for i, artist := range m.Authors {
		if artist != testManga.Artists[i] {
			t.Error("Error not obtain", artist, "but obtain", testManga.Artists[i])
		} else {
			t.Log("Author", i, "[OK]")
		}
	}
}

func TestManga_GetType(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetType()
	if err != nil {
		t.Error("Error to get type")
	}

	if m.Type != testManga.Type {
		t.Error("Error not obtain", testManga.Type, "but obtain", m.Type)
	} else {
		t.Log("Type [OK]")
	}
}

func TestManga_GetState(t *testing.T) {
	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetState()
	if err != nil {
		t.Error("Error to get state")
	}

	if m.State != testManga.State {
		t.Error("Error not obtain", testManga.State, "but obtain", m.State)
	} else {
		t.Log("State [OK]")
	}
}

func TestManga_GetVisual(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetVisual()
	if err != nil {
		t.Error("Error to get visual")
	}

	if m.Visual != testManga.Visual+1 {
		t.Error("Error not obtain", testManga.Visual+1, "but obtain", m.Visual)
	} else {
		t.Log("Visual [OK]")
	}
}

func TestManga_GetYearsStart(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetYearsStart()
	if err != nil {
		t.Error("Error to get year start")
	}

	if m.YearsStart != testManga.YearsStart {
		t.Error("Error not obtain", testManga.YearsStart, "but obtain", m.YearsStart)
	} else {
		t.Log("Year start [OK]")
	}
}

func TestManga_GetFansub(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetFansub()
	if err != nil {
		t.Error("Error to get fansub")
	}

	if m.Fansub.Name != testManga.Fansub.Name {
		t.Error("Error not obtain", testManga.Fansub.Name, "but obtain", m.Fansub.Name)
	} else {
		t.Log("Fansub name [OK]")
	}

	if m.Fansub.Url != testManga.Fansub.Url {
		t.Error("Error not obtain", testManga.Fansub.Url, "but obtain", m.Fansub.Url)
	} else {
		t.Log("Fansub url [OK]")
	}
}

func TestManga_GetAnimeworldUrl(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetAnimeworldUrl()
	if err != nil {
		t.Error("Error to get animeworld url")
	}

	if m.AnimeworldUrl != testManga.AnimeworldUrl {
		t.Error("Error not obtain", testManga.AnimeworldUrl, "but obtain", m.AnimeworldUrl)
	} else {
		t.Log("Animeworld url [OK]")
	}
}

func TestManga_GetMalUrl(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetMalUrl()
	if err != nil {
		t.Error("Error to get MAL url")
	}

	if m.MALUrl != testManga.AnilistUrl {
		t.Error("Error not obtain", testManga.AnilistUrl, "but obtain", m.MALUrl)
	} else {
		t.Log("MAL url [OK]")
	}
}

func TestManga_GetAnilistUrl(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetAnilistUrl()
	if err != nil {
		t.Error("Error to get anilist url")
	}

	if m.AnilistUrl != testManga.AnilistUrl {
		t.Error("Error not obtain", testManga.AnilistUrl, "but obtain", m.AnilistUrl)
	} else {
		t.Log("Anilist url [OK]")
	}
}

func TestManga_GetMangaUpdatesUrl(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetMangaUpdatesUrl()
	if err != nil {
		t.Error("Error to get MangaUpdates url")
	}

	if m.MangaUpdatesUrl != testManga.MangaUpdatesUrl {
		t.Error("Error not obtain", testManga.MangaUpdatesUrl, "but obtain", m.MangaUpdatesUrl)
	} else {
		t.Log("MangaUpdates url [OK]")
	}
}

func TestManga_GetPlot(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetPlot()
	if err != nil {
		t.Error("Error to get plot")
	}

	if m.Plot != testManga.Plot {
		t.Error("Error not obtain", testManga.Plot, "but obtain", m.Plot)
	} else {
		t.Log("Plot [OK]")
	}

}

func TestManga_GetVolumsNum(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetVolumsNum()
	if err != nil {
		t.Error("Error to get volums num")
	}

	if m.VolumsNum != testManga.VolumsNum {
		t.Error("Error not obtain", testManga.VolumsNum, "but obtain", m.VolumsNum)
	} else {
		t.Log("Volums num [OK]")
	}

}

func TestManga_GetChaptersNum(t *testing.T) {

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetChaptersNum()
	if err != nil {
		t.Error("Error to get chapters num")
	}

	if m.ChaptersNum != testManga.ChaptersNum {
		t.Error("Error not obtain", testManga.ChaptersNum, "but obtain", m.ChaptersNum)
	} else {
		t.Log("Chapters num [OK]")
	}

}

func TestManga_GetRelations(t *testing.T) {
	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetRelations()
	if err != nil {
		t.Error("Error to get relation")
	}

	for i, manga := range m.Relations {
		if testManga.Relations[i].Title != manga.Title {
			t.Error("Error not obtain", testManga.Relations[i].Title, "but obtain", manga.Title)
		} else {
			t.Log("Relation ", i, "title [OK]")
		}

		if testManga.Relations[i].CoverUrl != manga.CoverUrl {
			t.Error("Error not obtain", testManga.Relations[i].CoverUrl, "but obtain", manga.CoverUrl)
		} else {
			t.Log("Relation ", i, "cover url [OK]")
		}

		if testManga.Relations[i].YearsStart != manga.YearsStart {
			t.Error("Error not obtain", testManga.Relations[i].YearsStart, "but obtain", manga.YearsStart)
		} else {
			t.Log("Relation ", i, "years start [OK]")
		}

		if testManga.Relations[i].Type != manga.Type {
			t.Error("Error not obtain", testManga.Relations[i].Type, "but obtain", manga.Type)
		} else {
			t.Log("Relation ", i, "type [OK]")
		}
	}
}
