package mangaworld

import (
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

//Volum is a object with all volum information.
type Volum struct {
	Number   int
	Name     string
	Chapters []Chapter
	resp     *html.Node
}

//NewVolum is a construct of volum object.
func NewVolum(VolumNode *html.Node) (*Volum, error) {

	//Check
	_, err := htmlutils.QuerySelector(VolumNode, "div", "class", "volume w-100 py-2")
	if err != nil {
		return nil, err
	}

	return &Volum{resp: VolumNode}, nil
}

//Add name value to the object.
//EX: "Volume 01".
func (v *Volum) GetName() error {

	tagsP, err := htmlutils.QuerySelector(v.resp, "p", "class", "volume-name d-inline")
	if err != nil {
		return err
	}

	v.Name = string(htmlutils.GetNodeText(tagsP[0], "p"))
	return nil
}

//Add number of volum to the object.
func (v *Volum) GetNumber() error {

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
func (v *Volum) GetChapters() error {

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
