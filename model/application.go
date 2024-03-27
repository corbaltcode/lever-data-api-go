package model

// When a candidate applies to a job posting, an application is created.
//
// Lever is candidate-centric, meaning that candidates can exist in the system without being
// applied to a specific job posting. However, almost all candidates are applied to job postings,
// and thus almost all candidates have one or more applications.
//
// There are three different ways that applications can be created in Lever:
//
// * Through a posting: An application is created when a candidate applies to a job posting through
// your company's public or internal job site, or is submitted by an agency.
//
// * By a user: A team member at your company manually adds a job posting to a specific candidate
// in Lever.
//
// * As a referral: A team member at your company refers the candidate into Lever for a specific
// job posting.
//
// Candidates can be applied to multiple job postings, meaning that candidates can have multiple
// applications. A candidate or contact may have multiple applications, each of which will be on a
// unique Opportunity. An Opportunity will have no more than one Application.
type Application struct {
	// Application UID
	ID string `json:"id,omitempty"`

	// Opportunity profile associated with an application.
	OpportunityID string `json:"opportunityId,omitempty"`

	// Datestamp when application was created in Lever.
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// An application can be of type referral, user, or posting. Applications of type referral are
	// created when a user refers a candidate for a job posting. Applications have type user when
	// they are applied manually to a posting in Lever. Applications have type posting when a
	// candidate applies to a job posting through your company's jobs page.
	Type string `json:"type,omitempty"`

	// Job posting to apply to candidate.
	PostingID string `json:"posting,omitempty"`

	// Job posting to apply to candidate. Returned if expand=posting is specified.
	Posting *Posting `json:"-"`

	// The owner of the job posting at the time when the candidate applies to that job.
	PostingOwnerID string `json:"postingOwner,omitempty"`

	// The owner of the job posting at the time when the candidate applies to that job. Returned if
	// expand=postingOwner is specified.
	PostingOwner *User `json:"-"`

	// The hiring manager of the job posting at the time when the candidate applies to that job.
	PostingHiringManagerID string `json:"postingHiringManager,omitempty"`

	// The hiring manager of the job posting at the time when the candidate applies to that job.
	// Returned if expand=postingHiringManager is specified.
	PostingHiringManager *User `json:"-"`

	// If the application is of type referral, this is the user who made the referral.
	UserID string `json:"user,omitempty"`

	// If the application is of type referral, this is the user who made the referral. Returned if
	// expand=user is specified.
	User *User `json:"-"`

	// Name of candidate who applied.
	Name string `json:"name,omitempty"`

	// Candidate email
	Email string `json:"email,omitempty"`

	// Candidate phone number
	Phone *Phone `json:"phone,omitempty"`

	// Candidate's current company or organization
	Company string `json:"company,omitempty"`

	// List of candidate links (e.g. personal website, LinkedIn profile, etc.)
	Links []string `json:"links,omitempty"`

	// Any additional comments from candidate included in job application
	Comments string `json:"comments,omitempty"`

	// An array of customized forms. If the application is of type referral, the custom questions
	// will include a referral form. If the application is type posting, the custom questions
	// include customized posting forms.
	CustomQuestions []any `json:"customQuestions,omitempty"`

	// Application archived status
	Archived *Archived `json:"archived,omitempty"`

	// If the application was archived as hired against a requisition, this is the data related to
	// the requisition.
	RequisitionForHire *ApplicationRequisitionForHire `json:"requisitionForHire,omitempty"`
}

// Data related to an application requisition.
//
// TODO: Determine whether this struct should be shared with other models.
type ApplicationRequisitionForHire struct {
	// The Lever id of the requisition against which the application was hired
	ID string `json:"id,omitempty"`

	// The requisitionCode field from the requisition.
	RequisitionCode string `json:"requisitionCode,omitempty"`

	// The Lever user id of the hiring manager specified on the requisition, if any.
	HiringManagerIdOnHire string `json:"hiringManagerOnHire,omitempty"`
}
