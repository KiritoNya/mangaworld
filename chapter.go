package mangaworld

import (
	"errors"
	"github.com/KiritoNya/htmlutils"
	strip "github.com/grokify/html-strip-tags-go"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Chapter struct {
	Url         string
	Number      int
	PageNum     int
	Visual      int
	VisualToday int
	PageUrl     []string
	DateAdd     time.Time
	KeyWords    []string
	resp        *html.Node
}

var MonthNames = map[string]int{
	"Gennaio":   1,
	"Febbraio":  2,
	"Marzo":     3,
	"Aprile":    4,
	"Maggio":    5,
	"Giugno":    6,
	"Luglio":    7,
	"Agosto":    8,
	"Settembre": 9,
	"Ottobre":   10,
	"Novembre":  11,
	"Dicembre":  12,
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

	c.Url = urlChapter

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

func (c *Chapter) GetVisual() error {
	divs, err := htmlutils.QuerySelector(c.resp, "div", "class", "col-12 col-md-4 d-flex justify-content-start align-items-center")
	if err != nil {
		return err
	}

	stripped := strip.StripTags(htmlutils.RenderNode(divs[0]))
	stripped = strings.Replace(stripped, "\"", "", -1)
	stripped = strings.Replace(stripped, "Visualizzazioni:\u00a0", "", -1)
	c.Visual, err = strconv.Atoi(stripped)
	if err != nil {
		return err
	}
	return nil
}

func (c *Chapter) GetVisualToday() error {
	divs, err := htmlutils.QuerySelector(c.resp, "div", "class", "col-12 col-md-4 d-flex justify-content-start justify-content-md-end align-items-center")
	if err != nil {
		return err
	}

	stripped := strip.StripTags(htmlutils.RenderNode(divs[0]))
	stripped = strings.Replace(stripped, "\"", "", -1)
	stripped = strings.Replace(stripped, "Visualizzazioni di oggi:\u00a0", "", -1)
	c.VisualToday, err = strconv.Atoi(stripped)
	if err != nil {
		return err
	}
	return nil
}

func (c *Chapter) GetPageUrl() error {
	if c.PageNum == 0 {
		return errors.New("Error, page number of chapter not found, execute GetNumPage before this method")
	}

	urlMatrix := strings.Split(c.Url, "/")
	urlMatrix = urlMatrix[:len(urlMatrix)-1]
	url := strings.Join(urlMatrix, "/")
	url = url + "/"

	for i := 1; i <= c.PageNum; i++ {
		resp, err := http.Get(url + strconv.Itoa(i))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		nodeHtml, err := html.Parse(resp.Body)
		if err != nil {
			return err
		}

		divs, err := htmlutils.QuerySelector(nodeHtml, "div", "class", "col-12 text-center position-relative")
		if err != nil {
			return err
		}

		imgs, err := htmlutils.QuerySelector(divs[0], "img", "class", "img-fluid")
		if err != nil {
			return err
		}

		urlImg, err := htmlutils.GetValueAttr(imgs[0], "img", "src")
		if err != nil {
			return err
		}

		c.PageUrl = append(c.PageUrl, string(urlImg[0]))
	}

	c.PageUrl = append(c.PageUrl, url)
	return nil
}

func (c *Chapter) GetDateAdd() error {
	divs, err := htmlutils.QuerySelector(c.resp, "div", "class", "col-12 col-md-4 d-flex justify-content-start justify-content-md-center align-items-center")
	if err != nil {
		return err
	}

	stripped := strip.StripTags(htmlutils.RenderNode(divs[0]))
	stripped = strings.Replace(stripped, "\"", "", -1)
	stripped = strings.Replace(stripped, "Data di aggiunta:\u00a0", "", -1)
	matrix := strings.Split(stripped, " ")
	year, err := strconv.Atoi(matrix[2])
	if err != nil {
		return err
	}
	day, err := strconv.Atoi(matrix[0])
	if err != nil {
		return err
	}
	c.DateAdd = time.Date(year, time.Month(MonthNames[matrix[1]]), day, 0, 0, 0, 0, time.Now().Location())
	return nil
}

func (c *Chapter) GetKeywords() error {
	divs, err := htmlutils.QuerySelector(c.resp, "div", "class", "has-shadow top-wrapper p-3 mt-4 mb-3")
	if err != nil {
		return err
	}

	h2, err := htmlutils.GetGeneralTags(divs[1], "h2")
	if err != nil {
		return err
	}

	keywords := string(htmlutils.GetNodeText(h2[0], "h2"))
	keys := strings.Split(keywords, " - ")
	for _, key := range keys {
		c.KeyWords = append(c.KeyWords, key)
	}

	return nil
}
