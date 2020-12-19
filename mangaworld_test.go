package mangaworld

import "testing"

func TestSearchByName(t *testing.T) {

	var result = []Manga{
		{Url: "https://www.mangaworld.cc/manga/395/citrus"},
		{Url: "https://www.mangaworld.cc/manga/1876/citrus-1"},
	}

	mangas, err := SearchByName("citrus")
	if err != nil {
		t.Fatal("Error to search manga: ", err)
	}

	for i, manga := range mangas {
		if manga.Url != result[i].Url {
			t.Error("Error not obtain", result[i].Url, "but obtain", manga.Url)
		} else {
			t.Log("Search manga [OK]")
		}
	}
}

func TestSearchByGenre(t *testing.T) {

	var result = []Manga{
		{Url: "https://www.mangaworld.cc/manga/734/rain"},
		{Url: "https://www.mangaworld.cc/manga/817/6-hours-and-5-minutes"},
	}

	mangas, err := SearchByGenre("citrus")
	if err != nil {
		t.Fatal("Error to search manga: ", err)
	}

	mangas = mangas[:2]

	for i, manga := range mangas {
		if manga.Url != result[i].Url {
			t.Error("Error not obtain", result[i].Url, "but obtain", manga.Url)
		} else {
			t.Log("Search manga [OK]")
		}
	}
}
