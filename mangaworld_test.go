package mangaworld

import (
	"testing"
)

func TestSearchByGenre(t *testing.T) {
	var result = []Manga{
		{Url: "https://www.mangaworld.cc/manga/734/rain"},
		{Url: "https://www.mangaworld.cc/manga/817/6-hours-and-5-minutes"},
	}

	lm := NewListManga()

	err := lm.SearchByGenre([]Genre{Yuri})
	if err != nil {
		t.Fatal("Error to search manga: ", err)
	}

	lm.Mangas = lm.Mangas[:2]

	for i, manga := range lm.Mangas {
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

	lm := NewListManga()

	err := lm.SearchByName("citrus")
	if err != nil {
		t.Fatal("Error to search manga: ", err)
	}

	for i, manga := range lm.Mangas {
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

	lm := NewListManga()

	err := lm.SearchByType([]Type{Manga_type})
	if err != nil {
		t.Fatal("Error to search manga: ", err)
	}

	lm.Mangas = lm.Mangas[:2]

	for i, manga := range lm.Mangas {
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

	lm := NewListManga()

	err := lm.SearchByStatus([]State{Releasing})
	if err != nil {
		t.Fatal("Error to search manga: ", err)
	}

	lm.Mangas = lm.Mangas[:2]

	for i, manga := range lm.Mangas {
		if manga.Url != result[i].Url {
			t.Error("Error not obtain", result[i].Url, "but obtain", manga.Url)
		} else {
			t.Log("Search manga", i, "[OK]")
		}
	}
}

func TestMonthlyManga(t *testing.T) {

	var result = []Manga{
		{Url: "https://www.mangaworld.cc/manga/1972/martial-peak"},
		{Url: "https://www.mangaworld.cc/manga/1929/tales-of-demons-and-gods"},
		{Url: "https://www.mangaworld.cc/manga/1847/boku-no-hero-academia"},
		{Url: "https://www.mangaworld.cc/manga/1738/solo-leveling"},
		{Url: "https://www.mangaworld.cc/manga/1816/shingeki-no-kyojin"},
		{Url: "https://www.mangaworld.cc/manga/1708/one-piece"},
		{Url: "https://www.mangaworld.cc/manga/1674/jujutsu-kaisen"},
		{Url: "https://www.mangaworld.cc/manga/2278/berserk"},
		{Url: "https://www.mangaworld.cc/manga/716/kimetsu-no-yaiba"},
		{Url: "https://www.mangaworld.cc/manga/1747/black-clover"},
	}

	lm := NewListManga()

	err := lm.MonthlyManga()
	if err != nil {
		t.Error("Error to ge monthly manga:", err)
	}

	for i, manga := range lm.Mangas {
		if manga.Url != result[i].Url {
			t.Error("Error not obtain", result[i].Url, "but obtain", manga.Url)
		} else {
			t.Log("Monthly manga", i, "[OK]")
		}
	}
}

func TestChaptersNew(t *testing.T) {
	var result = []ChapterNew{
		{
			MangaNew: Manga{
				Url: "https://www.mangaworld.cc/manga/1984/creepy-cat/", //It changes constantly
			},
			Chapters: []Chapter{
				{Url: "https://www.mangaworld.cc/manga/1984/creepy-cat/read/5fe0df428d9b52437fffc98c"},
				{Url: "https://www.mangaworld.cc/manga/1984/creepy-cat/read/5fdf37356fe15b440da486de"},
				{Url: "https://www.mangaworld.cc/manga/1984/creepy-cat/read/5fde5f001ff3df43a02fe365"},
			},
		},
		{
			MangaNew: Manga{
				Url: "https://www.mangaworld.cc/manga/1661/shuumatsu-no-walkuere/", //It changes constantly
			},
			Chapters: []Chapter{
				{Url: "https://www.mangaworld.cc/manga/1661/shuumatsu-no-walkuere/read/5fe0dd5ff2ab2d434dcf488b/1"},
			},
		},
	}

	cn, err := ChaptersNew(2)
	if err != nil {
		t.Fatal("Error to get chapters new:", err)
	}

	for k, chapterNew := range cn {
		for i, chapter := range chapterNew.Chapters {
			if chapter.Url != result[k].Chapters[i].Url {
				t.Error("Error not obtain", result[k].Chapters[i].Url, "but obtain", chapter.Url)
			} else {
				t.Log("Chapter new", i, "[OK]")
			}
		}
	}
}
