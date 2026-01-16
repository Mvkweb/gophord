// Package main demonstrates a simple Discord bot using gophord.
//
// This example shows how to:
//   - Connect to Discord using the gateway (lxzan/gws)
//   - Handle basic events like messages and interactions
//   - Send messages with Components V2
//   - Manage Guilds (Kick members)
//   - Manage Webhooks (Create and Execute)
//   - Identify as a Mobile Client
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
	"strings"

	"github.com/gophord/gophord/pkg/client"
	"github.com/gophord/gophord/pkg/gateway"
	"github.com/gophord/gophord/pkg/types"
	godotenv "github.com/joho/godotenv"
)

func main() {
	// Load .env file from current directory
	godotenv.Load()
	token := flag.String("token", "", "Discord bot token")
	flag.Parse()

	if *token == "" {
		// Try environment variable
		*token = os.Getenv("DISCORD_BOT_TOKEN")
		if *token == "" {
			log.Fatal("Bot token is required. Use -token flag or DISCORD_BOT_TOKEN environment variable.")
		}
	}

	// Create Discord client with necessary intents
	bot := client.New(*token,
		client.WithIntents(
			types.IntentGuilds|
			types.IntentGuildMessages|
			types.IntentGuildMessageReactions|
			types.IntentDirectMessages|
			types.IntentMessageContent|
			types.IntentGuildMembers, // Required for management features
		),
		client.WithMobileStatus(true), // Enable mobile status (Discord Android)
	)

	// Register event handlers
	bot.OnReady(func(event *gateway.ReadyEvent) {
		log.Printf("Bot is ready! Logged in as %s#%s (ID: %s)",
			event.User.Username,
			event.User.Discriminator,
			event.User.ID)
		log.Printf("Connected to %d guilds", len(event.Guilds))
	})

	bot.OnMessageCreate(func(event *gateway.MessageCreateEvent) {
		// Ignore bot messages
		if event.Author.Bot {
			return
		}

		ctx := context.Background()
		content := event.Content
		args := strings.Split(content, " ")
		command := strings.ToLower(args[0])

		switch {
		case command == "!ping":
			handlePing(ctx, bot, event)
		case command == "!hello":
			handleHello(ctx, bot, event)
		case command == "!components":
			handleComponents(ctx, bot, event)
		case command == "!container":
			handleContainer(ctx, bot, event)
		case command == "!help":
			handleHelp(ctx, bot, event)
		case command == "!kick" && len(args) > 1:
			handleKick(ctx, bot, event, args[1])
		case command == "!webhook" && len(args) > 1:
			handleWebhookDemo(ctx, bot, event, args[1])
		}
	})

	bot.OnInteractionCreate(func(event *gateway.InteractionCreateEvent) {
		ctx := context.Background()
		handleInteraction(ctx, bot, &event.Interaction)
	})

	// Register a slash command
	go func() {
		ctx := context.Background()
		cmdParams := types.CreateApplicationCommandParams{
			Name:        "hello",
			Description: "Get a greeting from gophord",
		}

		guildIDStr := os.Getenv("DISCORD_GUILD_ID")
		if guildIDStr != "" {
			// Dev Mode: Register to specific guild for instant updates
			gID, err := types.ParseSnowflake(guildIDStr)
			if err != nil {
				log.Printf("Invalid DISCORD_GUILD_ID: %v", err)
				return
			}
			_, err = bot.RegisterGuildCommand(ctx, gID, cmdParams)
			if err != nil {
				log.Printf("Failed to register guild slash command: %v", err)
			} else {
				log.Printf("Successfully registered /hello slash command to guild %s (Instant Update)", gID)
			}
		} else {
			// Production: Register globally (can take up to 1 hour to update)
			_, err := bot.RegisterGlobalCommand(ctx, cmdParams)
			if err != nil {
				log.Printf("Failed to register global slash command: %v", err)
			} else {
				log.Println("Successfully registered /hello slash command globally (May take up to 1 hour to update)")
			}
		}
	}()

	// Connect to Discord
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown gracefully
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

	<-	stop
	log.Println("\nShutting down...")

	// Restore terminal before exit
	restoreTerminal()

	// Cancel context to stop gateway
	cancel()

	// Close bot connection
	bot.Close()
	log.Println("Bot stopped.")
}

func handlePing(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	_, err := bot.SendMessage(ctx, event.ChannelID, "Pong! 🏓")
	if err != nil {
		log.Printf("Failed to send ping response: %v", err)
	}
}

