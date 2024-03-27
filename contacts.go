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

// Lever data contact client interface
type ContactInterface interface {
	ClientInterface

	// Retrieve a single contact.
	GetContact(ctx context.Context, req *GetContactRequest) (*GetContactResponse, error)

	// Update a contact.
	UpdateContact(ctx context.Context, req *UpdateContactRequest) (*UpdateContactResponse, error)
}

// Parameters for retrieving a single contact.
type GetContactRequest struct {
	BaseRequest

	// The contact UID.
	ContactUID string
}

// Create a new GetContactRequest with the required fields.
func NewGetContactRequest(contactUID string) *GetContactRequest {
	return &GetContactRequest{
		ContactUID: contactUID,
	}
}

func (r *GetContactRequest) GetPath() string {
	return fmt.Sprintf("contacts/%s", url.PathEscape(r.ContactUID))
}

// Response for retrieving a single contact.
type GetContactResponse struct {
	BaseResponse

	// The contact record.
	Contact *model.Contact `json:"data"`
}

// Parameters for updating a contact.
type UpdateContactRequest struct {
	BaseRequest

	// The contact data. This is required.
	Contact *model.Contact
}

// Create a new UpdateContactRequest with the required fields.
func NewUpdateContactRequest(contact *model.Contact) *UpdateContactRequest {
	return &UpdateContactRequest{
		Contact: contact,
	}
}

func (r *UpdateContactRequest) GetPath() string {
	return fmt.Sprintf("contacts/%s", url.PathEscape(r.Contact.ID))
}

func (r *UpdateContactRequest) GetMethod() string {
	return http.MethodPut
}

func (r *UpdateContactRequest) GetBody() (io.Reader, error) {
	contactClone := *r.Contact
	contactClone.ID = ""

	result := bytes.Buffer{}
	encoder := json.NewEncoder(&result)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "")
	if err := encoder.Encode(&contactClone); err != nil {
		return nil, err
	}

	json.Marshal(&contactClone)
	return &result, nil
}

// Response for updating a contact.
type UpdateContactResponse struct {
	BaseResponse

	// The contact record.
	Contact *model.Contact `json:"data"`
}
