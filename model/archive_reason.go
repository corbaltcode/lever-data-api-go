package model

// Archive reasons provide granularity behind to candidates who have exited your active hiring
// pipeline.
//
// Candidates exit your active pipeline either due to being hired at your company or due to being
// rejected for a specific reason. These dispositions allow you to track each and every candidate
// who is no longer active within your pipeline.
type ArchiveReason struct {
	// Archive reason UID
	ID string `json:"id,omitempty"`

	// The name of the archive reason as shown in the Lever interface.
	Text string `json:"text,omitempty"`

	// The status of the archive reason. Can be either active or inactive.
	Status string `json:"status,omitempty"`

	// The type of the archive reason. Can be either hired or non-hired.
	Type string `json:"type,omitempty"`
}
