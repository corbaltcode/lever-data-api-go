package model

// Resumes are data about a candidate's work history and/or education. Some resumes may have files
// associated with them. Resumes can be added to an Opportunity in a number of ways. For example,
// when (1) the candidate applies to a job posting and uploads a resume (usually this will have a
// file associated with it), (2) when a recruiter uses the Chrome extension tool via Linkedin,
// Github, Twitter, etc., (3) the candidate is manually added by a recruiter or (4) a resume file
// is added directly to an Opportunity from within Lever.
type Resume struct {
	// Resume UID
	ID string `json:"id,omitempty"`

	// Datetime when resume was created in Lever. For candidates who applied to a job posting on
	// your website, the date and time when the candidate's resume was created in Lever is the
	// moment when the candidate clicked the "Attach Resume/CV" button on their application.
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// An object containing resume file metadata and download url.
	File *ResumeFile `json:"file,omitempty"`

	// The candidate's parsed resume, usually extracted from an attached PDF/Word document or
	// online profile (LinkedIn, GitHub, AngelList, etc.).
	ParsedData ResumeParsedData `json:"parsedData,omitempty"`
}

// An object containing resume file metadata and download url.
type ResumeFile struct {
	// Resume file download URL
	DownloadURL string `json:"downloadUrl,omitempty"`

	// Resume file extension
	Ext string `json:"ext,omitempty"`

	// Resume file name
	Name string `json:"name,omitempty"`

	// Datetime when file was uploaded in Lever
	UploadedAt *int64 `json:"uploadedAt,omitempty"`

	// The status of processing the file. Can be one of the following values: processing,
	// processed, unsupported, error, or null.
	Status string `json:"status,omitempty"`

	// The size of the file in bytes.
	Size *int64 `json:"size,omitempty"`
}

// The candidate's parsed resume, usually extracted from an attached PDF/Word document or online
// profile (LinkedIn, GitHub, AngelList, etc.).
type ResumeParsedData struct {
	// An array of objects containing various information about positions the candidate has held.
	// This includes company name, location, job title, job summary, and start and end date. This
	// information can be parsed from the candidate's resume, LinkedIn profile, or entered
	// manually.
	Positions []map[string]any `json:"positions,omitempty"`

	// An array of objects containing various information about schools the candidate has attended.
	// This includes the school name, degree, field of study, school summary, and start and end
	// date. This information can be parsed from the candidate's resume, LinkedIn profile, or
	// entered manually.
	Schools []map[string]any `json:"school,omitempty"`
}
