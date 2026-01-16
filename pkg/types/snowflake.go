// Package types provides Discord API type definitions.
package types

import (
	"strconv"
	"time"
)

// DiscordEpoch is the Discord epoch timestamp (January 1, 2015).
const DiscordEpoch int64 = 1420070400000

// Snowflake represents a Discord snowflake ID.
// Snowflakes are 64-bit unsigned integers used as unique identifiers.
type Snowflake uint64

// String returns the string representation of the snowflake.
func (s Snowflake) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

// Int64 returns the snowflake as an int64.
func (s Snowflake) Int64() int64 {
	return int64(s)
}

// UInt64 returns the snowflake as a uint64.
func (s Snowflake) UInt64() uint64 {
	return uint64(s)
}

// Timestamp extracts the timestamp from the snowflake.
func (s Snowflake) Timestamp() time.Time {
	ms := (int64(s) >> 22) + DiscordEpoch
	return time.UnixMilli(ms)
}

// WorkerID extracts the internal worker ID from the snowflake.
func (s Snowflake) WorkerID() uint8 {
	return uint8((s >> 17) & 0x1F)
}

// ProcessID extracts the internal process ID from the snowflake.
func (s Snowflake) ProcessID() uint8 {
	return uint8((s >> 12) & 0x1F)
}

// Increment extracts the increment from the snowflake.
func (s Snowflake) Increment() uint16 {
	return uint16(s & 0xFFF)
}

// IsZero returns true if the snowflake is zero (invalid/unset).
func (s Snowflake) IsZero() bool {
	return s == 0
}

// ParseSnowflake parses a string into a Snowflake.
func ParseSnowflake(s string) (Snowflake, error) {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return Snowflake(v), nil
}

// MustParseSnowflake parses a string into a Snowflake, panicking on error.
func MustParseSnowflake(s string) Snowflake {
	sf, err := ParseSnowflake(s)
	if err != nil {
		panic(err)
	}
	return sf
}

// NewSnowflake creates a snowflake from the given components.
func NewSnowflake(timestamp time.Time, workerID, processID uint8, increment uint16) Snowflake {
	ms := timestamp.UnixMilli() - DiscordEpoch
	return Snowflake((ms << 22) | (int64(workerID&0x1F) << 17) | (int64(processID&0x1F) << 12) | int64(increment&0xFFF))
}

// MarshalJSON implements json.Marshaler for Snowflake.
// Discord sends snowflakes as strings in JSON.
func (s Snowflake) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler for Snowflake.
// Handles both string and number representations.
func (s *Snowflake) UnmarshalJSON(data []byte) error {
	str := string(data)

	// Handle quoted strings
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}

	// Handle null
	if str == "null" {
		*s = 0
		return nil
	}

	v, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}
	*s = Snowflake(v)
	return nil
}
