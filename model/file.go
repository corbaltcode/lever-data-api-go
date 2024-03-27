package model

import "io"

// Reader provides the contents and the name of a file.
type Reader struct {
	// The file name
	Name string

	// The file contents
	Contents io.ReadCloser

	// The MIME type for the file.
	MIMEType string
}
