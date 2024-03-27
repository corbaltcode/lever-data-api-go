package model

// A source is the way that a candidate entered into your Lever account. The most common sources in
// your Lever account are:
//
// * Posting - Candidate applied to a posting on your careers page.
//
// * Referral - Candidate was referred by an employee at your company.
//
// * Add New - Candidate was added manually into your Lever account in the app.
//
// * Email Applicant - Candidate was added via applicant@hire.lever.co email address.
//
// * Email Lead - Candidate was added via lead@hire.lever.co email address.
//
// * LinkedIn - Candidate was added from LinkedIn using the Lever Chrome Extension.
//
// * GitHub - Candidate was added from GitHub using the Lever Chrome Extension.
//
// * AngelList - Candidate was added from AngelList using the Lever Chrome Extension.
type Source struct {
	// Source text
	Text string `json:"text,omitempty"`

	// Number of candidates tagged with this source
	Count int64 `json:"count,omitempty"`
}
