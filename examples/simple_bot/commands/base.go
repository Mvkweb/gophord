// Package commands provides prefix command handlers for the simple_bot example.
// Each file in this package handles a specific command or group of related commands.
//
// To add a new command:
// 1. Create a new file (e.g., mycommand.go)
// 2. Implement a handler function with signature: func HandleXXX(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent)
// 3. Add the command routing in message_create.go
//
// Command naming convention: HandleCommandName
package commands

import (
	"context"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
)

// CommandHandler defines the interface for all command handlers.
// Currently handlers are functions, but this allows for future extensibility.
type CommandHandler func(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent)

// MessageHandlerFunc is a type alias for message command handlers.
type MessageHandlerFunc func(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent)
