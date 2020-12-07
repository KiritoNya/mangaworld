package mangaworld

import (
	"time"
)

type Chapter struct {
	Number       int
	PageNum      int
	Visual       int
	Visual_today int
	Page_url     []string
	DateAdd      time.Time
	keywords     []string
}
