// Package utils provides shared helper functions for the simple_bot example.

package utils

import "github.com/Mvkweb/gophord/pkg/types"

// ToPointer returns a pointer to the given value.
// Useful for setting optional pointer fields in structs.
func ToPointer[T any](value T) *T {
	return &value
}

// GetUserName extracts the username from an interaction.
// Checks both User and Member fields since DM vs guild contexts differ.
func GetUserName(interaction *types.Interaction) string {
	if interaction.Member != nil && interaction.Member.User != nil {
		return interaction.Member.User.Username
	}
	if interaction.User != nil {
		return interaction.User.Username
	}
	return "User"
}

// GetChannelID extracts the channel ID from a message event.
func GetChannelID(event *types.Message) types.Snowflake {
	return event.ChannelID
}
