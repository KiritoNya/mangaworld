package v2_test

import (
	v2 "github.com/KiritoNya/mangaworld/v2"
	"testing"
)

const pageUrl string = "https://www.mangaworld.in/manga/2490/blue-box/read/61c5f93ed1503504ee9f5b94/1"

func TestNewPage(t *testing.T) {
	_, err := v2.NewPage(pageUrl)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Constructor [OK]")
}

func TestPage_GetNumber(t *testing.T) {
	p, err := v2.NewPage(pageUrl)
	if err != nil {
		t.Fatal(err)
	}

	err = p.GetNumber()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Page Number:", p.Number)
}

func TestPage_GetImage(t *testing.T) {
	p, err := v2.NewPage(pageUrl)
	if err != nil {
		t.Fatal(err)
	}

	err = p.GetImage()
	if err != nil {
		t.Fatal(err)
	}

	p.Image.GetName()
	t.Log("Image:", p.Image)
}
