package lever

import (
	"context"
	"fmt"
	"net/url"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// Lever stages client interface.
type StagesClientInterface interface {
	ClientInterface

	// Retrieve a single stage.
	GetStage(ctx context.Context, req *GetStageRequest) (*GetStageResponse, error)

	// Lists all pipeline stages in your Lever account.
	ListStages(ctx context.Context, req *ListStagesRequest) (*ListStagesResponse, error)
}

// Parameters for retrieving a single stage.
type GetStageRequest struct {
	BaseRequest

	// The ID of the stage to retrieve.
	ID string
}

// Create a new get stage request with the required fields.
func NewGetStageRequest(id string) *GetStageRequest {
	return &GetStageRequest{
		ID: id,
	}
}

func (r *GetStageRequest) GetPath() string {
	return fmt.Sprintf("stages/%s", url.PathEscape(r.ID))
}

// Response for retrieving a single stage.
type GetStageResponse struct {
	BaseResponse

	// The stage.
	Stage model.Stage `json:"data"`
}

// Parameters for listing stages.
type ListStagesRequest struct {
	BaseListRequest
}

// Create a new list stages request with the required fields.
func NewListStagesRequest() *ListStagesRequest {
	return &ListStagesRequest{}
}

func (r *ListStagesRequest) GetPath() string {
	return "stages"
}

// Response for listing stages.
type ListStagesResponse struct {
	BaseListResponse

	// The stages in your Lever account.
	Stages []model.Stage `json:"data"`
}

// Retrieve a single stage.
func (c *Client) GetStage(ctx context.Context, req *GetStageRequest) (*GetStageResponse, error) {
	var resp GetStageResponse

	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Lists all pipeline stages in your Lever account.
func (c *Client) ListStages(ctx context.Context, req *ListStagesRequest) (*ListStagesResponse, error) {
	var resp ListStagesResponse

	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
