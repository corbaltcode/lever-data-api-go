package lever

import (
	"context"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// Lever sources client interface
type SourcesClientInterface interface {
	ClientInterface

	// Lists all sources in your Lever account.
	ListSources(ctx context.Context, req *ListSourcesRequest) (*ListSourcesResponse, error)
}

// Parameters for listing sources.
type ListSourcesRequest struct {
	BaseListRequest
}

// Create a new list sources request with the required fields.
func NewListSourcesRequest() *ListSourcesRequest {
	return &ListSourcesRequest{}
}

func (r *ListSourcesRequest) GetPath() string {
	return "sources"
}

// Response for listing sources.
type ListSourcesResponse struct {
	BaseListResponse

	// The sources in your Lever account.
	Sources []model.Tag `json:"data"`
}

// Lists all sources in your Lever account.
func (c *Client) ListSources(ctx context.Context, req *ListSourcesRequest) (*ListSourcesResponse, error) {
	var resp ListSourcesResponse

	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
