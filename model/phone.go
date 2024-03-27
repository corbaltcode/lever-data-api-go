package model

// Phone number.
type Phone struct {
	// One of 'other', 'home', 'mobile', 'work', 'skype'
	Type string `json:"type,omitempty"`

	// The phone number.
	Value string `json:"value,omitempty"`
}
