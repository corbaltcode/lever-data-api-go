// Basic request/response types, interfaces, and helper methods.
package lever

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Interface that all requests must implement.
type RequestInterface interface {
	// Retrieve the path for the request, relative to the base URL.
	GetPath() string

	// Retrieve the HTTP method for the request.
	GetHTTPMethod() string

	// Retrieve the content type for the request body (if any).
	GetContentType() string

	// Retrieve the data to send in the request body.
	GetBody() (io.Reader, error)

	// Add standard query parameters.
	AddAPIQueryParams(query *url.Values)
}

// Base type for all requests.
// This adds the includes and expands parameters.
type BaseRequest struct {
	// Parameters to include the the response. This is optional.
	Include []string

	// Parameters to expand in the response. This is optional.
	Expand []string
}

// Add include= and expand= query parameters to a URL.
func (br *BaseRequest) AddAPIQueryParams(query *url.Values) {
	for _, inc := range br.Include {
		query.Add(paramInclude, inc)
	}

	for _, exp := range br.Expand {
		query.Add(paramExpand, exp)
	}
}

// Default method for HTTP request.
func (br *BaseRequest) GetHTTPMethod() string {
	return http.MethodGet
}

// Default body for HTTP request.
func (br *BaseRequest) GetBody() (io.Reader, error) {
	return nil, nil
}

// Default content type for HTTP request
func (br *BaseRequest) GetContentType() string {
	return mimeTypeApplicationJSON
}

// Base type for all list requests.
// This builds on BaseRequest by adding limit and offset parameters.
type BaseListRequest struct {
	BaseRequest

	// The number of items to include in the response. This is optional.
	Limit int

	// The pagination offset. This is optional.
	Offset string
}

// Add include=, expand=, limit=, and offset= query parameters to a URL.
func (blr *BaseListRequest) AddAPIQueryParams(query *url.Values) {
	blr.BaseRequest.AddAPIQueryParams(query)

	if blr.Offset != "" {
		query.Add(paramOffset, blr.Offset)
	}

	if blr.Limit > 0 {
		query.Add(paramLimit, fmt.Sprintf("%d", blr.Limit))
	}
}

// Interface that all responses must implement.
type ResponseInterface interface {
	// Set the HTTP response on this API response.
	SetHTTPResponse(resp *http.Response)
}

// Base type for all responses that includes the HTTP response.
type BaseResponse struct {
	HTTPResponse *http.Response `json:"-"`
}

// Set the HTTP response on this API response.
func (r *BaseResponse) SetHTTPResponse(resp *http.Response) {
	r.HTTPResponse = resp
}

// Base type for all list responses.
// This adds the next and hasNext parameters for pagination.
type BaseListResponse struct {
	BaseResponse

	// The next pagination offset.
	Next string `json:"next,omitempty"`

	// Whether there is a next page.
	HasNext bool `json:"hasNext,omitempty"`
}
