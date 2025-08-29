package cli

import (
	"log"
	"os"

	"github.com/artsadert/test_parser/internal/application/interfaces"
	"github.com/artsadert/test_parser/internal/interface/cli/dto"
	"github.com/spf13/cobra"
)

type SiteController struct {
	service interfaces.SiteService
	command *cobra.Command
}

func NewSiteController(service interfaces.SiteService) *SiteController {
	controller := &SiteController{service: service}
	controller.command = &cobra.Command{
		Use:   "",
		Short: "Downloads site",
		Long:  "Downloads site html and put file in .env workdir",
		Run: func(cmd *cobra.Command, args []string) {
			url, err := cmd.Flags().GetString("url")
			if err != nil {
				log.Fatal("url must not be empty")
			}
			keyword, err := cmd.Flags().GetString("keyword")
			if err != nil {
				log.Fatal("keyword must not be empty")
			}
			log.Println(url, keyword)

			err = controller.service.CreateSite(dto.ToCreateSiteCommand(url, keyword)).Result.Status
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	controller.command.Flags().StringP("url", "u", "https://google.com/", "set the url to parse page")
	controller.command.Flags().StringP("keyword", "k", "", "set the keyword to include on page")
	return controller
}

func (c SiteController) Execute() {
	if err := c.command.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
