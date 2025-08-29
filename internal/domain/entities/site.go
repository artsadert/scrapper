package entities

type Site struct {
	Url     string
	Keyword string
	Data    string
}

func NewSite(url string, keyword string) *Site {
	return &Site{Url: url, Keyword: keyword}
}
