# Session & Client

> Client creation, options, lifecycle, and initialization.

## Overview

The `client.Client` acts as the primary entry point for the Gophord wrapper. It provides a high-level, unified interface for interacting with Discord, handling both HTTP REST requests and WebSocket gateway events. Under the hood, it orchestrates the `pkg/rest` and `pkg/gateway` sub-clients.

## Basic Usage

```go
package main

import (
    "context"
    "log"

    "github.com/Mvkweb/gophord/pkg/client"
    "github.com/Mvkweb/gophord/pkg/types"
)

func main() {
    token := "YOUR_BOT_TOKEN_HERE"
    ctx := context.Background()

    // Create the client
    bot := client.New(token,
        client.WithIntents(types.IntentsDefault | types.IntentMessageContent),
        client.WithMobileStatus(false),
    )

    // Register handlers before connecting
    // bot.OnReady(...)

    // Connect to Discord
    if err := bot.Connect(ctx); err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }

    // Always defer closing the connection
    defer bot.Close()

    // Access underlying REST client directly if needed:
    // user, err := bot.REST.GetCurrentUser(ctx)
}
```

## Detailed Usage

### Customizing Options
When calling `client.New`, provide options as functional arguments.

```go
bot := client.New(token,
    // It's required to pass intents to receive events
    client.WithIntents(types.IntentGuilds | types.IntentGuildMessages),
    // Setting to true shows a mobile phone icon next to the bot's status
    client.WithMobileStatus(true),
)
```

## API Reference

| Method / Function | Signature | Description |
|---|---|---|
| `New` | `func New(token string, opts ...ClientOption) *Client` | Creates a new high-level Discord client. |
| `WithIntents` | `func WithIntents(intents types.IntentFlags) ClientOption` | Specifies which gateway events the bot asks Discord to send. |
| `WithMobileStatus` | `func WithMobileStatus(enabled bool) ClientOption` | Identifies the bot as using the mobile platform. |
| `Connect` | `func (c *Client) Connect(ctx context.Context) error` | Connects to the gateway and begins handling events. |
| `Close` | `func (c *Client) Close() error` | Closes the gateway connection and terminates the session. |

## Sub-Clients

The `client.Client` encapsulates both the REST and Gateway connections. If you need fine-grained control or lower-level API access, you can access these directly:

```go
// Direct REST API access
bot.REST.GetGuild(ctx, guildID)

// Direct Gateway access
bot.Gateway.UpdatePresence(&presenceUpdate)
```

## Related

- [Intents](intents.md) — Understand what intents are required for your bot
- [Gateway](gateway.md) — For lower-level gateway connections
- [REST API](rest.md) — For lower-level REST client access
