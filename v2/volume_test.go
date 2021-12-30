package v2_test

import (
	v2 "github.com/KiritoNya/mangaworld/v2"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"os"
	"testing"
)

const volumePath string = "./tests/volume/Blue Box Volume 3.html"

func TestVolume_GetChapters(t *testing.T) {
	v, err := newVolume()
	if err != nil {
		t.Fatal(err)
	}

	err = v.GetChapters()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Chapters:", v.Chapters)
}

func TestVolume_GetNumber(t *testing.T) {
	v, err := newVolume()
	if err != nil {
		t.Fatal(err)
	}

	err = v.GetNumber()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Number:", v.Number)
}

func newVolume() (*v2.Volume, error) {
	f, err := os.Open(volumePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	htmlNode, err := html.Parse(f)
	if err != nil {
		return nil, err
	}
	doc := goquery.NewDocumentFromNode(htmlNode)
	v := v2.NewVolume(doc.Find(".parse"))

	return v, nil
}
