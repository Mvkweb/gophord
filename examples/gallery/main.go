// Example: Media Gallery
// Display multiple images in a gallery layout

package main

import (
	"context"
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
)

// DOC:START
// sendGalleryDemo shows multiple images in a gallery
func sendGalleryDemo(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# 🖼️ Media Gallery").
		AddTextDisplay("Galleries can display multiple images in a grid:").
		Build()

	// Add media gallery with multiple items
	gallery := &types.MediaGallery{
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
			{
				Media:       types.UnfurledMediaItem{URL: "https://picsum.photos/400/300?random=4"},
				Description: "Random landscape image 4",
			},
		},
	}

	// Combine components
	fullComponents := append(components, gallery, &types.TextDisplay{
		Content: "-# Images from Lorem Picsum",
	})

	_, err := bot.SendMessageWithComponents(ctx, channelID, fullComponents)
	if err != nil {
		log.Printf("Failed to send gallery: %v", err)
	}
}

// DOC:END

func main() {
	bot := client.New("YOUR_TOKEN", client.WithIntents(
		types.IntentGuilds|types.IntentGuildMessages|types.IntentMessageContent,
	))

	bot.OnMessageCreate(func(event *gateway.MessageCreateEvent) {
		if event.Content == "!gallery" {
			sendGalleryDemo(context.Background(), bot, event.ChannelID)
		}
	})

	bot.Connect(context.Background())
}
