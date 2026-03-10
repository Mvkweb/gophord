// Package commands provides prefix command handlers for the simple_bot example.
// This file contains commands demonstrating Components V2 features:
// components, container, gallery, section, and fileupload.

package commands

import (
	"context"
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
)

// HandleComponents demonstrates various Components V2 features including:
// - Text Displays with markdown formatting
// - Separators
// - Interactive buttons with different styles
func HandleComponents(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
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

// HandleContainer demonstrates container components with accent colors.
// Containers group related content with a colored left border.
func HandleContainer(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
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

// HandleGallery demonstrates media galleries.
// Media galleries display multiple images in a grid.
// Note: The File component is for modal file uploads, not sending files in messages.
func HandleGallery(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# Media Gallery Demo").
		AddTextDisplay("Media galleries display multiple images in a grid.").
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
		Build()

	_, err := bot.SendMessageWithComponents(ctx, event.ChannelID, components)
	if err != nil {
		log.Printf("Failed to send gallery message: %v", err)
	}
}

// HandleSection demonstrates section components with thumbnail accessories.
// Sections organize content with optional media (images, thumbnails) on the side.
func HandleSection(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
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

	_, err := bot.SendMessageWithComponents(ctx, event.ChannelID, components)
	if err != nil {
		log.Printf("Failed to send section message: %v", err)
	}
}

// HandleFileUploadDemo demonstrates the file upload modal feature.
// Sends a message with a button that opens a modal for file uploads.
func HandleFileUploadDemo(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# 📎 File Upload Demo").
		AddTextDisplay("This demo shows how to use the **File Upload** component in modals.").
		AddTextDisplay("Click the button below to open a modal where you can upload files!").
		AddSeparator(true, types.SeparatorSpacingSmall).
		AddActionRow(
			client.NewPrimaryButton("open_file_modal", "Open File Upload Modal"),
		).
		Build()

	_, err := bot.SendMessageWithComponents(ctx, event.ChannelID, components)
	if err != nil {
		log.Printf("Failed to send file upload demo message: %v", err)
	}
}

// Helper function to create a boolean pointer.
// Required for some component fields that need pointer values.
func boolPtr(b bool) *bool {
	return &b
}
