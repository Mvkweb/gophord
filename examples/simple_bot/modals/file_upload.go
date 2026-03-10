// Package modals provides modal form handlers for the simple_bot example.
// This file handles the file upload modal demonstrations.

package modals

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/types"
)

// HandleFileUploadModal displays a modal with file upload component.
// Triggered by the /fileupload command or button click.
func HandleFileUploadModal(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
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
		log.Printf("Failed to show file upload modal: %v", err)
	}
}

// HandleFileUploadSubmit processes the submitted file upload modal.
// Displays information about uploaded files.
func HandleFileUploadSubmit(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
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
