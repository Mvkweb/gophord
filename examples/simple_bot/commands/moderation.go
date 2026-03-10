// Package commands provides prefix command handlers for the simple_bot example.
// This file contains moderation-related commands like kick and purge.

package commands

import (
	"context"
	"fmt"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/rest"
	"github.com/Mvkweb/gophord/pkg/types"
)

// HandleKick removes a member from the guild using their user ID.
// Usage: !kick <user_id>
// Requires the bot to have Kick Members permission and be in the same guild.
func HandleKick(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent, userIDStr string) {
	if event.GuildID == nil {
		bot.SendMessage(ctx, event.ChannelID, "This command can only be used in a server!")
		return
	}

	uID, err := types.ParseSnowflake(userIDStr)
	if err != nil {
		bot.SendMessage(ctx, event.ChannelID, "Invalid User ID!")
		return
	}

	err = bot.KickMember(ctx, *event.GuildID, uID)
	if err != nil {
		bot.SendMessage(ctx, event.ChannelID, fmt.Sprintf("Failed to kick member: %v", err))
		return
	}

	bot.SendMessage(ctx, event.ChannelID, fmt.Sprintf("Successfully kicked user <@%s>! 🥾", userIDStr))
}

// HandlePurge deletes multiple messages from a channel.
// Usage: !purge <count>
// Deletes between 1-100 messages at a time. The command message itself is also deleted.
func HandlePurge(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent, countStr string) {
	count := 0
	fmt.Sscanf(countStr, "%d", &count)
	if count <= 0 || count > 100 {
		bot.SendMessage(ctx, event.ChannelID, "Please specify a count between 1 and 100.")
		return
	}

	// Fetch messages (count + 1 to include the command message itself so we can skip it)
	messages, err := bot.REST.GetMessages(ctx, event.ChannelID, &rest.GetMessagesParams{Limit: count + 1})
	if err != nil {
		bot.SendMessage(ctx, event.ChannelID, fmt.Sprintf("Failed to fetch messages: %v", err))
		return
	}

	if len(messages) == 0 {
		return
	}

	// Collect IDs, skipping the command message (the first one)
	ids := make([]types.Snowflake, 0, len(messages))
	for i, m := range messages {
		// The first message in the list is the !purge command we just sent
		if i == 0 {
			// We delete the command message separately or just let it be purged
			// Actually, let's include it in the purge to keep chat clean
			ids = append(ids, m.ID)
			continue
		}
		ids = append(ids, m.ID)
	}

	// If we only have 1 ID (e.g. only the command itself), use DeleteMessage
	if len(ids) == 1 {
		err = bot.REST.DeleteMessage(ctx, event.ChannelID, ids[0])
	} else {
		// Use BulkDelete for 2-100 messages
		err = bot.REST.BulkDeleteMessages(ctx, event.ChannelID, ids)
	}

	if err != nil {
		bot.SendMessage(ctx, event.ChannelID, fmt.Sprintf("Failed to purge: %v", err))
		return
	}

	// Send a confirmation that auto-deletes or is silent
	bot.SendMessageSilent(ctx, event.ChannelID, fmt.Sprintf("🧹 Purged %d messages (including command).", len(ids)))
}
