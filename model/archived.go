package model

// Archived status.
type Archived struct {
	// Datetime when application was last archived in Lever
	ArchivedAt *int64 `json:"archivedAt,omitempty"`

	// Reason why application was last archived
	ReasonID string `json:"reason,omitempty"`
}
