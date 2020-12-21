package mangaworld

import (
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
	"net/http"
	"testing"
)

var testChapterNew = ChapterNew{
	MangaNew: Manga{
		Url: "https://www.mangaworld.cc/manga/1984/creepy-cat/", //It changes constantly
	},
	Chapters: []Chapter{
		{Url: "https://www.mangaworld.cc/manga/1984/creepy-cat/read/5fe0df428d9b52437fffc98c"},
		{Url: "https://www.mangaworld.cc/manga/1984/creepy-cat/read/5fdf37356fe15b440da486de"},
		{Url: "https://www.mangaworld.cc/manga/1984/creepy-cat/read/5fde5f001ff3df43a02fe365"},
	},
}

func TestNewChapterNew(t *testing.T) {
	resp, err := http.Get(UrlSite)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	grid, err := htmlutils.QuerySelector(node, "div", "class", "comics-grid")
	if err != nil {
		t.Fatal(err)
	}

	chDivs, err := htmlutils.QuerySelector(grid[0], "div", "class", "entry")
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewChapterNew(chDivs[0])
	if err != nil {
		t.Error("Error to create new chapter object: ", err)
	}

	t.Log("Object [OK]")
}

func TestChapterNew_GetManga(t *testing.T) {
	resp, err := http.Get(UrlSite)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	grid, err := htmlutils.QuerySelector(node, "div", "class", "comics-grid")
	if err != nil {
		t.Fatal(err)
	}

	chDivs, err := htmlutils.QuerySelector(grid[0], "div", "class", "entry")
	if err != nil {
		t.Fatal(err)
	}

	cn, err := NewChapterNew(chDivs[0])
	if err != nil {
		t.Error("Error to create new chapter object: ", err)
	}

	err = cn.GetManga()
	if err != nil {
		t.Error("Error to get manga of chapter new: ", err)
	}

	if cn.MangaNew.Url != testChapterNew.MangaNew.Url {
		t.Error("Error not obtain", testChapterNew.MangaNew.Url, "but obtain", cn.MangaNew.Url)
	} else {
		t.Log("Manga link chapter new [OK]")
	}
}

func TestChapterNew_GetChapter(t *testing.T) {
	resp, err := http.Get(UrlSite)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	grid, err := htmlutils.QuerySelector(node, "div", "class", "comics-grid")
	if err != nil {
		t.Fatal(err)
	}

	chDivs, err := htmlutils.QuerySelector(grid[0], "div", "class", "entry")
	if err != nil {
		t.Fatal(err)
	}

	cn, err := NewChapterNew(chDivs[0])
	if err != nil {
		t.Error("Error to create new chapter object: ", err)
	}

	err = cn.GetChapter()
	if err != nil {
		t.Error("Error to get chapter of chapter new: ", err)
	}

	for i, chapter := range cn.Chapters {
		if chapter.Url != testChapterNew.Chapters[i].Url {
			t.Error("Error not obtain", testChapterNew.Chapters[i].Url, "but obtain", chapter.Url)
		} else {
			t.Log("Chapter link chapter new", i, "[OK]")
		}
	}
}
