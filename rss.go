package animeworld

import (
	"github.com/mmcdole/gofeed"
	"time"
)

const rssLink string = "https://www.animeworld.tv/rss/episodes"

type RssAnimeworld struct {
	feed *gofeed.Feed
}

//NewRssAnimeworld create an object of type RssAnimeworld.
func NewRssAnimeworld() (RssAnimeworld, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rssLink)
	if err != nil {
		return RssAnimeworld{}, err
	}
	return RssAnimeworld{feed: feed}, nil
}

//GetNumberItems is a method of obtaining the number of items of the RSS file.
func (ra *RssAnimeworld) GetNumberItems() int {
	return len(ra.feed.Items)
}

//GetTitleRss is a method of obtaining the title of the RSS file.
func (ra *RssAnimeworld) GetTitleRss() string {
	return ra.feed.Title
}

//GetLinkRss is a method of obtaining the link of site.
func (ra *RssAnimeworld) GetLinkRss() string {
	return ra.feed.Link
}

//GetDescriptionRss is a method of obtaining the description of rss.
func (ra *RssAnimeworld) GetDescriptionRss() string {
	return ra.feed.Description
}

//GetPubDateRssRss is a method of obtaining the last modify date of site.
func (ra *RssAnimeworld) GetPubDateRss() time.Time {
	return *ra.feed.PublishedParsed
}

//GetCopyrightRss is a method of obtaining the copyright of rss.
func (ra *RssAnimeworld) GetCopyrightRss() string {
	return ra.feed.Copyright
}

//GetLanguageRss is a method of obtaining the language of rss.
func (ra *RssAnimeworld) GetLanguageRss() string {
	return ra.feed.Language
}

//GetItemsTitle is a method that returns a slice of string with the titles of the various items inside.
func (ra *RssAnimeworld) GetItemsTitle() (titles []string) {
	for _, item := range ra.feed.Items {
		titles = append(titles, item.Title)
	}
	return titles
}

//GetItemsDescription is a method that returns a slice of string with the description of the various items inside.
func (ra *RssAnimeworld) GetItemsDescription() (descriptions []string) {
	for _, item := range ra.feed.Items {
		descriptions = append(descriptions, item.Description)
	}
	return descriptions
}

//GetItemsLink is a method that returns a slice of string with the links of the various items inside.
func (ra *RssAnimeworld) GetItemsLinks() (links []string) {
	for _, item := range ra.feed.Items {
		links = append(links, item.Link)
	}
	return links
}

//GetItemsGuid is a method that returns a slice of string with the GUID of the various items inside.
func (ra *RssAnimeworld) GetItemsGuid() (guids []string) {
	for _, item := range ra.feed.Items {
		guids = append(guids, item.GUID)
	}
	return guids
}

//GetItemsPubDate is a method that returns a slice of string with the public dates of the various items inside.
func (ra *RssAnimeworld) GetItemsPubDate() (dates []time.Time) {
	for _, item := range ra.feed.Items {
		dates = append(dates, *item.PublishedParsed)
	}
	return dates
}

//GetItemsEpisodeNumber is a method that returns a slice of string with the episode number of the various items inside.
func (ra *RssAnimeworld) GetItemsEpisodeNumber() (epNumbers []string) {
	for _, item := range ra.feed.Items {
		for _, value := range item.Extensions["episode"]["number"] {
			epNumbers = append(epNumbers, value.Value)
		}
	}
	return epNumbers
}

//GetItemsEpisodeDouble is a method that returns a slice of bool which indicates if it's two episodes or not, of the various items inside.
func (ra *RssAnimeworld) GetItemsEpisodeDouble() (epDoubles []bool) {
	for _, item := range ra.feed.Items {
		for _, value := range item.Extensions["episode"]["double"] {
			if value.Value == "false" {
				epDoubles = append(epDoubles, false)
			} else {
				epDoubles = append(epDoubles, true)
			}
		}
	}
	return epDoubles
}

//GetItemsAnimeName is a method that returns a slice of string with the anime name of the various items inside.
func (ra *RssAnimeworld) GetItemsAnimeName() (animeNames []string) {
	for _, item := range ra.feed.Items {
		for _, value := range item.Extensions["episode"]["name"] {
				animeNames = append(animeNames, value.Value)
		}
	}
	return animeNames
}

//GetItemsJtitle is a method that returns a slice of string with the anime japan title of the various items inside.
func (ra *RssAnimeworld) GetItemsJtitle() (animeJtitles []string) {
	for _, item := range ra.feed.Items {
		for _, value := range item.Extensions["episode"]["jtitle"] {
				animeJtitles = append(animeJtitles, value.Value)
		}
	}
	return animeJtitles
}

//GetItemsAnimeEpisodes is a method that returns a slice of string with the number of max anime episode of the various items inside. If it doesn't exist it is "??".
func (ra *RssAnimeworld) GetItemsAnimeEpisodes() (numEpisodes []string) {
	for _, item := range ra.feed.Items {
		for _, value := range item.Extensions["episode"]["episodes"] {
			numEpisodes = append(numEpisodes, value.Value)
		}
	}
	return numEpisodes
}

