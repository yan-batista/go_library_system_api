package utils

import (
	"regexp"
	"strings"
)

func CreateSlug(name string) string {
	// Remove special characters
	slug := regexp.MustCompile(`[^a-zA-Z0-9\s]+`).ReplaceAllString(name, "")

	// Trim white space from both ends of the string
	slug = strings.TrimSpace(slug)

	// Replace any sequence of white space characters with a single hyphen
	slug = regexp.MustCompile(`\s+`).ReplaceAllString(slug, "-")

	// Convert to lowercase
	slug = strings.ToLower(slug)

	return slug
}