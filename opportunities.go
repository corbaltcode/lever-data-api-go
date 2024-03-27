package lever

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"path/filepath"
	"strconv"

	"github.com/corbaltcode/lever-data-api-go/internal/multimodel"
	"github.com/corbaltcode/lever-data-api-go/model"
)

// Lever opportunities client interface
type OpportunitiesClientInterface interface {
	// Retrieve a single opportunity
	GetOpportunity(ctx context.Context, req *GetOpportunityRequest) (*GetOpportunityResponse, error)

	// List all opportunities
	//
	// Lists all pipeline Opportunities for Contacts in your Lever account.
	ListOpportunities(ctx context.Context, req *ListOpportunitiesRequest) (*ListOpportunitiesResponse, error)

	// List deleted opportunities
	//
	// Lists all deleted Opportunities in your Lever account.
	ListDeletedOpportunities(ctx context.Context, req *ListDeletedOpportunitiesRequest) (*ListDeletedOpportunitiesResponse, error)

	// Create an opportunity
	//
	// This endpoint enables integrations to create candidates and opportunities in your Lever
	// account.
	//
	// If you want to apply a candidate to a job posting or create a custom job site, you should
	// use the Lever Postings API instead of the Lever Data API.
	//
	// We accept requests of type application/json and multipart/form-data. If you are including a
	// resume or other files, you must use the multipart/form-data type.
	//
	// There are many ways to create a candidate. Here are some examples:
	//
	// * Provide a JSON representation of a candidate with basic information like candidate name,
	// email, and phone number
	//
	// * Upload just a resume file and specify you'd like the resume parsed. Information parsed
	// from the resume will be used to create the candidate—their name, email, and phone number,
	// for example.
	//
	// Note: If you are creating a confidential opportunity, you must provide a posting UID for a
	// confidential job posting. Learn more about confidential data in the API.
	//
	// If candidate information is provided in the POST request and resume parsing is requested,
	// the manually provided information will always take precedence over the parsed information
	// from the candidate resume.
	//
	// All fields are optional, but an empty candidate is not particularly interesting. All query
	// parameters except the perform_as parameter are optional.
	//
	// If an email address is provided, we will always attempt to dedupe the candidate. If a match
	// is found, we will create a new Opportunity that is linked to the existing matching
	// candidate’s contact (i.e. we never create a new contact, or person, if a match has been
	// found). The existing candidate’s contact data will take precedence over new manually
	// provided information.
	//
	// If a contact already exists for a candidate, the ID of the existing contact may be provided
	// in the POST request to create an opportunity associated with the existing candidate's
	// contact (the candidate will be deduped). If additional contact details are included in the
	// request (emails, phones, tags, web links), these will be added to the existing candidate's
	// contact information.
	CreateOpportunity(ctx context.Context, req *CreateOpportunityRequest) (*CreateOpportunityResponse, error)

	// Update opportunity stage
	//
	// Change an Opportunity's current stage
	UpdateOpportunityStage(ctx context.Context, req *UpdateOpportunityStageRequest) (*UpdateOpportunityStageResponse, error)

	// Update opportunity archived state
	//
	// Update an Opportunity's archived state. If an Opportunity is already archived, its archive
	// reason can be changed or if null is specified as the reason, it will be unarchived. If an
	// Opportunity is active, it will be archived with the reason provided.
	//
	// The requisitionId is optional. If the provided reason maps to ‘Hired’ and a requisition is
	// provided, the Opportunity will be marked as Hired, the active offer is removed from the
	// requisition, and the hired count for the requisition will be incremented.
	//
	// If a requisition is specified and there are multiple active applications on the profile,
	// you will receive an error. If the specific requisition is closed, you will receive an error.
	// If there is an offer extended, it must be signed, and the offer must be associated with an
	// application for a posting linked to the provided requisition. You can hire a candidate
	// against a requisition without an offer.
	UpdateOpportunityArchivedState(ctx context.Context, req *UpdateOpportunityArchivedStateRequest) (*UpdateOpportunityArchivedStateResponse, error)

	// Add contact links to a contact by opportunity
	//
	// Add links to a Contact by an Opportunity
	AddOpportunityLinks(ctx context.Context, req *AddOpportunityLinksRequest) (*AddOpportunityLinksResponse, error)

	// Remove contact links from a contact by opportunity
	//
	// Remove links from a Contact by an Opportunity
	RemoveOpportunityLinks(ctx context.Context, req *RemoveOpportunityLinksRequest) (*RemoveOpportunityLinksResponse, error)

	// Add tags to an opportunity
	AddOpportunityTags(ctx context.Context, req *AddOpportunityTagsRequest) (*AddOpportunityTagsResponse, error)

	// Remove tags from an opportunity
	RemoveOpportunityTags(ctx context.Context, req *RemoveOpportunityTagsRequest) (*RemoveOpportunityTagsResponse, error)

	// Add sources to an opportunity
	AddOpportunitySources(ctx context.Context, req *AddOpportunitySourcesRequest) (*AddOpportunitySourcesResponse, error)

	// Remove sources from an opportunity
	RemoveOpportunitySources(ctx context.Context, req *RemoveOpportunitySourcesRequest) (*RemoveOpportunitySourcesResponse, error)
}

// Parameters for retrieving a single opportunity.
type GetOpportunityRequest struct {
	BaseRequest

	// The opportunity id. This is required.
	OpportunityID string
}

// Create a new GetOpportunityRequest with the required fields.
func NewGetOpportunityRequest(opportunityID string) *GetOpportunityRequest {
	return &GetOpportunityRequest{
		OpportunityID: opportunityID,
	}
}

