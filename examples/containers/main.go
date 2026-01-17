// Example: Containers
// Group components with accent colors and visual organization

package main

import (
	"context"
	"log"

	"github.com/gophord/gophord/pkg/client"
	"github.com/gophord/gophord/pkg/gateway"
	"github.com/gophord/gophord/pkg/types"
)

// DOC:START
// sendContainerDemo shows containers with accent colors
func sendContainerDemo(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
	// Containers group components with a colored accent bar
	accentColor := 0x5865F2 // Discord Blurple

	container := &types.Container{
		AccentColor: &accentColor,
		Components: types.ComponentList{
			&types.TextDisplay{Content: "# 📦 Container Example"},
			&types.TextDisplay{Content: "This content is wrapped in a container with a **blurple** accent bar."},
			&types.Separator{Divider: boolPtr(true), Spacing: types.SeparatorSpacingSmall},
			&types.TextDisplay{Content: "Containers help organize content and draw attention with colored bars."},
			&types.ActionRow{
				Components: types.ComponentList{
					client.NewPrimaryButton("container_confirm", "Confirm"),
					client.NewSecondaryButton("container_cancel", "Cancel"),
				},
			},
		},
	}

	_, err := bot.SendMessageWithComponents(ctx, channelID, types.ComponentList{container})
	if err != nil {
		log.Printf("Failed to send container: %v", err)
	}
}

// sendMultipleContainers shows different accent colors
func sendMultipleContainers(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
	success := 0x57F287 // Green
	warning := 0xFEE75C // Yellow
	danger := 0xED4245  // Red

	components := types.ComponentList{
		&types.Container{
			AccentColor: &success,
			Components: types.ComponentList{
				&types.TextDisplay{Content: "✅ **Success Container**\nOperation completed successfully!"},
			},
		},
		&types.Container{
			AccentColor: &warning,
			Components: types.ComponentList{
				&types.TextDisplay{Content: "⚠️ **Warning Container**\nPlease review before continuing."},
			},
		},
		&types.Container{
			AccentColor: &danger,
			Components: types.ComponentList{
				&types.TextDisplay{Content: "❌ **Error Container**\nSomething went wrong."},
			},
		},
	}

	_, err := bot.SendMessageWithComponents(ctx, channelID, components)
	if err != nil {
		log.Printf("Failed to send containers: %v", err)
	}
}

func boolPtr(b bool) *bool { return &b }

// DOC:END

func main() {
	bot := client.New("YOUR_TOKEN", client.WithIntents(
		types.IntentGuilds|types.IntentGuildMessages|types.IntentMessageContent,
	))

	bot.OnMessageCreate(func(event *gateway.MessageCreateEvent) {
		if event.Content == "!container" {
			sendContainerDemo(context.Background(), bot, event.ChannelID)
		}
		if event.Content == "!containers" {
			sendMultipleContainers(context.Background(), bot, event.ChannelID)
		}
	})

	bot.Connect(context.Background())
}
