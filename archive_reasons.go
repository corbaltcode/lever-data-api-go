package lever

import (
	"context"
	"fmt"
	"net/url"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// Lever data archive reasons client interface
type ArchiveReasonsClientInterface interface {
	ClientInterface

	// Retrieve a single archive reason.
	GetArchiveReason(ctx context.Context, req *GetArchiveReasonRequest) (*GetArchiveReasonResponse, error)

	// List all archive reasons
	//
	// Lists all archive reasons in your Lever account.
	ListArchiveReasons(ctx context.Context, req *ListArchiveReasonsRequest) (*ListArchiveReasonsResponse, error)
}

// Parameters for retrieving a single archive reason.
type GetArchiveReasonRequest struct {
	BaseRequest

	// The archive reason id. This is required.
	ArchiveReasonId string
}

// Create a new GetArchiveReasonRequest with the required fields.
func NewGetArchiveReasonRequest(archiveReasonId string) *GetArchiveReasonRequest {
	return &GetArchiveReasonRequest{
		ArchiveReasonId: archiveReasonId,
	}
}

func (r *GetArchiveReasonRequest) GetPath() string {
	return fmt.Sprintf("archive_reasons/%s", url.PathEscape(r.ArchiveReasonId))
}

// Response for retrieving a single archive reason.
type GetArchiveReasonResponse struct {
	BaseResponse

	// The archive reason record.
	ArchiveReason *model.ArchiveReason `json:"data"`
}

// Parameters for listing archive reasons.
type ListArchiveReasonsRequest struct {
	BaseListRequest
}

// Create a new ListArchiveReasonsRequest with the required fields.
func NewListArchiveReasonsRequest() *ListArchiveReasonsRequest {
	return &ListArchiveReasonsRequest{}
}

func (r *ListArchiveReasonsRequest) GetPath() string {
	return "archive_reasons"
}

// Response for listing archive reasons.
type ListArchiveReasonsResponse struct {
	BaseListResponse

	// The archive reason records.
	ArchiveReasons []model.ArchiveReason `json:"data"`
}

// Retrieve a single archive reason.
func (c *Client) GetArchiveReason(ctx context.Context, req *GetArchiveReasonRequest) (*GetArchiveReasonResponse, error) {
	var resp GetArchiveReasonResponse
	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// List all archive reasons
//
// Lists all archive reasons in your Lever account.
func (c *Client) ListArchiveReasons(ctx context.Context, req *ListArchiveReasonsRequest) (*ListArchiveReasonsResponse, error) {
	var resp ListArchiveReasonsResponse
	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
