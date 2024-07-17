// Package logs provides functionality for handling log entries emitted by different applications.
package logs

import (
	"strings"
	"unicode/utf8"
)

// applications maps specific Unicode runes to their corresponding application names.
var applications = map[rune]string{
	'\u2757':     "recommendation",
	'\U0001f50d': "search",
	'\u2600':     "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		if app, exists := applications[r]; exists {
			return app
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.Map(func(r rune) rune {
		if r == oldRune {
			return newRune
		}
		return r
	}, log)
}

// WithinLimit determines weather or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
