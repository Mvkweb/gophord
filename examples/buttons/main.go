// Example: Button Components
// Demonstrates all button styles available in Discord Components V2

package main

import (
	"context"
	"log"

	"github.com/gophord/gophord/pkg/client"
	"github.com/gophord/gophord/pkg/gateway"
	"github.com/gophord/gophord/pkg/types"
)

// DOC:START
// sendButtonsDemo demonstrates all button styles
func sendButtonsDemo(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# Button Styles Demo").
		AddTextDisplay("Here are all the available button styles:").
		AddSeparator(true, types.SeparatorSpacingSmall).
		AddActionRow(
			// Primary button - Blue, for main actions
			client.NewPrimaryButton("btn_primary", "Primary"),
			// Secondary button - Gray, for secondary actions
			client.NewSecondaryButton("btn_secondary", "Secondary"),
			// Success button - Green, for positive actions
			client.NewSuccessButton("btn_success", "Success"),
			// Danger button - Red, for destructive actions
			client.NewDangerButton("btn_danger", "Danger"),
		).
		AddActionRow(
			// Link button - Opens URL in browser
			client.NewLinkButton("https://github.com/gophord/gophord", "GitHub"),
			// Button with emoji
			&types.Button{
				Style:    types.ButtonStylePrimary,
				CustomID: "btn_emoji",
				Label:    "With Emoji",
				Emoji:    &types.PartialEmoji{Name: "🚀"},
			},
			// Disabled button
			&types.Button{
				Style:    types.ButtonStyleSecondary,
				CustomID: "btn_disabled",
				Label:    "Disabled",
				Disabled: true,
			},
		).
		Build()

	_, err := bot.SendMessageWithComponents(ctx, channelID, components)
	if err != nil {
		log.Printf("Failed to send buttons: %v", err)
	}
}

// DOC:END

// Handle button interactions
func handleButtonClick(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	customID := interaction.Data.CustomID

	var response string
	switch customID {
	case "btn_primary":
		response = "You clicked the Primary button!"
	case "btn_secondary":
		response = "You clicked the Secondary button!"
	case "btn_success":
		response = "Success! ✅"
	case "btn_danger":
		response = "Danger zone! ⚠️"
	default:
		response = "Button clicked: " + customID
	}

	bot.RespondWithMessage(ctx, interaction, response)
}

func main() {
	bot := client.New("YOUR_TOKEN", client.WithIntents(
		types.IntentGuilds|types.IntentGuildMessages|types.IntentMessageContent,
	))

	bot.OnMessageCreate(func(event *gateway.MessageCreateEvent) {
		if event.Content == "!buttons" {
			sendButtonsDemo(context.Background(), bot, event.ChannelID)
		}
	})

	bot.OnInteractionCreate(func(event *gateway.InteractionCreateEvent) {
		if event.Type == types.InteractionTypeMessageComponent {
			handleButtonClick(context.Background(), bot, &event.Interaction)
		}
	})

	bot.Connect(context.Background())
}
