// Package commands provides prefix command handlers for the simple_bot example.

package commands

import (
	"context"
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
)

// HandlePing responds with "Pong!" when users type !ping.
// This is the simplest possible command and serves as a connectivity test.
func HandlePing(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	_, err := bot.SendMessage(ctx, event.ChannelID, "Pong! 🏓")
	if err != nil {
		log.Printf("Failed to send ping response: %v", err)
	}
}
