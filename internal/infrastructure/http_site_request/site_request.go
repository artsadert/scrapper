package httpsiterequest

import (
	"fmt"
	"math/rand/v2"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/artsadert/test_parser/internal/domain/entities"
	"github.com/artsadert/test_parser/internal/domain/request"
)

type HttpSiteRequest struct {
}

func NewHttpSiteRequest() request.SiteRequest {
	return &HttpSiteRequest{}
}

func (r HttpSiteRequest) GetAndModifySite(entity *entities.Site) (*entities.Site, error) {
	resp, err := http.Get(entity.Url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is not 200, status code is %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	_, err = doc.Find("title").AppendHtml(entity.Keyword).Html()
	if err != nil {
		return nil, err
	}

	_, err = doc.Find("h1").AppendHtml(entity.Keyword).Html()
	if err != nil {
		return nil, err
	}
	replace := func(i int, s *goquery.Selection) {
		text := s.Text()
		if len(text) != 0 {
			random_pos := rand.IntN(len(text))
			s.SetHtml(text[:random_pos] + entity.Keyword + text[random_pos:])
		}
	}

	doc.Find("a, span, li, alt, p, noscript").Each(replace)

	data, err := doc.Html()
	if err != nil {
		return nil, err
	}
	entity.Data = data

	return entity, nil
}