//GetItemsAnimeLink is a method that returns a slice of string with the link of anime season of the various items inside.
func (ra *RssAnimeworld) GetItemsAnimeLink() (links []string) {
	for _, item := range ra.feed.Items {
		for _, value := range item.Extensions["episode"]["link"] {
			links = append(links, value.Value)
		}
	}
	return links
}

//GetItemsAnimeDub is a method that returns a slice of bool which indicates if it's dubs or not, of the various items inside.
func (ra *RssAnimeworld) GetItemsAnimeDub() (animeDubs []bool) {
	for _, item := range ra.feed.Items {
		for _, value := range item.Extensions["episode"]["dub"] {
			if value.Value == "false" {
				animeDubs = append(animeDubs, false)
			} else {
				animeDubs = append(animeDubs, true)
			}
		}
	}
	return animeDubs
}

//GetItemsAnimeImage is a method that returns a slice of strings with season anime image link of the various items inside.
func (ra *RssAnimeworld) GetItemsAnimeImage() (linksImage []string) {
	for _, item := range ra.feed.Items {
		for _, value := range item.Extensions["episode"]["image"] {
				linksImage = append(linksImage, value.Value)
		}
	}
	return linksImage
}

//GetItemsAnimeCoverImage is a method that returns a slice of strings with season anime cover image link of the various items inside.
func (ra *RssAnimeworld) GetItemsAnimeCoverImage() (linksImage []string) {
	for _, item := range ra.feed.Items {
		for _, value := range item.Extensions["episode"]["cover"] {
			linksImage = append(linksImage, value.Value)
		}
	}
	return linksImage
}

//GetItemTitle is a method that returns an string with the title of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemTitle(numItem int) string {
	return ra.feed.Items[numItem].Title
}

//GetItemDescription is a method that returns an string with the description of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemDescription(numItem int) string {
	return ra.feed.Items[numItem].Description
}

//GetItemLink is a method that returns an string with the link of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemLink(numItem int) string {
	return ra.feed.Items[numItem].Link
}

//GetItemGuid is a method that returns an string with the guid of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemGuid(numItem int) string {
	return ra.feed.Items[numItem].GUID
}

//GetItemPubDate is a method that returns an string with the link of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemPubDate(numItem int) time.Time {
	return *ra.feed.Items[numItem].PublishedParsed
}

//GetItemEpisodeNumber is a method that returns an string with the episode number of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemEpisodeNumber(numItem int) string {
	for _, value := range ra.feed.Items[numItem].Extensions["episode"]["number"] {
		return value.Value
	}
	return ""
}

//GetItemEpisodeDouble is a method that returns an bool which indicates if it's two episodes or not, of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemEpisodeDouble(numItem int) bool {
	for _, value := range ra.feed.Items[numItem].Extensions["episode"]["double"] {
		if value.Value == "false" {
			return false
		} else {
			return true
		}
	}
	return false
}

//GetItemAnimeName is a method that returns an string with the name of season anime of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemAnimeName(numItem int) string {
	for _, value := range ra.feed.Items[numItem].Extensions["episode"]["name"] {
		return value.Value
	}
	return ""
}

//GetItemJtitle is a method that returns an string with the japan name of season anime of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemJtitle(numItem int) string {
	for _, value := range ra.feed.Items[numItem].Extensions["episode"]["jtitle"] {
		return value.Value
	}
	return ""
}

//GetItemAnimeEpisodes is a method that returns an string with the of max anime season episodes of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemAnimeEpisodes(numItem int) string {
	for _, value := range ra.feed.Items[numItem].Extensions["episode"]["episodes"] {
		return value.Value
	}
	return ""
}

//GetItemAnimeLink is a method that returns an string with the season anime link of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemAnimeLink(numItem int) string {
	for _, value := range ra.feed.Items[numItem].Extensions["episode"]["link"] {
		return value.Value
	}
	return ""
}

//GetItemAnimeDub is a method that returns an bool which indicates if it's dubs or not, of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemAnimeDub(numItem int) bool {
	for _, value := range ra.feed.Items[numItem].Extensions["episode"]["dub"] {
		if value.Value == "false" {
			return false
		} else {
			return true
		}
	}
	return false
}

//GetItemAnimeImage is a method that returns an string with the season anime image link of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemAnimeImage(numItem int) string {
	for _, value := range ra.feed.Items[numItem].Extensions["episode"]["image"] {
		return value.Value
	}
	return ""
}

//GetItemAnimeCoverImage is a method that returns an string with the season anime cover image link of the item defined by the number of items sent as a parameter.
func (ra *RssAnimeworld) GetItemAnimeCoverImage(numItem int) string {
	for _, value := range ra.feed.Items[numItem].Extensions["episode"]["cover"] {
		return value.Value
	}
	return ""
}