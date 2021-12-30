package v2_test

import (
	v2 "github.com/KiritoNya/mangaworld/v2"
	"testing"
)

const trendingPath string = "./tests/trending/trending.html"

func TestNewTrendingManga(t *testing.T) {
	node, err := getHtmlFromFile(trendingPath)
	if err != nil {
		t.Fatal(err)
	}

	_, err = v2.NewTrendingManga(node)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Constructor [OK]")
}

func TestTrending_GetManga(t *testing.T) {
	node, err := getHtmlFromFile(trendingPath)
	if err != nil {
		t.Fatal(err)
	}

	tm, err := v2.NewTrendingManga(node)
	if err != nil {
		t.Fatal(err)
	}

	err = tm.GetManga()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Trending Manga:", tm.Manga)
}

func TestTrending_GetChapterNum(t *testing.T) {
	node, err := getHtmlFromFile(trendingPath)
	if err != nil {
		t.Fatal(err)
	}

	tm, err := v2.NewTrendingManga(node)
	if err != nil {
		t.Fatal(err)
	}

	err = tm.GetChapterNum()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ChapterNum:", tm.ChapterNum)
}