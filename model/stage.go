package model

// Stages are steps in the recruiting workflow of your hiring pipeline. All candidates belong to a
// stage and change stages as they progress through the recruiting pipeline, typically in a linear
// fashion.
type Stage struct {
	// Stage UID
	ID string `json:"id,omitempty"`

	// Title of the stage
	Text string `json:"text,omitempty"`
}

// Stage change for an opportunity
type StageChange struct {
	// Stage UID of the stage the candidate entered
	ToStageID string `json:"toStageId,omitempty"`

	// The index of the stage in the pipeline at the time the stage change occurred
	ToStageIndex int `json:"toStageIndex,omitempty"`

	// Time at which stage change occurred
	UpdatedAt *int64 `json:"updatedAt,omitempty"`

	// User UID
	UserID string `json:"userId,omitempty"`
}
