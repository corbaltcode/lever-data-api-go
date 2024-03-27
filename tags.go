package lever

import (
	"context"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// Lever tags client interface
type TagsClientInterface interface {
	ClientInterface

	// Lists all tags in your Lever account.
	ListTags(ctx context.Context, req *ListTagsRequest) (*ListTagsResponse, error)
}

// Parameters for listing tags.
type ListTagsRequest struct {
	BaseListRequest
}

// Create a new list tags request with the required fields.
func NewListTagsRequest() *ListTagsRequest {
	return &ListTagsRequest{}
}

func (r *ListTagsRequest) GetPath() string {
	return "tags"
}

// Response for listing tags.
type ListTagsResponse struct {
	BaseListResponse

	// The tags in your Lever account.
	Tags []model.Tag `json:"data"`
}

// Lists all tags in your Lever account.
func (c *Client) ListTags(ctx context.Context, req *ListTagsRequest) (*ListTagsResponse, error) {
	var resp ListTagsResponse

	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
