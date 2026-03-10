# Interactions

> Respond to slash commands, component clicks, and modal submissions.

## Overview

When a user uses an Application Command or interacts with a Message Component (like a button or select menu), Discord sends an Interaction payload via the Gateway. In Gophord, you handle this using `client.OnInteractionCreate`.

### Interaction Types

| Constant | Description |
|---|---|
| `types.InteractionTypePing` | Ping (usually handled automatically). |
| `types.InteractionTypeApplicationCommand` | Slash command, User command, Message command. |
| `types.InteractionTypeMessageComponent` | Button click, Select Menu selection. |
| `types.InteractionTypeApplicationCommandAutocomplete` | Autocomplete input. |
| `types.InteractionTypeModalSubmit` | Modal form submitted. |

## Detailed Usage

### Handling Slash Commands

```go
package main

import (
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
)

func handleInteractions(bot *client.Client) {
	bot.OnInteractionCreate(func(e *gateway.InteractionCreateEvent) {
		// e is an embedded types.Interaction
		i := e.Interaction

		if i.Type == types.InteractionTypeApplicationCommand {
			data := i.Data
			
			// Switch based on command name
			switch data.Name {
			case "ping":
				err := bot.RespondWithMessage(i.ID, i.Token, "Pong!")
				if err != nil {
					log.Printf("Error responding: %v", err)
				}
				
			case "echo":
				// Read command options
				var echoText string
				if data.Options != nil && len(data.Options) > 0 {
					echoText = data.Options[0].Value.(string)
				}
				
				bot.RespondWithMessage(i.ID, i.Token, "Echo: "+echoText)
			}
		}
	})
}
```

### Response Methods

The high-level `client.Client` provides helper methods to respond easily. Note that you MUST respond within 3 seconds, or acknowledge and then use the REST client or follow-up webhooks.

| Method | Signature | Description |
|---|---|---|
| `RespondWithMessage` | `func (c *Client) RespondWithMessage(id types.Snowflake, token string, content string)` | Send a simple text message. |
| `RespondWithEmbeds` | `func (c *Client) RespondWithEmbeds(id types.Snowflake, token string, embeds []types.Embed)` | Send an embed message. |
| `RespondWithComponents` | `func (c *Client) RespondWithComponents(id types.Snowflake, token string, content string, components []types.Component)` | Send a message with UI components. |
| `RespondRaw` | `func (c *Client) RespondRaw(id types.Snowflake, token string, resp *types.InteractionResponse)` | Send a raw, manually constructed response. |

### Raw Interaction Responses

If you need more control (e.g., ephemeral messages, modals), use `RespondRaw`:

```go
func sendEphemeralMessage(bot *client.Client, i *types.Interaction) {
	bot.RespondRaw(i.ID, i.Token, &types.InteractionResponse{
		Type: types.InteractionCallbackTypeChannelMessageWithSource,
		Data: &types.InteractionCallbackData{
			Content: "This message is only visible to you!",
			Flags:   64, // EPHEMERAL flag
		},
	})
}
```

## Interaction Context and Structs

The raw `gateway.InteractionCreateEvent` embeds `types.Interaction`:

```go
type Interaction struct {
	ID            Snowflake                 // The ID of the interaction
	ApplicationID Snowflake                 // The ID of the bot
	Type          InteractionType           // The type (Command, Component, etc.)
	Data          *InteractionData          // The payload (command name/options, button custom_id)
	GuildID       *Snowflake                // ID of the guild (if in a guild)
	ChannelID     *Snowflake                // ID of the channel
	Member        *GuildMember              // The member who invoked it (if in a guild)
	User          *User                     // The user who invoked it (if in a DM)
	Token         string                    // A continuation token for responding
	Version       int                       // 1
	Message       *Message                  // The message the component was attached to
}
```

## Related

- [Application Commands](commands.md) — How to register the commands you are handling here.
- [Components V2](components-v2.md) — Interactive UI elements like buttons.
- [Modals](modals.md) — Handling Modal Submissions.
