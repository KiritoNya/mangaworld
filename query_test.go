package mangaworld

import (
	"testing"
)

func TestQuery_SetMangaName(t *testing.T) {

	result := "yagate+kimi+ni+naru"

	q := NewQuery()
	q.SetMangaName("yagate kimi ni naru")

	if q.MangaName.Val == result {
		t.Log("Query set name [OK]")
	} else {
		t.Error("Error to generate query, not obtain", result, "but obtain", q.MangaName.Val)
	}
}

func TestQuery_SetArtist(t *testing.T) {
	result := "Yuki+Kaori"

	q := NewQuery()
	q.SetArtists([]string{"Yuki Kaori"})

	if q.Artist.Val[0] == result {
		t.Log("Query set artist [OK]")
	} else {
		t.Error("Error to generate query, not obtain", result, "but obtain", q.Artist.Val[0])
	}
}

func TestQuery_SetAuthors(t *testing.T) {
	result := "Yuki+Kaori"

	q := NewQuery()
	q.SetAuthors([]string{"Yuki Kaori"})

	if q.Author.Val[0] == result {
		t.Log("Query set author [OK]")
	} else {
		t.Error("Error to generate query, not obtain", result, "but obtain", q.Author.Val[0])
	}
}

func TestQuery_SetYears(t *testing.T) {
	result := "1995"

	q := NewQuery()
	q.SetYears([]string{"1995", "5555", "kkdkk"})

	for _, year := range q.Year.Val {
		if year == result {
			t.Log("Query set year [OK]")
		} else {
			t.Error("Error to generate query, not obtain", result, "but obtain", year)
		}
	}
}

func TestQuery_CreateQuery(t *testing.T) {

	var result = "https://www.mangaworld.cc/archive?keyword=citrus&genre=Yuri&genre=Shoujo-Ai&type=Manga&state=completed&year=2012&sort=a-z"

	q := NewQuery()

	q.SetMangaName("citrus")
	q.SetGenres([]Genre{Yuri, Shoujo_ai})
	q.SetMangaTypes([]Type{Manga_type})
	q.SetStatus([]State{Finish})
	q.SetYears([]string{"2012"})
	query := q.createQuery()

	if query == result {
		t.Log("Query generetor [OK]")
	} else {
		t.Error("Error to generate query, not obtain", result, "but obtain", query)
	}
}

func TestQuery_Do(t *testing.T) {
	var result = []Manga{
		{Url: "https://www.mangaworld.cc/manga/395/citrus"},
		{Url: "https://www.mangaworld.cc/manga/1876/citrus-1"},
	}

	q := NewQuery()
	q.SetMangaName("citrus")
	mangas, err := q.Do()
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
