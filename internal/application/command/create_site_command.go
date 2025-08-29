package command

import "github.com/artsadert/test_parser/internal/application/common"

type CreateSiteCommand struct {
	Url     string
	Keyword string
}

type CreateSiteCommandResult struct {
	Result *common.SiteResult
}
