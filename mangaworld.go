// Copyright KiritoNya.
// All Rights Reserved.

//Package to download manga and get related information from mangaworld.
package mangaworld

import (
	"fmt"
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
	"time"
)

//UrlSite is base URL of site MangaWorld.
const UrlSite = "https://www.mangaworld.io/"

//UrlSearch is base URL for query search.
const UrlSearch = "https://www.mangaworld.io/archive?"

//ListManga is a type that contain a slice of manga
type ListManga struct {
	Mangas []*Manga
}

type ListChapter struct {
	Chapters []Chapter
}

//NewListManga is an constructor of ListManga object
func NewListManga() *ListManga {
	return &ListManga{}
}

//SearchByName is a query with only the manga name. Add to the object ListManga a slice of Manga.
func (lm *ListManga) SearchByName(name string) error {
	q := NewQuery()
	q.SetMangaName(name)
	mangas, err := q.Do()
	if err != nil {
		return err
	}
	lm.Mangas = mangas
	return nil
}

//SearchByGenre is a query with only the manga genre. Add to the object ListManga a slice of Manga.
func (lm *ListManga) SearchByGenre(genres []Genre) error {
	q := NewQuery()
	q.SetGenres(genres)
	mangas, err := q.Do()
	if err != nil {
		return err
	}
	lm.Mangas = mangas
	return nil
}

//SearchByType is a query with only the manga type. Add to the object ListManga a slice of Manga.
func (lm *ListManga) SearchByType(types []Type) error {
	q := NewQuery()
	q.SetMangaTypes(types)
	mangas, err := q.Do()
	if err != nil {
		return err
	}
	lm.Mangas = mangas
	return nil
}

//SearchByStatus is a query with only the manga status. Add to the object ListManga a slice of Manga.
func (lm *ListManga) SearchByStatus(states []State) error {
	q := NewQuery()
	q.SetStatus(states)
	mangas, err := q.Do()
	if err != nil {
		return err
	}
	lm.Mangas = mangas
	return nil
}

//MonthlyManga add to the object a slice of manga with all the top 10 manga of the month.
func (lm *ListManga) MonthlyManga() error {
	resp, err := http.Get(UrlSite)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	divs, err := htmlutils.QuerySelector(node, "div", "class", "top-wrapper")
	if err != nil {
		return err
	}

	entries, err := htmlutils.QuerySelector(divs[1], "div", "class", "short")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		tagsA, err := htmlutils.GetGeneralTags(entry, "a")
		if err != nil {
			return err
		}

		url, err := htmlutils.GetValueAttr(tagsA[0], "a", "href")
		if err != nil {
			return err
		}

		lm.Mangas = append(lm.Mangas, &Manga{Url: string(url[0])})
	}
	return nil
}

//AddTitle add the title to the object Manga of the manga in the list.
func (lm *ListManga) AddTitles() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetTitle()
		if err != nil {
			return err
		}
		fmt.Println(lm.Mangas[i].Title)
	}
	return nil
}

//AddTitlesAlternatives add the alternative titles to the object Manga of the manga in the list.
func (lm *ListManga) AddTitlesAlternatives() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetAlternativeTitle()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddCoverUrls add the cover url to the object Manga of the manga in the list.
func (lm *ListManga) AddCoverUrls() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetCoverUrl()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddGenres add the genres to the object Manga of the manga in the list.
func (lm *ListManga) AddGenres() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetGenre()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddGenres add the authors to the object Manga of the manga in the list.
func (lm *ListManga) AddAuthors() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetAuthors()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddArtists add the artists to the object Manga of the manga in the list.
func (lm *ListManga) AddArtists() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetArtists()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddTypes add the type to the object Manga of the manga in the list.
func (lm *ListManga) AddTypes() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetType()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddStates add the state to the object Manga of the manga in the list.
func (lm *ListManga) AddStates() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetState()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddPlots add the plot to the object Manga of the manga in the list.
func (lm *ListManga) AddPlots() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetPlot()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddYearsStart add the year start to the object Manga of the manga in the list.
func (lm *ListManga) AddYearsStart() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetYearsStart()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddVolumsNum add the volume num to the object Manga of the manga in the list.
func (lm *ListManga) AddVolumesNum() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetVolumsNum()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddChaptersNum add the chapters num to the object Manga of the manga in the list.
func (lm *ListManga) AddChaptersNum() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetChaptersNum()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddChapters add the chapters to the object Manga of the manga in the list.
func (lm *ListManga) AddChapters(start, end int) (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetChapters(start, end)
		if err != nil {
			return err
		}
	}
	return nil
}

