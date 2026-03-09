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
	"github.com/gophord/gophord/pkg/rest"
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
		case command == "!gallery":
			handleGallery(ctx, bot, event)
		case command == "!section":
			handleSection(ctx, bot, event)
		case command == "!help":
			handleHelp(ctx, bot, event)
		case command == "!silent" && len(args) > 1:
			handleSilent(ctx, bot, event, strings.Join(args[1:], " "))
		case command == "!kick" && len(args) > 1:
			handleKick(ctx, bot, event, args[1])
		case command == "!webhook" && len(args) > 1:
			handleWebhookDemo(ctx, bot, event, args[1])
		case command == "!react" && len(args) > 2:
			handleReact(ctx, bot, event, args[1], args[2])
		case command == "!pin" && len(args) > 1:
			handlePin(ctx, bot, event, args[1])
		case command == "!purge" && len(args) > 1:
			handlePurge(ctx, bot, event, args[1])
		}
	})

	bot.OnInteractionCreate(func(event *gateway.InteractionCreateEvent) {
		ctx := context.Background()
		handleInteraction(ctx, bot, &event.Interaction)
	})

	// Register slash commands
	go func() {
		ctx := context.Background()
		commands := []types.CreateApplicationCommandParams{
			{
				Name:        "hello",
				Description: "Get a greeting from gophord",
			},
			{
				Name:        "ephemeral",
				Description: "Send a message only you can see",
			},
			{
				Name:        "modal",
				Description: "Open a test modal",
			},
			{
				Name:        "gallery",
				Description: "See a media gallery and file attachment demo",
			},
			{
				Name:        "section",
				Description: "See a section component demo",
			},
		}

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

	<-stop
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

func handleSilent(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent, content string) {
	_, err := bot.SendMessageSilent(ctx, event.ChannelID, fmt.Sprintf("🤫 %s", content))
	if err != nil {
		log.Printf("Failed to send silent message: %v", err)
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

func handleReact(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent, messageIDStr, emoji string) {
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

func handlePin(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent, messageIDStr string) {
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

func handlePurge(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent, countStr string) {
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

func handleModalDemo(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	// Comprehensive modal demo with all supported component types
	components := client.NewComponentBuilder().
		// Text Input (Short - Style 1)
		AddLabel(client.NewLabel("Your Name", "Please enter your name",
			client.NewTextInput("name_input", "", 1, client.WithRequired(true)),
		)).
		// Text Input (Paragraph - Style 2)
		AddLabel(client.NewLabel("Feedback", "Tell us what you think",
			client.NewTextInput("feedback_input", "", 2,
				client.WithPlaceholder("Share your thoughts..."),
				client.WithMinLength(10),
				client.WithMaxLength(500)),
		)).
		// String Select with options
		AddLabel(client.NewLabel("Favorite Bug", "Choose your favorite bug",
			client.NewStringSelect("bug_select",
				types.SelectOption{Label: "Ant", Value: "ant", Description: "(best option)", Emoji: &types.PartialEmoji{Name: "🐜"}},
				types.SelectOption{Label: "Butterfly", Value: "butterfly", Emoji: &types.PartialEmoji{Name: "🦋"}},
				types.SelectOption{Label: "Caterpillar", Value: "caterpillar", Emoji: &types.PartialEmoji{Name: "🐛"}},
			),
		)).
		// User Select
		AddLabel(client.NewLabel("Pick a User", "Select a user to mention",
			client.NewUserSelect("user_select"),
		)).
		// Channel Select
		AddLabel(client.NewLabel("Pick a Channel", "Select a channel",
			client.NewChannelSelect("channel_select"),
		)).
		Build()

	err := bot.ShowModal(ctx, interaction, "Components V2 Modal Demo", "demo_modal", components)
	if err != nil {
		log.Printf("Failed to show modal: %v", err)
	}
}

func handleModalSubmit(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	var results []string

	for _, comp := range interaction.Data.Components {
		if container, ok := comp.(*types.Label); ok {
			switch inner := container.Component.(type) {
			case *types.TextInput:
				results = append(results, fmt.Sprintf("**%s**: %s", inner.CustomID, inner.Value))
			case *types.StringSelect:
				results = append(results, fmt.Sprintf("**%s**: %v", inner.CustomID, inner.Values))
			case *types.UserSelect:
				results = append(results, fmt.Sprintf("**%s**: %v", inner.CustomID, inner.Values))
			case *types.RoleSelect:
				results = append(results, fmt.Sprintf("**%s**: %v", inner.CustomID, inner.Values))
			case *types.ChannelSelect:
				results = append(results, fmt.Sprintf("**%s**: %v", inner.CustomID, inner.Values))
			case *types.MentionableSelect:
				results = append(results, fmt.Sprintf("**%s**: %v", inner.CustomID, inner.Values))
			case *types.FileUpload:
				results = append(results, fmt.Sprintf("**%s**: %v", inner.CustomID, inner.Values))
			}
		}
	}

	response := "## Modal Submission Results\n" + strings.Join(results, "\n")
	err := bot.RespondWithEphemeral(ctx, interaction, response)
	if err != nil {
		log.Printf("Failed to respond to modal submit: %v", err)
	}
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

func handleGallery(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# Media Gallery & File Demo").
		AddMediaGallery(
			types.MediaGalleryItem{
				Media:       types.UnfurledMediaItem{URL: "https://http.cat/200.jpg"},
				Description: "HTTP 200 OK",
			},
			types.MediaGalleryItem{
				Media:       types.UnfurledMediaItem{URL: "https://http.cat/404.jpg"},
				Description: "HTTP 404 Not Found",
			},
		).
		AddSeparator(true, types.SeparatorSpacingLarge).
		Build()

	// Currently no helper for adding File components to Builder, so we manually append
	components = append(components, &types.File{
		File: types.UnfurledMediaItem{URL: "https://http.cat/500"},
		Name: "error.txt",
		Size: 1024,
	})

	_, err := bot.SendMessageWithComponents(ctx, event.ChannelID, components)
	if err != nil {
		log.Printf("Failed to send gallery message: %v", err)
	}
}

func handleSection(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# Section Demo").
		AddSection(
			"Here is an example of a **Section** layout component, accompanied by a Thumbnail accessory.",
			&types.Thumbnail{
				Media:       types.UnfurledMediaItem{URL: "https://http.cat/418.jpg"},
				Description: "I'm a teapot",
			},
		).
		Build()

	// Manually add a Premium button since there is no helper yet
	premiumSKU := types.Snowflake(123456789)
	components = append(components, &types.ActionRow{
		Components: types.ComponentList{
			&types.Button{
				Style: types.ButtonStylePremium,
				SKUID: &premiumSKU,
			},
		},
	})

	_, err := bot.SendMessageWithComponents(ctx, event.ChannelID, components)
	if err != nil {
		log.Printf("Failed to send section message: %v", err)
	}
}

func handleHelp(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# Gophord Bot Help").
		AddTextDisplay("Available commands:").
		AddTextDisplay("- `!ping` - Check if the bot is responsive\n- `!hello` - Get a greeting with interactive buttons\n- `!kick <user_id>` - Kick a user from the server\n- `!webhook <name>` - Create and test a webhook\n- `!components` - See Components V2 features demo\n- `!container` - See a container with accent color\n- `!gallery` - See a media gallery and file attachment demo\n- `!section` - See a section component demo\n- `!help` - Show this help message").
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
		} else if interaction.Data.Name == "ephemeral" {
			err := bot.RespondWithEphemeral(ctx, interaction, "This is an ephemeral message! Only you can see this. 👀")
			if err != nil {
				log.Printf("Failed to respond to /ephemeral interaction: %v", err)
			}
		} else if interaction.Data.Name == "modal" {
			handleModalDemo(ctx, bot, interaction)
		} else if interaction.Data.Name == "gallery" {
			components := client.NewComponentBuilder().
				AddTextDisplay("# Media Gallery & File Demo").
				AddMediaGallery(
					types.MediaGalleryItem{
						Media:       types.UnfurledMediaItem{URL: "https://http.cat/200.jpg"},
						Description: "HTTP 200 OK",
					},
					types.MediaGalleryItem{
						Media:       types.UnfurledMediaItem{URL: "https://http.cat/404.jpg"},
						Description: "HTTP 404 Not Found",
					},
				).
				AddSeparator(true, types.SeparatorSpacingLarge).
				Build()

			components = append(components, &types.File{
				File: types.UnfurledMediaItem{URL: "https://http.cat/500"},
				Name: "error.txt",
				Size: 1024,
			})

			err := bot.RespondWithComponents(ctx, interaction, components)
			if err != nil {
				log.Printf("Failed to respond to /gallery interaction: %v", err)
			}
		} else if interaction.Data.Name == "section" {
			components := client.NewComponentBuilder().
				AddTextDisplay("# Section Demo").
				AddSection(
					"Here is an example of a **Section** layout component, accompanied by a Thumbnail accessory.",
					&types.Thumbnail{
						Media:       types.UnfurledMediaItem{URL: "https://http.cat/418.jpg"},
						Description: "I'm a teapot",
					},
				).
				Build()

			premiumSKU := types.Snowflake(123456789)
			components = append(components, &types.ActionRow{
				Components: types.ComponentList{
					&types.Button{
						Style: types.ButtonStylePremium,
						SKUID: &premiumSKU,
					},
				},
			})

			err := bot.RespondWithComponents(ctx, interaction, components)
			if err != nil {
				log.Printf("Failed to respond to /section interaction: %v", err)
			}
		}
		return
	}

	// Handle modal submissions
	if interaction.Type == types.InteractionTypeModalSubmit {
		handleModalSubmit(ctx, bot, interaction)
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
	isEphemeral := false

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
			"• Uses lxzan/gws for WebSocket\n" +
			"• Uses bytedance/sonic for fast JSON\n" +
			"• Full Components V2 support\n" +
			"• Native Mobile Status preset"
		isEphemeral = true // Make info ephemeral to keep chat clean
	case "say_goodbye":
		response = "Goodbye! See you next time! 👋"
	case "btn_primary", "btn_secondary", "btn_success", "btn_danger":
		response = fmt.Sprintf("You clicked the **%s** button!", strings.TrimPrefix(customID, "btn_"))
	case "container_action":
		response = "Action taken! ✅"
	case "container_dismiss":
		response = "Dismissed! 👍"
		isEphemeral = true
	default:
		response = fmt.Sprintf("Button clicked: `%s`", customID)
	}

	var err error
	if isEphemeral {
		err = bot.RespondWithEphemeral(ctx, interaction, response)
	} else {
		err = bot.RespondWithMessage(ctx, interaction, response)
	}

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
