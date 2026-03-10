# Modals

> Popup forms that accept text input and select menu choices.

## Overview

Modals are popup windows that gather structured input from a user. They can only be triggered as a direct response to an interaction (like clicking a button or using a slash command). A modal can contain up to 5 `ActionRow`s.

Within a modal's Action Rows, you can place:
- `types.TextInput`
- `types.StringSelect` (and other selects, though less common)

**Important**: You must wrap inputs in a `types.Label` component to give them a visible name and description.

## Detailed Usage

To show a modal, use `bot.ShowModal(ctx, interaction, title, customID, components)`.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/types"
)

func showFeedbackModal(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	components := client.NewComponentBuilder().
		// Adding a Text Input. Must be wrapped in a Label!
		AddLabel(client.NewLabel("Your Name", "Enter your name",
			client.NewTextInput("name_input", "", 1, // 1 = Short Style
				client.WithRequired(true),
				client.WithPlaceholder("John Doe"),
			),
		)).
		// Adding a Multi-line Text Input
		AddLabel(client.NewLabel("Feedback", "Share your thoughts",
			client.NewTextInput("feedback_input", "", 2, // 2 = Paragraph Style
				client.WithPlaceholder("Tell us what you think..."),
				client.WithMinLength(10),
				client.WithMaxLength(500),
			),
		)).
		// Adding a Select Menu
		AddLabel(client.NewLabel("Rating", "How would you rate us?",
			client.NewStringSelect("rating_select",
				types.SelectOption{Label: "⭐⭐⭐⭐⭐ Excellent", Value: "5"},
				types.SelectOption{Label: "⭐ Terrible", Value: "1"},
			),
		)).
		Build()

	// ShowModal sends an InteractionCallbackTypeModal
	err := bot.ShowModal(ctx, interaction, "Feedback Form", "feedback_modal", components)
	if err != nil {
		log.Printf("Failed to show modal: %v", err)
	}
}
```

## Handling Submissions

When the user submits the form, Discord sends an `InteractionTypeModalSubmit`. 

```go
func handleModalSubmit(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	// The CustomID of the Modal itself
	if interaction.Data.CustomID != "feedback_modal" {
		return
	}

	var results []string

	// Iterate through the raw components sent back by Discord
	for _, comp := range interaction.Data.Components {
		if label, ok := comp.(*types.Label); ok {
			switch inner := label.Component.(type) {
			case *types.TextInput:
				results = append(results, fmt.Sprintf("**%s**: %s", inner.CustomID, inner.Value))
			case *types.StringSelect:
				results = append(results, fmt.Sprintf("**%s**: %s", inner.CustomID, strings.Join(inner.Values, ", ")))
			}
		}
	}

	response := "## Form Submitted!\n" + strings.Join(results, "\n")
	bot.RespondWithEphemeral(ctx, interaction, response)
}
```

## API Reference

### Response Method

| Method | Signature | Description |
|---|---|---|
| `ShowModal` | `func (c *Client) ShowModal(ctx context.Context, interaction *types.Interaction, title, customID string, components types.ComponentList) error` | Triggers the modal popup for the interaction's user. |

### ComponentBuilder Methods

| Method | Signature | Description |
|---|---|---|
| `AddLabel` | `func (b *ComponentBuilder) AddLabel(label *types.Label) *ComponentBuilder` | Helper to wrap inputs in the required Label format. |
| `NewLabel` | `func NewLabel(title, description string, child types.Component) *types.Label` | Factory for `types.Label`. |
| `NewTextInput` | `func NewTextInput(customID, value string, style types.TextInputStyle, opts ...TextInputOption) *types.TextInput` | Factory for `types.TextInput`. |

### Text Input Options

Passed to `NewTextInput` to customize the field:

| Option | Signature | Description |
|---|---|---|
| `WithPlaceholder` | `func WithPlaceholder(placeholder string) TextInputOption` | Text showing when empty. |
| `WithMinLength` | `func WithMinLength(min int) TextInputOption` | Validation. |
| `WithMaxLength` | `func WithMaxLength(max int) TextInputOption` | Validation. |
| `WithRequired` | `func WithRequired(req bool) TextInputOption` | Is completion mandatory? |

### Struct Reference

| Struct | Notes |
|---|---|
| `types.TextInput` | Represents the input itself. Values are read from `.Value`. |
| `types.Label` | The Discord API requires inputs in modals to be placed inside a specific wrapping structure (Label -> ActionRow -> Input). Gophord's `client.NewLabel` and `client.ComponentBuilder` handle this abstraction for you automatically. |

## Related

- [Components V2](components-v2.md)
- [Interactions](interactions.md)
