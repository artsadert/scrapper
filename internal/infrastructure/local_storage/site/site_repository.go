package site

import (
	"fmt"
	"os"
	"strings"

	"github.com/artsadert/test_parser/internal/domain/entities"
	"github.com/artsadert/test_parser/internal/domain/repository"
)

type LocalSiteRepository struct {
	WorkDir string
}

func NewLocalSiteRepository(work_dir string) repository.SiteRepo {
	return &LocalSiteRepository{WorkDir: work_dir}
}

func (r LocalSiteRepository) CreateSite(site *entities.Site) error {

	domain_name := strings.Split(site.Url, "/")[2]

	filename := fmt.Sprintf("%s%s.html", r.WorkDir, domain_name)

	err := os.WriteFile(filename, []byte(site.Data), 0644)
	return err
}
