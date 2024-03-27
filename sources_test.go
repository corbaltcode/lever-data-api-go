package lever

import (
	"context"
	"net/http"
	"testing"

	"github.com/corbaltcode/lever-data-api-go/internal/testclient"
	"github.com/stretchr/testify/assert"
)

func TestSources(t *testing.T) {
	s := testclient.NewExpectManyHandler(
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"text":"Gild","count":24},{"text":"Posting","count":51},{"text":"Referral","count":90},{"text":"Email Applicant","count":135},{"text":"Email Lead","count":83}]}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/sources"),
		),
	)

	httpClient := http.Client{
		Transport: s,
	}

	ta := assert.New(t)

	c := NewClient(WithHTTPClient(&httpClient))
	ctx := context.Background()

	// List sources
	listReq := NewListSourcesRequest()
	listResp, err := c.ListSources(ctx, listReq)

	if ta.NoError(err) {
		ta.Len(listResp.Sources, 5)
	}
}
