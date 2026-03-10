// Package events provides gateway event handlers for the simple_bot example.
// This file handles the READY event when the bot connects to Discord.

package events

import (
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
)

// OnReady sets up the handler for when the bot successfully connects to Discord.
// This is triggered once during initial connection and on successful reconnections.
func OnReady(bot *client.Client) {
	bot.OnReady(func(event *gateway.ReadyEvent) {
		log.Printf("Bot is ready! Logged in as %s#%s (ID: %s)",
			event.User.Username,
			event.User.Discriminator,
			event.User.ID)
		log.Printf("Connected to %d guilds", len(event.Guilds))
	})
}
