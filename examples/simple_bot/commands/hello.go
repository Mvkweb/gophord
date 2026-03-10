// Package commands provides prefix command handlers for the simple_bot example.

package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
)

// HandleHello sends a greeting message with interactive buttons.
// Demonstrates Components V2 features including text displays, separators, and buttons.
func HandleHello(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	components := client.NewComponentBuilder().
		AddTextDisplay(fmt.Sprintf("# Hello, %s! 👋", event.Author.Username)).
		AddTextDisplay("Welcome to the **gophord** Discord library demo!").
		AddSeparator(true, types.SeparatorSpacingSmall).
		AddTextDisplay("Click a button below to try out the interactive features:").
		AddActionRow(
			client.NewPrimaryButton("greet_back", "Greet Back"),
			client.NewSuccessButton("show_info", "Show Info"),
			client.NewDangerButton("say_goodbye", "Goodbye"),
		).
		Build()

	_, err := bot.SendMessageWithComponents(ctx, event.ChannelID, components)
	if err != nil {
		log.Printf("Failed to send hello message: %v", err)
	}
}
