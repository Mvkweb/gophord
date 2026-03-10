# Application Commands

> Create and manage slash commands, user commands, and message commands.

## Overview

Application Commands (often just "slash commands") are the primary way users interact with your bot. There are three types:
- **Chat Input** (`types.ApplicationCommandTypeChatInput`): Standard slash commands.
- **User** (`types.ApplicationCommandTypeUser`): Right-click a user -> Apps.
- **Message** (`types.ApplicationCommandTypeMessage`): Right-click a message -> Apps.

## Detailed Usage

To create or overwrite commands, use the REST API.

```go
package main

import (
	"log"

	"github.com/Mvkweb/gophord/pkg/rest"
	"github.com/Mvkweb/gophord/pkg/types"
)

func registerCommands(restClient *rest.Client, appID, guildID string) {
	cmdType := types.ApplicationCommandTypeChatInput

	commands := []types.CreateApplicationCommandParams{
		{
			Name:        "ping",
			Description: "Responds with pong!",
			Type:        &cmdType,
		},
		{
			Name:        "echo",
			Description: "Echoes your text",
			Type:        &cmdType,
			Options: []types.ApplicationCommandOption{
				{
					Type:        types.ApplicationCommandOptionTypeString,
					Name:        "text",
					Description: "The text to echo",
					Required:    true,
				},
			},
		},
	}

	// Bulk overwrite commands for a specific guild (instant update)
	// For global commands, pass an empty string "" for the guildID (takes up to an hour to propagate)
	registered, err := restClient.BulkOverwriteGuildApplicationCommands(appID, guildID, commands)
	if err != nil {
		log.Fatalf("Failed to register commands: %v", err)
	}

	log.Printf("Successfully registered %d commands", len(registered))
}
```

## API Reference

### Structs

| Struct | Description |
|---|---|
| `types.ApplicationCommand` | Represents an existing command returned from the API. |
| `types.CreateApplicationCommandParams` | Payload used when creating or bulk-overwriting commands. |
| `types.ApplicationCommandOption` | Represents a parameter for a Chat Input command. |

### Enums

| Type | Examples |
|---|---|
| `types.ApplicationCommandType` | `ApplicationCommandTypeChatInput`, `ApplicationCommandTypeUser`, `ApplicationCommandTypeMessage` |
| `types.ApplicationCommandOptionType` | `ApplicationCommandOptionTypeString`, `ApplicationCommandOptionTypeInteger`, `ApplicationCommandOptionTypeUser`, etc. |

## Related

- [Interactions](interactions.md) — How to handle these commands when users execute them.
- [REST Client](rest.md) — Documentation for the `rest` package used to register these commands.
