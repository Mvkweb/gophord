// Example: Modal Forms
// Demonstrates creating popup forms with text inputs and selects

package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
)

// DOC:START
// showFeedbackModal displays a modal with text inputs and selects
func showFeedbackModal(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	components := client.NewComponentBuilder().
		// Short text input (single line)
		AddLabel(client.NewLabel("Your Name", "Enter your name",
			client.NewTextInput("name_input", "", 1,
				client.WithRequired(true),
				client.WithPlaceholder("John Doe"),
			),
		)).
		// Paragraph text input (multi-line)
		AddLabel(client.NewLabel("Feedback", "Share your thoughts",
			client.NewTextInput("feedback_input", "", 2,
				client.WithPlaceholder("Tell us what you think..."),
				client.WithMinLength(10),
				client.WithMaxLength(500),
			),
		)).
		// String select inside modal
		AddLabel(client.NewLabel("Rating", "How would you rate us?",
			client.NewStringSelect("rating_select",
				types.SelectOption{Label: "⭐⭐⭐⭐⭐ Excellent", Value: "5"},
				types.SelectOption{Label: "⭐⭐⭐⭐ Good", Value: "4"},
				types.SelectOption{Label: "⭐⭐⭐ Average", Value: "3"},
				types.SelectOption{Label: "⭐⭐ Poor", Value: "2"},
				types.SelectOption{Label: "⭐ Terrible", Value: "1"},
			),
		)).
		Build()

	err := bot.ShowModal(ctx, interaction, "Feedback Form", "feedback_modal", components)
	if err != nil {
		log.Printf("Failed to show modal: %v", err)
	}
}

// handleModalSubmit processes the submitted modal data
func handleModalSubmit(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	var results []string

	for _, comp := range interaction.Data.Components {
		if label, ok := comp.(*types.Label); ok {
			switch inner := label.Component.(type) {
			case *types.TextInput:
				results = append(results, fmt.Sprintf("**%s**: %s", inner.CustomID, inner.Value))
			case *types.StringSelect:
				results = append(results, fmt.Sprintf("**%s**: %s", inner.CustomID, strings.Join(inner.Values, ", ")))
			}
		}
	}

	response := "## Form Submitted!\n" + strings.Join(results, "\n")
	bot.RespondWithEphemeral(ctx, interaction, response)
}

// DOC:END

func main() {
	bot := client.New("YOUR_TOKEN", client.WithIntents(
		types.IntentGuilds|types.IntentGuildMessages|types.IntentMessageContent,
	))

	// Register slash command to open modal
	bot.OnReady(func(event *gateway.ReadyEvent) {
		bot.RegisterGlobalCommand(context.Background(), types.CreateApplicationCommandParams{
			Name:        "feedback",
			Description: "Open the feedback form",
		})
	})

	bot.OnInteractionCreate(func(event *gateway.InteractionCreateEvent) {
		ctx := context.Background()

		switch event.Type {
		case types.InteractionTypeApplicationCommand:
			if event.Data.Name == "feedback" {
				showFeedbackModal(ctx, bot, &event.Interaction)
			}
		case types.InteractionTypeModalSubmit:
			handleModalSubmit(ctx, bot, &event.Interaction)
		}
	})

	bot.Connect(context.Background())
}
