// Example: Silent Messages
// Send messages without triggering notifications

package main

import (
	"context"
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
)

// DOC:START
// sendSilentMessage sends a message that doesn't notify users
func sendSilentMessage(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
	// Silent messages appear in the channel but don't trigger:
	// - Push notifications
	// - Desktop notifications
	// - Mobile alerts
	// - Unread indicators

	_, err := bot.SendMessageSilent(ctx, channelID,
		"🔕 **Silent Update**\n\n"+
			"This message was sent silently.\n"+
			"Users won't receive notifications for this message.",
	)
	if err != nil {
		log.Printf("Failed to send silent message: %v", err)
	}
}

// sendSilentWithFlags shows how to use MessageFlagSuppressNotifications
func sendSilentWithFlags(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
	// Using the REST client directly with flags
	_, err := bot.REST.CreateMessage(ctx, channelID, &rest.CreateMessageParams{
		Content: "🔕 Another silent message using flags",
		Flags:   types.MessageFlagSuppressNotifications,
	})
	if err != nil {
		log.Printf("Failed to send: %v", err)
	}
}

// DOC:END

func main() {
	bot := client.New("YOUR_TOKEN", client.WithIntents(
		types.IntentGuilds|types.IntentGuildMessages|types.IntentMessageContent,
	))

	bot.OnMessageCreate(func(event *gateway.MessageCreateEvent) {
		if event.Author.Bot {
			return
		}

		if event.Content == "!silent" {
			sendSilentMessage(context.Background(), bot, event.ChannelID)
		}
	})

	bot.Connect(context.Background())
}
