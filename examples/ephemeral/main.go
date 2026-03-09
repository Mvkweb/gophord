// Example: Ephemeral Messages
// Messages only visible to the target user

package main

import (
	"context"
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
)

// DOC:START
// handleEphemeralCommand sends a private message only the user can see
func handleEphemeralCommand(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	// Ephemeral messages are only visible to the user who triggered the interaction
	// Great for: sensitive data, confirmations, help messages
	err := bot.RespondWithEphemeral(ctx, interaction,
		"👀 **Secret Message**\n\n"+
			"This message is only visible to you!\n"+
			"Other users in the channel cannot see this.",
	)
	if err != nil {
		log.Printf("Failed to send ephemeral: %v", err)
	}
}

// sendEphemeralWithComponents sends ephemeral message with buttons
func sendEphemeralWithComponents(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# Private Actions").
		AddTextDisplay("Only you can see these options:").
		AddActionRow(
			client.NewPrimaryButton("private_action", "Take Action"),
			client.NewDangerButton("private_cancel", "Cancel"),
		).
		Build()

	err := bot.RespondWithEphemeralComponents(ctx, interaction, components)
	if err != nil {
		log.Printf("Failed to send ephemeral components: %v", err)
	}
}

// DOC:END

func main() {
	bot := client.New("YOUR_TOKEN", client.WithIntents(
		types.IntentGuilds|types.IntentGuildMessages,
	))

	// Register slash command
	bot.OnReady(func(event *gateway.ReadyEvent) {
		bot.RegisterGlobalCommand(context.Background(), types.CreateApplicationCommandParams{
			Name:        "secret",
			Description: "Get a secret message only you can see",
		})
	})

	bot.OnInteractionCreate(func(event *gateway.InteractionCreateEvent) {
		if event.Type == types.InteractionTypeApplicationCommand {
			if event.Data.Name == "secret" {
				handleEphemeralCommand(context.Background(), bot, &event.Interaction)
			}
		}
	})

	bot.Connect(context.Background())
}