func (r *GetOpportunityRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s", url.PathEscape(r.OpportunityID))
}

// Response for retrieving a single opportunity; returned to client users.
type GetOpportunityResponse struct {
	BaseResponse

	// The opportunity record.
	Opportunity *model.Opportunity `json:"data"`
}

// JSON response type for retrieving a single opportunity, with some field types dynamically determined.
type getOpportunityResponseJSON struct {
	BaseResponse

	// The opportunity record.
	Opportunity *multimodel.Opportunity `json:"data"`
}

// Parameters for listing opportunities.
type ListOpportunitiesRequest struct {
	BaseListRequest

	// If specified, filter Opportunities by tag (case sensitive). Results will include
	// Opportunities that contain the specified tag. Multiple tags can be specified and results
	// will include a union of result sets (i.e. Opportunities that have either tag).
	Tags []string

	// If specified, filter Opportunities by email addresses. Results will include Opportunities
	// for Contacts that contain the canonicalized email address.
	Emails []string

	// If specified, filter Opportunities by origin. Results will include Opportunities that
	// contain the specified origin. Multiple origins can be specified and results will include a
	// union of result sets (i.e. Opportunities from either origin).
	Origins []string

	// If specified, filter Opportunities by source. Results will include Opportunities that
	// contain the specified source tag. Multiple sources can be specified and results will include
	// a union of result sets (i.e. Opportunities from either source).
	Sources []string

	// If specified, filter opportunities by confidentiality. If unspecified, defaults to
	// non-confidential. To get both confidential and non-confidential opportunities you must
	// specify all. Learn more about confidential data in the API.
	Confidentiality []string

	// If specified, filter Opportunities by current stage. Results will include Opportunities that
	// are currently in the specified stage. Multiple stages can be specified and results will
	// include a union of result sets (i.e. Opportunities that are in either stage).
	StageIDs []string

	// If specified, filter Opportunities by posting. Results will include Opportunities that are
	// applied to the specified posting. Multiple postings can be specified and results will
	// include a union of result sets (i.e. Opportunities that are applied to either posting).
	PostingIDs []string

	// If specified, filter Opportunities by postings for which they have been archived. Results
	// will include opportunities for candidates that applied to the specified posting and then the
	// application was archived. Multiple postings can be specified and results will include a
	// union of result sets (i.e. Opportunities that were applied to either posting).
	ArchivedPostingIDs []string

	// Filter Opportunities by the timestamp they were created. If only CreatedAtStart is
	// Specified, all Opportunities created from that timestamp (inclusive) to the present will be
	// included. If only CreatedAtEnd is specified, all Opportunities created before that timestamp
	// (inclusive) are included.
	CreatedAtStart *int64
	CreatedAtEnd   *int64

	// Filter Opportunities by the timestamp they were last updated. If only UpdatedAtStart is
	// specified, all Opportunities updated from that timestamp (inclusive) to the present will be
	// included. If only UpdatedAtEnd is specified, all Opportunities updated before that timestamp
	// (inclusive) are included.
	UpdatedAtStart *int64
	UpdatedAtEnd   *int64

	// Filter Opportunities by the timestamp they were advanced to their current stage. If only
	// AdvancedAtStart is specified, all Opportunities advanced from that timestamp (inclusive) to
	// the present will be included. If only AdvancedAtEnd is specified, all Opportunities advanced
	// before that timestamp (inclusive) are included.
	AdvancedAtStart *int64
	AdvancedAtEnd   *int64

	// Filter Opportunities by the timestamp they were archived. If only ArchivedAtStart is
	// specified, all Opportunities archived from that timestamp (inclusive) to the present will be
	// included. If only ArchivedAtEnd is specified, all Opportunities archived before that
	// timestamp (inclusive) are included.
	ArchivedAtStart *int64
	ArchivedAtEnd   *int64

	// If specified, filter Opportunities by archive status. If unspecified, results include both
	// archived and unarchived Opportunities. If true, results only include archived Opportunities.
	// If false, results only include active Opportunities.
	Archived *bool

	// If specified, filter Opportunities by archive reason ID. Results will include Opportunities
	// that have been archived with the specified reason. Multiple archive reasons can be specified
	// and results will include a union of result sets (i.e. Opportunities that have been archived
	// for either reason).
	ArchiveReasonIDs []string

	// If specified, filter Opportunities by snoozed status. If unspecified, results include both
	// snoozed and unsnoozed Opportunities. If true, results only include snoozed Opportunities. If
	// false, results only include unsnoozed Opportunities.
	Snoozed *bool

	// If specified, filter Opportunities by contact ID. Results will include the Opportunities
	// that match the specified contact. Multiple contacts can be specified and results will
	// include a union of result sets (i.e. Opportunities that match each of the contacts).
	ContactIDs []string

	// If specified, filter opportunities by the posting location associated with the opportunity.
	// Results will include Opportunities that contain the specified opportunity location. Multiple
	// opportunity locations can be specified and results will include a union of result sets (i.e.
	// Opportunities that have either opportunity location).
	Locations []string
}

// Create a new ListOpportunitiesRequest with the required fields.
func NewListOpportunitiesRequest() *ListOpportunitiesRequest {
	return &ListOpportunitiesRequest{}
}

func (r *ListOpportunitiesRequest) GetPath() string {
	return "opportunities"
}