//AddRelations add the relations to the object Manga of the manga in the list.
func (lm *ListManga) AddRelations() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetRelations()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddVisuals add the visual to the object Manga of the manga in the list.
func (lm *ListManga) AddVisuals() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetVisual()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddFansubs add the fansub to the object Manga of the manga in the list.
func (lm *ListManga) AddFansubs() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetFansub()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddAnimeworldUrls add the animeworld url to the object Manga of the manga in the list.
func (lm *ListManga) AddAnimeworldUrls() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetAnimeworldUrl()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddAnilistUrls add the anilist url to the object Manga of the manga in the list.
func (lm *ListManga) AddAnilistUrls() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetAnilistUrl()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddMangaUpdatesUrls add the manga updates url to the object Manga of the manga in the list.
func (lm *ListManga) AddMangaUpdatesUrls() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetMangaUpdatesUrl()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddKeywords add the keywords to the object Manga of the manga in the list.
func (lm *ListManga) AddKeywords() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetKeywords()
		if err != nil {
			return err
		}
	}
	return nil
}

//GetUrls returns a slice of string with the urls of manga in the list.
func (lm *ListManga) GetUrls() (urls []string) {
	for _, manga := range lm.Mangas {
		urls = append(urls, manga.Title)
	}
	return urls
}

//GetTitle returns a slice of string with the titles of manga in the list.
func (lm *ListManga) GetTitles() (titles []string) {
	for _, manga := range lm.Mangas {
		titles = append(titles, manga.Title)
	}
	return titles
}

//GetAlternativeTitles returns a matrix of string with the alternative titles of manga in the list.
func (lm *ListManga) GetAlternativeTitles() (altTitles [][]string) {
	for _, manga := range lm.Mangas {
		altTitles = append(altTitles, manga.TitleAlternative)
	}
	return altTitles
}

//GetCoverUrls returns a slice of string with the cover urls of manga in the list.
func (lm *ListManga) GetCoverUrls() (coverUrls []string) {
	for _, manga := range lm.Mangas {
		coverUrls = append(coverUrls, manga.Url)
	}
	return coverUrls
}

//GetGenres returns a matrix of Genre with the genres of manga in the list.
func (lm *ListManga) GetGenres() (genres [][]Genre) {
	for _, manga := range lm.Mangas {
		genres = append(genres, manga.Genres)
	}
	return genres
}

//GetAuthors returns a matrix of string with the authors of manga in the list.
func (lm *ListManga) GetAuthors() (authors [][]string) {
	for _, manga := range lm.Mangas {
		authors = append(authors, manga.Authors)
	}
	return authors
}

//GetArtists returns a matrix of string with the artists of manga in the list.
func (lm *ListManga) GetArtists() (artists [][]string) {
	for _, manga := range lm.Mangas {
		artists = append(artists, manga.Authors)
	}
	return artists
}

//GetTypes returns a slice of Type with the types of manga in the list.
func (lm *ListManga) GetTypes() (types []Type) {
	for _, manga := range lm.Mangas {
		types = append(types, manga.Type)
	}
	return types
}

//GetStates returns a slice of State with the states of manga in the list.
func (lm *ListManga) GetStates() (states []State) {
	for _, manga := range lm.Mangas {
		states = append(states, manga.State)
	}
	return states
}

