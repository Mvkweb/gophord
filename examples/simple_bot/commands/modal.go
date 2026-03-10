// Package commands provides prefix command handlers for the simple_bot example.
// This file contains the modal demonstration command.

package commands

import (
	"context"

	"github.com/Mvkweb/gophord/examples/simple_bot/modals"
	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/types"
)

// HandleModalDemo opens a comprehensive modal demonstrating all modal component types.
// This includes text inputs, selects, and other form elements.
func HandleModalDemo(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	// Delegate to the modals package
	modals.HandleFeedbackModal(ctx, bot, interaction)
}
