package multimodel

import (
	"encoding/json"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// The Opportunity model, but with expandable fields left unparsed.
type Opportunity struct {
	// Opportunity UID
	ID string `json:"id,omitempty"`

	// Contact full name
	Name string `json:"name,omitempty"`

	// Contact headline, typically a list of previous companies where the contact has worked or
	// schools that the contact has attended
	Headline string `json:"headline,omitempty"`

	// Contact UID
	ContactID string `json:"contact,omitempty"`

	// The stage (ID or struct) of this Opportunity's current stage
	Stage json.RawMessage `json:"stage,omitempty"`

	// An array of historical stage changes for this Opportunity
	StageChanges []model.StageChange `json:"stageChanges,omitempty"`

	// The confidentiality of the opportunity. An opportunity can only be confidential if it is
	// associated with a confidential job posting. Learn more about confidential data in the API.
	// Can be one of the following values: non-confidential, confidential.
	Confidentiality string `json:"confidentiality,omitempty"`

	// Contact current location
	Location string `json:"location,omitempty"`

	// Contact phone number(s)
	Phones []model.Phone `json:"phones,omitempty"`

	// Contact emails
	Emails []string `json:"emails,omitempty"`

	// List of Contact links (e.g. personal website, LinkedIn profile, etc.)
	Links []string `json:"links,omitempty"`

	// Opportunity archived status
	Archived *model.Archived `json:"archived,omitempty"`

	// An array containing a list of tags for this Opportunity. Tags are specified as strings,
	// identical to the ones displayed in the Lever interface.
	Tags []string `json:"tags,omitempty"`

	// An array of source strings for this Opportunity.
	Sources []string `json:"sources,omitempty"`

	// The user (ID or struct) of the user who created the opportunity.
	SourcedBy json.RawMessage `json:"sourcedBy,omitempty"`

	// The way this Opportunity was added to Levermodel. Can be one of the following values: "agency",
	// "applied", "internal", "referred", "sourced", "university"
	Origin string `json:"origin,omitempty"`

	// The user (ID or struct) of the owner of this Opportunity.
	Owner json.RawMessage `json:"owner,omitempty"`

	// An array of users (IDs or structs) of the followers of this Opportunity.
	Followers json.RawMessage `json:"followers,omitempty"`

	// An array, containing up to one Application ID (can be either an active or archived
	// Application). Each Opportunity can only have up to one application.
	ApplicationIDs []string `json:"applications,omitempty"`

	// Datetime when this Opportunity was created in Levermodel. For candidates who applied to a job
	// posting on your website, the date and time when the Opportunity was created in Lever is the
	// moment when the candidate clicked the "Apply" button on their application.
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// Datetime when this Opportunity was updated in Levermodel. This property is updated when the
	// following fields are modified: applications, archived, confidentiality, contact,
	// dataProtection, emails, followers, headline, isAnonymized, lastAdvancedAt,
	// lastInteractionAt, links, location, name, origin, owner, phones, snoozedUntil, sourcedBy,
	// sources, stage, stageChanges, tags. It is also updated when the following fields on the
	// expanded applications object are modified: archived, candidateId, comments, company,
	// customQuestions, email, links, name, opportunityId, phone, postingId, postingHiringManager,
	// postingOwner, primarySource, requisitionForHire, secondarySources, type, user.
	//
	// WARNING: The dataProtection status is based on candidate-provided consent and applicable
	// data policy regulations which can change according to many factors. The updatedAt field is
	// only updated when the candidate-provided consent changes.
	//
	// This value is null when the updatedAt property has not been previously set. This is likely
	// to occur for opportunities that were created prior to the introduction of this property,
	// and have not since been updated.
	UpdatedAt *int64 `json:"updatedAt,omitempty"`

	// Datetime when the last interaction with this Opportunity profile occurred.
	LastInteractionAt *int64 `json:"lastInteractionAt,omitempty"`

	// Datetime when the candidate advanced to the pipeline stage where they are currently located
	// in your hiring process for this Opportunity
	LastAdvancedAt *int64 `json:"lastAdvancedAt,omitempty"`

	// If this Opportunity is snoozed, the timestamp will reflect the datetime when the snooze
	// period ends
	SnoozedUntil *int64 `json:"snoozedUntil,omitempty"`

	// An object containing the list and show urls for this Opportunity.
	URLs *model.OpportunityURLs `json:"urls,omitempty"`

	// An object containing a candidate's data protection status based on candidate-provided
	// consent and applicable data policy regulations. If there is no policy in place or if no
	// policies apply to the candidate, value is null. (shared by contact)
	DataProtection *model.OpportunityDataProtection `json:"dataProtection,omitempty"`

	// Indicates whether an Opportunity has been anonymized. When all of a contact’s Opportunities
	// have been anonymized, the contact is fully anonymized and their personal information is
	// removed. Non-personal metadata may remain for accurate reporting purposes.
	IsAnonymized bool `json:"isAnonymized,omitempty"`

	// User ID of the user who deleted the Opportunity. Note that this attribute only appears for
	// deleted Opportunities.
	DeletedByID string `json:"deletedBy,omitempty"`

	// Timestamp for when the Opportunity was deleted. Note that this attribute only appears for
	// deleted Opportunities.
	DeletedAt *int64 `json:"deletedAt,omitempty"`

	// The posting location associated with the opportunity. Can be “unspecified” if not selected,
	// or absent for opportunities not associated with a posting, or for a posting that currently
	// has no location.
	OpportunityLocation string `json:"oppoLocation,omitempty"`
}

// Populate a regular [model.Opportunity] from this [multimodel.Opportunity].
func (o *Opportunity) ToModel(result *model.Opportunity) error {
	// Fields that map 1:1
	result.ID = o.ID
	result.Name = o.Name
	result.Headline = o.Headline
	result.ContactID = o.ContactID
	result.StageChanges = o.StageChanges
	result.Confidentiality = o.Confidentiality
	result.Location = o.Location
	result.Phones = o.Phones
	result.Emails = o.Emails
	result.Links = o.Links
	result.Archived = o.Archived
	result.Tags = o.Tags
	result.Sources = o.Sources
	result.Origin = o.Origin
	result.ApplicationIDs = o.ApplicationIDs
	result.CreatedAt = o.CreatedAt
	result.UpdatedAt = o.UpdatedAt
	result.LastInteractionAt = o.LastInteractionAt
	result.LastAdvancedAt = o.LastAdvancedAt
	result.SnoozedUntil = o.SnoozedUntil
	result.URLs = o.URLs
	result.DataProtection = o.DataProtection
	result.IsAnonymized = o.IsAnonymized
	result.DeletedByID = o.DeletedByID
	result.DeletedAt = o.DeletedAt
	result.OpportunityLocation = o.OpportunityLocation

	stageID, stage, err := unmarshalStageOrID(o.Stage)
	if err != nil {
		return err
	}

	result.StageID = stageID
	result.Stage = stage

	sourcedByID, sourcedBy, err := unmarshalUserOrID(o.SourcedBy)
	if err != nil {
		return err
	}
	result.SourcedByID = sourcedByID
	result.SourcedBy = sourcedBy

	ownerID, owner, err := unmarshalUserOrID(o.Owner)
	if err != nil {
		return err
	}
	result.OwnerID = ownerID
	result.Owner = owner

	followerIDs, followers, err := unmarshalArrayOfUsersOrIDs(o.Followers)
	if err != nil {
		return err
	}
	result.FollowerIDs = followerIDs
	result.Followers = followers

	return nil
}
