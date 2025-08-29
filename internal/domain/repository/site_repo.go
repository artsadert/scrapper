package repository

import (
	"github.com/artsadert/test_parser/internal/domain/entities"
)

type SiteRepo interface {
	CreateSite(*entities.Site) error
}
