package lever

import (
	"context"
	"fmt"
	"net/url"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// Lever data application client interface
type ApplicationClientInterface interface {
	ClientInterface

	// Retrieve a single application.
	//
	// This method returns the full application record for a single application.
	//
	// WARNING: This endpoint is deprecated but maintained for backwards compatibility. Use the
	// OpportunityClient GetOpportunity() method, specifying the expand parameter to include
	// applications.
	GetApplication(ctx context.Context, req *GetApplicationRequest) (*GetApplicationResponse, error)

	// Lists all applications for a candidate.
	//
	// WARNING: This endpoint is deprecated but maintained for backwards compatibility. Use the
	// OpportunityClient ListAllOpportunities() method, specifying the relevant contact UID in the
	// contact_id parameter and specifying the expand parameter to include applications.
	ListApplications(ctx context.Context, req *ListApplicationsRequest) (*ListApplicationsResponse, error)
}

// Parameters for retrieving a single application.
type GetApplicationRequest struct {
	BaseRequest

	// The opportunity id. This is required.
	OpportunityId string

	// The application id. This is required.
	ApplicationId string
}

// Create a new GetApplicationRequest with the required fields.
func NewGetApplicationRequest(opportunityId, applicationId string) *GetApplicationRequest {
	return &GetApplicationRequest{
		OpportunityId: opportunityId,
		ApplicationId: applicationId,
	}
}

func (r *GetApplicationRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/applications/%s", url.PathEscape(r.OpportunityId), url.PathEscape(r.ApplicationId))
}

// Response for retrieving a single application.
type GetApplicationResponse struct {
	BaseResponse

	// The application record.
	Application *model.Application `json:"data"`
}

// Parameters for listing applications for a candidate.
type ListApplicationsRequest struct {
	BaseListRequest

	// The opportunity id. This is required.
	OpportunityId string
}

// Create a new ListApplications with the required fields.
func NewListApplicationsRequest(opportunityId string) *ListApplicationsRequest {
	return &ListApplicationsRequest{
		OpportunityId: opportunityId,
	}
}

func (r *ListApplicationsRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/applications", url.PathEscape(r.OpportunityId))
}

// Response for listing applications for a candidate.
type ListApplicationsResponse struct {
	BaseListResponse

	// The application records.
	Applications []*model.Application `json:"data"`
}

// Retrieve a single application.
//
// This method returns the full application record for a single application.
//
// WARNING: This endpoint is deprecated but maintained for backwards compatibility. Use the
// OpportunityClient GetOpportunity() method, specifying the expand parameter to include
// applications.
func (c *Client) GetApplication(ctx context.Context, req *GetApplicationRequest) (*GetApplicationResponse, error) {
	var resp GetApplicationResponse
	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Lists all applications for a candidate.
//
// WARNING: This endpoint is deprecated but maintained for backwards compatibility. Use the
// OpportunityClient ListAllOpportunities() method, specifying the relevant contact UID in the
// contact_id parameter and specifying the expand parameter to include applications.
func (c *Client) ListApplications(ctx context.Context, req *ListApplicationsRequest) (*ListApplicationsResponse, error) {
	var resp ListApplicationsResponse
	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
