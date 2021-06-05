package mangaworld

import (
	"testing"
	"time"
)

const urlTest = "https://www.mangaworld.io/manga/1983/1-3-sanbun-no-ichi/read/5fa8c1fa552b377191b08ba2/1"

var testChapter = Chapter{
	Volume:      3,
	Number:      "4",
	PageNum:     22,
	Visual:      1196, //changes constantly
	VisualToday: 33,
	PageUrl: []string{
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/1.jpg",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/2.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/3.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/4.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/5.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/6.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/7.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/8.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/9.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/10.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/11.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/12.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/13.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/14.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/15.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/16.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/17.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/18.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/19.png",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/20.jpg",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/21.jpg",
		"https://cdn.mangaworld.cc/chapters/citrus-5fa4d1817a7a701817de97e1/volume-03-5fbbfab01c9bb544acdbbabd/capitolo-16-5fbbfab01c9bb544acdbbac0/22.jpg",
	},
	DateAdd: time.Date(2020, time.Month(11), 23, 0, 0, 0, 0, time.Now().Location()),
	KeyWords: []string{
		"Citrus+ Capitolo 16 Scan ITA",
		"Citrus+ Capitolo 16 Scan ITALIANE",
		"Citrus+ Capitolo 16 MangaWorld",
		"Citrus+ Capitolo 16 MangaDex",
		"Citrus+ Capitolo 16 MangaEden",
		"Citrus+ Capitolo 16 Leggi online",
		"Citrus+ Capitolo 16 Reader ITA",
		"Citrus+ Capitolo 16 Scan online",
		"Citrus+ Capitolo 16 Read online",
		"Citrus+ Scan ITA",
		"Citrus+ Scan ITALIANE",
		"Citrus+ MangaWorld",
		"Citrus+ MangaDex",
		"Citrus+ MangaEden",
		"Citrus+ SUB ITA",
		"Citrus+ Leggi online",
		"Citrus+ Reader ITA",
		"Citrus+ Scan online",
		"Citrus+ Read online",
		"Citrus+ Online",
		"Citrus+ Manga ITA",
		"Citrus+ Manga Scan",
		"Citrus+ Scan manga online",
		"Citrus+ ITA Scan",
		"MangaWorld Citrus+",
		"Scan ITA Citrus+",
		"Read online Citrus+",
		"Citrus+ MangaWorld",
	},
	resp: nil,
}

func TestChapter_GetVolume(t *testing.T) {
	c, err := NewChapter(urlTest)
	if err != nil {
		t.Fatalf("Error to create object")
	}

	err = c.GetVolume()
	if err != nil {
		t.Error("Error to get chapter volume: ", err)
	}

	if c.Volume != testChapter.Volume {
		t.Error("Error not obtain", testChapter.Volume, "but obtain", c.Volume)
	} else {
		t.Log("Chapter volume [OK]")
	}
}

func TestChapter_GetNumber(t *testing.T) {

	c, err := NewChapter(urlTest)
	if err != nil {
		t.Fatalf("Error to create object")
	}

	err = c.GetNumber()
	if err != nil {
		t.Error("Error to get chapter number: ", err)
	}

	if c.Number != "04" {
		t.Error("Error not obtain", "04", "but obtain", c.Number)
	} else {
		t.Log("Chapter Number [OK]")
	}
}

func TestChapter_GetPageNum(t *testing.T) {

	c, err := NewChapter(urlTest)
	if err != nil {
		t.Fatalf("Error to create object")
	}

	err = c.GetPageNum()
	if err != nil {
		t.Error("Error to get chapter page number: ", err)
	}

	if c.PageNum != testChapter.PageNum {
		t.Error("Error not obtain", testChapter.PageNum, "but obtain", c.PageNum)
	} else {
		t.Log("Chapter Page Number [OK]")
	}
}

func TestChapter_GetVisual(t *testing.T) {
	c, err := NewChapter(urlTest)
	if err != nil {
		t.Fatalf("Error to create object")
	}

	err = c.GetVisual()
	if err != nil {
		t.Error("Error to get visual: ", err)
	}

	if c.Visual != testChapter.Visual+1 {
		t.Error("Error not obtain", testChapter.Visual, "but obtain", c.Visual)
	} else {
		t.Log("Visual [OK]")
	}
}

func TestChapter_GetVisualToday(t *testing.T) {
	c, err := NewChapter(urlTest)
	if err != nil {
		t.Fatalf("Error to create object")
	}

	err = c.GetVisualToday()
	if err != nil {
		t.Error("Error to get visual today: ", err)
	}

	if c.VisualToday != testChapter.VisualToday+1 {
		t.Error("Error not obtain", testChapter.VisualToday, "but obtain", c.VisualToday)
	} else {
		t.Log("Visual Today [OK]")
	}
}

func TestChapter_GetPageUrl(t *testing.T) {
	c, err := NewChapter(urlTest)
	if err != nil {
		t.Fatalf("Error to create object")
	}

	t.Log(c.Url)

	err = c.GetPageNum()
	if err != nil {
		t.Error("Error to get chapter page number: ", err)
	}

	err = c.GetPageUrl()
	if err != nil {
		t.Error("Error to get chapter page url: ", err)
	}

	for i, page := range c.PageUrl {
		if page != testChapter.PageUrl[i] {
			t.Error("Error not obtain", testChapter.PageUrl[i], "but obtain", page)
		} else {
			t.Log("Chapter Page url", i, "[OK]")
		}
	}
}

func TestChapter_GetDateAdd(t *testing.T) {
	c, err := NewChapter(urlTest)
	if err != nil {
		t.Fatalf("Error to create object")
	}

	err = c.GetDateAdd()
	if err != nil {
		t.Error("Error to get chapter date release: ", err)
	}

	if c.DateAdd != testChapter.DateAdd {
		t.Error("Error not obtain", testChapter.DateAdd, "but obtain", c.DateAdd)
	} else {
		t.Log("Chapter release date [OK]")
	}
}

func TestChapter_GetKeywords(t *testing.T) {
	c, err := NewChapter(urlTest)
	if err != nil {
		t.Fatalf("Error to create object")
	}

	err = c.GetKeywords()
	if err != nil {
		t.Error("Error to get chapter keywords: ", err)
	}

	for i, key := range c.PageUrl {
		if key != testChapter.KeyWords[i] {
			t.Error("Error not obtain", testChapter.KeyWords[i], "but obtain", key)
		} else {
			t.Log("Chapter key", i, "[OK]")
		}
	}

}

func TestChapter_Download(t *testing.T) {
	c, err := NewChapter("https://www.mangaworld.cc/manga/1876/citrus-1/read/5fbbfab01c9bb544acdbbaac/1")
	if err != nil {
		t.Error(err)
	}

	c.Download("C:\\Users\\KiritoNya\\Desktop\\Nuova")

}
