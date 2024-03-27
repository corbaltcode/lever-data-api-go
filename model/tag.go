package model

// Tags provide additional information or context to a candidate within your pipeline. Tags serve
// as a way of grouping candidates for easy viewing of individuals with specific attributes.
type Tag struct {
	// Tag text
	Text string `json:"text,omitempty"`

	// Number of candidates tagged
	Count int64 `json:"count,omitempty"`
}
