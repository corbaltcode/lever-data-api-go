package model

// Users in Lever include anyone who has been invited to join in on recruiting efforts. There are
// five different access roles in Lever. From most to least access, these roles are Super admin,
// Admin, Team member, Limited team member, and Interviewer.
type User struct {
	// User UID
	ID string `json:"id,omitempty"`

	// User's preferred name
	Name string `json:"name,omitempty"`

	// Username, extracted from user's email address
	Username string `json:"username,omitempty"`

	// User's email address
	Email string `json:"email,omitempty"`

	// Datetime when user was created
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// Datetime when user was deactivated, null for an active user
	DeactivatedAt *int64 `json:"deactivatedAt,omitempty"`

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