func handleHello(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	// Send a message with Components V2
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

func handleKick(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent, userIDStr string) {
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

func handleWebhookDemo(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent, name string) {
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

func handleComponents(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	// Demonstrate various Components V2 features
	components := client.NewComponentBuilder().
		AddTextDisplay("# Components V2 Demo").
		AddTextDisplay("This message showcases the new Discord Components V2 features.").
		AddSeparator(true, types.SeparatorSpacingLarge).
		AddTextDisplay("## Text Display").
		AddTextDisplay("Text displays support **markdown** formatting, including:\n- Bold and *italic* text\n- `code blocks`\n- And more!").
		AddSeparator(true, types.SeparatorSpacingSmall).
		AddTextDisplay("## Interactive Buttons").
		AddActionRow(
			client.NewPrimaryButton("btn_primary", "Primary"),
			client.NewSecondaryButton("btn_secondary", "Secondary"),
			client.NewSuccessButton("btn_success", "Success"),
			client.NewDangerButton("btn_danger", "Danger"),
			client.NewLinkButton("https://discord.com", "Link"),
		).
		Build()

	_, err := bot.SendMessageWithComponents(ctx, event.ChannelID, components)
	if err != nil {
		log.Printf("Failed to send components message: %v", err)
	}
}

func handleContainer(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	// Demonstrate containers with accent colors
	accentColor := 0x5865F2 // Discord blurple

	container := &types.Container{
		AccentColor: &accentColor,
		Components: types.ComponentList{
			&types.TextDisplay{Content: "# Contained Message"},
			&types.TextDisplay{Content: "This message is wrapped in a **container** with an accent color bar!"},
			&types.Separator{Divider: boolPtr(true), Spacing: types.SeparatorSpacingSmall},
			&types.TextDisplay{Content: "Containers help organize and visually group related content."},
			&types.ActionRow{
				Components: types.ComponentList{
					client.NewPrimaryButton("container_action", "Take Action"),
					client.NewSecondaryButton("container_dismiss", "Dismiss"),
				},
			},
		},
	}

	components := types.ComponentList{container}

	_, err := bot.SendMessageWithComponents(ctx, event.ChannelID, components)
	if err != nil {
		log.Printf("Failed to send container message: %v", err)
	}
}

func handleHelp(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# Gophord Bot Help").
		AddTextDisplay("Available commands:").
		AddTextDisplay("- `!ping` - Check if the bot is responsive\n- `!hello` - Get a greeting with interactive buttons\n- `!kick <user_id>` - Kick a user from the server\n- `!webhook <name>` - Create and test a webhook\n- `!components` - See Components V2 features demo\n- `!container` - See a container with accent color\n- `!help` - Show this help message").
		AddSeparator(true, types.SeparatorSpacingSmall).
		AddTextDisplay("-# Powered by gophord - A high-performance Go Discord library").
		Build()

	_, err := bot.SendMessageWithComponents(ctx, event.ChannelID, components)
	if err != nil {
		log.Printf("Failed to send help message: %v", err)
	}
}

func handleInteraction(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	// Handle slash commands
	if interaction.Type == types.InteractionTypeApplicationCommand {
		if interaction.Data == nil {
			log.Println("Interaction data is nil for application command")
			return
		}
		if interaction.Data.Name == "hello" {
			userName := "User"
			if interaction.User != nil {
				userName = interaction.User.Username
			} else if interaction.Member != nil && interaction.Member.User != nil {
				userName = interaction.Member.User.Username
			}

			// Reuse the handleHello logic but for interaction response
			components := client.NewComponentBuilder().
				AddTextDisplay(fmt.Sprintf("# Hello, %s! 👋", userName)).
				AddTextDisplay("Welcome to the **gophord** Discord library demo!").
				AddSeparator(true, types.SeparatorSpacingSmall).
				AddTextDisplay("Click a button below to try out the interactive features:").
				AddActionRow(
					client.NewPrimaryButton("greet_back", "Greet Back"),
					client.NewSuccessButton("show_info", "Show Info"),
					client.NewDangerButton("say_goodbye", "Goodbye"),
				).
				Build()

			err := bot.RespondWithComponents(ctx, interaction, components)
			if err != nil {
				log.Printf("Failed to respond to /hello interaction: %v", err)
			}
		}
		return
	}

	// Only handle component interactions
	if interaction.Type != types.InteractionTypeMessageComponent {
		return
	}

	if interaction.Data == nil {
		log.Println("Interaction data is nil for message component")
		return
	}

	customID := interaction.Data.CustomID

	var response string
	switch customID {
	case "greet_back":
		userName := "friend"
		if interaction.Member != nil && interaction.Member.User != nil {
			userName = interaction.Member.User.Username
		} else if interaction.User != nil {
			userName = interaction.User.Username
		}
		response = fmt.Sprintf("Hello to you too, %s! 😊", userName)
	case "show_info":
		response = "**Gophord Info**\n\n" +
			"• High-performance Go Discord library\n" +
			"• Uses bytedance/sonic for fast JSON\n" +
			"• Full Components V2 support\n" +
			"• Idiomatic Go design"
	case "say_goodbye":
		response = "Goodbye! See you next time! 👋"
	case "btn_primary", "btn_secondary", "btn_success", "btn_danger":
		response = fmt.Sprintf("You clicked the **%s** button!", strings.TrimPrefix(customID, "btn_"))
	case "container_action":
		response = "Action taken! ✅"
	case "container_dismiss":
		response = "Dismissed! 👍"
	default:
		response = fmt.Sprintf("Button clicked: `%s`", customID)
	}

	err := bot.RespondWithMessage(ctx, interaction, response)
	if err != nil {
		log.Printf("Failed to respond to interaction: %v", err)
	}
}

func boolPtr(b bool) *bool {
	return &b
}

func restoreTerminal() {
	if runtime.GOOS == "windows" {
		// On Windows, syscall constants vary by version
		// The fmt.Scanln() approach provides a simple way to exit
		// If terminal is stuck, close the terminal window and open a new one
	}
}