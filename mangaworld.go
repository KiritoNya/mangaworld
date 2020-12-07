package mangaworld

type Mangaworld struct {
	Url string
}

func New(url string) *Mangaworld {
	return &Mangaworld{Url: url}
}

