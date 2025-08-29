package main

import (
	"log"

	"github.com/artsadert/test_parser/internal/application/services"
	httpsiterequest "github.com/artsadert/test_parser/internal/infrastructure/http_site_request"
	localstorage "github.com/artsadert/test_parser/internal/infrastructure/local_storage"
	"github.com/artsadert/test_parser/internal/infrastructure/local_storage/site"
	"github.com/artsadert/test_parser/internal/interface/cli"
)

func main() {
	cfg, err := localstorage.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	if cfg.WorkDir == "" {
		log.Panic("workdir is empty")
	}
	repo := site.NewLocalSiteRepository(cfg.WorkDir)
	req := httpsiterequest.NewHttpSiteRequest()

	service := services.NewSiteService(repo, req)

	controller := cli.NewSiteController(service)

	controller.Execute()
}
