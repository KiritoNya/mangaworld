package mangaworld

import "testing"

const (
	link = "https://www.mangaworld.cc/manga/1876/citrus-1"
)

func TestManga_GetTitle(t *testing.T) {

	const title = "Citrus+"

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetTitle()
	if err != nil {
		t.Error("Error to get title")
	}

	if m.Title != title {
		t.Error("Error not obtain", title, "but obtain", m.Title)
	} else {
		t.Log("Title [OK]")
	}
}

func TestManga_GetAlternativeTitle(t *testing.T) {

	var alternativeName = []string{
		"Citrus Plus",
		"シトラスプラス",
	}

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetAlternativeTitle()
	if err != nil {
		t.Error("Error to get alternative title")
	}

	for i, title := range m.TitleAlternative {
		if title != alternativeName[i] {
			t.Error("Error not obtain", title, "but obtain", alternativeName[i])
		} else {
			t.Log("Alternative title", i, "[OK]")
		}
	}
}

func TestManga_CoverUrl(t *testing.T) {
	const imageUrl = "https://cdn.mangaworld.cc/mangas/5fa4d1817a7a701817de97e1.jpg"

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetCoverUrl()
	if err != nil {
		t.Error("Error to get cover")
	}

	if m.CoverUrl != imageUrl {
		t.Error("Error not obtain", imageUrl, "but obtain", m.CoverUrl)
	} else {
		t.Log("Title [OK]")
	}
}

func TestManga_GetGenre(t *testing.T) {
	var genresCheck = []Genre{
		"Drammatico",
		"Romantico",
		"Scolastico",
		"Slice of Life",
		"Yuri",
	}

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetAlternativeTitle()
	if err != nil {
		t.Error("Error to get genres")
	}

	for i, genre := range m.Genres {
		if genre != genresCheck[i] {
			t.Error("Error not obtain", genre, "but obtain", genresCheck[i])
		} else {
			t.Log("Alternative title", i, "[OK]")
		}
	}
}

func TestManga_GetAuthor(t *testing.T) {
	var authorsCheck = []string{
		"Saburouta",
	}

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetAuthors()
	if err != nil {
		t.Error("Error to get authors")
	}

	for i, auth := range m.Authors {
		if auth != authorsCheck[i] {
			t.Error("Error not obtain", auth, "but obtain", authorsCheck[i])
		} else {
			t.Log("Author", i, "[OK]")
		}
	}
}

func TestManga_GetArtists(t *testing.T) {
	var artistCheck = []string{
		"Saburouta",
	}

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetAuthors()
	if err != nil {
		t.Error("Error to get artist")
	}

	for i, artist := range m.Authors {
		if artist != artistCheck[i] {
			t.Error("Error not obtain", artist, "but obtain", artistCheck[i])
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

	if m.Type != Manga_type {
		t.Error("Error not obtain", Manga_type, "but obtain", m.Type)
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

	if m.State != Releasing {
		t.Error("Error not obtain", Releasing, "but obtain", m.State)
	} else {
		t.Log("State [OK]")
	}
}

func TestManga_GetVisual(t *testing.T) {
	const visual = 4215

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetVisual()
	if err != nil {
		t.Error("Error to get visual")
	}

	if m.Visual != visual+1 {
		t.Error("Error not obtain", visual+1, "but obtain", m.Visual)
	} else {
		t.Log("Visual [OK]")
	}
}

func TestManga_GetYearsStart(t *testing.T) {
	const year = "2018"

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetYearsStart()
	if err != nil {
		t.Error("Error to get year start")
	}

	if m.YearsStart != year {
		t.Error("Error not obtain", year, "but obtain", m.YearsStart)
	} else {
		t.Log("Year start [OK]")
	}
}

func TestManga_GetFansub(t *testing.T) {
	var fansub = Fansub{
		Name: "Phoenix Scans",
		Url:  "https://www.phantomreader.com/",
	}

	m, err := NewManga(link)
	if err != nil {
		t.Error("Error to create object")
	}

	err = m.GetFansub()
	if err != nil {
		t.Error("Error to get fansub")
	}

	if m.Fansub.Name != fansub.Name {
		t.Error("Error not obtain", fansub.Name, "but obtain", m.Fansub.Name)
	} else {
		t.Log("Fansub name [OK]")
	}

	if m.Fansub.Url != fansub.Url {
		t.Error("Error not obtain", fansub.Url, "but obtain", m.Fansub.Url)
	} else {
		t.Log("Fansub url [OK]")
	}
}
