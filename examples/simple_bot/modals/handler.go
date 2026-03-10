// Package modals provides modal form handlers for the simple_bot example.
// This file contains the main modal router that dispatches to specific modal handlers.

package modals

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/types"
)

// HandleModalSubmit is the main entry point for all modal submission interactions.
// Routes to appropriate handler based on the modal's custom_id.
func HandleModalSubmit(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	switch interaction.Data.CustomID {
	case "demo_modal":
		HandleFeedbackModal(ctx, bot, interaction)
	case "file_upload_modal":
		HandleFileUploadSubmit(ctx, bot, interaction)
	default:
		log.Printf("Unknown modal submitted: %s", interaction.Data.CustomID)
	}
}

// HandleFeedbackModal displays the feedback form modal.
// This is a comprehensive modal with text inputs and selects.
func HandleFeedbackModal(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	components := client.NewComponentBuilder().
		// Text Input (Short - Style 1)
		AddLabel(client.NewLabel("Your Name", "Please enter your name",
			client.NewTextInput("name_input", "", 1, client.WithRequired(true)),
		)).
		// Text Input (Paragraph - Style 2)
		AddLabel(client.NewLabel("Feedback", "Tell us what you think",
			client.NewTextInput("feedback_input", "", 2,
				client.WithPlaceholder("Share your thoughts..."),
				client.WithMinLength(10),
				client.WithMaxLength(500)),
		)).
		// String Select with options
		AddLabel(client.NewLabel("Favorite Bug", "Choose your favorite bug",
			client.NewStringSelect("bug_select",
				types.SelectOption{Label: "Ant", Value: "ant", Description: "(best option)", Emoji: &types.PartialEmoji{Name: "🐜"}},
				types.SelectOption{Label: "Butterfly", Value: "butterfly", Emoji: &types.PartialEmoji{Name: "🦋"}},
				types.SelectOption{Label: "Caterpillar", Value: "caterpillar", Emoji: &types.PartialEmoji{Name: "🐛"}},
			),
		)).
		// User Select
		AddLabel(client.NewLabel("Pick a User", "Select a user to mention",
			client.NewUserSelect("user_select"),
		)).
		// Channel Select
		AddLabel(client.NewLabel("Pick a Channel", "Select a channel",
			client.NewChannelSelect("channel_select"),
		)).
		Build()

	err := bot.ShowModal(ctx, interaction, "Components V2 Modal Demo", "demo_modal", components)
	if err != nil {
		log.Printf("Failed to show modal: %v", err)
	}
}

// HandleFeedbackSubmit processes the feedback modal submission.
// Extracts all input values and displays them to the user.
func HandleFeedbackSubmit(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	var results []string

	for _, comp := range interaction.Data.Components {
		if container, ok := comp.(*types.Label); ok {
			switch inner := container.Component.(type) {
			case *types.TextInput:
				results = append(results, fmt.Sprintf("**%s**: %s", inner.CustomID, inner.Value))
			case *types.StringSelect:
				results = append(results, fmt.Sprintf("**%s**: %v", inner.CustomID, inner.Values))
			case *types.UserSelect:
				results = append(results, fmt.Sprintf("**%s**: %v", inner.CustomID, inner.Values))
			case *types.RoleSelect:
				results = append(results, fmt.Sprintf("**%s**: %v", inner.CustomID, inner.Values))
			case *types.ChannelSelect:
				results = append(results, fmt.Sprintf("**%s**: %v", inner.CustomID, inner.Values))
			case *types.MentionableSelect:
				results = append(results, fmt.Sprintf("**%s**: %v", inner.CustomID, inner.Values))
			case *types.FileUpload:
				fileCount := len(inner.Values)
				if fileCount > 0 {
					results = append(results, fmt.Sprintf("**%s**: %d file(s) uploaded - IDs: `%v`", inner.CustomID, fileCount, inner.Values))
				} else {
					results = append(results, fmt.Sprintf("**%s**: No files uploaded", inner.CustomID))
				}
			}
		}
	}

	response := "## Modal Submission Results\n" + strings.Join(results, "\n")
	err := bot.RespondWithEphemeral(ctx, interaction, response)
	if err != nil {
		log.Printf("Failed to respond to modal submit: %v", err)
	}
}
