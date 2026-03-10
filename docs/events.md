# Events

> Register handlers and respond to Discord gateway events.

## Overview

Events are triggered when things happen on Discord (e.g., a message is sent, a user joins, a button is clicked). Gophord allows you to register handlers on the high-level `client.Client` or process raw events directly from the `gateway.Client`.

## Detailed Usage

The easiest way to handle events is using the typed methods on the `client.Client`. This requires setting up before connection.

```go
package main

import (
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
)

func registerHandlers(bot *client.Client) {
	// Fired when the bot initially connects and is ready
	bot.OnReady(func(event *gateway.ReadyEvent) {
		log.Printf("Ready! Session ID: %s", event.SessionID)
	})

	// Fired when a message is created in a text channel
	bot.OnMessageCreate(func(event *gateway.MessageCreateEvent) {
		log.Printf("New message from %s: %s", event.Author.Username, event.Content)
	})

	// Fired when an interaction (slash command, button, modal) is triggered
	bot.OnInteractionCreate(func(event *gateway.InteractionCreateEvent) {
		log.Printf("Interaction received of type %d", event.Type)
	})
	
	// Fired for any event (useful for generic logic or unimplemented events)
	bot.On(gateway.EventGuildMemberAdd, func(data interface{}) {
		log.Printf("Raw event data: %v", data)
	})
}
```

## Advanced: Raw Gateway Events

If you are using the `gateway.Client` directly (without the high-level `client.Client`), you can listen to the `Events()` channel and use `gateway.ParseEvent` to cast the raw JSON into typed structs.

```go
package main

import (
	"log"

	"github.com/Mvkweb/gophord/pkg/gateway"
)

func processRawEvents(gw *gateway.Client) {
	for {
		select {
		case event := <-gw.Events():
			// Parse the raw JSON payload into a typed struct
			parsed, err := gateway.ParseEvent(event.Type, event.Data)
			if err != nil {
				log.Printf("Failed to parse event %s: %v", event.Type, err)
				continue
			}

			// Type switch on the parsed event
			switch e := parsed.(type) {
			case *gateway.ReadyEvent:
				log.Printf("Connected as %s", e.User.Username)
			case *gateway.MessageCreateEvent:
				log.Printf("Message: %s", e.Content)
			}
			
		case err := <-gw.Errors():
			log.Printf("Gateway error: %v", err)
		}
	}
}
```

## API Reference

### Client Handlers

| Method | Signature | Description |
|---|---|---|
| `On` | `func (c *Client) On(eventType string, handler EventHandler)` | Registers a generic handler for any event type string. |
| `OnReady` | `func (c *Client) OnReady(handler func(*gateway.ReadyEvent))` | Registers a typed handler for `READY`. |
| `OnMessageCreate` | `func (c *Client) OnMessageCreate(handler func(*gateway.MessageCreateEvent))` | Registers a typed handler for `MESSAGE_CREATE`. |
| `OnInteractionCreate` | `func (c *Client) OnInteractionCreate(handler func(*gateway.InteractionCreateEvent))` | Registers a typed handler for `INTERACTION_CREATE`. |

### Common Event Types

| Struct | Description |
|---|---|
| `gateway.ReadyEvent` | Fired on initial connection completion. Contains user info and session ID. |
| `gateway.MessageCreateEvent` | Fired when a message is sent. Embeds `types.Message`. |
| `gateway.InteractionCreateEvent` | Fired for slash commands, components, and modals. Embeds `types.Interaction`. |
| `gateway.TypingStartEvent` | Fired when a user starts typing. |
| `gateway.MessageReactionAddEvent` | Fired when a reaction is added to a message. |

*Note: You must request the correct intents (e.g., `types.IntentGuildMessages`) to receive corresponding events.*

## Related

- [Interactions](interactions.md) — How to handle slash commands and component clicks
- [Intents](intents.md) — What intents are required for what events
- [Gateway](gateway.md) — Low level gateway management
