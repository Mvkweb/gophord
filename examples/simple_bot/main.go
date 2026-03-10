// Package main demonstrates a simple Discord bot using gophord.
//
// This example shows how to:
//   - Connect to Discord using the gateway
//   - Handle basic events like messages and interactions
//   - Send messages with Components V2
//   - Organize code in a scalable, maintainable way
//
// The code is organized into packages:
//   - commands/ - Prefix command handlers (!ping, !kick, etc.)
//   - interactions/ - Slash command and button handlers
//   - modals/ - Modal form handlers
//   - events/ - Gateway event registration
//   - utils/ - Shared constants and helpers
//
// To run this example:
//
//	go run main.go -token YOUR_BOT_TOKEN
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"

	"github.com/Mvkweb/gophord/examples/simple_bot/events"
	"github.com/Mvkweb/gophord/examples/simple_bot/interactions"
	"github.com/Mvkweb/gophord/examples/simple_bot/modals"
	"github.com/Mvkweb/gophord/examples/simple_bot/utils"
	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
	godotenv "github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()

	// Parse command line flags
	token := flag.String("token", "", "Discord bot token")
	flag.Parse()

	// Fall back to environment variable if not provided via flag
	if *token == "" {
		*token = os.Getenv("DISCORD_BOT_TOKEN")
		if *token == "" {
			log.Fatal("Bot token is required. Use -token flag or DISCORD_BOT_TOKEN environment variable.")
		}
	}

	// Create the Discord client with required intents
	// Mobile status makes the bot appear as online from a mobile device
	bot := client.New(*token,
		client.WithIntents(utils.RequiredIntents),
		client.WithMobileStatus(true),
	)

	// Register event handlers
	// These are set up once at startup and called whenever events occur
	events.OnReady(bot)
	events.OnMessageCreate(bot)

	// Register interaction handler for slash commands, buttons, and modals
	bot.OnInteractionCreate(func(event *gateway.InteractionCreateEvent) {
		ctx := context.Background()
		interaction := event.Interaction

		switch interaction.Type {
		case types.InteractionTypeApplicationCommand:
			interactions.HandleSlashCommand(ctx, bot, &interaction)
		case types.InteractionTypeMessageComponent:
			interactions.HandleButtonClick(ctx, bot, &interaction)
		case types.InteractionTypeModalSubmit:
			modals.HandleModalSubmit(ctx, bot, &interaction)
		}
	})

	// Register slash commands in a background goroutine
	// This runs concurrently with the main bot
	go registerSlashCommands(bot)

	// Connect to Discord
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown gracefully on Ctrl+C or interrupt
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		if err := bot.Connect(ctx); err != nil {
			log.Printf("Connection error: %v", err)
		}
	}()

	log.Println("Bot is running. Press ENTER or Ctrl+C to stop.")

	// Wait for interrupt signal or ENTER key
	go func() {
		fmt.Scanln()
		cancel()
	}()

	<-stop
	log.Println("\nShutting down...")

	// Clean up terminal before exit
	restoreTerminal()

	// Stop gateway and close connection
	cancel()
	bot.Close()
	log.Println("Bot stopped.")
}

// registerSlashCommands registers all slash commands with Discord.
// If a guild ID is set in DISCORD_GUILD_ID, commands are registered to that guild.
// Otherwise, commands are registered globally (may take up to an hour to propagate).
func registerSlashCommands(bot *client.Client) {
	ctx := context.Background()
	commands := interactions.GetSlashCommands()

	guildIDStr := os.Getenv("DISCORD_GUILD_ID")
	if guildIDStr != "" {
		gID, err := types.ParseSnowflake(guildIDStr)
		if err != nil {
			log.Printf("Invalid DISCORD_GUILD_ID: %v", err)
			return
		}
		for _, cmd := range commands {
			_, err = bot.RegisterGuildCommand(ctx, gID, cmd)
			if err != nil {
				log.Printf("Failed to register guild slash command %s: %v", cmd.Name, err)
			}
		}
		log.Printf("Successfully registered slash commands to guild %s", gID)
	} else {
		for _, cmd := range commands {
			_, err := bot.RegisterGlobalCommand(ctx, cmd)
			if err != nil {
				log.Printf("Failed to register global slash command %s: %v", cmd.Name, err)
			}
		}
		log.Println("Successfully registered slash commands globally")
	}
}

// restoreTerminal performs any necessary terminal cleanup on shutdown.
// Currently handles Windows-specific issues with console input.
func restoreTerminal() {
	if runtime.GOOS == "windows" {
		// On Windows, syscall constants vary by version
		// The fmt.Scanln() approach provides a simple way to exit
		// If terminal is stuck, close the terminal window and open a new one
	}
}
