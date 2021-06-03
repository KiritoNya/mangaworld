package mangaworld

import (
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
	"net/http"
	"testing"
)

var testTrending = Trending{
	Manga: Manga{
		Url: "https://www.mangaworld.cc/manga/1738/solo-leveling/",
	},
	Chapter: Chapter{
		Number: "131",
	},
}

func TestNewTrendingManga(t *testing.T) {
	resp, err := http.Get(UrlSite)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	track, err := htmlutils.QuerySelector(node, "div", "id", "chapters-slide")
	if err != nil {
		t.Fatal(err)
	}

	divs, err := htmlutils.QuerySelector(track[1], "div", "class", "entry vertical")
	if err != nil {
		t.Error(err)
	}

	_, err = NewTrendingManga(divs[0])
	if err != nil {
		t.Error(err)
	}
}

func TestTrending_GetManga(t *testing.T) {
	resp, err := http.Get(UrlSite)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	track, err := htmlutils.QuerySelector(node, "div", "id", "chapters-slide")
	if err != nil {
		t.Fatal(err)
	}

	divs, err := htmlutils.QuerySelector(track[1], "div", "class", "entry vertical")
	if err != nil {
		t.Error(err)
	}

	tm, err := NewTrendingManga(divs[0])
	if err != nil {
		t.Error(err)
	}

	err = tm.GetManga()
	if err != nil {
		t.Error("Error to get trending manga: ", err)
	}

	if tm.Manga.Url != testTrending.Manga.Url {
		t.Error("Error not obtain", testTrending.Manga.Url, "but obtain", tm.Manga.Url)
	} else {
		t.Log("Trending manga [OK]")
	}
}

func TestTrending_GetChapter(t *testing.T) {
	resp, err := http.Get(UrlSite)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	track, err := htmlutils.QuerySelector(node, "div", "id", "chapters-slide")
	if err != nil {
		t.Fatal(err)
	}

	divs, err := htmlutils.QuerySelector(track[1], "div", "class", "entry vertical")
	if err != nil {
		t.Error(err)
	}

	tm, err := NewTrendingManga(divs[0])
	if err != nil {
		t.Error(err)
	}

	err = tm.GetChapter()
	if err != nil {
		t.Error("Error to get trending chapter: ", err)
	}

	if tm.Chapter.Number != testTrending.Chapter.Number {
		t.Error("Error not obtain", testTrending.Chapter.Number, "but obtain", tm.Chapter.Number)
	} else {
		t.Log("Trending number chapter [OK]")
	}
}
