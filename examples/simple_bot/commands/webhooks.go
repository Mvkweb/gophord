// Package commands provides prefix command handlers for the simple_bot example.
// This file contains webhook management commands.

package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
)

// HandleWebhookDemo creates a webhook with the given name and immediately executes it.
// Usage: !webhook <name>
// Demonstrates both creating and executing webhooks programmatically.
func HandleWebhookDemo(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent, name string) {
	// Create a webhook
	webhook, err := bot.CreateWebhook(ctx, event.ChannelID, name)
	if err != nil {
		bot.SendMessage(ctx, event.ChannelID, fmt.Sprintf("Failed to create webhook: %v", err))
		return
	}

	// Execute it immediately
	err = bot.ExecuteWebhook(ctx, webhook.ID, webhook.Token, fmt.Sprintf("Hello! I am the **%s** webhook, powered by gophord! 🦫", webhook.Name))
	if err != nil {
		log.Printf("Failed to execute webhook: %v", err)
	}

	bot.SendMessage(ctx, event.ChannelID, fmt.Sprintf("Created and executed webhook: **%s**", name))
}
