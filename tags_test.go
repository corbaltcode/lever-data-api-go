package lever

import (
	"context"
	"net/http"
	"testing"

	"github.com/corbaltcode/lever-data-api-go/internal/testclient"
	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {
	s := testclient.NewExpectManyHandler(
		testclient.NewExpectHandler(
			http.StatusOK,
			`{"data":[{"text":"Infrastructure Engineer","count":23},{"text":"Customer Success Manager","count":15},{"text":"Customer Success","count":31},{"text":"Full-time","count":66},{"text":"San Francisco","count":70}]}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/v1/tags"),
		),
	)

	httpClient := http.Client{
		Transport: s,
	}

	ta := assert.New(t)

	c := NewClient(WithHTTPClient(&httpClient))
	ctx := context.Background()

	// List tags
	listReq := NewListTagsRequest()
	listResp, err := c.ListTags(ctx, listReq)

	if ta.NoError(err) {
		if ta.NotEmpty(listResp.Tags) {
			ta.Len(listResp.Tags, 5)
		}
	}
}
