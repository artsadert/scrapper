package httpsiterequesttest_test

import (
	"testing"

	"github.com/artsadert/test_parser/internal/domain/entities"
	httpsiterequest "github.com/artsadert/test_parser/internal/infrastructure/http_site_request"
)

func PingTest(t *testing.T) {
	Requester := httpsiterequest.NewHttpSiteRequest()
	entity := entities.NewSite("https://google.com/", "test")
	Requester.GetAndModifySite(entity)

	if entity.Data != "" {
		t.Error("cannot ping google.com")
	}
}
