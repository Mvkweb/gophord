# Gateway

> Direct WebSocket gateway connection handling, heartbeating, and raw events.

## Overview

The `gateway` package manages the WebSocket connection to Discord. It handles connection, authentication, automatic heartbeating, reconnecting after drops, and identifying. Most applications will use `client.Client` which wraps this, but building custom high-performance shards may require using `gateway.Client` directly.

## Detailed Usage

Creating and managing a raw gateway client manually requires consuming events from channels rather than registering callback functions.

```go
package main

import (
    "context"
    "log"

    "github.com/Mvkweb/gophord/pkg/gateway"
    "github.com/Mvkweb/gophord/pkg/types"
)

func runGateway(token string) {
    // 1. Create gateway client with options
    gw := gateway.New(token,
        gateway.WithIntents(types.IntentsDefault),
        gateway.WithMobileStatus(false),
    )

    // 2. Connect
    if err := gw.Connect(context.Background()); err != nil {
        log.Fatal(err)
    }
    defer gw.Close()

    // 3. Process events from the Event and Error channels
    for {
        select {
        case event := <-gw.Events():
            // Basic raw event data
            // To make it useful, parse it:
            parsed, err := gateway.ParseEvent(event.Type, event.Data)
            if err != nil {
                log.Printf("Parse error: %v", err)
                continue
            }

            switch e := parsed.(type) {
            case *gateway.ReadyEvent:
                log.Printf("Ready! Session: %s", e.SessionID)
            case *gateway.MessageCreateEvent:
                log.Printf("Message from %s: %s", e.Author.Username, e.Content)
            }

        case err := <-gw.Errors():
            log.Printf("Gateway error: %v", err)
        }
    }
}
```

## Advanced Patterns

### Updating Presence

You can update the bot's status and activity directly via the gateway.

```go
func updatePresence(gw *gateway.Client) {
    gw.UpdatePresence(&gateway.PresenceUpdate{
        Status: "online", // "online", "idle", "dnd", "invisible"
        Activities: []gateway.Activity{
            {
                Name: "with the Discord API",
                Type: 0, // 0 = Playing, 1 = Streaming, 2 = Listening, 3 = Watching, 4 = Custom Status, 5 = Competing
            },
        },
        AFK: false,
    })
}
```

### Setting Custom Status

Bots can now set custom status (like user custom statuses). Use activity type `4` with the `State` field:

```go
func setCustomStatus(gw *gateway.Client) {
    gw.UpdatePresence(&gateway.PresenceUpdate{
        Status: "online",
        Activities: []gateway.Activity{
            {
                Type:  4, // Custom Status
                State: "evelith.dev",
            },
        },
        AFK: false,
    })
}
```

- **Type 4** = Custom Status
- **State** = the text displayed (e.g., "evelith.dev")

## API Reference

### Client Instantiation & Lifecycle

| Method / Function | Signature | Description |
|---|---|---|
| `New` | `func New(token string, opts ...ClientOption) *Client` | Creates a new gateway client. |
| `WithIntents` | `func WithIntents(intents types.IntentFlags) ClientOption` | Configures the events Discord sends over the WS connection. |
| `WithMobileStatus` | `func WithMobileStatus(enabled bool) ClientOption` | Mimics Discord Android to get a phone status icon. |
| `Connect` | `func (c *Client) Connect(ctx context.Context) error` | Establishes the WebSocket connection using `bytedance/gws`. |
| `Close` | `func (c *Client) Close() error` | Closes the connection. |

### Event Channels

| Method | Signature | Description |
|---|---|---|
| `Events` | `func (c *Client) Events() <-chan *Event` | Returns a read-only channel for raw `gateway.Event` objects. |
| `Errors` | `func (c *Client) Errors() <-chan error` | Returns a read-only channel for gateway errors (e.g., missed heartbeats, disconnects). |
| `ParseEvent` | `func ParseEvent(eventType string, data []byte) (interface{}, error)` | Used to deserialize a generic event payload into a specific struct (e.g., `*gateway.ReadyEvent`). |

### Client Methods

| Method | Signature | Description |
|---|---|---|
| `UpdatePresence` | `func (c *Client) UpdatePresence(presence *PresenceUpdate) error` | Send a presence update payload to Discord. |

## Related

- [Events](events.md) — What to do with the messages coming through `Events()`.
- [Session & Client](session.md) — The higher-level abstraction.
