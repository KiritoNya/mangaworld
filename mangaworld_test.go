package mangaworld

import "testing"

func TestSearchByGenre(t *testing.T) {
	var result = []Manga{
		{Url: "https://www.mangaworld.cc/manga/734/rain"},
		{Url: "https://www.mangaworld.cc/manga/817/6-hours-and-5-minutes"},
	}

	mangas, err := SearchByGenre([]Genre{Yuri})
	if err != nil {
		t.Fatal("Error to search manga: ", err)
	}

	mangas = mangas[:2]

	for i, manga := range mangas {
		if manga.Url != result[i].Url {
			t.Error("Error not obtain", result[i].Url, "but obtain", manga.Url)
		} else {
			t.Log("Search manga", i, "[OK]")
		}
	}
}

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
			t.Log("Search manga", i, "[OK]")
		}
	}
}

func TestSearchByType(t *testing.T) {
	var result = []Manga{
		{Url: "https://www.mangaworld.cc/manga/426/1-2-cigarette"},
		{Url: "https://www.mangaworld.cc/manga/1151/1-3-romantica"},
	}

	mangas, err := SearchByType([]Type{Manga_type})
	if err != nil {
		t.Fatal("Error to search manga: ", err)
	}

	mangas = mangas[:2]

	for i, manga := range mangas {
		if manga.Url != result[i].Url {
			t.Error("Error not obtain", result[i].Url, "but obtain", manga.Url)
		} else {
			t.Log("Search manga", i, "[OK]")
		}
	}
}

func TestSearchByStatus(t *testing.T) {
	var result = []Manga{
		{Url: "https://www.mangaworld.cc/manga/1983/1-3-sanbun-no-ichi"},
		{Url: "https://www.mangaworld.cc/manga/1895/100-nichigo-ni-kuufuku-de-taore-maid-ni-naru-onna-no-ko"},
	}

	mangas, err := SearchByStatus([]State{Releasing})
	if err != nil {
		t.Fatal("Error to search manga: ", err)
	}

	mangas = mangas[:2]

	for i, manga := range mangas {
		if manga.Url != result[i].Url {
			t.Error("Error not obtain", result[i].Url, "but obtain", manga.Url)
		} else {
			t.Log("Search manga", i, "[OK]")
		}
	}
}
