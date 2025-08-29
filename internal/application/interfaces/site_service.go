package interfaces

import "github.com/artsadert/test_parser/internal/application/command"

type SiteService interface {
	CreateSite(*command.CreateSiteCommand) *command.CreateSiteCommandResult
}
