// Package main demonstrates Message Components V2 usage with gophord.
//
// This example shows:
//   - Creating various component types (buttons, selects, text displays)
//   - Using containers with accent colors
//   - Building media galleries
//   - Handling component interactions
//
// To run this example:
//
//	go run main.go -token YOUR_BOT_TOKEN
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/rest"
	"github.com/Mvkweb/gophord/pkg/types"
)

func main() {
	token := flag.String("token", "", "Discord bot token")
	flag.Parse()

	if *token == "" {
		*token = os.Getenv("DISCORD_BOT_TOKEN")
		if *token == "" {
			log.Fatal("Bot token is required")
		}
	}

	bot := client.New(*token, client.WithIntents(
		types.IntentGuilds|
			types.IntentGuildMessages|
			types.IntentMessageContent,
	))

	bot.OnReady(func(event *gateway.ReadyEvent) {
		log.Printf("Ready as %s", event.User.Username)
	})

	bot.OnMessageCreate(func(event *gateway.MessageCreateEvent) {
		if event.Author.Bot {
			return
		}

		ctx := context.Background()
		content := strings.ToLower(event.Content)

		switch content {
		case "!buttons":
			sendButtonsDemo(ctx, bot.REST, event.ChannelID)
		case "!selects":
			sendSelectsDemo(ctx, bot.REST, event.ChannelID)
		case "!gallery":
			sendGalleryDemo(ctx, bot.REST, event.ChannelID)
		case "!section":
			sendSectionDemo(ctx, bot.REST, event.ChannelID)
		case "!full":
			sendFullDemo(ctx, bot.REST, event.ChannelID)
		}
	})

	bot.OnInteractionCreate(func(event *gateway.InteractionCreateEvent) {
		ctx := context.Background()

		if event.Type == types.InteractionTypeMessageComponent {
			handleComponentInteraction(ctx, bot, &event.Interaction)
		}
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := bot.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer bot.Close()

	log.Println("Components demo bot running. Commands: !buttons, !selects, !gallery, !section, !full")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
}

func sendButtonsDemo(ctx context.Context, restClient *rest.Client, channelID types.Snowflake) {
	// Demonstrate all button styles
	components := []types.Component{
		&types.TextDisplay{Content: "# Button Styles Demo"},
		&types.TextDisplay{Content: "Here are all the available button styles in Discord:"},
		&types.Separator{Divider: boolPtr(true), Spacing: types.SeparatorSpacingSmall},
		&types.ActionRow{
			Components: []types.Component{
				&types.Button{Style: types.ButtonStylePrimary, CustomID: "demo_primary", Label: "Primary"},
				&types.Button{Style: types.ButtonStyleSecondary, CustomID: "demo_secondary", Label: "Secondary"},
				&types.Button{Style: types.ButtonStyleSuccess, CustomID: "demo_success", Label: "Success"},
				&types.Button{Style: types.ButtonStyleDanger, CustomID: "demo_danger", Label: "Danger"},
			},
		},
		&types.ActionRow{
			Components: []types.Component{
				&types.Button{Style: types.ButtonStyleLink, URL: "https://github.com/Mvkweb/gophord", Label: "GitHub"},
				&types.Button{Style: types.ButtonStylePrimary, CustomID: "demo_emoji", Label: "With Emoji", Emoji: &types.PartialEmoji{Name: "🚀"}},
				&types.Button{Style: types.ButtonStyleSecondary, CustomID: "demo_disabled", Label: "Disabled", Disabled: true},
			},
		},
	}

	_, err := restClient.CreateMessage(ctx, channelID, &rest.CreateMessageParams{
		Components: components,
		Flags:      types.MessageFlagIsComponentsV2,
	})
	if err != nil {
		log.Printf("Failed to send buttons demo: %v", err)
	}
}

func sendSelectsDemo(ctx context.Context, restClient *rest.Client, channelID types.Snowflake) {
	components := []types.Component{
		&types.TextDisplay{Content: "# Select Menus Demo"},
		&types.TextDisplay{Content: "Select menus allow users to choose from a list of options:"},
		&types.Separator{Divider: boolPtr(true), Spacing: types.SeparatorSpacingSmall},
		&types.TextDisplay{Content: "## String Select"},
		&types.ActionRow{
			Components: []types.Component{
				&types.StringSelect{
					CustomID:    "demo_string_select",
					Placeholder: "Choose your favorite language",
					Options: []types.SelectOption{
						{Label: "Go", Value: "go", Description: "Simple and efficient", Emoji: &types.PartialEmoji{Name: "🐹"}},
						{Label: "Rust", Value: "rust", Description: "Memory safety", Emoji: &types.PartialEmoji{Name: "🦀"}},
						{Label: "Python", Value: "python", Description: "Easy to learn", Emoji: &types.PartialEmoji{Name: "🐍"}},
						{Label: "JavaScript", Value: "js", Description: "Web everywhere", Emoji: &types.PartialEmoji{Name: "📜"}},
					},
				},
			},
		},
		&types.Separator{Divider: boolPtr(true), Spacing: types.SeparatorSpacingSmall},
		&types.TextDisplay{Content: "## Auto-populated Selects"},
		&types.ActionRow{
			Components: []types.Component{
				&types.UserSelect{
					CustomID:    "demo_user_select",
					Placeholder: "Select a user",
				},
			},
		},
		&types.ActionRow{
			Components: []types.Component{
				&types.ChannelSelect{
					CustomID:     "demo_channel_select",
					Placeholder:  "Select a text channel",
					ChannelTypes: []types.ChannelType{types.ChannelTypeGuildText},
				},
			},
		},
	}

	_, err := restClient.CreateMessage(ctx, channelID, &rest.CreateMessageParams{
		Components: components,
		Flags:      types.MessageFlagIsComponentsV2,
	})
	if err != nil {
		log.Printf("Failed to send selects demo: %v", err)
	}
}

func sendGalleryDemo(ctx context.Context, restClient *rest.Client, channelID types.Snowflake) {
	components := []types.Component{
		&types.TextDisplay{Content: "# Media Gallery Demo"},
		&types.TextDisplay{Content: "Media galleries can display multiple images:"},
		&types.MediaGallery{
			Items: []types.MediaGalleryItem{
				{
					Media:       types.UnfurledMediaItem{URL: "https://picsum.photos/400/300?random=1"},
					Description: "Random landscape image 1",
				},
				{
					Media:       types.UnfurledMediaItem{URL: "https://picsum.photos/400/300?random=2"},
					Description: "Random landscape image 2",
				},
				{
					Media:       types.UnfurledMediaItem{URL: "https://picsum.photos/400/300?random=3"},
					Description: "Random landscape image 3",
				},
			},
		},
		&types.TextDisplay{Content: "-# Images from Lorem Picsum"},
	}

	_, err := restClient.CreateMessage(ctx, channelID, &rest.CreateMessageParams{
		Components: components,
		Flags:      types.MessageFlagIsComponentsV2,
	})
	if err != nil {
		log.Printf("Failed to send gallery demo: %v", err)
	}
}

func sendSectionDemo(ctx context.Context, restClient *rest.Client, channelID types.Snowflake) {
	components := []types.Component{
		&types.TextDisplay{Content: "# Sections Demo"},
		&types.TextDisplay{Content: "Sections associate text with an accessory (button or thumbnail):"},
		&types.Separator{Divider: boolPtr(true), Spacing: types.SeparatorSpacingSmall},
		&types.Section{
			Components: []types.Component{
				&types.TextDisplay{Content: "**With Thumbnail Accessory**\nThis section has a thumbnail image on the right side."},
			},
			Accessory: &types.Thumbnail{
				Media:       types.UnfurledMediaItem{URL: "https://picsum.photos/80/80?random=4"},
				Description: "Section thumbnail",
			},
		},
		&types.Separator{Divider: boolPtr(true), Spacing: types.SeparatorSpacingSmall},
		&types.Section{
			Components: []types.Component{
				&types.TextDisplay{Content: "**With Button Accessory**\nThis section has an action button on the right side."},
			},
			Accessory: &types.Button{
				Style:    types.ButtonStylePrimary,
				CustomID: "section_action",
				Label:    "Action",
			},
		},
	}

	_, err := restClient.CreateMessage(ctx, channelID, &rest.CreateMessageParams{
		Components: components,
		Flags:      types.MessageFlagIsComponentsV2,
	})
	if err != nil {
		log.Printf("Failed to send section demo: %v", err)
	}
}

func sendFullDemo(ctx context.Context, restClient *rest.Client, channelID types.Snowflake) {
	// Demonstrate a comprehensive message with multiple component types
	accentColor := 0x57F287 // Green

	container := &types.Container{
		AccentColor: &accentColor,
		Components: []types.Component{
			&types.TextDisplay{Content: "# 🎮 Game Night Registration"},
			&types.TextDisplay{Content: "Join us for our weekly game night! Fill out the form below to register."},
			&types.Separator{Divider: boolPtr(true), Spacing: types.SeparatorSpacingSmall},
			&types.Section{
				Components: []types.Component{
					&types.TextDisplay{Content: "**Date:** Saturday, 8 PM UTC\n**Duration:** ~3 hours\n**Voice Channel:** #game-night"},
				},
				Accessory: &types.Thumbnail{
					Media: types.UnfurledMediaItem{URL: "https://picsum.photos/80/80?random=5"},
				},
			},
			&types.Separator{Divider: boolPtr(true), Spacing: types.SeparatorSpacingSmall},
			&types.TextDisplay{Content: "## Select Your Games"},
			&types.ActionRow{
				Components: []types.Component{
					&types.StringSelect{
						CustomID:    "game_select",
						Placeholder: "Choose games you want to play",
						MinValues:   intPtr(1),
						MaxValues:   intPtr(3),
						Options: []types.SelectOption{
							{Label: "Among Us", Value: "among_us", Emoji: &types.PartialEmoji{Name: "🚀"}},
							{Label: "Jackbox Games", Value: "jackbox", Emoji: &types.PartialEmoji{Name: "🎭"}},
							{Label: "Minecraft", Value: "minecraft", Emoji: &types.PartialEmoji{Name: "⛏️"}},
							{Label: "Gartic Phone", Value: "gartic", Emoji: &types.PartialEmoji{Name: "📞"}},
							{Label: "Skribbl.io", Value: "skribbl", Emoji: &types.PartialEmoji{Name: "🎨"}},
						},
					},
				},
			},
			&types.Separator{Divider: boolPtr(false), Spacing: types.SeparatorSpacingSmall},
			&types.ActionRow{
				Components: []types.Component{
					&types.Button{Style: types.ButtonStyleSuccess, CustomID: "register", Label: "Register", Emoji: &types.PartialEmoji{Name: "✅"}},
					&types.Button{Style: types.ButtonStyleSecondary, CustomID: "maybe", Label: "Maybe"},
					&types.Button{Style: types.ButtonStyleDanger, CustomID: "cant_attend", Label: "Can't Attend"},
				},
			},
		},
	}

	components := []types.Component{container}

	_, err := restClient.CreateMessage(ctx, channelID, &rest.CreateMessageParams{
		Components: components,
		Flags:      types.MessageFlagIsComponentsV2,
	})
	if err != nil {
		log.Printf("Failed to send full demo: %v", err)
	}
}

func handleComponentInteraction(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	customID := interaction.Data.CustomID
	var response string

	switch {
	case customID == "demo_string_select":
		if len(interaction.Data.Values) > 0 {
			response = "You selected: " + strings.Join(interaction.Data.Values, ", ")
		}
	case customID == "demo_user_select":
		response = "User select interaction received!"
	case customID == "demo_channel_select":
		response = "Channel select interaction received!"
	case customID == "game_select":
		response = "You want to play: " + strings.Join(interaction.Data.Values, ", ")
	case customID == "register":
		response = "✅ You're registered for game night!"
	case customID == "maybe":
		response = "📝 Marked as maybe attending"
	case customID == "cant_attend":
		response = "😢 Sorry you can't make it!"
	case customID == "section_action":
		response = "Section action clicked!"
	case strings.HasPrefix(customID, "demo_"):
		response = "Button clicked: " + strings.TrimPrefix(customID, "demo_")
	default:
		response = "Interaction received: " + customID
	}

	if response != "" {
		bot.RespondWithMessage(ctx, interaction, response)
	}
}

func boolPtr(b bool) *bool {
	return &b
}

func intPtr(i int) *int {
	return &i
}
