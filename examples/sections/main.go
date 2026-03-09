// Example: Sections
// Create content with thumbnail or button accessories

package main

import (
	"context"
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
)

// DOC:START
// sendSectionDemo shows sections with accessories
func sendSectionDemo(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# 📋 Sections Demo").
		AddTextDisplay("Sections pair content with an accessory (thumbnail or button):").
		AddSeparator(true, types.SeparatorSpacingSmall).
		Build()

	// Section with thumbnail accessory
	thumbnailSection := &types.Section{
		Components: []types.Component{
			&types.TextDisplay{Content: "**With Thumbnail**\nThis section has an image on the right side."},
		},
		Accessory: &types.Thumbnail{
			Media:       types.UnfurledMediaItem{URL: "https://picsum.photos/80/80?random=1"},
			Description: "Section thumbnail",
		},
	}

	// Section with button accessory
	buttonSection := &types.Section{
		Components: []types.Component{
			&types.TextDisplay{Content: "**With Button**\nThis section has an action button on the right."},
		},
		Accessory: &types.Button{
			Style:    types.ButtonStylePrimary,
			CustomID: "section_action",
			Label:    "Action",
		},
	}

	// Product-style section
	productSection := &types.Section{
		Components: []types.Component{
			&types.TextDisplay{Content: "**Gophord Pro** 🚀\nHigh-performance Discord API wrapper\n• Blazing fast JSON parsing\n• Native Components V2\n• Zero-alloc design"},
		},
		Accessory: &types.Thumbnail{
			Media: types.UnfurledMediaItem{URL: "https://picsum.photos/80/80?random=2"},
		},
	}

	separator := &types.Separator{Divider: boolPtr(true), Spacing: types.SeparatorSpacingSmall}

	fullComponents := append(components,
		thumbnailSection,
		separator,
		buttonSection,
		separator,
		productSection,
	)

	_, err := bot.SendMessageWithComponents(ctx, channelID, fullComponents)
	if err != nil {
		log.Printf("Failed to send sections: %v", err)
	}
}

func boolPtr(b bool) *bool { return &b }

// DOC:END

func main() {
	bot := client.New("YOUR_TOKEN", client.WithIntents(
		types.IntentGuilds|types.IntentGuildMessages|types.IntentMessageContent,
	))

	bot.OnMessageCreate(func(event *gateway.MessageCreateEvent) {
		if event.Content == "!sections" {
			sendSectionDemo(context.Background(), bot, event.ChannelID)
		}
	})

	bot.OnInteractionCreate(func(event *gateway.InteractionCreateEvent) {
		if event.Type == types.InteractionTypeMessageComponent {
			if event.Data.CustomID == "section_action" {
				bot.RespondWithMessage(context.Background(), &event.Interaction, "Section action clicked!")
			}
		}
	})

	bot.Connect(context.Background())
}
