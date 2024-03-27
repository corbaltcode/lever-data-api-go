package multimodel

import (
	"encoding/json"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// The Interview model, but with expandable fields left unparsed.
type Interview struct {
	// Interview UID
	ID string `json:"id,omitempty"`

	// Interview Panel UID
	PanelID string `json:"panel,omitempty"`

	// Interview subject
	Subject string `json:"subject,omitempty"`

	// Interview note
	Note string `json:"note,omitempty"`

	// Array of interviewers
	Interviewers []model.Interviewer `json:"interviewers,omitempty"`

	// Name of timezone in which interview was scheduled to occur.
	Timezone string `json:"timezone,omitempty"`

	// Datetime when interview was created.
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// Datetime when interview is scheduled to occur.
	Date *int64 `json:"date,omitempty"`

	// Interview duration in minutes
	Duration int `json:"duration,omitempty"`

	// Interview location. Usually the name of a booked conference room but can also be a phone
	// number to call.
	Location string `json:"location,omitempty"`

	// The feedback form template selected for this interview.
	FeedbackTemplateID string `json:"feedbackTemplate,omitempty"`

	// The feedback forms (IDs or structs) associated with this interview.
	FeedbackForms json.RawMessage `json:"feedbackForms,omitempty"`

	// Frequency of feedback reminders (i.e. once, daily, frequently, none). Defaults to
	// 'frequently' which is every 6 hours.
	FeedbackReminder string `json:"feedbackReminder,omitempty"`

	// The user (ID or struct) who created the interview.
	User json.RawMessage `json:"user,omitempty"`

	// The stage (ID or struct) in which the candidate resided when this interview was scheduled.
	Stage json.RawMessage `json:"stage,omitempty"`

	// Datetime when interview was canceled. Value is nil if interview was never canceled.
	CanceledAt *int64 `json:"canceledAt,omitempty"`

	// List of job posting IDs that the interview is associated with
	PostingIDs []string `json:"postings,omitempty"`

	// List of job postings that the interview is associated with. Returned if expand=postings is
	// specified.
	Postings []model.Posting `json:"-"`
}

// Populate a regular [model.Interview] from this [multimodel.Interview].
func (o *Interview) ToModel(result *model.Interview) error {
	// Fields that map 1:1
	result.ID = o.ID
	result.PanelID = o.PanelID
	result.Subject = o.Subject
	result.Note = o.Note
	result.Interviewers = o.Interviewers
	result.Timezone = o.Timezone
	result.CreatedAt = o.CreatedAt
	result.Date = o.Date
	result.Duration = o.Duration
	result.Location = o.Location
	result.FeedbackTemplateID = o.FeedbackTemplateID
	result.FeedbackReminder = o.FeedbackReminder
	result.CanceledAt = o.CanceledAt
	result.PostingIDs = o.PostingIDs
	result.Postings = o.Postings

	feedbackFormIDs, feedbackForms, err := unmarshalArrayOfFeedbackFormsOrIDs(o.FeedbackForms)
	if err != nil {
		return err
	}

	result.FeedbackFormIDs = feedbackFormIDs
	result.FeedbackForms = feedbackForms

	userID, user, err := unmarshalUserOrID(o.User)
	if err != nil {
		return err
	}

	result.UserID = userID
	result.User = user

	stageID, stage, err := unmarshalStageOrID(o.Stage)
	if err != nil {
		return err
	}

	result.StageID = stageID
	result.Stage = stage

	return nil
}
