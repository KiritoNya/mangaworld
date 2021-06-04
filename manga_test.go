package mangaworld

import (
	"testing"
)

const (
	link = "https://www.mangaworld.io/manga/508/gokujou-drops"
)

var testManga = Manga{
	Title: "Kishuku Gakkou no Juliet",
	TitleAlternative: []string{
		"Boarding School Juliet",
		"Juliet of Boarding School",
		"Kushuku Gakkou no Alice",
		"Romio VS Juliet (French)",
		"รักลับๆ ข้ามหอ ของนายหมากับน้องแมว (Thai)",
		"寄宿学校のジュリエット",
		"寄宿学校的朱丽叶",
		"기숙학교의 줄리엣",
	},
	CoverUrl: "https://cdn.mangaworld.cc/mangas/5fa62d90552b377191b013ec.jpg",
	Genres: []Genre{
		"Azione",
		"Commedia",
		"Ecchi",
		"Harem",
		"Romantico",
		"Scolastico",
		"Shounen",
	},
	Authors: []string{
		"KANEDA Yousuke",
	},
	Artists: []string{
		"KANEDA Yousuke",
	},
	Type:       Manga_type,
	State:      Releasing,
	Visual:     80627,
	YearsStart: "2015",
	Fansub: Fansub{
		Name: "Phoenix Scans",
		Url:  "https://www.phantomreader.com/",
	},
	AnimeworldUrl:   "https://www.animeworld.tv/play/kishuku-gakkou-no-juliet.k8c0S/",
	MALUrl:          "https://myanimelist.net/manga/91514",
	AnilistUrl:      "https://anilist.co/manga/86279",
	MangaUpdatesUrl: "https://www.mangaupdates.com/series.html?id=122123",
	Plot:            "Juliet Persia e Romio Inuzuka sono a capo di due dormitori rivali presso la Dahlia Academy Boarding School, in cui risiedono gli studenti appartenenti a due paesi differenti, i “Black Dogs” e i “White Cats”. Le due fazioni sono in guerra tra loro per contendersi il controllo dell’isola in cui vivono. Juliet e Romio si dimostrano agli occhi dei loro compagni di dormitorio come acerrimi nemici, ma in realtà sono segretamente innamorati l’uno dell’altra e costretti a mantenere nascosta la loro relazione per non subire ripercussioni.",
	VolumsNum:       11,
	ChaptersNum:     26,
	Chapters: []*Chapter{
		{Url: "https://www.mangaworld.cc/manga/1919/kishuku-gakkou-no-juliet/read/601ab540f0ba3f6b7ca7b8e3"},
		{Url: "https://www.mangaworld.cc/manga/1919/kishuku-gakkou-no-juliet/read/601d8fb51116ba6b71afa17c"},
		{Url: "https://www.mangaworld.cc/manga/1919/kishuku-gakkou-no-juliet/read/60242979b6bece4364bc46f4"},
	},
	Relations: []Manga{
		{
			Title:      "",
			CoverUrl:   "",
			YearsStart: "",
			Type:       "",
		},
	},
	Keywords: []string{
		"Kishuku Gakkou no Juliet Scan ITA",
		"Kishuku Gakkou no Juliet Scan ITALIANE",
		"Kishuku Gakkou no Juliet MangaWorld",
		"Kishuku Gakkou no Juliet MangaDex",
		"Kishuku Gakkou no Juliet MangaEden",
		"Kishuku Gakkou no Juliet SUB ITA",
		"Kishuku Gakkou no Juliet Leggi online",
		"Kishuku Gakkou no Juliet Reader ITA",
		"Kishuku Gakkou no Juliet Scan online",
		"Kishuku Gakkou no Juliet Read online",
		"Kishuku Gakkou no Juliet Online",
		"Kishuku Gakkou no Juliet Manga ITA",
		"Kishuku Gakkou no Juliet Manga Scan",
		"Kishuku Gakkou no Juliet Scan manga online",
		"Kishuku Gakkou no Juliet ITA Scan",
		"MangaWorld Kishuku Gakkou no Juliet",
		"Scan ITA Kishuku Gakkou no Juliet",
		"Read online Kishuku Gakkou no Juliet",
	},
}

func TestNewManga(t *testing.T) {
	_, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	t.Log("Object create [OK]")
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
		t.Error("Error to get fansub: ", err)
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

	if m.MALUrl != testManga.MALUrl {
		t.Error("Error not obtain", testManga.MALUrl, "but obtain", m.MALUrl)
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

func TestManga_GetChapters(t *testing.T) {
	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetChapters(66, 68)
	if err != nil {
		t.Error("Error to get chapters: ", err)
	}

	//m.Chapters = m.Chapters[:3]

	for i, chapter := range m.Chapters {
		if chapter.Url != testManga.Chapters[i].Url {
			t.Error("Error not obtain", testManga.Chapters[i].Url, "but obtain", chapter.Url)
		} else {
			t.Log("Chapters ", i, "[OK]")
		}
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

func TestManga_GetKeyword(t *testing.T) {
	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetKeywords()
	if err != nil {
		t.Error("Error to get keywords")
	}

	for i, key := range m.Keywords {
		if key != testManga.Keywords[i] {
			t.Error("Error not obtain", testManga.Keywords[i], "but obtain", key)
		} else {
			t.Log("Key", i, "[OK]")
		}
	}

}
