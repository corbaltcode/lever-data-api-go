package model

import (
	"fmt"
	"net/http"
)

// Lever error response
type LeverError struct {
	HTTPResponse *http.Response `json:"-"`
	Code         string         `json:"code"`
	Message      string         `json:"message"`
}

func (e *LeverError) Error() string {
	return fmt.Sprintf("LeverError: %03d %s: %s", e.HTTPResponse.StatusCode, e.Code, e.Message)
}
