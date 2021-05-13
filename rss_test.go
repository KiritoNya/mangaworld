package animeworld

import (
	"log"
	"testing"
)

func TestNewRssAnimeworld(t *testing.T) {
	_, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
}

func TestRssAnimeworld_GetTitleRss(t *testing.T) {
	result := "AnimeWorld Feed RSS degli episodi"

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	title := rm.GetTitleRss()

	if title != result {
		t.Error("not obtain", result, "but obtain", title)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetDescriptionRss(t *testing.T) {
	result := "Le feed RSS ufficiali per gli episodi di AnimeWorld."

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	title := rm.GetDescriptionRss()

	if title != result {
		t.Error("not obtain", result, "but obtain", title)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetLinkRss(t *testing.T) {
	result := "https://www.animeworld.tv"

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	title := rm.GetLinkRss()

	if title != result {
		t.Error("not obtain", result, "but obtain", title)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetPubDateRss(t *testing.T) {
	result := "2021-01-03 20:19:36 +0000 UTC"

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	date := rm.GetPubDateRss()

	if date.String() != result {
		t.Error("not obtain", result, "but obtain", date.String())
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetCopyrightRss(t *testing.T) {
	result := "Copyright © animeworld.tv. All Rights Reserved"

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	copy := rm.GetCopyrightRss()

	if copy != result {
		t.Error("not obtain", result, "but obtain", copy)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetLanguageRss(t *testing.T) {
	result := "it"

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	lang := rm.GetLanguageRss()

	if lang != result {
		t.Error("not obtain", result, "but obtain", lang)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetItemsTitle(t *testing.T) {
	result := []string{
		"Uma Musume: Pretty Derby 2 Ep. 1", //Changes constantly
		"Mushoku Tensei: Isekai Ittara Honki Dasu Ep. 1-2",     //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	titles := rm.GetItemsTitle()

	titles = titles[:2]

	for i, title := range titles {
		if title != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", title)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemsDescription(t *testing.T) {
	result := []string{
		"NUOVO EPISODIO: Uma Musume: Pretty Derby 2 Ep. 1", //Changes constantly
		"NUOVO EPISODIO: Mushoku Tensei: Isekai Ittara Honki Dasu Ep. 1-2",     //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	slice := rm.GetItemsDescription()

	slice = slice[:2]

	for i, element := range slice {
		if element != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", element)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemsLinks(t *testing.T) {
	result := []string{
		"https://www.animeworld.tv/play/uma-musume-pretty-derby-2.jeUb5/OrEZso", //Changes constantly
		"https://www.animeworld.tv/play/mushoku-tensei-isekai-ittara-honki-dasu.FwbNc/SKUFKi",     //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	slice := rm.GetItemsLinks()

	slice = slice[:2]

	for i, element := range slice {
		if element != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", element)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemsGuid(t *testing.T) {
	result := []string{
		"53601", //Changes constantly
		"53562",     //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	slice := rm.GetItemsGuid()

	slice = slice[:2]

	for i, element := range slice {
		if element != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", element)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemsPubDate(t *testing.T) {
	result := []string{
		"2021-01-04 17:58:51 +0000 UTC", //Changes constantly
		"2021-01-03 20:19:36 +0000 UTC", //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	slice := rm.GetItemsPubDate()

	slice = slice[:2]

	for i, element := range slice {
		if element.String() != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", element.String())
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemsEpisodeNumber(t *testing.T) {
	result := []string{
		"1",   //Changes constantly
		"1-2", //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	for _, item := range rm.feed.Items {
		for key, value := range item.Extensions["episode"] {
			log.Println("KEY:", key, " VALUE", value)
		}
		log.Println("")
	}
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	slice := rm.GetItemsEpisodeNumber()

	slice = slice[:2]

	for i, element := range slice {
		if element != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", element)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemsEpisodeDouble(t *testing.T) {
	result := []bool{
		false,   //Changes constantly
		true, //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	slice := rm.GetItemsEpisodeDouble()

	slice = slice[:2]

	for i, element := range slice {
		if element != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", element)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemsAnimeName(t *testing.T) {
	result := []string{
		"Uma Musume: Pretty Derby 2",   //Changes constantly
		"Mushoku Tensei: Isekai Ittara Honki Dasu", //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	slice := rm.GetItemsAnimeName()

	slice = slice[:2]

	for i, element := range slice {
		if element != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", element)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemsJtitle(t *testing.T) {
	result := []string{
		"Uma Musume: Pretty Derby Season 2",   //Changes constantly
		"Mushoku Tensei: Isekai Ittara Honki Dasu", //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	slice := rm.GetItemsJtitle()

	slice = slice[:2]

	for i, element := range slice {
		if element != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", element)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemsAnimeEpisodes(t *testing.T) {
	result := []string{
		"??",   //Changes constantly
		"??", //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	slice := rm.GetItemsAnimeEpisodes()

	slice = slice[:2]

	for i, element := range slice {
		if element != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", element)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemsAnimeDub(t *testing.T) {
	result := []bool{
		false,   //Changes constantly
		false, //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	slice := rm.GetItemsAnimeDub()

	slice = slice[:2]

	for i, element := range slice {
		if element != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", element)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemsAnimeImage(t *testing.T) {
	result := []string{
		"https://img.animeworld.tv/locandine/jeUb5.jpg",   //Changes constantly
		"https://img.animeworld.tv/locandine/FwbNc.jpg", //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	slice := rm.GetItemsAnimeImage()

	slice = slice[:2]

	for i, element := range slice {
		if element != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", element)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemsAnimeCoverImage(t *testing.T) {
	result := []string{
		"https://img.animeworld.tv/copertine/0utVc.jpg",   //Changes constantly
		"https://img.animeworld.tv/copertine/jeUb5.png", //Changes constantly
	}

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}
	slice := rm.GetItemsAnimeCoverImage()

	slice = slice[:2]

	for i, element := range slice {
		if element != result[i] {
			t.Fatal("not obtain", result[i], "but obtain", element)
		} else {
			t.Log("item", i, "[OK]")
		}
	}
}

func TestRssAnimeworld_GetItemTitle(t *testing.T) {
	result := "That is the Bottleneck Ep. 12" //Changes continuously

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	element := rm.GetItemTitle(0) //last items

	if element != result {
		t.Error("not obtain", result, "but obtain", element)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetItemDescription(t *testing.T) {
	result := "NUOVO EPISODIO: That is the Bottleneck Ep. 12" //Changes continuously

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	element := rm.GetItemDescription(0) //last items

	if element != result {
		t.Error("not obtain", result, "but obtain", element)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetItemLink(t *testing.T) {
	result := "https://www.animeworld.tv/play/sore-dake-ga-neck.0utVc/dnB2ed" //Changes continuously

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	element := rm.GetItemLink(0) //last items

	if element != result {
		t.Error("not obtain", result, "but obtain", element)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetItemGuid(t *testing.T) {
	result := "53602" //Changes continuously

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	element := rm.GetItemGuid(0) //last items

	if element != result {
		t.Error("not obtain", result, "but obtain", element)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetItemPubDate(t *testing.T) {
	result := "2021-01-04 18:56:57 +0000 UTC" //Changes continuously

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	element := rm.GetItemPubDate(0) //last items

	if element.String() != result {
		t.Error("not obtain", result, "but obtain", element.String())
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetEpisodeNumber(t *testing.T) {
	result := "13" //Changes continuously

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	element := rm.GetItemEpisodeNumber(0) //last items

	if element != result {
		t.Error("not obtain", result, "but obtain", element)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetItemEpisodeDouble(t *testing.T) {
	result := false //Changes continuously

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	element := rm.GetItemEpisodeDouble(0) //last items

	if element != result {
		t.Error("not obtain", result, "but obtain", element)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetItemJtitle(t *testing.T) {
	result := "Mr. Osomatsu 3" //Changes continuously

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	element := rm.GetItemJtitle(0) //last items

	if element != result {
		t.Error("not obtain", result, "but obtain", element)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetItemAnimeEpisodes(t *testing.T) {
	result := "25" //Changes continuously

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	element := rm.GetItemAnimeEpisodes(0) //last items

	if element != result {
		t.Error("not obtain", result, "but obtain", element)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetItemAnimeDub(t *testing.T) {
	result := false //Changes continuously

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	element := rm.GetItemAnimeDub(0) //last items

	if element != result {
		t.Error("not obtain", result, "but obtain", element)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetItemAnimeImage(t *testing.T) {
	result := "https://img.animeworld.tv/locandine/TOTFF.jpg" //Changes continuously

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	element := rm.GetItemAnimeImage(0) //last items

	if element != result {
		t.Error("not obtain", result, "but obtain", element)
	} else {
		t.Log("[OK]")
	}
}

func TestRssAnimeworld_GetItemAnimeCoverImage(t *testing.T) {
	result := "https://img.animeworld.tv/copertine/TOTFF.jpg?000" //Changes continuously

	rm, err := NewRssAnimeworld()
	if err != nil {
		t.Fatal("Error to create object RssMagnet:", err)
	}

	element := rm.GetItemAnimeCoverImage(0) //last items

	if element != result {
		t.Error("not obtain", result, "but obtain", element)
	} else {
		t.Log("[OK]")
	}
}
