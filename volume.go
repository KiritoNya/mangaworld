package mangaworld

import (
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

//Volume is a object with all volum information.
type Volume struct {
	Number   int
	Name     string
	Chapters []Chapter
	resp     *html.Node
}

//NewVolume is a construct of volum object.
//You have to pass as a parameter the section dedicated to volumes on the manga page.
func NewVolume(VolumNode *html.Node) (*Volume, error) {

	//Check
	_, err := htmlutils.QuerySelector(VolumNode, "div", "class", "volume w-100 py-2")
	if err != nil {
		return nil, err
	}

	return &Volume{resp: VolumNode}, nil
}

//Add name value to the object.
//EX: "Volume 01".
func (v *Volume) GetName() error {

	tagsP, err := htmlutils.QuerySelector(v.resp, "p", "class", "volume-name d-inline")
	if err != nil {
		return err
	}

	v.Name = string(htmlutils.GetNodeText(tagsP[0], "p"))
	return nil
}

//Add number of volume to the object.
func (v *Volume) GetNumber() error {

	err := v.GetName()
	if err != nil {
		return err
	}

	v.Number, err = strconv.Atoi(strings.Split(v.Name, " ")[1])
	if err != nil {
		return err
	}

	return nil
}

//Add object Chapter to the object.
func (v *Volume) GetChapters() error {

	chContain, err := htmlutils.QuerySelector(v.resp, "div", "class", "volume-chapters pl-2")
	if err != nil {
		return err
	}

	divsCh, err := htmlutils.QuerySelector(chContain[0], "div", "class", "chapter")
	if err != nil {
		return err
	}

	for _, divCh := range divsCh {
		tagsA, err := htmlutils.GetGeneralTags(divCh, "a")
		if err != nil {
			return err
		}

		urlChapter, err := htmlutils.GetValueAttr(tagsA[0], "a", "href")
		if err != nil {
			return err
		}

		v.Chapters = append(v.Chapters, Chapter{Url: string(urlChapter[0])})
	}

	return nil
}
