// Package utils provides shared constants and helper functions for the simple_bot example.
// This centralizes common values to make the codebase easier to maintain and update.

package utils

import "github.com/Mvkweb/gophord/pkg/types"

// Command prefixes used throughout the bot.
const (
	// Prefix is the character used to invoke bot commands.
	Prefix = "!"
)

// Slash command names - keep these in sync with command registration.
const (
	CmdHello      = "hello"
	CmdEphemeral  = "ephemeral"
	CmdModal      = "modal"
	CmdGallery    = "gallery"
	CmdSection    = "section"
	CmdFileUpload = "fileupload"
)

// Button custom IDs - used to identify which button was clicked.
const (
	// Greeting buttons
	ButtonGreetBack  = "greet_back"
	ButtonShowInfo   = "show_info"
	ButtonSayGoodbye = "say_goodbye"

	// Component demo buttons
	ButtonPrimary   = "btn_primary"
	ButtonSecondary = "btn_secondary"
	ButtonSuccess   = "btn_success"
	ButtonDanger    = "btn_danger"

	// Container buttons
	ButtonContainerAction  = "container_action"
	ButtonContainerDismiss = "container_dismiss"

	// File upload
	ButtonOpenFileModal = "open_file_modal"
)

// Modal custom IDs - used to identify which modal was submitted.
const (
	ModalDemo       = "demo_modal"
	ModalFileUpload = "file_upload_modal"
)

// Response messages used throughout the bot.
var (
	// Bot info displayed when using the info button.
	BotInfo = "**Gophord Info**\n\n" +
		"• High-performance Go Discord library\n" +
		"• Uses lxzan/gws for WebSocket\n" +
		"• Uses bytedance/sonic for fast JSON\n" +
		"• Full Components V2 support\n" +
		"• Native Mobile Status preset"

	// Help footer text.
	HelpFooter = "-# Powered by gophord - A high-performance Go Discord library"
)

// Required intents for the bot to function properly.
// These must be requested when creating the client.
var RequiredIntents = types.IntentGuilds |
	types.IntentGuildMessages |
	types.IntentGuildMessageReactions |
	types.IntentDirectMessages |
	types.IntentMessageContent |
	types.IntentGuildMembers
