// Example: Select Menus
// Demonstrates string, user, role, and channel select components

package main

import (
	"context"
	"log"
	"strings"

	"github.com/gophord/gophord/pkg/client"
	"github.com/gophord/gophord/pkg/gateway"
	"github.com/gophord/gophord/pkg/types"
)

// DOC:START
// sendSelectsDemo shows different types of select menus
func sendSelectsDemo(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# Select Menus Demo").
		AddTextDisplay("Select menus let users choose from a list:").
		AddSeparator(true, types.SeparatorSpacingSmall).
		// String Select - Custom options
		AddTextDisplay("## String Select").
		AddActionRow(&types.StringSelect{
			CustomID:    "language_select",
			Placeholder: "Choose your favorite language",
			Options: []types.SelectOption{
				{Label: "Go", Value: "go", Description: "Simple and efficient", Emoji: &types.PartialEmoji{Name: "🐹"}},
				{Label: "Rust", Value: "rust", Description: "Memory safety", Emoji: &types.PartialEmoji{Name: "🦀"}},
				{Label: "Python", Value: "python", Description: "Easy to learn", Emoji: &types.PartialEmoji{Name: "🐍"}},
				{Label: "TypeScript", Value: "ts", Description: "Typed JavaScript", Emoji: &types.PartialEmoji{Name: "📘"}},
			},
		}).
		AddSeparator(true, types.SeparatorSpacingSmall).
		// User Select - Auto-populated with server members
		AddTextDisplay("## User Select").
		AddActionRow(&types.UserSelect{
			CustomID:    "user_select",
			Placeholder: "Select a user",
		}).
		// Channel Select - Filter by channel type
		AddTextDisplay("## Channel Select").
		AddActionRow(&types.ChannelSelect{
			CustomID:     "channel_select",
			Placeholder:  "Select a text channel",
			ChannelTypes: []types.ChannelType{types.ChannelTypeGuildText},
		}).
		// Role Select
		AddTextDisplay("## Role Select").
		AddActionRow(&types.RoleSelect{
			CustomID:    "role_select",
			Placeholder: "Select a role",
		}).
		Build()

	_, err := bot.SendMessageWithComponents(ctx, channelID, components)
	if err != nil {
		log.Printf("Failed to send selects: %v", err)
	}
}

// DOC:END

// Handle select interactions
func handleSelectInteraction(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	customID := interaction.Data.CustomID
	values := interaction.Data.Values

	var response string
	switch customID {
	case "language_select":
		response = "You selected: " + strings.Join(values, ", ")
	case "user_select":
		response = "User selected! ID: " + strings.Join(values, ", ")
	case "channel_select":
		response = "Channel selected! ID: " + strings.Join(values, ", ")
	case "role_select":
		response = "Role selected! ID: " + strings.Join(values, ", ")
	}

	bot.RespondWithMessage(ctx, interaction, response)
}

func main() {
	bot := client.New("YOUR_TOKEN", client.WithIntents(
		types.IntentGuilds|types.IntentGuildMessages|types.IntentMessageContent,
	))

	bot.OnMessageCreate(func(event *gateway.MessageCreateEvent) {
		if event.Content == "!selects" {
			sendSelectsDemo(context.Background(), bot, event.ChannelID)
		}
	})

	bot.OnInteractionCreate(func(event *gateway.InteractionCreateEvent) {
		if event.Type == types.InteractionTypeMessageComponent {
			handleSelectInteraction(context.Background(), bot, &event.Interaction)
		}
	})

	bot.Connect(context.Background())
}
