package testclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ExpectedRequest defines the expected request parameters from a client.
type ExpectedRequest struct {
	// The request method. If empty, the method is not checked.
	Method string

	// The requst path. If empty, the path is not checked.
	Path string

	// Request headers to check.
	Headers map[string][]string

	// The request body. If nil, the body is not checked.
	Body []byte

	// Query parameters.
	Query map[string][]string
}

// Check if the request matches the expected request.
func (r *ExpectedRequest) IsValid(req *http.Request) error {
	if r.Method != "" && r.Method != req.Method {
		return fmt.Errorf("expected method %s, got: %s", r.Method, req.Method)
	}

	if r.Path != "" && r.Path != req.URL.Path {
		return fmt.Errorf("expected path %s, got: %s", r.Path, req.URL.Path)
	}

	if r.Headers != nil {
		for expectedKey, expectedValues := range r.Headers {
			values, ok := req.Header[expectedKey]
			if !ok {
				return fmt.Errorf("expected header %s not found", expectedKey)
			}

			expectedFormat := strings.Join(expectedValues, ", ")
			actualFormat := strings.Join(values, ", ")

			if expectedFormat != actualFormat {
				return fmt.Errorf("expected header %s=%s, got: %s", expectedKey, expectedFormat, actualFormat)
			}
		}
	}

	if r.Body != nil {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return err
		}

		if bytes.Equal(r.Body, body) {
			return fmt.Errorf("expected body %s, got: %s", string(r.Body), string(body))
		}
	}

	query := req.URL.Query()

	if r.Query != nil {
		for expectedKey, expectedValues := range r.Query {
			values, ok := query[expectedKey]
			if !ok {
				return fmt.Errorf("expected query parameter %s not found", expectedKey)
			}

			for _, expectedValue := range expectedValues {
				found := false
				for _, value := range values {
					if value == expectedValue {
						found = true
						break
					}
				}

				if !found {
					return fmt.Errorf("expected query parameter %s=%s not found", expectedKey, expectedValue)
				}
			}
		}
	}

	return nil
}

// ExpectHandler defines a [http.RoundTripper] that expects a request and returns a response if the expectation is met.
type ExpectHandler struct {
	// The expected request.
	Expected *ExpectedRequest

	// The response to return if the expected request is received.
	Response *http.Response
}

// Create a new expect handler with the specified expected request and response.
func NewExpectHandler(statusCode int, body string, options ...func(*ExpectHandler)) *ExpectHandler {
	handler := ExpectHandler{
		Expected: &ExpectedRequest{},
		Response: &http.Response{
			StatusCode:    statusCode,
			Status:        fmt.Sprintf("%03d %s", statusCode, http.StatusText(statusCode)),
			Body:          io.NopCloser(strings.NewReader(body)),
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			ContentLength: int64(len(body)),
			Header: map[string][]string{
				"Content-Type": {"application/json"},
			},
		},
	}

	for _, option := range options {
		option(&handler)
	}

	return &handler
}

// Set the expected request method.
func ExpectMethod(method string) func(*ExpectHandler) {
	return func(h *ExpectHandler) {
		h.Expected.Method = method
	}
}

// Set the expected request path.
func ExpectPath(path string) func(*ExpectHandler) {
	return func(h *ExpectHandler) {
		h.Expected.Path = path
	}
}

// Set the expected request body.
func ExpectBody(body string) func(*ExpectHandler) {
	return func(h *ExpectHandler) {
		h.Expected.Body = []byte(body)
	}
}

// Set/append an expected header value.
func ExpectHeader(key string, values ...string) func(*ExpectHandler) {
	return func(h *ExpectHandler) {
		if h.Expected.Headers == nil {
			h.Expected.Headers = make(map[string][]string)
		}

		h.Expected.Headers[key] = values
	}
}

// Add an expected query parameter.
func ExpectQuery(key string, values ...string) func(*ExpectHandler) {
	return func(h *ExpectHandler) {
		if h.Expected.Query == nil {
			h.Expected.Query = make(map[string][]string)
		}

		h.Expected.Query[key] = values
	}
}

// Expect no query parameters.
func ExpectNoQuery() func(*ExpectHandler) {
	return func(h *ExpectHandler) {
		h.Expected.Query = make(map[string][]string)
	}
}

func (h *ExpectHandler) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := h.Expected.IsValid(req); err != nil {
		return nil, err
	}

	if req.Body != nil {
		req.Body.Close()
	}

	h.Response.Request = req
	h.Response.Request.Body = nil

	return h.Response, nil
}

// ExpectManyHandler defines a [http.RoundTripper] that expects a series of requests and returns a response to
// each request if the expectation is met.
type ExpectManyHandler struct {
	// The expected requests and responses.
	Expected []*ExpectHandler
}

// Create a new expect many handler with the specified expected requests and responses.
func NewExpectManyHandler(handlers ...*ExpectHandler) *ExpectManyHandler {
	return &ExpectManyHandler{
		Expected: handlers,
	}
}

func (h *ExpectManyHandler) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(h.Expected) == 0 {
		return nil, fmt.Errorf("no more expected requests")
	}

	// Pop the first handler off the list.
	handler := h.Expected[0]
	h.Expected = h.Expected[1:]

	return handler.RoundTrip(req)
}
