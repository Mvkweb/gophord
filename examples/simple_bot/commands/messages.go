// Package commands provides prefix command handlers for the simple_bot example.
// This file contains message-related commands like react and pin.

package commands

import (
	"context"
	"fmt"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
)

// HandleReact adds a reaction emoji to a message.
// Usage: !react <message_id> <emoji>
// Example: !react 123456789012345678 🔥
func HandleReact(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent, messageIDStr, emoji string) {
	mID, err := types.ParseSnowflake(messageIDStr)
	if err != nil {
		bot.SendMessage(ctx, event.ChannelID, "Invalid Message ID!")
		return
	}

	err = bot.React(ctx, event.ChannelID, mID, emoji)
	if err != nil {
		bot.SendMessage(ctx, event.ChannelID, fmt.Sprintf("Failed to react: %v", err))
		return
	}
}

// HandlePin pins a message to the channel.
// Usage: !pin <message_id>
// Pinned messages can be viewed by clicking the pin icon in Discord clients.
func HandlePin(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent, messageIDStr string) {
	mID, err := types.ParseSnowflake(messageIDStr)
	if err != nil {
		bot.SendMessage(ctx, event.ChannelID, "Invalid Message ID!")
		return
	}

	err = bot.Pin(ctx, event.ChannelID, mID)
	if err != nil {
		bot.SendMessage(ctx, event.ChannelID, fmt.Sprintf("Failed to pin: %v", err))
		return
	}

	bot.SendMessage(ctx, event.ChannelID, "Pinned message! 📌")
}
