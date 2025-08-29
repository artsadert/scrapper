package services

import (
	"github.com/artsadert/test_parser/internal/application/command"
	"github.com/artsadert/test_parser/internal/application/common"
	"github.com/artsadert/test_parser/internal/application/interfaces"
	"github.com/artsadert/test_parser/internal/domain/entities"
	"github.com/artsadert/test_parser/internal/domain/repository"
	"github.com/artsadert/test_parser/internal/domain/request"
)

type SiteService struct {
	siteRepo  repository.SiteRepo
	requester request.SiteRequest
}

func NewSiteService(repo repository.SiteRepo, requester request.SiteRequest) interfaces.SiteService {
	return &SiteService{siteRepo: repo, requester: requester}
}

func (s SiteService) CreateSite(site *command.CreateSiteCommand) *command.CreateSiteCommandResult {
	// Create entity without data
	entity := entities.NewSite(site.Url, site.Keyword)

	// Add data and modify it
	s.requester.GetAndModifySite(entity)

	// Write changed data
	err := s.siteRepo.CreateSite(entity)
	return &command.CreateSiteCommandResult{Result: &common.SiteResult{Status: err}}
}
