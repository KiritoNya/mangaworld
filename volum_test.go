package mangaworld

import (
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
	"net/http"
	"testing"
)

const urlVolume = "https://www.mangaworld.cc/manga/1876/citrus-1/"

var testVolum = Volum{
	Name:   "Volume 03",
	Number: 3,
	Chapters: []Chapter{
		{Url: "https://www.mangaworld.cc/manga/1876/citrus-1/read/5fbbfab01c9bb544acdbbac0"},
		{Url: "https://www.mangaworld.cc/manga/1876/citrus-1/read/5fbbfab01c9bb544acdbbac0"},
		{Url: "https://www.mangaworld.cc/manga/1876/citrus-1/read/5fbbfab01c9bb544acdbbabe"},
	},
}

func TestNewVolum(t *testing.T) {

	resp, err := http.Get(urlVolume)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	htmlBody, err := html.Parse(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	divs, err := htmlutils.QuerySelector(htmlBody, "div", "class", "volume-element pl-2")
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewVolum(divs[0])
	if err != nil {
		t.Fatal("Error to create object Volum: ", err)
	}
}

func TestVolum_GetName(t *testing.T) {
	resp, err := http.Get(urlVolume)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	htmlBody, err := html.Parse(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	divs, err := htmlutils.QuerySelector(htmlBody, "div", "class", "volume-element pl-2")
	if err != nil {
		t.Fatal(err)
	}

	v, err := NewVolum(divs[0])
	if err != nil {
		t.Fatal("Error to create object Volum: ", err)
	}

	err = v.GetName()
	if err != nil {
		t.Error("Error to get Volume name: ", err)
	}

	if v.Name != testVolum.Name {
		t.Error("Error not obtain", testVolum.Name, "but obtain", v.Name)
	} else {
		t.Log("Volum name [OK]")
	}
}

func TestVolum_GetNumber(t *testing.T) {
	resp, err := http.Get(urlVolume)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	htmlBody, err := html.Parse(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	divs, err := htmlutils.QuerySelector(htmlBody, "div", "class", "volume-element pl-2")
	if err != nil {
		t.Fatal(err)
	}

	v, err := NewVolum(divs[0])
	if err != nil {
		t.Fatal("Error to create object Volum: ", err)
	}

	err = v.GetNumber()
	if err != nil {
		t.Error("Error to get Volum number: ", err)
	}

	if v.Number != testVolum.Number {
		t.Error("Error not obtain", testVolum.Number, "but obtain", v.Number)
	} else {
		t.Log("Volum number [OK]")
	}
}

func TestVolum_GetChapters(t *testing.T) {
	resp, err := http.Get(urlVolume)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	htmlBody, err := html.Parse(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	divs, err := htmlutils.QuerySelector(htmlBody, "div", "class", "volume-element pl-2")
	if err != nil {
		t.Fatal(err)
	}

	v, err := NewVolum(divs[0])
	if err != nil {
		t.Fatal("Error to create object Volum: ", err)
	}

	err = v.GetChapters()
	if err != nil {
		t.Error("Error to get Chapters: ", err)
	}

	for i, chapter := range v.Chapters {
		if chapter.Url != testVolum.Chapters[i].Url {
			t.Error("Error not obtain", testVolum.Chapters[i].Url, "but obtain", chapter.Url)
		} else {
			t.Log("Chapter ", i, "[OK]")
		}
	}
}
