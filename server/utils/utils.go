package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func CreateSlug(name, author string) string {
	// Remove special characters
	slug_format := fmt.Sprintf("%s - %s", name, author)
	slug := regexp.MustCompile(`[^a-zA-Z0-9\s]+`).ReplaceAllString(slug_format, "")

	// Trim white space from both ends of the string
	slug = strings.TrimSpace(slug)

	// Replace any sequence of white space characters with a single hyphen
	slug = regexp.MustCompile(`\s+`).ReplaceAllString(slug, "-")

	// Convert to lowercase
	slug = strings.ToLower(slug)

	return slug
}

func IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&’*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)
    isValid := emailRegex.MatchString(email)
	return isValid
}