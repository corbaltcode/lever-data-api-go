package model

// Individual interviews can be returned on this endpoint. Interviews are also contained in
// interview panels, and can be updated via the Panels endpoint.
type Interview struct {
	// Interview UID
	ID string

	// Interview Panel UID
	PanelID string

	// Interview subject
	Subject string

	// Interview note
	Note string

	// Array of interviewers
	Interviewers []Interviewer

	// Name of timezone in which interview was scheduled to occur.
	Timezone string

	// Datetime when interview was created.
	CreatedAt *int64

	// Datetime when interview is scheduled to occur.
	Date *int64

	// Interview duration in minutes
	Duration int

	// Interview location. Usually the name of a booked conference room but can also be a phone
	// number to call.
	Location string

	// The feedback form template selected for this interview.
	FeedbackTemplateID string

	// The feedback forms IDs associated with this interview.
	FeedbackFormIDs []string

	// The feedback forms associated with this interview. Returned if expand=feedbackForms is
	// specified.
	FeedbackForms []FeedbackForm

	// Frequency of feedback reminders (i.e. once, daily, frequently, none). Defaults to
	// 'frequently' which is every 6 hours.
	FeedbackReminder string

	// The user ID who created the interview.
	UserID string

	// The user who created the interview. Returned if expand=user is specified.
	User *User

	// The stage ID in which the candidate resided when this interview was scheduled.
	StageID string

	// The stage in which the candidate resided when this interview was scheduled. Returned if
	// expand=stage is specified.
	Stage *Stage

	// Datetime when interview was canceled. Value is nil if interview was never canceled.
	CanceledAt *int64

	// List of job posting IDs that the interview is associated with
	PostingIDs []string

	// List of job postings that the interview is associated with. Returned if expand=postings is
	// specified.
	Postings []Posting
}

// Interviewer is a user who is scheduled to conduct an interview.
type Interviewer struct {
	// User UID
	ID string `json:"id,omitempty"`

	// User's name
	Name string `json:"name,omitempty"`

	// User's email
	Email string `json:"email,omitempty"`
}
