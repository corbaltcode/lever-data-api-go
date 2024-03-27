package lever

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// Lever resumes client interface
type ResumesClientInterface interface {
	ClientInterface

	// Retrieve information about a single resume.
	//
	// This endpoint retrieves the metadata for a single resume. To download a resume, see the
	// resume download endpoint.
	GetResume(ctx context.Context, req *GetResumeRequest) (*GetResumeResponse, error)

	// Download a resume file.
	//
	// Downloads a resume file if it exists
	//
	// The caller is responsible for closing the response body after consuming it.
	DownloadResume(ctx context.Context, req *DownloadResumeRequest) (*http.Response, error)

	// Lists all resumes associated with an opportunity in your Lever account.
	ListResumes(ctx context.Context, req *ListResumesRequest) (*ListResumesResponse, error)
}

// Parameters for retrieving a single resume.
type GetResumeRequest struct {
	BaseRequest

	// The ID of the opportunity associated with the resume.
	OpportunityID string

	// The ID of the resume to retrieve.
	ID string
}

// Create a new get resume request with the required fields.
func NewGetResumeRequest(opportunityID, id string) *GetResumeRequest {
	return &GetResumeRequest{
		OpportunityID: opportunityID,
		ID:            id,
	}
}

func (r *GetResumeRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/resumes/%s", url.PathEscape(r.OpportunityID), url.PathEscape(r.ID))
}

// Response for retrieving a single resume.
type GetResumeResponse struct {
	BaseResponse

	// The resume.
	Data model.Resume `json:"data"`
}

// Parameters for downloading a single resume.
type DownloadResumeRequest struct {
	BaseRequest

	// The ID of the opportunity associated with the resume.
	OpportunityID string

	// The ID of the resume to retrieve.
	ID string
}

// Create a new download resume request with the required fields.
func NewDownloadResumeRequest(opportunityID, id string) *DownloadResumeRequest {
	return &DownloadResumeRequest{
		OpportunityID: opportunityID,
		ID:            id,
	}
}

func (r *DownloadResumeRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/resumes/%s/download", url.PathEscape(r.OpportunityID), url.PathEscape(r.ID))
}

// Parameters for listing resumes.
type ListResumesRequest struct {
	BaseListRequest

	// The opportunity id associated with the resumes to list.
	OpportunityID string

	// If set, filter resumes by the timestamp they were uploaded at. If only UploadedAtStart is
	// specified, all resumes uploaded from that timestamp (inclusive) to the present will be
	// included. If only UploadedAtEnd is specified, all resumes uploaded before that timestamp
	// (inclusive) are included. If either value is not a proper timestamp a 400 error will be
	// returned for a malformed request. If there is no uploadedAt date on the resume (for example,
	// resumes parsed from online sources such as LinkedIn or GitHub) the createdAt date will be
	// used instead.
	UploadedAtStart *int64
	UploadedAtEnd   *int64
}

// Create a new list resumes request with the required fields.
func NewListResumesRequest(opportunityID string) *ListResumesRequest {
	return &ListResumesRequest{
		OpportunityID: opportunityID,
	}
}

func (r *ListResumesRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/resumes", url.PathEscape(r.OpportunityID))
}

func (r *ListResumesRequest) AddAPIQueryParams(v *url.Values) {
	r.BaseListRequest.AddAPIQueryParams(v)

	if r.UploadedAtStart != nil {
		v.Add("uploadedAtStart", fmt.Sprintf("%d", *r.UploadedAtStart))
	}

	if r.UploadedAtEnd != nil {
		v.Add("uploadedAtEnd", fmt.Sprintf("%d", *r.UploadedAtEnd))
	}
}

// Response for listing resumes.
type ListResumesResponse struct {
	BaseListResponse

	// The resumes in your Lever account.
	Data []model.Resume `json:"data"`
}

// Retrieve information about a single resume.
//
// This endpoint retrieves the metadata for a single resume. To download a resume, see the
// resume download endpoint.
func (c *Client) GetResume(ctx context.Context, req *GetResumeRequest) (*GetResumeResponse, error) {
	var resp GetResumeResponse

	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Download a resume file.
//
// Downloads a resume file if it exists
func (c *Client) DownloadResume(ctx context.Context, req *DownloadResumeRequest) (*http.Response, error) {
	httpResp, err := c.send(ctx, req)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode >= 300 {
		defer httpResp.Body.Close()
		decoder := json.NewDecoder(httpResp.Body)
		leverError := &model.LeverError{
			HTTPResponse: httpResp,
		}

		if decoder.Decode(leverError) != nil {
			leverError.Code = httpResp.Status
			leverError.Message = fmt.Sprintf("Unexpected HTTP response status code: %d", httpResp.StatusCode)
		}

		return nil, leverError
	}

	return httpResp, nil
}

// Lists all resumes associated with an opportunity in your Lever account.
func (c *Client) ListResumes(ctx context.Context, req *ListResumesRequest) (*ListResumesResponse, error) {
	var resp ListResumesResponse

	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
