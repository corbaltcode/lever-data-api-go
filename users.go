package lever

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/corbaltcode/lever-data-api-go/model"
)

// Lever users client interface
type UsersClientInterface interface {
	ClientInterface

	// Retrieve a single user.
	//
	// This method returns the full user record for a single user.
	GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error)

	// List users
	//
	// Lists the users in your Lever account. Only active users are returned by default.
	ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error)

	// Create a user
	//
	// This endpoint enables integrations to create users in your Lever account.
	//
	// Users will be created with the Interviewer access role by default. Users may be created with
	// Interviewer, Limited Team Member, Team Member, Admin, or Super Admin access.
	//
	// Note: This will not send an invite to the user, so direct auth users will need to go through
	// the direct auth password flow.
	CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)

	// Update a user
	//
	// When you update a user, Lever expects you to send the entire resource. Every field will be
	// overwritten by the body of the request. If you don't include a field, it will be deleted or
	// reset to its default. Be sure to include all fields you still want to be populated. name,
	// email, and accessRole are required fields. Note that resetting accessRole to interviewer
	// will result in a user losing all of their followed profiles.
	UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error)

	// Deactivate a user
	//
	// Deactivated users remain in the system for historical record keeping, but can no longer log
	// in and use Lever.
	DeactivateUser(ctx context.Context, req *DeactivateUserRequest) (*DeactivateUserResponse, error)

	// Reactivate a user
	//
	// Reactivate a user that has been previously deactivated
	ReactivateUser(ctx context.Context, req *ReactivateUserRequest) (*ReactivateUserResponse, error)
}

// Parameters for retrieving a single user.
type GetUserRequest struct {
	BaseRequest

	// The user id. This is required.
	UserId string
}

// Create a new GetUserRequest with the required fields.
func NewGetUserRequest(userId string) *GetUserRequest {
	return &GetUserRequest{
		UserId: userId,
	}
}

func (r *GetUserRequest) GetPath() string {
	return fmt.Sprintf("users/%s", url.PathEscape(r.UserId))
}

// Response for retrieving a single user.
type GetUserResponse struct {
	BaseResponse

	// The user record.
	User *model.User `json:"data"`
}

// Parameters for listing users.
type ListUsersRequest struct {
	BaseListRequest

	// If set, filter results to users that match an email address. Provided email must exactly
	// match the canonicalized version of the user's email.
	Email []string

	// If set, filter by access role. One of: 'super admin', 'admin', 'team member',
	// 'limited team member', 'interviewer', or the ID for a custom role listed on your roles page.
	AccessRole []string

	// If set, include deactivated users along with activated users.
	IncludeDeactivated bool

	// If set, filter results that match a user's external directory id.
	ExternalDirectoryID []string
}

// Create a new ListUsersRequest with the required fields.
func NewListUsersRequest() *ListUsersRequest {
	return &ListUsersRequest{}
}

func (r *ListUsersRequest) GetPath() string {
	return "users"
}

func (r *ListUsersRequest) AddAPIQueryParams(query *url.Values) {
	r.BaseListRequest.AddAPIQueryParams(query)

	for _, email := range r.Email {
		query.Add(paramEmail, email)
	}

	for _, accessRole := range r.AccessRole {
		query.Add(paramAccessRole, accessRole)
	}

	if r.IncludeDeactivated {
		query.Add(paramIncludeDeactivated, "true")
	}

	for _, externalDirectoryId := range r.ExternalDirectoryID {
		query.Add(paramExternalDirectoryId, externalDirectoryId)
	}
}

// Response for listing users.
type ListUsersResponse struct {
	BaseListResponse

	// The user records.
	Users []model.User `json:"data"`
}

// Parameters for creating a user.
type CreateUserRequest struct {
	BaseRequest

	// User's preferred name
	Name string `json:"name"`

	// User's email address
	Email string `json:"email"`

	// User's access role. One of: 'super admin', 'admin', 'team member', 'limited team member',
	// 'interviewer'
	AccessRole string `json:"accessRole,omitempty"`

	// Unique id for user in external HR directory
	ExternalDirectoryID string `json:"externalDirectoryId,omitempty"`

	// User's job title
	JobTitle string `json:"jobTitle,omitempty"`

	// User's manager ID
	ManagerID string `json:"manager,omitempty"`
}

// Create a new CreateUserRequest with the required fields.
func NewCreateUserRequest(name, email string) *CreateUserRequest {
	return &CreateUserRequest{
		Name:  name,
		Email: email,
	}
}

func (r *CreateUserRequest) GetPath() string {
	return "users"
}

func (r *CreateUserRequest) GetHTTPMethod() string {
	return http.MethodPost
}

func (r *CreateUserRequest) GetBody() (io.Reader, error) {
	result := bytes.Buffer{}
	encoder := json.NewEncoder(&result)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "")
	if err := encoder.Encode(r); err != nil {
		return nil, err
	}

	return &result, nil
}

// Response for creating a user.
type CreateUserResponse struct {
	BaseResponse

	// The user record.
	User *model.User `json:"data"`
}

