// Package types provides Discord API type definitions.
package types

import (
	"time"
)

// Timestamp represents a Discord timestamp in ISO8601 format.
type Timestamp string

// Time parses the timestamp into a time.Time value.
func (t Timestamp) Time() (time.Time, error) {
	return time.Parse(time.RFC3339, string(t))
}

// MustTime parses the timestamp, panicking on error.
func (t Timestamp) MustTime() time.Time {
	tm, err := t.Time()
	if err != nil {
		panic(err)
	}
	return tm
}

// String returns the timestamp as a string.
func (t Timestamp) String() string {
	return string(t)
}

// IsZero returns true if the timestamp is empty.
func (t Timestamp) IsZero() bool {
	return t == ""
}

// NewTimestamp creates a Timestamp from a time.Time value.
func NewTimestamp(t time.Time) Timestamp {
	return Timestamp(t.Format(time.RFC3339))
}

// Now returns the current time as a Timestamp.
func Now() Timestamp {
	return NewTimestamp(time.Now())
}