func (r *ListOpportunitiesRequest) AddAPIQueryParams(query *url.Values) {
	r.BaseListRequest.AddAPIQueryParams(query)

	for _, tag := range r.Tags {
		query.Add(paramTag, tag)
	}

	for _, email := range r.Emails {
		query.Add(paramEmail, email)
	}

	for _, origin := range r.Origins {
		query.Add(paramOrigin, origin)
	}

	for _, source := range r.Sources {
		query.Add(paramSource, source)
	}

	for _, confidentiality := range r.Confidentiality {
		query.Add(paramConfidentiality, confidentiality)
	}

	for _, stageID := range r.StageIDs {
		query.Add(paramStageID, stageID)
	}

	for _, postingID := range r.PostingIDs {
		query.Add(paramPostingID, postingID)
	}

	for _, archivedPostingID := range r.ArchivedPostingIDs {
		query.Add(paramArchivedPostingID, archivedPostingID)
	}

	if r.CreatedAtStart != nil {
		query.Add(paramCreatedAtStart, fmt.Sprint(*r.CreatedAtStart))
	}

	if r.CreatedAtEnd != nil {
		query.Add(paramCreatedAtEnd, fmt.Sprint(*r.CreatedAtEnd))
	}

	if r.UpdatedAtStart != nil {
		query.Add(paramUpdatedAtStart, fmt.Sprint(*r.UpdatedAtStart))
	}

	if r.UpdatedAtEnd != nil {
		query.Add(paramUpdatedAtEnd, fmt.Sprint(*r.UpdatedAtEnd))
	}

	if r.AdvancedAtStart != nil {
		query.Add(paramAdvancedAtStart, fmt.Sprint(*r.AdvancedAtStart))
	}

	if r.AdvancedAtEnd != nil {
		query.Add(paramAdvancedAtEnd, fmt.Sprint(*r.AdvancedAtEnd))
	}

	if r.ArchivedAtStart != nil {
		query.Add(paramArchivedAtStart, fmt.Sprint(*r.ArchivedAtStart))
	}

	if r.ArchivedAtEnd != nil {
		query.Add(paramArchivedAtEnd, fmt.Sprint(*r.ArchivedAtEnd))
	}

	if r.Archived != nil {
		query.Add(paramArchived, fmt.Sprint(*r.Archived))
	}

	for _, archiveReasonID := range r.ArchiveReasonIDs {
		query.Add(paramArchiveReasonID, archiveReasonID)
	}

	if r.Snoozed != nil {
		query.Add(paramSnoozed, fmt.Sprint(*r.Snoozed))
	}

	for _, contactID := range r.ContactIDs {
		query.Add(paramContactID, contactID)
	}

	for _, location := range r.Locations {
		query.Add(paramLocation, location)
	}
}

// Response for listing opportunities; returned to client users.
type ListOpportunitiesResponse struct {
	BaseListResponse

	// The opportunity records.
	Opportunities []model.Opportunity
}

// JSON response type for listing opportunities, with some field types dynamically determined.
type listOpportunitiesResponseJSON struct {
	BaseListResponse

	// The opportunity records.
	Opportunities []multimodel.Opportunity `json:"data"`
}

// Parameters for listing deleted opportunities.
type ListDeletedOpportunitiesRequest struct {
	BaseListRequest

	// If specified, filter deleted Opportunities by the timestamp they were deleted. If only
	// DeletedAtStart is specified, all Opportunities deleted from that timestamp (inclusive) to
	// the present will be included. If only DeletedAtEnd is specified, all Opportunities deleted
	// before that timestamp (inclusive) are included.
	DeletedAtStart *int64
	DeletedAtEnd   *int64
}

// Create a new ListDeletedOpportunitiesRequest with the required fields.
func NewListDeletedOpportunitiesRequest() *ListDeletedOpportunitiesRequest {
	return &ListDeletedOpportunitiesRequest{}
}

func (r *ListDeletedOpportunitiesRequest) GetPath() string {
	return "opportunities/deleted"
}

func (r *ListDeletedOpportunitiesRequest) AddAPIQueryParams(query *url.Values) {
	r.BaseListRequest.AddAPIQueryParams(query)

	if r.DeletedAtStart != nil {
		query.Add(paramDeletedAtStart, fmt.Sprint(*r.DeletedAtStart))
	}

	if r.DeletedAtEnd != nil {
		query.Add(paramDeletedAtEnd, fmt.Sprint(*r.DeletedAtEnd))
	}
}

// Response for listing deleted opportunities; returned to client users.
type ListDeletedOpportunitiesResponse struct {
	BaseListResponse

	// The opportunity records.
	Opportunities []model.Opportunity `json:"data"`
}

// JSON response type for listing deleted opportunities, with some field types dynamically determined.
type listDeletedOpportunitiesResponseJSON struct {
	BaseListResponse

	// The opportunity records.
	Opportunities []multimodel.Opportunity `json:"data"`
}

