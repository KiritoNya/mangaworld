package v2_test

import (
	v2 "github.com/KiritoNya/mangaworld/v2"
	"testing"
)

func TestListManga_MonthlyManga(t *testing.T) {
	lm := v2.NewListManga()

	err := lm.MonthlyManga()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("MANGA:", lm.Mangas)
}
