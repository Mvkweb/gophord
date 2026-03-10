# Messages

> Send text and content to text channels and users.

## Overview

The `types.Message` struct represents a message on Discord. While the Gateway receives messages via events, you send messages using the REST API (or as a response to an Interaction).

## Detailed Usage

To send messages proactively (not as an immediate response to an interaction within 3 seconds), use the REST Client.

```go
package main

import (
    "log"

    "github.com/Mvkweb/gophord/pkg/rest"
    "github.com/Mvkweb/gophord/pkg/types"
)

func sendMessages(restClient *rest.Client, channelID string) {
    // 1. Simple text message
    msg, err := restClient.CreateMessage(channelID, rest.CreateMessageParams{
        Content: "Hello, Discord!",
    })
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    log.Printf("Sent message %s", msg.ID)

    // 2. Message replying to another message
    _, err = restClient.CreateMessage(channelID, rest.CreateMessageParams{
        Content: "This is a reply",
        MessageReference: &types.MessageReference{
            MessageID: msg.ID,
            ChannelID: types.Snowflake(channelID),
            FailIfNotExists: true,
        },
    })

    // 3. Message with mentions disabled (silent)
    _, err = restClient.CreateMessage(channelID, rest.CreateMessageParams{
        Content: "Testing silent mentions @everyone",
        Flags:   4096, // MessageFlagSuppressNotifications
        AllowedMentions: &types.AllowedMentions{
            Parse: []string{}, // Empty array parses no mentions
        },
    })
}
```

## API Reference

### Structure

| Struct | Description |
|---|---|
| `types.Message` | A message object received from Discord. |
| `rest.CreateMessageParams` | The payload used to create a new message via REST. |
| `types.MessageReference` | Reference to another message (for replies or forwards). |
| `types.AllowedMentions` | Fine-grained control over who gets pinged. |

### Sending Messages

| Method | Signature | Description |
|---|---|---|
| `CreateMessage` | `func (c *Client) CreateMessage(channelID string, params CreateMessageParams) (*types.Message, error)` | Send a message to a channel. |
| `EditMessage` | `func (c *Client) EditMessage(channelID, messageID string, params EditMessageParams) (*types.Message, error)` | Edit an existing message. |
| `DeleteMessage` | `func (c *Client) DeleteMessage(channelID, messageID string) error` | Delete a message. |

### Common Flags

| Flag Value | Description |
|---|---|
| `1 << 6` (64) | **Ephemeral** (Interaction responses only): Only the user who invoked the interaction sees this message. |
| `1 << 2` (4) | **Suppress Embeds**: Do not include any auto-generated embeds from URLs. |
| `1 << 12` (4096) | **Suppress Notifications**: Send silently without pinging users (but still visible). |

## Related

- [Embeds](embeds.md) — How to add rich embeds to your messages.
- [Components V2](components-v2.md) — Attaching interactive elements like buttons.
- [Interactions](interactions.md) — Sending messages as a response to a command.