// Parameters for updating a user.
type UpdateUserRequest struct {
	BaseRequest

	// User UID
	ID string `json:"id,omitempty"`

	// User's preferred name
	Name string `json:"name,omitempty"`

	// User's email address
	Email string `json:"email,omitempty"`

	// User's access role. One of: 'super admin', 'admin', 'team member', 'limited team member',
	// 'interviewer'
	AccessRole string `json:"accessRole,omitempty"`

	// URL for user's gravatar, if enabled
	Photo string `json:"photo,omitempty"`

	// Unique id for user in external HR directory
	ExternalDirectoryID string `json:"externalDirectoryId,omitempty"`

	// An array of contact IDs which helps identify all contacts associated with a User. This can
	// be used to control User access to any Opportunities linked to a User.
	LinkedContactIds []string `json:"linkedContactIds,omitempty"`

	// User's job title
	JobTitle string `json:"jobTitle,omitempty"`

	// User's manager ID
	ManagerID string `json:"manager,omitempty"`
}

// Create a new UpdateUserRequest with the required fields.
func NewUpdateUserRequest(id, name, email, accessRole string) *UpdateUserRequest {
	return &UpdateUserRequest{
		ID:         id,
		Name:       name,
		Email:      email,
		AccessRole: accessRole,
	}
}

// Create a new UpdateUserRequest based on an existing User struct.
func NewUpdateUserRequestFromUser(user *model.User) *UpdateUserRequest {
	return &UpdateUserRequest{
		ID:                  user.ID,
		Name:                user.Name,
		Email:               user.Email,
		AccessRole:          user.AccessRole,
		Photo:               user.Photo,
		ExternalDirectoryID: user.ExternalDirectoryID,
		LinkedContactIds:    user.LinkedContactIds,
		JobTitle:            user.JobTitle,
		ManagerID:           user.ManagerID,
	}
}

func (r *UpdateUserRequest) GetPath() string {
	return fmt.Sprintf("users/%s", url.PathEscape(r.ID))
}

func (r *UpdateUserRequest) GetHTTPMethod() string {
	return http.MethodPut
}

func (r *UpdateUserRequest) GetBody() (io.Reader, error) {
	result := bytes.Buffer{}
	encoder := json.NewEncoder(&result)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "")
	if err := encoder.Encode(r); err != nil {
		return nil, err
	}

	return &result, nil
}

// Response for updating a user.
type UpdateUserResponse struct {
	BaseResponse

	// The user record.
	User *model.User `json:"data"`
}

// Parameters for deactivating a user.
type DeactivateUserRequest struct {
	BaseRequest

	// The user id. This is required.
	UserId string
}

// Create a new DeactivateUserRequest with the required fields.
func NewDeactivateUserRequest(userId string) *DeactivateUserRequest {
	return &DeactivateUserRequest{
		UserId: userId,
	}
}

func (r *DeactivateUserRequest) GetPath() string {
	return fmt.Sprintf("users/%s/deactivate", url.PathEscape(r.UserId))
}

func (r *DeactivateUserRequest) GetHTTPMethod() string {
	return http.MethodPost
}

// Response for deactivating a user.
type DeactivateUserResponse struct {
	BaseResponse

	// The user record.
	User *model.User `json:"data"`
}

// Parameters for reactivating a user.
type ReactivateUserRequest struct {
	BaseRequest

	// The user id. This is required.
	UserId string
}

// Create a new ReactivateUserRequest with the required fields.
func NewReactivateUserRequest(userId string) *ReactivateUserRequest {
	return &ReactivateUserRequest{
		UserId: userId,
	}
}

func (r *ReactivateUserRequest) GetPath() string {
	return fmt.Sprintf("users/%s/reactivate", url.PathEscape(r.UserId))
}

func (r *ReactivateUserRequest) GetHTTPMethod() string {
	return http.MethodPost
}

// Response for reactivating a user.
type ReactivateUserResponse struct {
	BaseResponse

	// The user record.
	User *model.User `json:"data"`
}

// Retrieve a single user.
//
// This method returns the full user record for a single user.
func (c *Client) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	var resp GetUserResponse

	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// List users
//
// Lists the users in your Lever account. Only active users are returned by default.
func (c *Client) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	var resp ListUsersResponse

	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Create a user
//
// This endpoint enables integrations to create users in your Lever account.
//
// Users will be created with the Interviewer access role by default. Users may be created with
// Interviewer, Limited Team Member, Team Member, Admin, or Super Admin access.
//
// Note: This will not send an invite to the user, so direct auth users will need to go through
// the direct auth password flow.
func (c *Client) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	var resp CreateUserResponse

	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Update a user
//
// When you update a user, Lever expects you to send the entire resource. Every field will be
// overwritten by the body of the request. If you don't include a field, it will be deleted or
// reset to its default. Be sure to include all fields you still want to be populated. name,
// email, and accessRole are required fields. Note that resetting accessRole to interviewer
// will result in a user losing all of their followed profiles.
func (c *Client) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error) {
	var resp UpdateUserResponse

	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// Deactivate a user
//
// Deactivated users remain in the system for historical record keeping, but can no longer log
// in and use Lever.
func (c *Client) DeactivateUser(ctx context.Context, req *DeactivateUserRequest) (*DeactivateUserResponse, error) {
	var resp DeactivateUserResponse

	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Reactivate a user
//
// Reactivate a user that has been previously deactivated
func (c *Client) ReactivateUser(ctx context.Context, req *ReactivateUserRequest) (*ReactivateUserResponse, error) {
	var resp ReactivateUserResponse

	if err := c.exec(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