// Parameters for creating an opportunity.
type CreateOpportunityRequest struct {
	BaseRequest

	// Perform this create on behalf of a specified user. The creator and the owner of this
	// Opportunity will default to the PerformAsID user. The owner can be explicitly specified in
	// the request body if you want the owner to be a different person.
	PerformAsID string

	// If unspecified, assumed to be false. If set to true and a resume file is provided, the
	// resume will be parsed and extracted data will be used to autofill information about the
	// contact such as email and phone number. Any fields manually passed to the endpoint take
	// precedence over any parsed data.
	Parse bool

	// If unspecified, assumed to be false and the Opportunity owner will default to the
	// PerformAsID user. If set to true, an array containing a single posting UID must be
	// passed in via the postings field. The Opportunity owner will be set to that of the posting
	// owner for the single posting. If the posting does not have an owner, the Opportunity owner
	// will default to the PerformAsUserID user.
	PerformAsPostingOwner bool

	// Contact full name
	Name string

	// Contact headline, typically a list of previous companies where the contact has worked or
	// schools that the contact has attended This field can also be populated by parsing a provided
	// resume file.
	Headline string

	// The stage ID of this Opportunity's current stage If omitted, the Opportunity will be placed
	// into the "New Lead" stage.
	StageID string

	// Contact current location
	Location string

	// Contact phone number(s)
	Phones []model.Phone

	// Contact emails
	Emails []string

	// List of Contact links (e.g. personal website, LinkedIn profile, etc.)
	Links []string

	// An array containing a list of tags to apply to this Opportunity. Tags are specified as
	// strings, identical to the ones displayed in the Lever interface. If you specify a tag that
	// does not exist yet, it will be created.
	Tags []string

	// An array containing a list of sources to apply to this Opportunity. Sources are specified as
	// strings, identical to the ones displayed in the Lever interface. If you specify a source that
	// does not exist yet, it will be created.
	Sources []string

	// The way this Opportunity was added to Lever. Can be one of the following values: agency,
	// applied, internal, referred, sourced, university
	Origin string

	// The user ID of the owner of this Opportunity. If not specified, Opportunity owner defaults
	// to the PerformAsID user.
	OwnerID string

	// An array of user IDs that should be added as followers to this Opportunity. The Opportunity
	// creator will always be added as a follower.
	FollowerIDs []string

	// Resume file for this Opportunity.
	ResumeFile *model.Reader

	// File(s) relating to this Opportunity.
	Files []model.Reader

	// Posting ID for this opportunity.
	PostingID string

	// To create a historical Opportunity, set the create time for an Opportunity in the past.
	// Default is current datetime. Note that time travel in the other direction is not permitted;
	// you cannot create a candidate in the future.
	CreatedAt *int64

	// Opportunity archived status. You must specify this field if you would like the candidate to
	// be archived for the created Opportunity. This is useful if you'd like to import historical
	// data into Lever and the Opportunities you are creating are not active. The archive reason
	// must be specified for an archived Opportunity (if you just set the ArchivedAt we will ignore
	// it). If you only specify an archive reason, archivedAt defaults to the current datetime. If
	// you specify an ArchivedAt datetime, you must specify a createdAt datetime that occurs before
	// ArchivedAt.
	Archived *model.Archived

	// The contact ID of an existing candidate's contact to be associated with an opportunity. If
	// specified, the created opportunity will be linked to the existing candidate's contact. If
	// not specified, the attempt to dedupe a candidate by finding a match to the email provided in
	// the POST request will be done.
	ContactID string

	// Content type value for the multipart/form-data boundary string
	contentType string
}

// Create a new CreateOpportunityRequest with the required fields.
func NewCreateOpportunityRequest(performAsID string) *CreateOpportunityRequest {
	return &CreateOpportunityRequest{
		PerformAsID: performAsID,
	}
}

func (r *CreateOpportunityRequest) GetPath() string {
	return "opportunities"
}

func (r *CreateOpportunityRequest) GetHTTPMethod() string {
	return http.MethodPost
}

func (r *CreateOpportunityRequest) AddAPIQueryParams(query *url.Values) {
	r.BaseRequest.AddAPIQueryParams(query)

	if r.PerformAsID != "" {
		query.Add(paramPerformAs, r.PerformAsID)
	}

	if r.Parse {
		query.Add(paramParse, "true")
	}

	if r.PerformAsPostingOwner {
		query.Add(paramPerformAsPostingOwner, "true")
	}
}

func (r *CreateOpportunityRequest) GetBody() (io.Reader, error) {
	reader, pipeWriter := io.Pipe()
	writer := multipart.NewWriter(pipeWriter)
	r.contentType = writer.FormDataContentType()

	go r.writeBody(writer)
	return reader, nil
}

func (r *CreateOpportunityRequest) GetContentType() string {
	return r.contentType
}

// writeBody writes the body of the request to the provided writer.
func (r *CreateOpportunityRequest) writeBody(w *multipart.Writer) {
	defer w.Close()

	if r.Name != "" {
		w.WriteField("name", r.Name)
	}

	if r.Headline != "" {
		w.WriteField("headline", r.Headline)
	}

	if r.StageID != "" {
		w.WriteField("stage", r.StageID)
	}

	if r.Location != "" {
		w.WriteField("location", r.Location)
	}

	for i, phone := range r.Phones {
		if phone.Type != "" {
			w.WriteField(fmt.Sprintf("phones[%d][type]", i), phone.Type)
		}
		if phone.Value != "" {
			w.WriteField(fmt.Sprintf("phones[%d][value]", i), phone.Value)
		}
	}

	for _, email := range r.Emails {
		w.WriteField("emails", email)
	}

	for _, link := range r.Links {
		w.WriteField("links", link)
	}

	for _, tag := range r.Tags {
		w.WriteField("tags", tag)
	}

	for _, source := range r.Sources {
		w.WriteField("sources", source)
	}

	if r.Origin != "" {
		w.WriteField("origin", r.Origin)
	}

	if r.OwnerID != "" {
		w.WriteField("owner", r.OwnerID)
	}

	for _, followerID := range r.FollowerIDs {
		w.WriteField("followers", followerID)
	}

	if r.PostingID != "" {
		w.WriteField("posting", r.PostingID)
	}

	if r.CreatedAt != nil {
		w.WriteField("createdAt", strconv.FormatInt(*r.CreatedAt, 10))
	}

	if r.Archived != nil {
		if r.Archived.ArchivedAt != nil {
			w.WriteField("archived[archivedAt]", strconv.FormatInt(*r.Archived.ArchivedAt, 10))
		}

		if r.Archived.ReasonID != "" {
			w.WriteField("archived[reason]", r.Archived.ReasonID)
		}
	}

	if r.ContactID != "" {
		w.WriteField("contact", r.ContactID)
	}

	if r.ResumeFile != nil {
		if err := writeMultipartFile(w, "resume", r.ResumeFile); err != nil {
			return
		}
	}

	for i, file := range r.Files {
		if err := writeMultipartFile(w, fmt.Sprintf("files[%d]", i), &file); err != nil {
			return
		}
	}
}