//GetPlots returns a slice of string with the plots of manga in the list.
func (lm *ListManga) GetPlots() (plots []string) {
	for _, manga := range lm.Mangas {
		plots = append(plots, manga.Plot)
	}
	return plots
}

//GetYearsStart returns a slice of string with the years start of manga in the list.
func (lm *ListManga) GetYearsStart() (years []string) {
	for _, manga := range lm.Mangas {
		years = append(years, manga.Plot)
	}
	return years
}

//GetVolumsNum returns a slice of int with the number of volums of manga in the list.
func (lm *ListManga) GetVolumsNum() (numVolums []int) {
	for _, manga := range lm.Mangas {
		numVolums = append(numVolums, manga.VolumsNum)
	}
	return numVolums
}

//GetChaptersNum returns a slice of int with the number of chapters of manga in the list.
func (lm *ListManga) GetChaptersNum() (numChapters []int) {
	for _, manga := range lm.Mangas {
		numChapters = append(numChapters, manga.VolumsNum)
	}
	return numChapters
}

//GetChapters returns a matrix of Chapters with the chapters of manga in the list.
func (lm *ListManga) GetChapters() (chapters [][]*Chapter) {
	for _, manga := range lm.Mangas {
		chapters = append(chapters, manga.Chapters)
	}
	return chapters
}

//GetVisuals returns a slice of int with the number of visuals of manga in the list.
func (lm *ListManga) GetVisuals() (visuals []int) {
	for _, manga := range lm.Mangas {
		visuals = append(visuals, manga.Visual)
	}
	return visuals
}

//GetFansubs returns a slice of Fansub with the fansubs of manga in the list.
func (lm *ListManga) GetFansubs() (fansubs []Fansub) {
	for _, manga := range lm.Mangas {
		fansubs = append(fansubs, manga.Fansub)
	}
	return fansubs
}

//GetAnimeworldUrls returns a slice of string with the animeworld urls of manga in the list.
func (lm *ListManga) GetAnimeworldUrls() (animeUrls []string) {
	for _, manga := range lm.Mangas {
		animeUrls = append(animeUrls, manga.AnimeworldUrl)
	}
	return animeUrls
}

//GetAnilistUrls returns a slice of string with the anilist urls of manga in the list.
func (lm *ListManga) GetAnilistUrls() (anilistUrls []string) {
	for _, manga := range lm.Mangas {
		anilistUrls = append(anilistUrls, manga.AnilistUrl)
	}
	return anilistUrls
}

//GetMangaUpdatesUrls returns a slice of string with the manga updates urls of manga in the list.
func (lm *ListManga) GetMangaUpdatesUrls() (mangaUpUrls []string) {
	for _, manga := range lm.Mangas {
		mangaUpUrls = append(mangaUpUrls, manga.MangaUpdatesUrl)
	}
	return mangaUpUrls
}

//GetKeywords returns a matrix of string with the keywords of manga in the list.
func (lm *ListManga) GetKeywords() (keywords [][]string) {
	for _, manga := range lm.Mangas {
		keywords = append(keywords, manga.Keywords)
	}
	return keywords
}

//NewChapterList is an constructor of ListChapter object.
func NewChapterList() *ListChapter {
	return &ListChapter{}
}

