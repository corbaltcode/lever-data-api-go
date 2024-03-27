package lever

import "strings"

// From https://github.com/golang/go/blob/0c5612092deb0a50c5a3d67babc1249049595558/src/mime/multipart/writer.go#L132-L136
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}