// Write a file to the request body.
func writeMultipartFile(w *multipart.Writer, fieldName string, file *model.Reader) error {
	defer file.Contents.Close()

	h := make(textproto.MIMEHeader)
	h.Set(
		headerContentDisposition,
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes(fieldName), escapeQuotes(file.Name)),
	)

	// Set the MIME type for the file, guessing it if necessary.
	if file.MIMEType != "" {
		h.Set(headerContentType, file.MIMEType)
	} else {
		contentType := mime.TypeByExtension(filepath.Ext(file.Name))
		if contentType == "" {
			contentType = mimeTypeApplicationOctetStream
		}

		h.Set(headerContentType, contentType)
	}

	fileWriter, err := w.CreatePart(h)
	if err != nil {
		return err
	}

	_, err = io.Copy(fileWriter, file.Contents)
	return err
}

// Response for creating an opportunity.
type CreateOpportunityResponse struct {
	BaseResponse

	// Whether this opportunity was deduplicated.
	Deduped bool

	// The opportunity record.
	Opportunity *model.Opportunity
}

// JSON response type for creating an opportunity, with some field types dynamically determined.
type createOpportunityResponseJSON struct {
	BaseResponse

	// Whether this opportunity was deduplicated.
	Deduped bool `json:"deduped"`

	// The opportunity record.
	Opportunity *multimodel.Opportunity `json:"data"`
}

// Parameters for updating an opportunity stage.
type UpdateOpportunityStageRequest struct {
	BaseRequest

	// The opportunity id. This is required.
	OpportunityID string

	// The stage id of this Opportunity's current stage.
	StageID string

	// Perform this update on behalf of a specified user.
	PerformAsID string
}

// Create a new UpdateOpportunityStageRequest with the required fields.
func NewUpdateOpportunityStageRequest(opportunityID, stageID string) *UpdateOpportunityStageRequest {
	return &UpdateOpportunityStageRequest{
		OpportunityID: opportunityID,
		StageID:       stageID,
	}
}

func (r *UpdateOpportunityStageRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/stage", url.PathEscape(r.OpportunityID))
}

func (r *UpdateOpportunityStageRequest) GetMethod() string {
	return http.MethodPut
}

func (r *UpdateOpportunityStageRequest) AddAPIQueryParams(query *url.Values) {
	r.BaseRequest.AddAPIQueryParams(query)

	if r.PerformAsID != "" {
		query.Add(paramPerformAs, r.PerformAsID)
	}
}

func (r *UpdateOpportunityStageRequest) GetBody() (io.Reader, error) {
	body := make(map[string]any)
	body[paramStage] = r.StageID

	result := bytes.Buffer{}
	if err := json.NewEncoder(&result).Encode(body); err != nil {
		return nil, err
	}

	return &result, nil
}

// Response for updating an opportunity stage.
type UpdateOpportunityStageResponse struct {
	BaseResponse
}

// Parameters for updating an opportunity archived state.
type UpdateOpportunityArchivedStateRequest struct {
	BaseRequest

	// The opportunity id. This is required.
	OpportunityID string

	// The archive reason id why this candidate is archived for this Opportunity
	ReasonID string

	// Remove pending interviews from Opportunity when it is archived.
	CleanInterviews bool

	// Hire a candidate for the Opportunity against the specific requisition. The active offer on
	// the profile must be associated with an application for a posting linked to this requisition.
	RequisitionID string

	// Perform this update on behalf of a specified user.
	PerformAsID string
}

// Create a new UpdateOpportunityStageRequest with the required fields.
func NewUpdateOpportunityArchivedStateRequest(opportunityID, reasonID string) *UpdateOpportunityArchivedStateRequest {
	return &UpdateOpportunityArchivedStateRequest{
		OpportunityID: opportunityID,
		ReasonID:      reasonID,
	}
}

func (r *UpdateOpportunityArchivedStateRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/archived", url.PathEscape(r.OpportunityID))
}

func (r *UpdateOpportunityArchivedStateRequest) GetMethod() string {
	return http.MethodPut
}

func (r *UpdateOpportunityArchivedStateRequest) AddAPIQueryParams(query *url.Values) {
	r.BaseRequest.AddAPIQueryParams(query)

	if r.PerformAsID != "" {
		query.Add(paramPerformAs, r.PerformAsID)
	}
}

func (r *UpdateOpportunityArchivedStateRequest) GetBody() (io.Reader, error) {
	body := make(map[string]any)
	body[paramReason] = r.ReasonID
	body[paramCleanInterviews] = r.CleanInterviews

	if r.RequisitionID != "" {
		body[paramRequisitionID] = r.RequisitionID
	}

	result := bytes.Buffer{}
	if err := json.NewEncoder(&result).Encode(body); err != nil {
		return nil, err
	}

	return &result, nil
}

// Response for updating an opportunity stage.
type UpdateOpportunityArchivedStateResponse struct {
	BaseResponse
}

// Parameters for adding contact links to a contact by opportunity.
type AddOpportunityLinksRequest struct {
	BaseRequest

	// The opportunity id. This is required.
	OpportunityID string

	// Array of links to add to the contact
	Links []string

	// Perform this update on behalf of a specified user.
	PerformAsID string
}

