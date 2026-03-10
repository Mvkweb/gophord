// Example: File Uploads in Modals
// Demonstrates using the FileUpload component to accept file attachments in modals

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
	godotenv "github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("Bot token required. Set DISCORD_BOT_TOKEN environment variable.")
	}

	bot := client.New(token, client.WithIntents(
		types.IntentGuilds|types.IntentGuildMessages|types.IntentMessageContent,
	))

	bot.OnReady(func(event *gateway.ReadyEvent) {
		log.Printf("Bot ready as %s#%s", event.User.Username, event.User.Discriminator)

		bot.RegisterGlobalCommand(context.Background(), types.CreateApplicationCommandParams{
			Name:        "fileupload",
			Description: "Open a modal with file upload component",
		})
	})

	bot.OnInteractionCreate(func(event *gateway.InteractionCreateEvent) {
		ctx := context.Background()

		switch event.Type {
		case types.InteractionTypeApplicationCommand:
			if event.Data.Name == "fileupload" {
				showFileUploadModal(ctx, bot, &event.Interaction)
			}
		case types.InteractionTypeModalSubmit:
			handleFileUploadSubmit(ctx, bot, &event.Interaction)
		}
	})

	bot.Connect(context.Background())
}

func showFileUploadModal(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	components := client.NewComponentBuilder().
		AddLabel(client.NewLabel("Title", "Enter a title for your file",
			client.NewTextInput("title", "", 1,
				client.WithPlaceholder("My awesome file"),
				client.WithRequired(true),
			),
		)).
		AddLabel(client.NewLabel("Description", "Describe what you're uploading",
			client.NewTextInput("description", "", 2,
				client.WithPlaceholder("This file contains..."),
				client.WithMaxLength(500),
			),
		)).
		AddLabel(client.NewLabel("Upload File", "Select a file to upload",
			client.NewFileUpload("file_upload"),
		)).
		Build()

	err := bot.ShowModal(ctx, interaction, "File Upload Demo", "file_upload_modal", components)
	if err != nil {
		log.Printf("Failed to show modal: %v", err)
	}
}

func handleFileUploadSubmit(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	var results []string
	var fileCount int

	for _, comp := range interaction.Data.Components {
		if label, ok := comp.(*types.Label); ok {
			switch inner := label.Component.(type) {
			case *types.TextInput:
				results = append(results, fmt.Sprintf("**%s**: %s", inner.CustomID, inner.Value))
			case *types.FileUpload:
				fileCount = len(inner.Values)
				if fileCount > 0 {
					results = append(results, fmt.Sprintf("**%s**: %d file(s) uploaded - IDs: `%v`",
						inner.CustomID, fileCount, inner.Values))
				} else {
					results = append(results, fmt.Sprintf("**%s**: No files uploaded", inner.CustomID))
				}
			}
		}
	}

	if fileCount == 0 {
		results = append(results, "\n⚠️ No files were attached. Please try again with a file!")
	}

	response := "## 📎 File Upload Received!\n" + strings.Join(results, "\n")
	err := bot.RespondWithEphemeral(ctx, interaction, response)
	if err != nil {
		log.Printf("Failed to respond: %v", err)
	}
}
