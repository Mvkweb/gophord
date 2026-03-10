// Package commands provides prefix command handlers for the simple_bot example.

package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/Mvkweb/gophord/examples/simple_bot/utils"
	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
)

// HandleSilent sends an ephemeral (silent) message that only the sender can see.
// Useful for sensitive information or bot status messages.
func HandleSilent(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent, content string) {
	_, err := bot.SendMessageSilent(ctx, event.ChannelID, fmt.Sprintf("🤫 %s", content))
	if err != nil {
		log.Printf("Failed to send silent message: %v", err)
	}
}

// HandleHelp displays a help message showing all available commands.
// Uses Components V2 text displays for formatted output.
func HandleHelp(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# Gophord Bot Help").
		AddTextDisplay("Available commands:").
		AddTextDisplay(
			"- `!ping` - Check if the bot is responsive\n"+
				"- `!hello` - Get a greeting with interactive buttons\n"+
				"- `!kick <user_id>` - Kick a user from the server\n"+
				"- `!webhook <name>` - Create and test a webhook\n"+
				"- `!components` - See Components V2 features demo\n"+
				"- `!container` - See a container with accent color\n"+
				"- `!gallery` - See a media gallery demo\n"+
				"- `!section` - See a section component demo\n"+
				"- `!fileupload` - See file upload modal demo\n"+
				"- `!premium` - See premium button demo\n"+
				"- `!help` - Show this help message",
		).
		AddSeparator(true, types.SeparatorSpacingSmall).
		AddTextDisplay(utils.HelpFooter).
		Build()

	_, err := bot.SendMessageWithComponents(ctx, event.ChannelID, components)
	if err != nil {
		log.Printf("Failed to send help message: %v", err)
	}
}
