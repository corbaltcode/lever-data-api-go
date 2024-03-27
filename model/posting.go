package model

// Job postings organize candidates based on the specific roles that they may fit into on your
// growing team.
//
// There are four different states of job postings: published, internal, closed, and draft.
// NOTE: In the Lever app, we refer to internal postings as “unlisted” postings. For organizations
// that have enabled job posting approvals, there are two additional states: pending and rejected.
//
// * published job postings are postings that appear on your public job site, internal job site, or
// both, depending on which distribution channels you have configured. You'll want to mark a job
// posting as published when you would like the posting to appear on one or more of your job sites.
//
// * internal job postings are postings that will NOT appear on your public or internal job sites.
// In the Lever app, internal postings are referred to as “unlisted” postings. You can, however,
// apply candidates to these internal postings. In addition, anyone with the link to an internal
// posting can apply, and people can refer candidates to any internal postings.
//
// * closed job postings are postings that will NOT appear on your public or internal job sites,
// likely because they are roles that you've already filled! You'll want to keep a job posting in
// the closed state instead of deleting it so that you can easily build reports on this specific
// job posting. If you and your team re-open the role, you can always move a closed posting back
// to published or internal.
//
// * draft job postings are postings that you are still in the process of finalizing. draft
// postings will NOT appear on your public or internal job sites. Once you've finalized a draft job
// posting, you'll want to change the state of the posting to published or internal.
//
// * pending job postings are postings that are awaiting approval. This state is applicable for
// organizations that are using the job posting approval workflows in Lever.
//
// * rejected job postings are postings that were submitted for approval and rejected. This state
// is applicable for organizations that are using the job posting approval workflows in Lever.
type Posting struct {
	// Posting UID
	ID string `json:"id,omitempty"`

	// Title of the job posting
	Text string `json:"text,omitempty"`

	// Datetime when posting was created in Lever
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// Datetime when posting was last updated
	UpdatedAt *int64 `json:"updatedAt,omitempty"`

	// Posting's current status
	State string `json:"state,omitempty"`

	// Array of job sites that a published posting appears on.
	DistributionChannels []string `json:"distributionChannels,omitempty"`

	// The confidentiality of the posting. It is not possible to update a posting’s
	// confidentiality. Can be one of the following values: non-confidential, confidential.
	Confidentiality string `json:"confidentiality,omitempty"`

	// The user ID of the user who created the posting.
	UserID string `json:"user,omitempty"`

	// The user who created the posting. Returned if expand=user is specified.
	User *User `json:"-"`

	// The user ID of the posting owner. The posting owner is the individual who is directly
	// responsible for managing all candidates who are applied to that role.
	OwnerID string `json:"owner,omitempty"`

	// The posting owner. Returned if expand=owner is specified.
	Owner *User `json:"-"`

	// The user ID of the hiring manager for the job posting.
	HiringManagerID string `json:"hiringManager,omitempty"`

	// The hiring manager for the job posting. Returned if expand=hiringManager is specified.
	HiringManager *User `json:"-"`

	// An object containing the tags of various categories.
	Categories *PostingCategories `json:"categories,omitempty"`

	// An array of additional posting tags.
	Tags []string `json:"tags,omitempty"`

	// Content of the job posting including any custom questions that you've built into the job
	// application.
	Content *PostingContent `json:"content,omitempty"`

	// An ISO 3166-1 alpha-2 code for a country / territory
	Country string `json:"country,omitempty"`

	// An array of user IDs of the followers of this posting.
	FollowerIds []string `json:"followers,omitempty"`

	// An array of users who are following this posting. Returned if expand=followers is specified.
	Followers []User `json:"-"`

	// Requisition code associated with this posting.
	// WARNING: This field is deprecated but maintained for backwards compatibility.
	RequisitionCode string `json:"requisitionCode,omitempty"`

	// Array of requisition codes associated with this posting.
	RequisitionCodes []string `json:"requisitionCodes,omitempty"`

	// An object containing the list, show and apply URLs for the job posting.
	URLs *PostingURLs `json:"urls,omitempty"`

	// Workplace type of this posting. Defaults to 'unspecified'. Can be one of the following
	// values: onsite, remote, hybrid
	WorkplaceType string `json:"workplaceType,omitempty"`
}

// An object containing the tags of various categories.
type PostingCategories struct {
	// Tag for the team to which the job posting belongs (e.g. Sales, Engineering)
	Team string `json:"team,omitempty"`

	// Tag for the department to which the job posting's team belongs, if present
	Department string `json:"department,omitempty"`

	// Tag for job position location
	Location string `json:"location,omitempty"`

	// An array of strings containing all the locations associated with the posting.
	AllLocations []string `json:"allLocations,omitempty"`

	// Tag for job position work type (e.g. Full-time, Part-time, Internship)
	Commitment string `json:"commitment,omitempty"`

	// Tag for job posting level (e.g. Senior, Junior).
	// Deprecated but currently maintained for backward compatibility.
	Level string `json:"level,omitempty"`
}

// Content of the job posting including any custom questions that you've built into the job
// application.
//
// TODO: Determine whether this struct should be shared with other models.
type PostingContent struct {
	// Job posting description that is shown at the top of the jobs pagem, as plaintext.
	Description string `json:"description,omitempty"`

	// Job posting description that is shown at the top of the jobs page.
	DescriptionHtml string `json:"descriptionHtml,omitempty"`

	// Lists of requirements, responsibilities, etc. that have been added to this posting
	Lists []struct {
		// Title of the list.
		Title string `json:"title,omitempty"`

		// Content of the list, as styled HTML.
		Content string `json:"content,omitempty"`
	} `json:"lists,omitempty"`

	// Closing statement on job posting, as plaintext.
	Closing string `json:"closing,omitempty"`

	// Closing statement on job posting, as styled HTML.
	ClosingHtml string `json:"closingHtml,omitempty"`
}

// An object containing the list, show and apply urls for the job posting.
//
// TODO: Determine whether this struct should be shared with other models.
type PostingURLs struct {
	// URL that points to the account's list of job postings
	List string `json:"list,omitempty"`

	// URL that points to the posting's information page
	Show string `json:"show,omitempty"`

	// URL that points to the posting's apply page
	Apply string `json:"apply,omitempty"`
}
