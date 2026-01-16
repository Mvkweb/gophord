// Package types provides Discord API type definitions.
//
// This package contains all the core data structures used to interact
// with the Discord API v10+, including:
//
//   - Snowflake IDs and timestamps
//   - Message Components V2 (buttons, selects, containers, sections, etc.)
//   - Users, guilds, channels, and messages
//   - Interactions and responses
//
// All types use idiomatic Go naming conventions and include comprehensive
// documentation for external documentation generation tools.
package types

// APIVersion is the Discord API version this library targets.
const APIVersion = "10"

// BaseURL is the base URL for the Discord API.
const BaseURL = "https://discord.com/api/v" + APIVersion

// GatewayURL is the base URL for the Discord Gateway.
const GatewayURL = "wss://gateway.discord.gg/?v=" + APIVersion + "&encoding=json"

// CDN URLs
const (
	// CDNURL is the base URL for Discord's CDN.
	CDNURL = "https://cdn.discordapp.com"
	// CDNEmojiURL is the URL format for emoji images.
	CDNEmojiURL = CDNURL + "/emojis/"
	// CDNAvatarURL is the URL format for user avatars.
	CDNAvatarURL = CDNURL + "/avatars/"
	// CDNIconURL is the URL format for guild icons.
	CDNIconURL = CDNURL + "/icons/"
)

// IntentFlags represents gateway intent flags.
type IntentFlags int

const (
	// IntentGuilds enables guild-related events.
	IntentGuilds IntentFlags = 1 << 0
	// IntentGuildMembers enables member-related events (privileged).
	IntentGuildMembers IntentFlags = 1 << 1
	// IntentGuildModeration enables moderation events.
	IntentGuildModeration IntentFlags = 1 << 2
	// IntentGuildEmojisAndStickers enables emoji/sticker events.
	IntentGuildEmojisAndStickers IntentFlags = 1 << 3
	// IntentGuildIntegrations enables integration events.
	IntentGuildIntegrations IntentFlags = 1 << 4
	// IntentGuildWebhooks enables webhook events.
	IntentGuildWebhooks IntentFlags = 1 << 5
	// IntentGuildInvites enables invite events.
	IntentGuildInvites IntentFlags = 1 << 6
	// IntentGuildVoiceStates enables voice state events.
	IntentGuildVoiceStates IntentFlags = 1 << 7
	// IntentGuildPresences enables presence events (privileged).
	IntentGuildPresences IntentFlags = 1 << 8
	// IntentGuildMessages enables guild message events.
	IntentGuildMessages IntentFlags = 1 << 9
	// IntentGuildMessageReactions enables guild reaction events.
	IntentGuildMessageReactions IntentFlags = 1 << 10
	// IntentGuildMessageTyping enables guild typing events.
	IntentGuildMessageTyping IntentFlags = 1 << 11
	// IntentDirectMessages enables DM message events.
	IntentDirectMessages IntentFlags = 1 << 12
	// IntentDirectMessageReactions enables DM reaction events.
	IntentDirectMessageReactions IntentFlags = 1 << 13
	// IntentDirectMessageTyping enables DM typing events.
	IntentDirectMessageTyping IntentFlags = 1 << 14
	// IntentMessageContent enables message content access (privileged).
	IntentMessageContent IntentFlags = 1 << 15
	// IntentGuildScheduledEvents enables scheduled event events.
	IntentGuildScheduledEvents IntentFlags = 1 << 16
	// IntentAutoModerationConfiguration enables auto-mod config events.
	IntentAutoModerationConfiguration IntentFlags = 1 << 20
	// IntentAutoModerationExecution enables auto-mod execution events.
	IntentAutoModerationExecution IntentFlags = 1 << 21
	// IntentGuildMessagePolls enables guild poll events.
	IntentGuildMessagePolls IntentFlags = 1 << 24
	// IntentDirectMessagePolls enables DM poll events.
	IntentDirectMessagePolls IntentFlags = 1 << 25
)

// IntentsAll combines all non-privileged intents.
var IntentsAll = IntentGuilds |
	IntentGuildModeration |
	IntentGuildEmojisAndStickers |
	IntentGuildIntegrations |
	IntentGuildWebhooks |
	IntentGuildInvites |
	IntentGuildVoiceStates |
	IntentGuildMessages |
	IntentGuildMessageReactions |
	IntentGuildMessageTyping |
	IntentDirectMessages |
	IntentDirectMessageReactions |
	IntentDirectMessageTyping |
	IntentGuildScheduledEvents |
	IntentAutoModerationConfiguration |
	IntentAutoModerationExecution |
	IntentGuildMessagePolls |
	IntentDirectMessagePolls

// IntentsPrivileged combines all privileged intents.
var IntentsPrivileged = IntentGuildMembers |
	IntentGuildPresences |
	IntentMessageContent

// IntentsDefault provides a sensible default for most bots.
var IntentsDefault = IntentGuilds |
	IntentGuildMessages |
	IntentGuildMessageReactions |
	IntentDirectMessages
