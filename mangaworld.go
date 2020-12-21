package mangaworld

import (
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
)

const UrlSite = "https://www.mangaworld.cc/"
const UrlSearch = "https://www.mangaworld.cc/archive?"

func SearchByName(name string) (manga []Manga, err error) {
	q := NewQuery()
	q.SetMangaName(name)
	return q.Do()
}

func SearchByGenre(genres []Genre) (manga []Manga, err error) {
	q := NewQuery()
	q.SetGenres(genres)
	return q.Do()
}

func SearchByType(types []Type) (manga []Manga, err error) {
	q := NewQuery()
	q.SetMangaTypes(types)
	return q.Do()
}

func SearchByStatus(states []State) (manga []Manga, err error) {
	q := NewQuery()
	q.SetStatus(states)
	return q.Do()
}

func TrendingManga() (mangaTrend []Trending, err error) {
	resp, err := http.Get(UrlSite)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	track, err := htmlutils.QuerySelector(node, "div", "id", "chapters-slide")
	if err != nil {
		return nil, err
	}

	divs, err := htmlutils.QuerySelector(track[1], "div", "class", "entry vertical")
	if err != nil {
		return nil, err
	}

	for _, div := range divs {
		tm, err := NewTrendingManga(div)
		if err != nil {
			return nil, err
		}

		err = tm.GetManga()
		if err != nil {
			return nil, err
		}

		err = tm.GetChapter()
		if err != nil {
			return nil, err
		}

		mangaTrend = append(mangaTrend, *tm)
	}
	return mangaTrend, err
}

func MonthlyManga() (mangas []Manga, err error) {
	resp, err := http.Get(UrlSite)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	divs, err := htmlutils.QuerySelector(node, "div", "class", "top-wrapper")
	if err != nil {
		return nil, err
	}

	entries, err := htmlutils.QuerySelector(divs[1], "div", "class", "short")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		tagsA, err := htmlutils.GetGeneralTags(entry, "a")
		if err != nil {
			return nil, err
		}

		url, err := htmlutils.GetValueAttr(tagsA[0], "a", "href")
		if err != nil {
			return nil, err
		}

		mangas = append(mangas, Manga{Url: string(url[0])})
	}
	return mangas, nil
}

func ChaptersNew(num int) (chapters []ChapterNew, err error) {

	for k := 0; k <= (num / 17); k++ {

		resp, err := http.Get(UrlSite + "?page=" + strconv.Itoa(k+1))
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		node, err := html.Parse(resp.Body)
		if err != nil {
			return nil, err
		}

		grid, err := htmlutils.QuerySelector(node, "div", "class", "comics-grid")
		if err != nil {
			return nil, err
		}

		chDivs, err := htmlutils.QuerySelector(grid[0], "div", "class", "entry")
		if err != nil {
			return nil, err
		}

		for i := 0; i < num; i++ {
			cn, err := NewChapterNew(chDivs[i])
			if err != nil {
				return nil, err
			}

			err = cn.GetChapter()
			if err != nil {
				return nil, err
			}

			err = cn.GetManga()
			if err != nil {
				return nil, err
			}

			chapters = append(chapters, *cn)
		}
	}
	return nil, nil
}
