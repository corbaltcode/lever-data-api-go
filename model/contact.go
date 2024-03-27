package model

// A contact represents the person associated with one or more opportunities and the various
// methods to contact them.
type Contact struct {
	// Contact UID
	ID string `json:"id,omitempty"`

	// Name of the contact
	Name string `json:"name,omitempty"`

	// Contact headline, typically a list of previous companies where the contact has worked or
	// schools that the contact has attended
	Headline string `json:"headline,omitempty"`

	// The current location of the contact
	Location *ContactLocation `json:"location,omitempty"`

	// Emails that the contact can be reached at
	Emails []string `json:"emails,omitempty"`

	// Indicates whether a contact has been anonymized. Anonymized contacts have their personal
	// information removed. Non-personal metadata may remain for accurate reporting purposes.
	IsAnonymized bool `json:"isAnonymized,omitempty"`

	// Phone numbers
	Phones []Phone `json:"phones,omitempty"`
}

// The current location of the contact.
//
// TODO: Determine whether this struct should be shared with other models.
type ContactLocation struct {
	// A single line description of the location
	Name string `json:"name,omitempty"`
}
