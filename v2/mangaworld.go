package v2

import (
	"errors"
	"fmt"
	"github.com/KiritoNya/gotaku"
	"github.com/KiritoNya/gotaku/manga"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//UrlSite is base URL of site MangaWorld.
const UrlSite = "https://www.mangaworld.in/"

//UrlSearch is base URL for query search.
const UrlSearch = "https://www.mangaworld.in/archive?"

//ClientHttp for the http requests. Default is http.DefaultClient.
var ClientHttp *http.Client

func init() {
	ClientHttp = http.DefaultClient
}

//ListManga is a type that contain a slice of manga
type ListManga struct {
	Mangas []*Manga
}

//NewListManga is an constructor of ListManga object
func NewListManga() *ListManga {
	return &ListManga{}
}

//ListChapter is a type that
type ListChapter struct {
	Chapters []Chapter
}

//MonthlyManga add to the object a slice of manga with all the top 10 manga of the month.
func (lm *ListManga) MonthlyManga() error {
	node, err := doRequest(UrlSite)
	if err != nil {
		return err
	}

	var entryErrors []string
	node.Find("div.top-wrapper").Find("div.short").EachWithBreak(func(i int, entry *goquery.Selection) bool {
		url, found := entry.Find("a").Attr("href")
		if !found {
			entryErrors = append(entryErrors, "Error with entry ", strconv.Itoa(i), ": url not found")
			return false
		}

		m, err := NewManga(url)
		if err != nil {
			entryErrors = append(entryErrors, err.Error())
			return false
		}

		lm.Mangas = append(lm.Mangas, m)
		return true
	})

	// Check errors
	if entryErrors != nil {
		return errors.New(strings.Join(entryErrors, "\n"))
	}

	return nil
}

//AddTitles is a funtion that adds the title to the object Manga of the manga in the list.
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

//AddCoverUrls is a function that adds the cover url to the object Manga of the manga in the list.
func (lm *ListManga) AddCoverUrls() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetCoverUrl()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddGenres is a function that adds the genres to the object Manga of the manga in the list.
func (lm *ListManga) AddGenres() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetGenres()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddAuthors is a function that adds the authors to the object Manga of the manga in the list.
func (lm *ListManga) AddAuthors() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetAuthors()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddArtists is a function that adds the artists to the object Manga of the manga in the list.
func (lm *ListManga) AddArtists() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetArtists()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddTypes is a function that adds the type to the object Manga of the manga in the list.
func (lm *ListManga) AddTypes() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetType()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddStates is a function that adds the status to the object Manga of the manga in the list.
func (lm *ListManga) AddStates() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetStatus()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddPlots is a function that adds the plot to the object Manga of the manga in the list.
func (lm *ListManga) AddPlots() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetPlot()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddReleaseDates is a function that adds the year start to the object Manga of the manga in the list.
func (lm *ListManga) AddReleaseDates() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetReleaseDate()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddVolumesNum is a function that adds the volume num to the object Manga of the manga in the list.
func (lm *ListManga) AddVolumesNum() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetVolumesNum()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddChaptersNum is a function that adds the chapters num to the object Manga of the manga in the list.
func (lm *ListManga) AddChaptersNum() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetChaptersNum()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddChapters is a function that adds the chapters to the object Manga of the manga in the list.
func (lm *ListManga) AddChapters(start, end int) (err error) {
	for i, _ := range lm.Mangas {
		if lm.Mangas[i].Volumes == nil {
			err := lm.Mangas[i].GetVolumes()
			if err != nil {
				return err
			}
		}
		for k, _ := range lm.Mangas[i].Volumes {
			err := lm.Mangas[i].Volumes[k].GetChapters()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//AddRelations is a function that adds the relations to the object Manga of the manga in the list.
func (lm *ListManga) AddRelations() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetRelations()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddFansubs is a function that adds the fansub to the object Manga of the manga in the list.
func (lm *ListManga) AddFansubs() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetFansub()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddAnilistUrls is a function that adds the anilist url to the object Manga of the manga in the list.
func (lm *ListManga) AddAnilistUrls() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetAnilist()
		if err != nil {
			return err
		}
	}
	return nil
}

//AddMangaUpdatesUrls is a function that adds the manga updates url to the object Manga of the manga in the list.
func (lm *ListManga) AddMangaUpdatesUrls() (err error) {
	for i, _ := range lm.Mangas {
		err = lm.Mangas[i].GetMangaUpdates()
		if err != nil {
			return err
		}
	}
	return nil
}

//GetUrls is a function that returns a slice of string with the urls of manga in the list.
func (lm *ListManga) GetUrls() (urls []string) {
	for _, manga := range lm.Mangas {
		urls = append(urls, manga.Url)
	}
	return urls
}

//GetTitles is a function that returns a slice of string with the titles of manga in the list.
func (lm *ListManga) GetTitles() (titles []string) {
	for _, manga := range lm.Mangas {
		titles = append(titles, manga.Title.Pretty)
	}
	return titles
}

//GetCoverUrls is a function that returns a slice of string with the cover urls of manga in the list.
func (lm *ListManga) GetCoverUrls() (coverUrls []string) {
	for _, manga := range lm.Mangas {
		coverUrls = append(coverUrls, manga.CoverImage.ExtraLarge.Url)
	}
	return coverUrls
}

//GetGenres is a function that returns a matrix of Genre with the genres of manga in the list.
func (lm *ListManga) GetGenres() (genres [][]Genre) {
	for _, manga := range lm.Mangas {
		var mangaGenres []Genre
		for _, genre := range manga.Genres {
			mangaGenres = append(mangaGenres, Genre(genre.Slug))
		}
		genres = append(genres, mangaGenres)
	}
	return genres
}

//GetAuthors is a function that returns a matrix of string with the authors of manga in the list.
func (lm *ListManga) GetAuthors() (authors [][]string) {
	for _, manga := range lm.Mangas {
		var mangaAuthors []string
		for _, staff := range manga.Staff {
			if staff.Role == "Author" {
				if staff.Name.Last != "" {
					mangaAuthors = append(mangaAuthors, staff.Name.First + " " + staff.Name.Last)
				} else {
					mangaAuthors = append(mangaAuthors, staff.Name.First)
				}
			}
		}
		authors = append(authors, mangaAuthors)
	}
	return authors
}

//GetArtists is a function that returns a matrix of string with the artists of manga in the list.
func (lm *ListManga) GetArtists() (artists [][]string) {
	for _, manga := range lm.Mangas {
		var mangaArtists []string
		for _, staff := range manga.Staff {
			if staff.Role == "Artist" {
				if staff.Name.Last != "" {
					mangaArtists = append(mangaArtists, staff.Name.First + " " + staff.Name.Last)
				} else {
					mangaArtists = append(mangaArtists, staff.Name.First)
				}
			}
		}
		artists = append(artists, mangaArtists)
	}
	return artists
}

//GetTypes is a function that returns a slice of Type with the types of manga in the list.
func (lm *ListManga) GetTypes() (types []gotaku.MediaType) {
	for _, manga := range lm.Mangas {
		types = append(types, manga.Type)
	}
	return types
}

//GetStates is a function that returns a slice of State with the states of manga in the list.
func (lm *ListManga) GetStates() (states []gotaku.Status) {
	for _, manga := range lm.Mangas {
		states = append(states, manga.Status)
	}
	return states
}

//GetPlots returns a slice of string with the plots of manga in the list.
func (lm *ListManga) GetPlots() (plots []string) {
	for _, manga := range lm.Mangas {
		plots = append(plots, manga.Description)
	}
	return plots
}

//GetReleaseDates is a function that returns a slice of time.Time with the release date of manga in the list.
func (lm *ListManga) GetReleaseDates() (years []time.Time) {
	for _, manga := range lm.Mangas {
		years = append(years, manga.StartDate)
	}
	return years
}

//GetVolumesNum is a function that returns a slice of int with the number of volums of manga in the list.
func (lm *ListManga) GetVolumesNum() (numVolums []int) {
	for _, manga := range lm.Mangas {
		numVolums = append(numVolums, manga.NumVolumes)
	}
	return numVolums
}

//GetChaptersNum is a function that returns a slice of int with the number of chapters of manga in the list.
func (lm *ListManga) GetChaptersNum() (numChapters []int) {
	for _, manga := range lm.Mangas {
		numChapters = append(numChapters, manga.NumChapters)
	}
	return numChapters
}

//GetFansubs is a function that returns a slice of Fansub with the fansubs of manga in the list.
func (lm *ListManga) GetFansubs() (fansubs []manga.Fansub) {
	for _, manga := range lm.Mangas {
		fansubs = append(fansubs, manga.Fansub)
	}
	return fansubs
}

//GetAnilistUrls is a function that returns a slice of string with the anilist urls of manga in the list.
func (lm *ListManga) GetAnilistUrls() (anilistUrls []string) {
	for _, manga := range lm.Mangas {
		for _, link := range manga.ExternalLinks {
			if link.Site == "anilist" {
				anilistUrls = append(anilistUrls, link.Url)
			}
		}
	}
	return anilistUrls
}

//GetMangaUpdatesUrls is a function that returns a slice of string with the manga updates urls of manga in the list.
func (lm *ListManga) GetMangaUpdatesUrls() (mangaUpUrls []string) {
	for _, manga := range lm.Mangas {
		for _, link := range manga.ExternalLinks {
			if link.Site == "mangaupdatesurl" {
				mangaUpUrls = append(mangaUpUrls, link.Url)
			}
		}
	}
	return mangaUpUrls
}

//TrendingManga is a function that returns a trending slice with all manga trending and relative chapter.
func TrendingManga() (mangaTrend []Trending, err error) {
	node, err := doRequest(UrlSite)
	if err != nil {
		return nil, err
	}

	var entryErrors []string
	node.Find("#chapters-slide").Find("div.entry.vertical").EachWithBreak(func(i int, entry *goquery.Selection) bool {
		tm, err := NewTrendingManga(entry.Nodes[0])
		if err != nil {
			entryErrors = append(entryErrors, err.Error())
			return false
		}

		err = tm.GetManga()
		if err != nil {
			entryErrors = append(entryErrors, err.Error())
			return false
		}

		err = tm.GetChapterNum()
		if err != nil {
			entryErrors = append(entryErrors, err.Error())
			return false
		}

		mangaTrend = append(mangaTrend, *tm)
		return true
	})

	// Check error
	if entryErrors != nil {
		return nil, errors.New(strings.Join(entryErrors, "\n"))
	}

	return mangaTrend, err
}

//ChaptersNew is a function that returns a slice of chapters with the chapters just released.
//It accepts as a parameter the number of new manga you want to get.
func ChaptersNew(num int) (chapters []ChapterNew, err error) {
	for k := 0; k <= (num / 17); k++ {
		//Make request
		node, err := doRequest(UrlSite + "?page=" + strconv.Itoa(k+1))
		if err != nil {
			return nil, err
		}

		var entryErrors []string
		node.Find("div.comics-grid").Find("div.entry").EachWithBreak(func(i int, entry *goquery.Selection) bool {
			cn, err := NewChapterNew(entry.Nodes[0])
			if err != nil {
				entryErrors = append(entryErrors, err.Error())
				return false
			}

			err = cn.GetChapters()
			if err != nil {
				entryErrors = append(entryErrors, err.Error())
				return false
			}

			err = cn.GetManga()
			if err != nil {
				entryErrors = append(entryErrors, err.Error())
				return false
			}

			chapters = append(chapters, *cn)
			return true
		})


		// Check errors
		if entryErrors != nil {
			return nil, errors.New(strings.Join(entryErrors, "\n"))
		}

	}

	return chapters, nil
}