//AddNumVolumes add number of volumes of chapters to the Chapter objects.
func (lc *ListChapter) AddNumVolumes() error {
	for i, _ := range lc.Chapters {
		err := lc.Chapters[i].GetVolume()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddNumber add numbers of chapters to the Chapter objects.
func (lc *ListChapter) AddNumber() error {
	for i, _ := range lc.Chapters {
		err := lc.Chapters[i].GetNumber()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddPageNum add number of pages of chapters to the Chapter objects.
func (lc *ListChapter) AddPageNums() error {
	for i, _ := range lc.Chapters {
		err := lc.Chapters[i].GetPageNum()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddVisual add number of visuals of chapters to the Chapter objects.
func (lc *ListChapter) AddVisuals() error {
	for i, _ := range lc.Chapters {
		err := lc.Chapters[i].GetVisual()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddVisualToday add number of daily visuals of chapters to the Chapter objects.
func (lc *ListChapter) AddVisualsToday() error {
	for i, _ := range lc.Chapters {
		err := lc.Chapters[i].GetVisualToday()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddPagesUrls add pages urls of chapters to the Chapter objects.
func (lc *ListChapter) AddPagesUrls() error {
	for i, _ := range lc.Chapters {
		err := lc.Chapters[i].GetPageUrl()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddDateAdd add date of addition of chapters to the Chapter objects.
func (lc *ListChapter) AddDateAdd() error {
	for i, _ := range lc.Chapters {
		err := lc.Chapters[i].GetDateAdd()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddKeywords add keywords of chapters to the Chapter objects.
func (lc *ListChapter) AddKeywords() error {
	for i, _ := range lc.Chapters {
		err := lc.Chapters[i].GetKeywords()
		if err != nil {
			return err
		}
	}
	return nil
}

//GetUrls returns a slice of string with urls of the chapters contained in the ListChapter.
func (lc *ListChapter) GetUrls() (urls []string) {
	for i, _ := range lc.Chapters {
		urls = append(urls, lc.Chapters[i].Url)
	}
	return urls
}

//GetNumVolumes returns a slice of int with number of volumes of the chapters contained in the ListChapter.
func (lc *ListChapter) GetNumVolumes() (numVolumes []int) {
	for i, _ := range lc.Chapters {
		numVolumes = append(numVolumes, lc.Chapters[i].Volume)
	}
	return numVolumes
}

//GetUrls returns a slice of int with numbers of the chapters contained in the ListChapter.
func (lc *ListChapter) GetNumber() (numbers []string) {
	for i, _ := range lc.Chapters {
		numbers = append(numbers, lc.Chapters[i].Number)
	}
	return numbers
}

//GetPageNums returns a slice of int with pages numbers of the chapters contained in the ListChapter.
func (lc *ListChapter) GetPageNums() (pageNumbers []int) {
	for i, _ := range lc.Chapters {
		pageNumbers = append(pageNumbers, lc.Chapters[i].PageNum)
	}
	return pageNumbers
}

//GetVisuals returns a slice of int with visuals of the chapters contained in the ListChapter.
func (lc *ListChapter) GetVisuals() (visuals []int) {
	for i, _ := range lc.Chapters {
		visuals = append(visuals, lc.Chapters[i].Visual)
	}
	return visuals
}

//GetVisualsToday returns a slice of int with daily visuals of the chapters contained in the ListChapter.
func (lc *ListChapter) GetVisualsToday() (dailyVisuals []int) {
	for i, _ := range lc.Chapters {
		dailyVisuals = append(dailyVisuals, lc.Chapters[i].VisualToday)
	}
	return dailyVisuals
}

//GetPagesUrls returns a matrix of string with pages urls of the chapters contained in the ListChapter.
func (lc *ListChapter) GetPagesUrls() (pagesUrls [][]string) {
	for i, _ := range lc.Chapters {
		pagesUrls = append(pagesUrls, lc.Chapters[i].PageUrl)
	}
	return pagesUrls
}

//GetDatesAdd returns a slice of time.Time with the dates add of the chapters contained in the ListChapter.
func (lc *ListChapter) GetDatesAdd() (dates []time.Time) {
	for i, _ := range lc.Chapters {
		dates = append(dates, lc.Chapters[i].DateAdd)
	}
	return dates
}

//GetKeywords returns a matrix of string with the keywords of the chapters contained in the ListChapter.
func (lc *ListChapter) GetKeywords() (keywords [][]string) {
	for i, _ := range lc.Chapters {
		keywords = append(keywords, lc.Chapters[i].KeyWords)
	}
	return keywords
}

//TrendingManga returns a trending slice with all manga trending and relative chapter.
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

//ChaptersNew returns a slice of chapters with the chapters just released.
//It accepts as a parameter the number of new manga you want to get.
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
