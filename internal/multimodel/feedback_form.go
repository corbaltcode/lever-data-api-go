package multimodel

import (
	"encoding/json"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// The FeedbackForm model, but with expandable fields left unparsed.
type FeedbackForm struct {
	// Form UID.
	ID string `json:"id,omitempty"`

	// Form type. Feedback forms are of type interview.
	Type string `json:"type,omitempty"`

	// Form title. This can be edited in Feedback and Form Settings.
	Text string `json:"text,omitempty"`

	// Form instructions.
	Instructions string `json:"instructions,omitempty"`

	// Form template UID. This form represents a completed form template.
	BaseTemplateID string `json:"baseTemplate,omitempty"`

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
	Forms []any `json:"forms,omitempty"`

	// The user (ID or struct) who completed and submitted the feedback.
	UserID json.RawMessage `json:"user,omitempty"`

	// The interview panel that the feedback is associated with, if the feedback is associated with
	// an interview.
	PanelID string `json:"panel,omitempty"`

	// The interview for which the feedback was submitted. Manually added feedback forms will not
	// be associated with an interview.
	InterviewID string `json:"interview,omitempty"`

	// Datetime when form was created.
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// This value is null when the updatedAt property has not been previously set. This is likely
	// to occur for feedback that were created prior to the introduction of this property, and have
	// not since been updated.
	UpdatedAt *int64 `json:"updatedAt,omitempty"`

	// Datetime when form was completed.
	CompletedAt *int64 `json:"completedAt,omitempty"`

	// Datetime when form was deleted.
	DeletedAt *int64 `json:"deletedAt,omitempty"`
}

// Populate a regular [model.FeedbackForm] from this [FeedbackForm].
func (f *FeedbackForm) ToModel(result *model.FeedbackForm) error {
	// Fields that map 1:1
	result.ID = f.ID
	result.Type = f.Type
	result.Text = f.Text
	result.Instructions = f.Instructions
	result.BaseTemplateID = f.BaseTemplateID
	result.Forms = f.Forms
	result.PanelID = f.PanelID
	result.InterviewID = f.InterviewID
	result.CreatedAt = f.CreatedAt
	result.UpdatedAt = f.UpdatedAt
	result.CompletedAt = f.CompletedAt
	result.DeletedAt = f.DeletedAt

	// Parse the user field
	userID, user, err := unmarshalUserOrID(f.UserID)
	if err != nil {
		return err
	}

	result.UserID = userID
	result.User = user

	return nil
}

// Unmarshal a feedback form ID or feedback form.
//   - If the raw message is empty, returns ("", nil, nil).
//   - If the raw message is a string, returns (id, nil, nil).
//   - If the raw message is a feedback form, returns (feedbackForm.ID, &feedbackForm, nil).
// func unmarshalFeedbackFormOrID(raw json.RawMessage) (string, *model.FeedbackForm, error) {
// 	if len(raw) == 0 {
// 		return "", nil, nil
// 	}

// 	if raw[0] == '"' {
// 		var id string
// 		if err := json.Unmarshal(raw, &id); err != nil {
// 			return "", nil, err
// 		}
// 		return id, nil, nil
// 	}

// 	var feedbackForm FeedbackForm
// 	if err := json.Unmarshal(raw, &feedbackForm); err != nil {
// 		return "", nil, err
// 	}

// 	var result model.FeedbackForm
// 	if err := feedbackForm.ToModel(&result); err != nil {
// 		return "", nil, err
// 	}

// 	return result.ID, &result, nil
// }

// Unmarshal an array of feedback form IDs or an array of feedback forms.
//   - If the raw message is empty, returns (nil, nil, nil).
//   - If the raw message is an array of strings, returns (ids, nil, nil).
//   - If the raw message is an array of feedback forms, returns (ids, feedbackForms, nil).
func unmarshalArrayOfFeedbackFormsOrIDs(raw json.RawMessage) ([]string, []model.FeedbackForm, error) {
	if len(raw) == 0 {
		return nil, nil, nil
	}

	// Try unmarshalling as an array of feedback forms first
	var feedbackForms []model.FeedbackForm
	var ids []string

	if err := json.Unmarshal(raw, &feedbackForms); err != nil {
		// Can't unmarshal as feedback forms; try unmarshalling as IDs.
		if err := json.Unmarshal(raw, &ids); err != nil {
			return nil, nil, err
		}

		return ids, nil, nil
	}

	for _, feedbackForm := range feedbackForms {
		ids = append(ids, feedbackForm.ID)
	}

	return ids, feedbackForms, nil
}
