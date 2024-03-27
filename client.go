package lever

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// Common client interface
type ClientInterface interface {
	// Returns the base URL for the client.
	GetBaseURL() string
}

// Actual client implementation
type Client struct {
	// The base URL for the client.
	baseURL string

	// The HTTP client used to make requests.
	httpClient *http.Client

	// Functions to call to modify the request.
	preSend []func(ctx context.Context, req *http.Request) error
}

// Create a new client with the given options.
func NewClient(opts ...func(*Client)) *Client {
	httpClient := &http.Client{}

	c := &Client{
		baseURL:    defaultBaseURL,
		httpClient: httpClient,
		preSend:    nil,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Option for setting the base URL for the client.
func WithBaseURL(baseURL string) func(*Client) {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// Option for setting the HTTP client for the client.
func WithHTTPClient(httpClient *http.Client) func(*Client) {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// Option for setting the API key for the client.
func WithAPIKey(apiKey string) func(*Client) {
	var authValue = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", apiKey))))
	return WithHeader(headerAuthorization, authValue)
}

// Option for setting the user-agent for the client.
func WithUserAgent(userAgent string) func(*Client) {
	return WithHeader(headerUserAgent, userAgent)
}

// Option for setting an arbitrary header.
func WithHeader(header, value string) func(*Client) {
	return func(c *Client) {
		c.preSend = append(c.preSend,
			func(ctx context.Context, req *http.Request) error {
				req.Header.Set(header, value)
				return nil
			})
	}
}

// Returns the base URL for the client.
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

// Send a request. This performs the following steps:
//  1. The URL is constructed from the base URL and the request's path ([RequestInterface.GetPath]).
//  2. Query parameters are added to the URL ([RequestInterface.AddAPIQueryParams]).
//  3. The HTTP method is obtained from the request ([RequestInterface.GetHTTPMethod]).
//  4. The request body, if any, is obtained from the request ([RequestInterface.GetBody]).
//  5. A request is constructed with a default Accept and User-Agent header.
//  6. Pre-send functions are called to modify the request.
//  7. The request is sent using [http.Client.Do].
func (c *Client) send(ctx context.Context, req RequestInterface) (*http.Response, error) {
	reqURLStr := fmt.Sprintf("%s/%s", c.baseURL, req.GetPath())
	reqURL, err := url.Parse(reqURLStr)
	if err != nil {
		return nil, err
	}

	// Add include=, expand=, limit=, and offset= query parameters to the URL.
	query := reqURL.Query()
	req.AddAPIQueryParams(&query)
	reqURL.RawQuery = query.Encode()

	reqURLStr = reqURL.String()
	method := req.GetHTTPMethod()
	body, err := req.GetBody()
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, method, reqURLStr, body)
	if err != nil {
		return nil, err
	}

	// Set default headers
	httpReq.Header.Set(headerAccept, defaultAccept)
	httpReq.Header.Set(headerUserAgent, defaultUserAgent)

	// Is there a body?
	if body != nil {
		httpReq.Header.Set(headerContentType, req.GetContentType())
	}

	// Call preSend functions to modify the request.
	for _, preSend := range c.preSend {
		if err := preSend(ctx, httpReq); err != nil {
			return nil, err
		}
	}

	// Send the request.
	return c.httpClient.Do(httpReq)
}

// Execute an API request and return the response.
//
// This calls [Client.send] to send the request, then decodes the response into the given response.
func (c *Client) exec(ctx context.Context, req RequestInterface, resp ResponseInterface) error {
	httpResp, err := c.send(ctx, req)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()
	decoder := json.NewDecoder(httpResp.Body)

	if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		if err := decoder.Decode(resp); err != nil {
			return err
		}

		resp.SetHTTPResponse(httpResp)

		return nil
	}

	leverError := &model.LeverError{}
	if err := decoder.Decode(leverError); err != nil {
		return err
	}

	leverError.HTTPResponse = httpResp

	return leverError
}
