package animeworld

import (
	"testing"
	"time"
)

var testSeason = Season{
	Name:          "Violet Evergarden",
	Category:      "",
	Audio:         "",
	RelaseData:    time.Time{},
	ReleaseSeason: ReleasingSeason{},
	Vote:          0,
	Duration:      0,
	TotEpisode:    0,
	Visuals:       0,
	Trama:         "",
	Keywords:      nil,
	Relations:     nil,
	Suggestions:   nil,
	Episodes:      nil,
	MyAnimeList:   "",
	Anilist:       "https://anilist.co/anime/21827",
	Youtube:       "",
	node:          nil,
}

func TestNewSeason(t *testing.T) {
	_, err := NewSeason("https://www.animeworld.tv/play/violet-evergarden.Bqh2M")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log("[OK]")
}

func TestSeason_GetName(t *testing.T) {

	s, err := NewSeason("https://www.animeworld.tv/play/violet-evergarden.Bqh2M")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = s.GetName()
	if err != nil {
		t.Fatal("Error to get name of season")
	}

	if s.Name != testSeason.Name {
		t.Fatal("Not obtain", testSeason.Name, "but obtain", s.Name)
	}
	t.Log("[OK] Season name")

}

func TestSeason_GetAnilist(t *testing.T) {

	season, err := NewSeason("https://www.animeworld.tv/play/violet-evergarden.Bqh2M")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = season.GetAnilist()
	if err != nil {
		t.Error(err.Error())
	}

	if season.Anilist != testSeason.Anilist {
		t.Error("Not obtain ", testSeason.Anilist, " but obtain ", season.Anilist)
	}

	t.Log("[OK] Anilist")

}

func TestSeason_GetEpisodes(t *testing.T) {

	season, err := NewSeason("https://www.animeworld.tv/play/violet-evergarden.Bqh2M")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = season.GetEpisodes()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("[OK] Episodes")
}
