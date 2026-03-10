// Package events provides gateway event handlers for the simple_bot example.
// This file handles the MESSAGE_CREATE event for processing prefix commands.

package events

import (
	"context"
	"strings"

	"github.com/Mvkweb/gophord/examples/simple_bot/commands"
	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
)

// OnMessageCreate sets up the handler for when a message is created in a channel.
// This is where we process traditional prefix commands (!ping, !help, etc.)
func OnMessageCreate(bot *client.Client) {
	bot.OnMessageCreate(func(event *gateway.MessageCreateEvent) {
		// Ignore bot messages to prevent infinite loops
		if event.Author.Bot {
			return
		}

		ctx := context.Background()
		content := event.Content
		args := strings.Split(content, " ")
		command := strings.ToLower(args[0])

		// Route to appropriate command handler
		// Each case calls the corresponding handler function from the commands package
		switch {
		case command == "!ping":
			commands.HandlePing(ctx, bot, event)
		case command == "!hello":
			commands.HandleHello(ctx, bot, event)
		case command == "!components":
			commands.HandleComponents(ctx, bot, event)
		case command == "!container":
			commands.HandleContainer(ctx, bot, event)
		case command == "!gallery":
			commands.HandleGallery(ctx, bot, event)
		case command == "!section":
			commands.HandleSection(ctx, bot, event)
		case command == "!help":
			commands.HandleHelp(ctx, bot, event)
		case command == "!silent" && len(args) > 1:
			commands.HandleSilent(ctx, bot, event, strings.Join(args[1:], " "))
		case command == "!kick" && len(args) > 1:
			commands.HandleKick(ctx, bot, event, args[1])
		case command == "!webhook" && len(args) > 1:
			commands.HandleWebhookDemo(ctx, bot, event, args[1])
		case command == "!react" && len(args) > 2:
			commands.HandleReact(ctx, bot, event, args[1], args[2])
		case command == "!pin" && len(args) > 1:
			commands.HandlePin(ctx, bot, event, args[1])
		case command == "!purge" && len(args) > 1:
			commands.HandlePurge(ctx, bot, event, args[1])
		case command == "!fileupload":
			commands.HandleFileUploadDemo(ctx, bot, event)
		case command == "!premium":
			commands.HandlePremium(ctx, bot, event)
		}
	})
}
