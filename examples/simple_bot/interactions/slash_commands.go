// Package interactions provides slash command and component interaction handlers.
// This file defines all slash commands that can be registered with Discord.

package interactions

import (
	"github.com/Mvkweb/gophord/pkg/types"
)

// GetSlashCommands returns all slash command definitions for the bot.
// These are registered with Discord on bot startup.
func GetSlashCommands() []types.CreateApplicationCommandParams {
	return []types.CreateApplicationCommandParams{
		{
			Name:        "hello",
			Description: "Get a greeting from gophord",
		},
		{
			Name:        "ephemeral",
			Description: "Send a message only you can see",
		},
		{
			Name:        "modal",
			Description: "Open a test modal",
		},
		{
			Name:        "gallery",
			Description: "See a media gallery demo",
		},
		{
			Name:        "section",
			Description: "See a section component demo",
		},
		{
			Name:        "fileupload",
			Description: "Open a modal with file upload component",
		},
		{
			Name:        "premium",
			Description: "See a premium button demo (requires SKU ID)",
		},
	}
}
