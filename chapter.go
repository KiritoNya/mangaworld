package mangaworld

import (
	"github.com/KiritoNya/htmlutils"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Chapter struct {
	Number       int
	PageNum      int
	Visual       int
	Visual_today int
	Page_url     []string
	DateAdd      time.Time
	KeyWords     []string
	resp         *html.Node
}

func NewChapter(urlChapter string) (*Chapter, error) {
	var c Chapter

	resp, err := http.Get(urlChapter)
	if err != nil {
		return &Chapter{}, err
	}
	defer resp.Body.Close()

	c.resp, err = html.Parse(resp.Body)
	if err != nil {
		return &Chapter{}, err
	}

	return &c, nil
}

func (c *Chapter) GetNumber() error {
	sel, err := htmlutils.QuerySelector(c.resp, "select", "class", "chapter custom-select")
	if err != nil {
		return err
	}

	options, err := htmlutils.GetGeneralTags(sel[0], "option")
	if err != nil {
		return err
	}

	for _, option := range options {
		if strings.Contains(htmlutils.RenderNode(option), "selected") {
			chapterString := string(htmlutils.GetNodeText(option, "option"))
			chapterString = strings.Replace(chapterString, "Capitolo ", "", -1)
			c.Number, err = strconv.Atoi(chapterString)
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}

func (c *Chapter) GetPageNum() error {
	sel, err := htmlutils.QuerySelector(c.resp, "select", "class", "page custom-select")
	if err != nil {
		return err
	}

	options, err := htmlutils.GetGeneralTags(sel[0], "option")
	if err != nil {
		return err
	}

	txtNode := string(htmlutils.GetNodeText(options[0], "option"))
	txtNode = strings.Split(txtNode, "/")[1]
	c.PageNum, err = strconv.Atoi(txtNode)
	if err != nil {
		return err
	}

	return nil
}
