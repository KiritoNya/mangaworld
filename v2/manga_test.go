package v2_test

import (
	mangaworld "github.com/KiritoNya/mangaworld/v2"
	"testing"
)

var UrlTest string = "https://www.mangaworld.in/manga/2490/blue-box/"

func TestManga_GetTitle(t *testing.T) {
	manga, err := mangaworld.NewManga(UrlTest)
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetTitle()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Title:", manga.Title.Pretty)
}

func TestManga_GetAlternativeTitle(t *testing.T) {
	manga, err := mangaworld.NewManga(UrlTest)
	if err != nil {
		t.Fatal(err)
	}

	titles, err := manga.GetAlternativeTitle()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Titles:", titles)
}

func TestManga_GetCoverUrl(t *testing.T) {
	manga, err := mangaworld.NewManga(UrlTest)
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetCoverUrl()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("CoverUrl:", manga.CoverImage)
}

func TestManga_GetGenres(t *testing.T) {
	manga, err := mangaworld.NewManga(UrlTest)
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetGenres()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Genres:", manga.Genres)
}

func TestManga_GetAuthors(t *testing.T) {
	manga, err := mangaworld.NewManga(UrlTest)
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetAuthors()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Authors:", manga.Staff)
}

func TestManga_GetArtists(t *testing.T) {
	manga, err := mangaworld.NewManga(UrlTest)
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetArtists()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Artist:", manga.Staff)
}

func TestManga_GetType(t *testing.T) {
	manga, err := mangaworld.NewManga(UrlTest)
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetType()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Type:", manga.Type)
}

func TestManga_GetStatus(t *testing.T) {
	manga, err := mangaworld.NewManga(UrlTest)
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetStatus()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Status:", manga.Status)
}

func TestManga_GetReleaseDate(t *testing.T) {
	manga, err := mangaworld.NewManga(UrlTest)
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetReleaseDate()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Release Date:", manga.StartDate)
}

func TestManga_GetFansub(t *testing.T) {
	manga, err := mangaworld.NewManga(UrlTest)
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetFansub()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Fansub:", manga.Fansub)
}

func TestManga_GetAnilist(t *testing.T) {
	manga, err := mangaworld.NewMangaFromFile("./tests/manga/Blue Box.html")
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetAnilist()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Anilist:", manga.ExternalLinks)
}

func TestManga_GetMyAnimeList(t *testing.T) {
	manga, err := mangaworld.NewMangaFromFile("./tests/manga/Blue Box.html")
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetMyAnimeList()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("MAL:", manga.ExternalLinks)
}

func TestManga_GetMangaUpdates(t *testing.T) {
	manga, err := mangaworld.NewMangaFromFile("./tests/manga/Blue Box.html")
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetMangaUpdates()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("MangaUpdates:", manga.ExternalLinks)
}

func TestManga_GetExternalLinks(t *testing.T) {
	manga, err := mangaworld.NewMangaFromFile("./tests/manga/Blue Box.html")
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetExternalLinks()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ExternalLinks:", manga.ExternalLinks)
}

func TestManga_GetPlot(t *testing.T) {
	manga, err := mangaworld.NewMangaFromFile("./tests/manga/Blue Box.html")
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetPlot()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Plot:", manga.Description)
}

func TestManga_GetVolumesNum(t *testing.T) {
	manga, err := mangaworld.NewMangaFromFile("./tests/manga/Blue Box.html")
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetVolumesNum()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("VolumesNum:", manga.NumVolumes)
}

func TestManga_GetChaptersNum(t *testing.T) {
	manga, err := mangaworld.NewMangaFromFile("./tests/manga/Blue Box.html")
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetChaptersNum()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ChaptersNum:", manga.NumChapters)
}

func TestManga_GetVolumes(t *testing.T) {
	manga, err := mangaworld.NewMangaFromFile("./tests/manga/Blue Box.html")
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetVolumes()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Volumes:", manga.Volumes)
}

func TestManga_GetRelations(t *testing.T) {
	manga, err := mangaworld.NewMangaFromFile("./tests/manga/Blue Box.html")
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetRelations()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Relations:")
	for _, rel := range manga.Relations {
		t.Log(rel.Url)
	}
}

func TestManga_GetKeyboards(t *testing.T) {
	manga, err := mangaworld.NewMangaFromFile("./tests/manga/Blue Box.html")
	if err != nil {
		t.Fatal(err)
	}

	err = manga.GetKeyboards()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Keyboards:", manga.Keywords)
}