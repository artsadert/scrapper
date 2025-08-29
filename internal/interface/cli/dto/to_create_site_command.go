package dto

import "github.com/artsadert/test_parser/internal/application/command"

func ToCreateSiteCommand(url string, keyword string) *command.CreateSiteCommand {
	return &command.CreateSiteCommand{Url: url, Keyword: keyword}
}
