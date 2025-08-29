package request

import (
	"github.com/artsadert/test_parser/internal/domain/entities"
)

type SiteRequest interface {
	GetAndModifySite(*entities.Site) (*entities.Site, error)
}
