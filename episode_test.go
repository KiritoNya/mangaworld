package animeworld

import (
	"fmt"
	"testing"
)

var testEpisode = Episode{
	Url:          "",
	StreamLinks:  map[string]string {
		"AnimeWorld Server": "https://server17.streamingaw.online/DDL/ANIME/JujutsuKaisenSUBITA/JujutsuKaisen_Ep_01_SUB_ITA.mp4",
		"Streamtape": "https://streamtape.com/e/KLMGkKGlvdF0Rry/JujutsuKaisen_Ep_01_SUB_ITA.mp4",
		"PlayTube": "https://playtube.ws/embed-piato97hkydy/JujutsuKaisen_Ep_01_SUB_ITA.mp4.html",
		"DoodStream": "https://dood.so/e/vnupcz3z02mx",
		"Userload": "https://userload.co/embed/a4f872ff5fd1/JujutsuKaisen_Ep_01_SUB_ITA.mp4",
		"NinjaStream": "https://ninjastream.to/watch/xvrBA1gK3Zb8o/jujutsukaisen-ep-01-sub-ita-mp4",
		"VUP": "https://vupload.com/e/pvlzk1dsidfc",
		"AW (1080p)": "https://player2.animeworld.tv/player?link=https://www.proxyworld1.org/DDL/ANIME/JujutsuKaisenFHD/01/playlist.m3u8",
		"MixDrop": "https://mixdrop.sx/e/9nwv048lcnz7op",
	},
	Number:       []float64{20.0},
	DownloadLink: "",
	htmlPage:     nil,
}

func TestNewEpisode(t *testing.T) {
	_, err := NewEpisode("https://www.animeworld.tv/play/jujutsu-kaisen.L4vvx/L91EOR")
	if err != nil {
		t.Error("error to create object:", err)
	}
	t.Log("[OK]")
}

func TestEpisode_GetNumber(t *testing.T) {

	ep, err := NewEpisode("https://www.animeworld.tv/play/x-ova-ita.rvNqx/yIGBNM")
	if err != nil {
		t.Error("error to create object:", err)
	}

	err = ep.GetNumber()
	if err != nil {
		t.Fatal("Error to get number: ", err.Error())
	}

	for i, tmp := range ep.Number {
		t.Log(ep.Number)
		if tmp == testEpisode.Number[i] {
			fmt.Println(tmp, ":", testEpisode.Number[i])
			t.Log("[OK] Get Number ", i)
		} else {
			t.Fatal("Not obtain ", testEpisode.Number[i], " but obtain ", tmp)
		}
	}

}

func TestEpisode_GetStreamLinks(t *testing.T) {

	ep, err := NewEpisode("https://www.animeworld.tv/play/x-ova-ita.rvNqx/yIGBNM")
	if err != nil {
		t.Error("error to create object:", err)
	}


	err = ep.GetStreamLinks()
	if err != nil {
		t.Fatal(err)
	}

	/*for key, _ := range ep.StreamLinks {
		if ep.StreamLinks[key] == testEpisode.StreamLinks[key] {
			t.Log("[OK] Get stream link of server", key)
		} else {
			t.Fatal("Not obtain ", testEpisode.StreamLinks[key], " but obtain ", ep.StreamLinks[key])
		}
	}*/
	t.Log(ep.StreamLinks)
}

func TestEpisode_GetDirectLinkServerBeta(t *testing.T) {

	var ep Episode

	ep.StreamLinks = make(map[string]string)

	ep.StreamLinks = map[string]string {
		"Beta Server": "https://www.animeworld.biz/v/6mo2pm6mwdv",
	}

	err := ep.GetDirectLinkServerBeta()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ep.StreamLinks)

}

func TestNewRssAnimeworld2(t *testing.T) {

	result := "http://www.thehiddendungeononlyicanenter.net/download-file.php?id=DLL/ANIME/MushokuTensei/MushokuTensei_Ep_01-02_SUB_ITA_v0.mp4"

	ep, err := NewEpisode("https://www.animeworld.tv/play/mushoku-tensei-isekai-ittara-honki-dasu.FwbNc/SKUFKi")
	if err != nil {
		t.Error("error to create object:", err)
	}

	err = ep.GetDownloadLink()
	if err != nil {
		t.Error("error to obtain direct link:", err)
	}

	if ep.DownloadLink != result {
		t.Error("not obtain", result, "but obtain", ep.DownloadLink)
	} else {
		t.Log("[OK]")
	}
}