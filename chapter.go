package mangaworld

import (
	"errors"
	"github.com/KiritoNya/htmlutils"
	pb "github.com/cheggaaa/pb/v3"
	strip "github.com/grokify/html-strip-tags-go"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//Chapter is an object with all chapters information.
type Chapter struct {
	Url         string
	Volume      int
	Number      string
	PageNum     int
	Visual      int
	VisualToday int
	PageUrl     []string
	DateAdd     time.Time
	KeyWords    []string
	resp        *html.Node
}

//MonthNames is a map with the value int of italian months.
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

//NewChapter is a chapter constructor.
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

	urlMatrix := strings.Split(urlChapter, "/")
	//Check if the link contain the number of page
	match, err := regexp.MatchString("^[0-9]$", urlMatrix[len(urlMatrix)-1])
	if err != nil {
		return nil, err
	}

	if match == true {
		urlMatrix = urlMatrix[:len(urlMatrix)-1]
		urlChapter = strings.Join(urlMatrix, "/")
	}

	//Remove last / if exist
	if urlChapter[len(urlChapter)-1] == '/' {
		urlChapter = urlChapter[:len(urlChapter)-1]
	}

	c.Url = urlChapter

	return &c, nil
}

//Add number of volume of chapter to the object.
func (c *Chapter) GetVolume() error {

	sel, err := htmlutils.QuerySelector(c.resp, "select", "class", "volume custom-select")
	if err != nil {
		return err
	}

	options, err := htmlutils.GetGeneralTags(sel[0], "option")
	if err != nil {
		return err
	}

	for _, option := range options {
		if strings.Contains(htmlutils.RenderNode(option), "selected") {
			volumeString := string(htmlutils.GetNodeText(option, "option"))
			volumeString = strings.Replace(volumeString, "Volume ", "", -1)
			c.Volume, err = strconv.Atoi(volumeString)
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}

//Add number of chapter to the object.
func (c *Chapter) GetNumber() error {

	head, err := htmlutils.GetGeneralTags(c.resp, "head")
	if err != nil {
		return err
	}

	title, err := htmlutils.GetGeneralTags(head[0], "title")
	if err != nil {
		return err
	}

	titleChap := htmlutils.GetNodeText(title[0], "title")

	//CASE ... Capitolo ...
	if strings.Contains(string(titleChap), "Capitolo ") {

		matrix := strings.Split(string(titleChap), "Capitolo ")
		tmp := strings.Split(matrix[1], " ")

		//CASE ... Capitolo Extra ...
		if tmp[0] == "Extra" {
			c.Number = tmp[0] + " " + tmp[1]
			return nil
		}

		//CASE ... Capitolo [0-9]+ ...
		c.Number  = tmp[0]
		return nil

	}

	if strings.Contains(string(titleChap), "CApitolo "){
		matrix := strings.Split(string(titleChap), "CApitolo ")
		tmp := strings.Split(matrix[1], " ")

		//CASE ... Capitolo Extra ...
		if tmp[0] == "Extra" {
			c.Number = tmp[0] + " " + tmp[1]
			return nil
		}

		//CASE ... Capitolo [0-9]+ ...
		c.Number  = tmp[0]
		return nil
	}

	//CASE ... Oneshot ...
	if strings.Contains(string(titleChap), "Oneshot ") {
		c.Number = "Oneshot"
		return nil
	}

	return errors.New("Type of manga number not found")
}

//Add the number of pages of chapter to the object.
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

//Add the visual of chapter to the object.
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

//Add the daily visual to the object.
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

//Add urls chapter pages to the object.
func (c *Chapter) GetPageUrl() error {

	if c.PageNum == 0 {
		return errors.New("Error, page number of chapter not found, execute GetNumPage before this method")
	}

	for i := 1; i <= c.PageNum; i++ {
		resp, err := http.Get(c.Url + "/" + strconv.Itoa(i))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		nodeHtml, err := html.Parse(resp.Body)
		if err != nil {
			return err
		}

		divs, err := htmlutils.QuerySelector(nodeHtml, "div", "id", "page")
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

	return nil
}

//Add the date added to the object.
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

//Add keywords to the object.
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

//Download all pages of chapter in a folder defined by the dest parameter.
func (c *Chapter) Download(dest string) error {

	//pageNum not set
	if c.PageNum == 0 {
		c.GetPageNum()
	}

	//pageUrl not set
	if c.PageUrl == nil {
		c.GetPageUrl()
	}

	//number not set
	if c.Number == "" {
		c.GetNumber()
	}

	for _, page := range c.PageUrl {
		req, err := http.NewRequest("GET", page, nil)
		if err != nil {
			return err
		}
		name := path.Base(req.URL.Path)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		writer, err := os.OpenFile(dest+string(os.PathSeparator)+createNameFile(name, c.PageNum), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer writer.Close()

		numBytes := resp.ContentLength
		reader := io.LimitReader(resp.Body, numBytes)
		tmpl := `{{ magenta "{prefix}"}} {{ bar . (magenta "[") "◼" (cycle . "□" ) "□" "]"}} {{speed . | magenta }} {{percent . | magenta}}`
		tmpl = strings.Replace(tmpl, "{prefix}", name, -1)

		bar := pb.ProgressBarTemplate(tmpl).Start64(numBytes)
		bar.Set(pb.Bytes, true)
		bar.Set(pb.SIBytesPrefix, true)
		barReader := bar.NewProxyReader(reader)

		io.Copy(writer, barReader)

		bar.Finish()
	}
	return nil
}

func createNameFile(nameFile string, maxPag int) string {
	var str string
	numString := strings.Split(nameFile, ".") //1.jpg
	num, _ := strconv.Atoi(numString[0])
	if num < 10 {
		str += "0"
	}
	if maxPag > 99 && num < 100 {
		str += "0"
	}
	return str + strconv.Itoa(num) + "." + numString[1]
}
