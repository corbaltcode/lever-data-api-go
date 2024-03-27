package model

// Opportunities
//
// "Candidates" are individuals who have been added to your Lever account as potential fits for
// your open job positions. "Opportunities" represent each of an individual’s unique candidacies or
// journeys through your pipeline for a given job position, meaning a single Candidate can be
// associated with multiple Opportunities. A “Contact” is a unique individual who may or may not
// have multiple candidacies or Opportunities.
//
// Candidates enter your pipeline for a new Opportunity by:
//
// * Applying to a posting on your jobs site
//
// * Being added by an external recruiting agency
//
// * Being referred by an employee
//
// * Being manually added by a Lever user, or
//
// * Being sourced from an online profile.
//
// Each Opportunity can have their own notes, feedback, interview schedules, and additional forms.
// An opportunity may be “confidential” if it is moving through your pipeline for a job posting
// that has been created as confidential. Opportunities exit your pipeline by being archived for
// one of two reasons: (1) The candidate was rejected for the opportunity, or (2) The candidate was
// hired for the opportunity.
//
// A "Contact" is an object that our application uses internally to identify an individual person
// and their personal or contact information, even though they may have multiple opportunities.
// From this API, the "Contact" is exposed via the contact field, which returns the unique ID for a
// Contact across your account. Contact information will be shared and consistent across an
// individual person's opportunities, and will continue to be aggregated onto individual
// opportunities in the responses to all GET and POST requests to /opportunities.
//
// WARNING: These Opportunities endpoints should be used instead of the now deprecated Candidates
// endpoints. Prior to the migration, for any given Candidate, the candidateId you would use for a
// request to a Candidates endpoint can be used as the opportunityId in a request to the
// corresponding Opportunities endpoint.
//
// Going forward, the contact field is the unique identifier for a Contact or an individual person
// in Lever, so all integrations should be built and updated using the contact as the unique person
// identifier and opportunityId as a specific opportunity or candidacy moving through the pipeline.
type Opportunity struct {
	// Opportunity UID
	ID string

	// Contact full name
	Name string

	// Contact headline, typically a list of previous companies where the contact has worked or
	// schools that the contact has attended
	Headline string

	// Contact UID
	ContactID string

	// The stage ID of this Opportunity's current stage
	StageID string

	// The stage of this Opportunity's current stage. Returned if expand=stage is specified.
	Stage *Stage

	// An array of historical stage changes for this Opportunity
	StageChanges []StageChange

	// The confidentiality of the opportunity. An opportunity can only be confidential if it is
	// associated with a confidential job posting. Learn more about confidential data in the API.
	// Can be one of the following values: non-confidential, confidential.
	Confidentiality string

	// Contact current location
	Location string

	// Contact phone number(s)
	Phones []Phone

	// Contact emails
	Emails []string

	// List of Contact links (e.g. personal website, LinkedIn profile, etc.)
	Links []string

	// Opportunity archived status
	Archived *Archived

	// An array containing a list of tags for this Opportunity. Tags are specified as strings,
	// identical to the ones displayed in the Lever interface.
	Tags []string

	// An array of source strings for this Opportunity.
	Sources []string

	// The user ID of the user who created the opportunity.
	SourcedByID string

	// The user who created the opportunity. Returned if expand=sourcedBy is specified.
	SourcedBy *User

	// The way this Opportunity was added to Lever. Can be one of the following values: "agency",
	// "applied", "internal", "referred", "sourced", "university"
	Origin string

	// The user ID of the owner of this Opportunity.
	OwnerID string

	// The owner of this Opportunity. Returned if expand=owner is specified.
	Owner *User

	// An array of user IDs of the followers of this Opportunity.
	FollowerIDs []string

	// An array of users who are following this Opportunity. Returned if expand=followers is
	// specified.
	Followers []User

	// An array, containing up to one Application ID (can be either an active or archived
	// Application). Each Opportunity can only have up to one application.
	ApplicationIDs []string

	// Datetime when this Opportunity was created in Lever. For candidates who applied to a job
	// posting on your website, the date and time when the Opportunity was created in Lever is the
	// moment when the candidate clicked the "Apply" button on their application.
	CreatedAt *int64

	// Datetime when this Opportunity was updated in Lever. This property is updated when the
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
	UpdatedAt *int64

	// Datetime when the last interaction with this Opportunity profile occurred.
	LastInteractionAt *int64

	// Datetime when the candidate advanced to the pipeline stage where they are currently located
	// in your hiring process for this Opportunity
	LastAdvancedAt *int64

	// If this Opportunity is snoozed, the timestamp will reflect the datetime when the snooze
	// period ends
	SnoozedUntil *int64

	// An object containing the list and show urls for this Opportunity.
	URLs *OpportunityURLs

	// An object containing a candidate's data protection status based on candidate-provided
	// consent and applicable data policy regulations. If there is no policy in place or if no
	// policies apply to the candidate, value is null. (shared by contact)
	DataProtection *OpportunityDataProtection

	// Indicates whether an Opportunity has been anonymized. When all of a contact’s Opportunities
	// have been anonymized, the contact is fully anonymized and their personal information is
	// removed. Non-personal metadata may remain for accurate reporting purposes.
	IsAnonymized bool

	// User ID of the user who deleted the Opportunity. Note that this attribute only appears for
	// deleted Opportunities.
	DeletedByID string

	// Timestamp for when the Opportunity was deleted. Note that this attribute only appears for
	// deleted Opportunities.
	DeletedAt *int64

	// The posting location associated with the opportunity. Can be “unspecified” if not selected,
	// or absent for opportunities not associated with a posting, or for a posting that currently
	// has no location.
	OpportunityLocation string
}

// An object containing a candidate's data protection status based on candidate-provided consent
// and applicable data policy regulations.
//
// TODO: Determine whether this struct should be shared with other models.
type OpportunityDataProtection struct {
	// The candidate's consent status for contacting them.
	Contact *OpportunityDataProtectionConsent `json:"contact,omitempty"`

	// The candidate's consent status for storing their data
	Store *OpportunityDataProtectionConsent `json:"store,omitempty"`
}

// An object representing the consent status for a processing activity
//
// TODO: Determine whether this struct should be shared with other models.
type OpportunityDataProtectionConsent struct {
	// True if the applicable data policy regulation allows for storage of this record.
	Allowed bool `json:"allowed,omitempty"`

	// Timestamp of when this permission expires.
	ExpiresAt *int64 `json:"expiresAt,omitempty"`
}

// An object containing the list and show urls for this Opportunity.
//
// TODO: Determine whether this struct should be shared with other models.
type OpportunityURLs struct {
	// URL that points to the account's list of candidates
	List string `json:"list,omitempty"`

	// URL that points to the candidate's profile page for this Opportunity
	Show string `json:"show,omitempty"`
}
