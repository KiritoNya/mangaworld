package mangaworld

import (
	"testing"
	"time"
)

const url = "https://www.mangaworld.cc/manga/1876/citrus-1/read/5fbbfab01c9bb544acdbbac0/1"

var testChapter = Chapter{
	Number:       16,
	PageNum:      22,
	Visual:       0,
	Visual_today: 0,
	Page_url:     nil,
	DateAdd:      time.Time{},
	KeyWords:     nil,
	resp:         nil,
}

func TestChapter_GetNumber(t *testing.T) {

	c, err := NewChapter(url)
	if err != nil {
		t.Fatalf("Error to create object")
	}

	err = c.GetNumber()
	if err != nil {
		t.Errorf("Error to get chapter number")
	}

	if c.Number != testChapter.Number {
		t.Error("Error not obtain", testChapter.Number, "but obtain", c.Number)
	} else {
		t.Log("Chapter Number [OK]")
	}
}

func TestChapter_GetPageNum(t *testing.T) {

	c, err := NewChapter(url)
	if err != nil {
		t.Fatalf("Error to create object")
	}

	err = c.GetPageNum()
	if err != nil {
		t.Errorf("Error to get chapter page number")
	}

	if c.PageNum != testChapter.PageNum {
		t.Error("Error not obtain", testChapter.PageNum, "but obtain", c.PageNum)
	} else {
		t.Log("Chapter Page Number [OK]")
	}
}