// Create a new AddOpportunityLinksRequest with the required fields.
func NewAddOpportunityLinksRequest(opportunityID string, links []string) *AddOpportunityLinksRequest {
	linkCopy := make([]string, len(links))
	copy(linkCopy, links)

	return &AddOpportunityLinksRequest{
		OpportunityID: opportunityID,
		Links:         linkCopy,
	}
}

func (r *AddOpportunityLinksRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/addLinks", url.PathEscape(r.OpportunityID))
}

func (r *AddOpportunityLinksRequest) GetMethod() string {
	return http.MethodPost
}

func (r *AddOpportunityLinksRequest) AddAPIQueryParams(query *url.Values) {
	r.BaseRequest.AddAPIQueryParams(query)

	if r.PerformAsID != "" {
		query.Add(paramPerformAs, r.PerformAsID)
	}
}

func (r *AddOpportunityLinksRequest) GetBody() (io.Reader, error) {
	body := make(map[string]any)
	body[paramLinks] = r.Links

	result := bytes.Buffer{}
	if err := json.NewEncoder(&result).Encode(body); err != nil {
		return nil, err
	}

	return &result, nil
}

// Response for adding contact links to a contact by opportunity.
type AddOpportunityLinksResponse struct {
	BaseResponse
}

// Parameters for removeing contact links to a contact by opportunity.
type RemoveOpportunityLinksRequest struct {
	BaseRequest

	// The opportunity id. This is required.
	OpportunityID string

	// Array of links to remove to the contact
	Links []string

	// Perform this update on behalf of a specified user.
	PerformAsID string
}

// Create a new RemoveOpportunityLinksRequest with the required fields.
func NewRemoveOpportunityLinksRequest(opportunityID string, links []string) *RemoveOpportunityLinksRequest {
	linkCopy := make([]string, len(links))
	copy(linkCopy, links)

	return &RemoveOpportunityLinksRequest{
		OpportunityID: opportunityID,
		Links:         linkCopy,
	}
}

func (r *RemoveOpportunityLinksRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/removeLinks", url.PathEscape(r.OpportunityID))
}

func (r *RemoveOpportunityLinksRequest) GetMethod() string {
	return http.MethodPost
}

func (r *RemoveOpportunityLinksRequest) RemoveAPIQueryParams(query *url.Values) {
	r.BaseRequest.AddAPIQueryParams(query)

	if r.PerformAsID != "" {
		query.Add(paramPerformAs, r.PerformAsID)
	}
}

func (r *RemoveOpportunityLinksRequest) GetBody() (io.Reader, error) {
	body := make(map[string]any)
	body[paramLinks] = r.Links

	result := bytes.Buffer{}
	if err := json.NewEncoder(&result).Encode(body); err != nil {
		return nil, err
	}

	return &result, nil
}

// Response for removeing contact links to a contact by opportunity.
type RemoveOpportunityLinksResponse struct {
	BaseResponse
}

// Parameters for adding contact tags to a contact by opportunity.
type AddOpportunityTagsRequest struct {
	BaseRequest

	// The opportunity id. This is required.
	OpportunityID string

	// Array of tags to add to the contact
	Tags []string

	// Perform this update on behalf of a specified user.
	PerformAsID string
}

// Create a new AddOpportunityTagsRequest with the required fields.
func NewAddOpportunityTagsRequest(opportunityID string, tags []string) *AddOpportunityTagsRequest {
	tagCopy := make([]string, len(tags))
	copy(tagCopy, tags)

	return &AddOpportunityTagsRequest{
		OpportunityID: opportunityID,
		Tags:          tagCopy,
	}
}

func (r *AddOpportunityTagsRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/addTags", url.PathEscape(r.OpportunityID))
}

func (r *AddOpportunityTagsRequest) GetMethod() string {
	return http.MethodPost
}

func (r *AddOpportunityTagsRequest) AddAPIQueryParams(query *url.Values) {
	r.BaseRequest.AddAPIQueryParams(query)

	if r.PerformAsID != "" {
		query.Add(paramPerformAs, r.PerformAsID)
	}
}

func (r *AddOpportunityTagsRequest) GetBody() (io.Reader, error) {
	body := make(map[string]any)
	body[paramTags] = r.Tags

	result := bytes.Buffer{}
	if err := json.NewEncoder(&result).Encode(body); err != nil {
		return nil, err
	}

	return &result, nil
}

// Response for adding contact tags to a contact by opportunity.
type AddOpportunityTagsResponse struct {
	BaseResponse
}

// Parameters for removeing contact tags to a contact by opportunity.
type RemoveOpportunityTagsRequest struct {
	BaseRequest

	// The opportunity id. This is required.
	OpportunityID string

	// Array of tags to remove to the contact
	Tags []string

	// Perform this update on behalf of a specified user.
	PerformAsID string
}

// Create a new RemoveOpportunityTagsRequest with the required fields.
func NewRemoveOpportunityTagsRequest(opportunityID string, tags []string) *RemoveOpportunityTagsRequest {
	tagCopy := make([]string, len(tags))
	copy(tagCopy, tags)

	return &RemoveOpportunityTagsRequest{
		OpportunityID: opportunityID,
		Tags:          tagCopy,
	}
}

func (r *RemoveOpportunityTagsRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/removeTags", url.PathEscape(r.OpportunityID))
}

func (r *RemoveOpportunityTagsRequest) GetMethod() string {
	return http.MethodPost
}

func (r *RemoveOpportunityTagsRequest) RemoveAPIQueryParams(query *url.Values) {
	r.BaseRequest.AddAPIQueryParams(query)

	if r.PerformAsID != "" {
		query.Add(paramPerformAs, r.PerformAsID)
	}
}

