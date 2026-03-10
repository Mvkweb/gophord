# Getting Started

> Build and run your first Gophord Discord bot quickly.

## Overview

Gophord is designed to make creating Discord bots in Go straightforward and idiomatic. This guide walks you through setting up a minimal bot that connects to the Discord gateway, logs when it is ready, and responds to a simple message.

## Setup

First, install Gophord:

```bash
go get github.com/Mvkweb/gophord
```

## Minimal Bot Example

The following code demonstrates how to create a bot, configure its required intents, listen for events, and properly shut down on exit.

```go
package main

import (
    "context"
    "log"
    "os"
    "os/signal"

    "github.com/Mvkweb/gophord/pkg/client"
    "github.com/Mvkweb/gophord/pkg/gateway"
    "github.com/Mvkweb/gophord/pkg/types"
)

func main() {
    token := os.Getenv("DISCORD_BOT_TOKEN")
    if token == "" {
        log.Fatal("DISCORD_BOT_TOKEN is required")
    }

    // 1. Create client with intents
    // To read message content in a guild, the IntentMessageContent flag is required.
    bot := client.New(token,
        client.WithIntents(types.IntentGuilds|types.IntentGuildMessages|types.IntentMessageContent),
        client.WithMobileStatus(true), // Appear as mobile client
    )

    // 2. Register event handlers
    bot.OnReady(func(event *gateway.ReadyEvent) {
        log.Printf("Logged in as %s#%s", event.User.Username, event.User.Discriminator)
    })

    bot.OnMessageCreate(func(event *gateway.MessageCreateEvent) {
        // Ignore messages from other bots
        if event.Author.Bot {
            return
        }

        if event.Content == "!ping" {
            // bot.SendMessage is for sending normal channel messages.
            // If replying to an Interaction, use bot.RespondWithMessage instead.
            bot.SendMessage(context.Background(), event.ChannelID, "Pong!")
        }
    })

    // 3. Connect to Discord
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    if err := bot.Connect(ctx); err != nil {
        log.Fatal(err)
    }

    // 4. Wait for interrupt
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    <-stop

    // 5. Clean up connection
    bot.Close()
}
```

## Related

- [Session & Client](session.md) — For advanced client configuration and lifecycle management
- [Events](events.md) — For handling additional gateway events
- [Intents](intents.md) — For understanding standard vs privileged intents
