// Package gophord provides a high-performance Go library for interacting with the Discord API.
//
// Gophord is designed for production use with:
//   - High-performance JSON using bytedance/sonic
//   - Full Discord API v10+ support including Components V2
//   - Idiomatic Go design with comprehensive documentation
//   - Gateway WebSocket connection with automatic reconnection
//   - REST API client with rate limiting
//
// # Quick Start
//
// Create a new client and connect to Discord:
//
//	package main
//
//	import (
//	    "context"
//	    "log"
//
//	    "github.com/gophord/gophord/pkg/client"
//	    "github.com/gophord/gophord/pkg/gateway"
//	    "github.com/gophord/gophord/pkg/types"
//	)
//
//	func main() {
//	    bot := client.New("your-bot-token", client.WithIntents(types.IntentsDefault))
//
//	    bot.OnReady(func(event *gateway.ReadyEvent) {
//	        log.Printf("Logged in as %s", event.User.Username)
//	    })
//
//	    bot.OnMessageCreate(func(event *gateway.MessageCreateEvent) {
//	        if event.Content == "!ping" {
//	            bot.SendMessage(context.Background(), event.ChannelID, "Pong!")
//	        }
//	    })
//
//	    ctx := context.Background()
//	    if err := bot.Connect(ctx); err != nil {
//	        log.Fatal(err)
//	    }
//
//	    // Keep running...
//	    select {}
//	}
//
// # Components V2
//
// Gophord fully supports Discord's Components V2 system:
//
//	components := client.NewComponentBuilder().
//	    AddTextDisplay("# Hello World").
//	    AddSeparator(true, types.SeparatorSpacingSmall).
//	    AddActionRow(
//	        client.NewPrimaryButton("btn_click", "Click Me"),
//	    ).
//	    Build()
//
//	bot.SendMessageWithComponents(ctx, channelID, components)
//
// # Package Structure
//
// The library is organized into several packages:
//
//   - [github.com/gophord/gophord/pkg/types] - Core Discord type definitions
//   - [github.com/gophord/gophord/pkg/client] - High-level Discord client
//   - [github.com/gophord/gophord/pkg/rest] - REST API client
//   - [github.com/gophord/gophord/pkg/gateway] - WebSocket gateway client
//   - [github.com/gophord/gophord/pkg/json] - High-performance JSON utilities
//
// # Documentation
//
// This library is designed to work with external documentation generators.
// All public types and functions include comprehensive documentation comments.
//
// For more examples, see the examples directory in the repository.
package gophord

import (
	// Re-export main packages for convenience
	_ "github.com/gophord/gophord/pkg/client"
	_ "github.com/gophord/gophord/pkg/gateway"
	_ "github.com/gophord/gophord/pkg/json"
	_ "github.com/gophord/gophord/pkg/rest"
	_ "github.com/gophord/gophord/pkg/types"
)

// Version is the library version.
const Version = "1.0.0"
