package lever

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/corbaltcode/lever-data-api-go/internal/testclient"
	"github.com/stretchr/testify/assert"
)

const testApiKey = "EXAMPLE_API_KEY"

type testRequest struct {
	BaseRequest
	path   string `json:"-"`
	method string `json:"-"`
}

func (tr *testRequest) GetPath() string {
	if tr.path == "" {
		return "req_path"
	}
	return tr.path
}

func (tr *testRequest) GetHTTPMethod() string {
	if tr.method == "" {
		return http.MethodGet
	}
	return tr.method
}

func (tr *testRequest) GetBody() (io.Reader, error) {
	return nil, nil
}

type testResponse struct {
	BaseResponse
}

func TestClientOptions(t *testing.T) {
	s := testclient.NewExpectManyHandler(
		testclient.NewExpectHandler(
			http.StatusOK,
			`{}`,
			testclient.ExpectMethod(http.MethodGet),
			testclient.ExpectPath("/custom_path/req_path"),
			testclient.ExpectHeader("Authorization", "Basic RVhBTVBMRV9BUElfS0VZOg=="),
			testclient.ExpectHeader("User-Agent", "TestClient"),
		),
	)

	httpClient := http.Client{
		Transport: s,
	}

	c := NewClient(WithHTTPClient(&httpClient), WithBaseURL("http://localhost/custom_path"), WithAPIKey(testApiKey), WithUserAgent("TestClient"))

	ctx := context.Background()
	ta := assert.New(t)

	// Make a request with the custom path.
	req := testRequest{}
	resp := testResponse{}
	err := c.exec(ctx, &req, &resp)

	ta.NoError(err)
}
