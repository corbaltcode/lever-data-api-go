package model

// Feedback forms are added to Opportunities as they are completed after interviews by interviewers
// or they can be manually added directly to the profile.
type FeedbackForm struct {
	// Form UID.
	ID string

	// Form type. Feedback forms are of type interview.
	Type string

	// Form title. This can be edited in Feedback and Form Settings.
	Text string

	// Form instructions.
	Instructions string

	// Form template UID. This form represents a completed form template.
	BaseTemplateID string

	// An array of form fields. Feedback forms support the follow field types:
	//     - code - for programming questions
	//     - date - special field for dates
	//     - dropdown - a dropdown menu
	//     - multiple choice - choose only one
	//     - multiple select - choose 1 or more
	//     - score system - overall candidate rating
	//     - score - thumbs up / thumbs down format
	//     - scorecard - customized evaluation for multiple skills
	//     - text - single line answer
	//     - textarea - longer form answer
	//     - yes/no - a yes or no question
	Forms []any

	// The user ID who completed and submitted the feedback.
	UserID string

	// The user who completed and submitted the feedback. Returned if expand=user is specified.
	User *User

	// The interview panel that the feedback is associated with, if the feedback is associated with
	// an interview.
	PanelID string

	// The interview for which the feedback was submitted. Manually added feedback forms will not
	// be associated with an interview.
	InterviewID string

	// Datetime when form was created.
	CreatedAt *int64

	// This value is null when the updatedAt property has not been previously set. This is likely
	// to occur for feedback that were created prior to the introduction of this property, and have
	// not since been updated.
	UpdatedAt *int64

	// Datetime when form was completed.
	CompletedAt *int64

	// Datetime when form was deleted.
	DeletedAt *int64
}
