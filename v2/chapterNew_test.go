package v2_test

import (
	v2 "github.com/KiritoNya/mangaworld/v2"
	"golang.org/x/net/html"
	"os"
	"testing"
)

const chapterNewFile string = "./tests/chapterNew/chapterNew.html"

func TestNewChapterNew(t *testing.T) {
	node, err := getHtmlFromFile(chapterNewFile)
	if err != nil {
		t.Fatal(err)
	}

	_, err = v2.NewChapterNew(node)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Constructor [OK]")
}

func TestChapterNew_GetManga(t *testing.T) {
	node, err := getHtmlFromFile(chapterNewFile)
	if err != nil {
		t.Fatal(err)
	}

	cn, err := v2.NewChapterNew(node)
	if err != nil {
		t.Fatal(err)
	}

	err = cn.GetManga()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Manga:", cn.Manga)
}

func TestChapterNew_GetChapters(t *testing.T) {
	node, err := getHtmlFromFile(chapterNewFile)
	if err != nil {
		t.Fatal(err)
	}

	cn, err := v2.NewChapterNew(node)
	if err != nil {
		t.Fatal(err)
	}

	err = cn.GetChapters()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Chapters:", cn.Chapters)
}

func getHtmlFromFile(file string) (*html.Node, error) {
	f, err := os.Open(chapterNewFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	content, err := html.Parse(f)
	if err != nil {
		return nil, err
	}

	return content, nil
}