func (r *RemoveOpportunityTagsRequest) GetBody() (io.Reader, error) {
	body := make(map[string]any)
	body[paramTags] = r.Tags

	result := bytes.Buffer{}
	if err := json.NewEncoder(&result).Encode(body); err != nil {
		return nil, err
	}

	return &result, nil
}

// Response for removeing contact tags to a contact by opportunity.
type RemoveOpportunityTagsResponse struct {
	BaseResponse
}

// Parameters for adding contact sources to a contact by opportunity.
type AddOpportunitySourcesRequest struct {
	BaseRequest

	// The opportunity id. This is required.
	OpportunityID string

	// Array of sources to add to the contact
	Sources []string

	// Perform this update on behalf of a specified user.
	PerformAsID string
}

// Create a new AddOpportunitySourcesRequest with the required fields.
func NewAddOpportunitySourcesRequest(opportunityID string, sources []string) *AddOpportunitySourcesRequest {
	sourceCopy := make([]string, len(sources))
	copy(sourceCopy, sources)

	return &AddOpportunitySourcesRequest{
		OpportunityID: opportunityID,
		Sources:       sourceCopy,
	}
}

func (r *AddOpportunitySourcesRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/addSources", url.PathEscape(r.OpportunityID))
}

func (r *AddOpportunitySourcesRequest) GetMethod() string {
	return http.MethodPost
}

func (r *AddOpportunitySourcesRequest) AddAPIQueryParams(query *url.Values) {
	r.BaseRequest.AddAPIQueryParams(query)

	if r.PerformAsID != "" {
		query.Add(paramPerformAs, r.PerformAsID)
	}
}

func (r *AddOpportunitySourcesRequest) GetBody() (io.Reader, error) {
	body := make(map[string]any)
	body[paramSources] = r.Sources

	result := bytes.Buffer{}
	if err := json.NewEncoder(&result).Encode(body); err != nil {
		return nil, err
	}

	return &result, nil
}

// Response for adding contact sources to a contact by opportunity.
type AddOpportunitySourcesResponse struct {
	BaseResponse
}

// Parameters for removeing contact sources to a contact by opportunity.
type RemoveOpportunitySourcesRequest struct {
	BaseRequest

	// The opportunity id. This is required.
	OpportunityID string

	// Array of sources to remove to the contact
	Sources []string

	// Perform this update on behalf of a specified user.
	PerformAsID string
}

// Create a new RemoveOpportunitySourcesRequest with the required fields.
func NewRemoveOpportunitySourcesRequest(opportunityID string, sources []string) *RemoveOpportunitySourcesRequest {
	sourceCopy := make([]string, len(sources))
	copy(sourceCopy, sources)

	return &RemoveOpportunitySourcesRequest{
		OpportunityID: opportunityID,
		Sources:       sourceCopy,
	}
}

func (r *RemoveOpportunitySourcesRequest) GetPath() string {
	return fmt.Sprintf("opportunities/%s/removeSources", url.PathEscape(r.OpportunityID))
}

func (r *RemoveOpportunitySourcesRequest) GetMethod() string {
	return http.MethodPost
}

func (r *RemoveOpportunitySourcesRequest) RemoveAPIQueryParams(query *url.Values) {
	r.BaseRequest.AddAPIQueryParams(query)

	if r.PerformAsID != "" {
		query.Add(paramPerformAs, r.PerformAsID)
	}
}

func (r *RemoveOpportunitySourcesRequest) GetBody() (io.Reader, error) {
	body := make(map[string]any)
	body[paramSources] = r.Sources

	result := bytes.Buffer{}
	if err := json.NewEncoder(&result).Encode(body); err != nil {
		return nil, err
	}

	return &result, nil
}

// Response for removeing contact sources to a contact by opportunity.
type RemoveOpportunitySourcesResponse struct {
	BaseResponse
}

// Retrieve a single opportunity
func (c *Client) GetOpportunity(ctx context.Context, req *GetOpportunityRequest) (*GetOpportunityResponse, error) {
	var respJSON getOpportunityResponseJSON
	if err := c.exec(ctx, req, &respJSON); err != nil {
		return nil, err
	}

	// Convert the response to the client type
	var opportunity model.Opportunity
	err := respJSON.Opportunity.ToModel(&opportunity)
	if err != nil {
		return nil, err
	}

	resp := GetOpportunityResponse{
		BaseResponse: respJSON.BaseResponse,
		Opportunity:  &opportunity,
	}

	return &resp, nil
}

// List all opportunities
//
// Lists all pipeline Opportunities for Contacts in your Lever account.
func (c *Client) ListOpportunities(ctx context.Context, req *ListOpportunitiesRequest) (*ListOpportunitiesResponse, error) {
	var respJSON listOpportunitiesResponseJSON
	if err := c.exec(ctx, req, &respJSON); err != nil {
		return nil, err
	}

	// Convert the response to the client type
	opportunities := make([]model.Opportunity, len(respJSON.Opportunities))
	for i := range respJSON.Opportunities {
		err := respJSON.Opportunities[i].ToModel(&opportunities[i])
		if err != nil {
			return nil, err
		}
	}

	resp := ListOpportunitiesResponse{
		BaseListResponse: respJSON.BaseListResponse,
		Opportunities:    opportunities,
	}

	return &resp, nil
}

