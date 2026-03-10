// Package commands provides prefix command handlers for the simple_bot example.
// This file contains premium button demo commands.

package commands

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
)

// HandlePremium demonstrates premium button functionality.
// Premium buttons prompt users to purchase SKUs. They don't send interactions when clicked.
// Requires DISCORD_SKU_ID environment variable to be set.
func HandlePremium(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	skuIDStr := os.Getenv("DISCORD_SKU_ID")
	if skuIDStr == "" {
		_, err := bot.SendMessage(ctx, event.ChannelID, "⚠️ Premium button requires DISCORD_SKU_ID to be set in environment.")
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
		return
	}

	skuID, err := strconv.ParseUint(skuIDStr, 10, 64)
	if err != nil {
		_, err := bot.SendMessage(ctx, event.ChannelID, "⚠️ Invalid SKU ID format.")
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
		return
	}

	components := client.NewComponentBuilder().
		AddTextDisplay("# Premium Button Demo").
		AddTextDisplay("Premium buttons prompt users to purchase items or subscriptions.").
		AddTextDisplay("Click the button below to open the Discord store!").
		AddSeparator(true, types.SeparatorSpacingSmall).
		AddTextDisplay("-# Note: Premium buttons don't trigger interactions when clicked").
		AddActionRow(
			client.NewPremiumButton(types.Snowflake(skuID)),
		).
		Build()

	_, err = bot.SendMessageWithComponents(ctx, event.ChannelID, components)
	if err != nil {
		log.Printf("Failed to send premium message: %v", err)
	}
}