// List deleted opportunities
//
// Lists all deleted Opportunities in your Lever account.
func (c *Client) ListDeletedOpportunities(ctx context.Context, req *ListDeletedOpportunitiesRequest) (*ListDeletedOpportunitiesResponse, error) {
	var respJSON listDeletedOpportunitiesResponseJSON
	if err := c.exec(ctx, req, &respJSON); err != nil {
		return nil, err
	}

	// Convert the response to the client type
	opportunities := make([]model.Opportunity, len(respJSON.Opportunities))
	for i := range respJSON.Opportunities {
		err := respJSON.Opportunities[i].ToModel(&opportunities[i])
		if err != nil {
			return nil, err
		}
	}

	resp := ListDeletedOpportunitiesResponse{
		BaseListResponse: respJSON.BaseListResponse,
		Opportunities:    opportunities,
	}

	return &resp, nil
}

// Create an opportunity
//
// This endpoint enables integrations to create candidates and opportunities in your Lever
// account.
//
// If you want to apply a candidate to a job posting or create a custom job site, you should
// use the Lever Postings API instead of the Lever Data API.
//
// We accept requests of type application/json and multipart/form-data. If you are including a
// resume or other files, you must use the multipart/form-data type.
//
// There are many ways to create a candidate. Here are some examples:
//
// * Provide a JSON representation of a candidate with basic information like candidate name,
// email, and phone number
//
// * Upload just a resume file and specify you'd like the resume parsed. Information parsed
// from the resume will be used to create the candidate—their name, email, and phone number,
// for example.
//
// Note: If you are creating a confidential opportunity, you must provide a posting UID for a
// confidential job posting. Learn more about confidential data in the API.
//
// If candidate information is provided in the POST request and resume parsing is requested,
// the manually provided information will always take precedence over the parsed information
// from the candidate resume.
//
// All fields are optional, but an empty candidate is not particularly interesting. All query
// parameters except the perform_as parameter are optional.
//
// If an email address is provided, we will always attempt to dedupe the candidate. If a match
// is found, we will create a new Opportunity that is linked to the existing matching
// candidate’s contact (i.e. we never create a new contact, or person, if a match has been
// found). The existing candidate’s contact data will take precedence over new manually
// provided information.
//
// If a contact already exists for a candidate, the ID of the existing contact may be provided
// in the POST request to create an opportunity associated with the existing candidate's
// contact (the candidate will be deduped). If additional contact details are included in the
// request (emails, phones, tags, web links), these will be added to the existing candidate's
// contact information.

func (c *Client) CreateOpportunity(ctx context.Context, req *CreateOpportunityRequest) (*CreateOpportunityResponse, error) {
	var respJSON createOpportunityResponseJSON
	if err := c.exec(ctx, req, &respJSON); err != nil {
		return nil, err
	}

	// Convert the response to the client type.
	var opportunity model.Opportunity
	err := respJSON.Opportunity.ToModel(&opportunity)
	if err != nil {
		return nil, err
	}

	resp := CreateOpportunityResponse{
		BaseResponse: respJSON.BaseResponse,
		Deduped:      respJSON.Deduped,
		Opportunity:  &opportunity,
	}

	return &resp, nil
}

// Update opportunity stage
//
// Change an Opportunity's current stage
func (c *Client) UpdateOpportunityStage(ctx context.Context, req *UpdateOpportunityStageRequest) (*UpdateOpportunityStageResponse, error) {
	var resp UpdateOpportunityStageResponse
	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Update opportunity archived state
//
// Update an Opportunity's archived state. If an Opportunity is already archived, its archive
// reason can be changed or if null is specified as the reason, it will be unarchived. If an
// Opportunity is active, it will be archived with the reason provided.
//
// The requisitionId is optional. If the provided reason maps to ‘Hired’ and a requisition is
// provided, the Opportunity will be marked as Hired, the active offer is removed from the
// requisition, and the hired count for the requisition will be incremented.
//
// If a requisition is specified and there are multiple active applications on the profile,
// you will receive an error. If the specific requisition is closed, you will receive an error.
// If there is an offer extended, it must be signed, and the offer must be associated with an
// application for a posting linked to the provided requisition. You can hire a candidate
// against a requisition without an offer.
func (c *Client) UpdateOpportunityArchivedState(ctx context.Context, req *UpdateOpportunityArchivedStateRequest) (*UpdateOpportunityArchivedStateResponse, error) {
	var resp UpdateOpportunityArchivedStateResponse
	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Add contact links to a contact by opportunity
//
// Add links to a Contact by an Opportunity
func (c *Client) AddOpportunityLinks(ctx context.Context, req *AddOpportunityLinksRequest) (*AddOpportunityLinksResponse, error) {
	var resp AddOpportunityLinksResponse
	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Remove contact links from a contact by opportunity
//
// Remove links from a Contact by an Opportunity
func (c *Client) RemoveOpportunityLinks(ctx context.Context, req *RemoveOpportunityLinksRequest) (*RemoveOpportunityLinksResponse, error) {
	var resp RemoveOpportunityLinksResponse
	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Add tags to an opportunity
func (c *Client) AddOpportunityTags(ctx context.Context, req *AddOpportunityTagsRequest) (*AddOpportunityTagsResponse, error) {
	var resp AddOpportunityTagsResponse
	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Remove tags from an opportunity
func (c *Client) RemoveOpportunityTags(ctx context.Context, req *RemoveOpportunityTagsRequest) (*RemoveOpportunityTagsResponse, error) {
	var resp RemoveOpportunityTagsResponse
	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Add sources to an opportunity
func (c *Client) AddOpportunitySources(ctx context.Context, req *AddOpportunitySourcesRequest) (*AddOpportunitySourcesResponse, error) {
	var resp AddOpportunitySourcesResponse
	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Remove sources from an opportunity
func (c *Client) RemoveOpportunitySources(ctx context.Context, req *RemoveOpportunitySourcesRequest) (*RemoveOpportunitySourcesResponse, error) {
	var resp RemoveOpportunitySourcesResponse
	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
