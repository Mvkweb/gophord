// Package gateway provides a WebSocket client for the Discord Gateway.
package gateway

import (
	"github.com/gophord/gophord/pkg/json"
	"github.com/gophord/gophord/pkg/types"
)

// Event type constants for gateway events.
const (
	EventReady                               = "READY"
	EventResumed                             = "RESUMED"
	EventApplicationCommandPermissionsUpdate = "APPLICATION_COMMAND_PERMISSIONS_UPDATE"
	EventAutoModerationRuleCreate            = "AUTO_MODERATION_RULE_CREATE"
	EventAutoModerationRuleUpdate            = "AUTO_MODERATION_RULE_UPDATE"
	EventAutoModerationRuleDelete            = "AUTO_MODERATION_RULE_DELETE"
	EventAutoModerationActionExecution       = "AUTO_MODERATION_ACTION_EXECUTION"
	EventChannelCreate                       = "CHANNEL_CREATE"
	EventChannelUpdate                       = "CHANNEL_UPDATE"
	EventChannelDelete                       = "CHANNEL_DELETE"
	EventChannelPinsUpdate                   = "CHANNEL_PINS_UPDATE"
	EventThreadCreate                        = "THREAD_CREATE"
	EventThreadUpdate                        = "THREAD_UPDATE"
	EventThreadDelete                        = "THREAD_DELETE"
	EventThreadListSync                      = "THREAD_LIST_SYNC"
	EventThreadMemberUpdate                  = "THREAD_MEMBER_UPDATE"
	EventThreadMembersUpdate                 = "THREAD_MEMBERS_UPDATE"
	EventGuildCreate                         = "GUILD_CREATE"
	EventGuildUpdate                         = "GUILD_UPDATE"
	EventGuildDelete                         = "GUILD_DELETE"
	EventGuildAuditLogEntryCreate            = "GUILD_AUDIT_LOG_ENTRY_CREATE"
	EventGuildBanAdd                         = "GUILD_BAN_ADD"
	EventGuildBanRemove                      = "GUILD_BAN_REMOVE"
	EventGuildEmojisUpdate                   = "GUILD_EMOJIS_UPDATE"
	EventGuildStickersUpdate                 = "GUILD_STICKERS_UPDATE"
	EventGuildIntegrationsUpdate             = "GUILD_INTEGRATIONS_UPDATE"
	EventGuildMemberAdd                      = "GUILD_MEMBER_ADD"
	EventGuildMemberRemove                   = "GUILD_MEMBER_REMOVE"
	EventGuildMemberUpdate                   = "GUILD_MEMBER_UPDATE"
	EventGuildMembersChunk                   = "GUILD_MEMBERS_CHUNK"
	EventGuildRoleCreate                     = "GUILD_ROLE_CREATE"
	EventGuildRoleUpdate                     = "GUILD_ROLE_UPDATE"
	EventGuildRoleDelete                     = "GUILD_ROLE_DELETE"
	EventGuildScheduledEventCreate           = "GUILD_SCHEDULED_EVENT_CREATE"
	EventGuildScheduledEventUpdate           = "GUILD_SCHEDULED_EVENT_UPDATE"
	EventGuildScheduledEventDelete           = "GUILD_SCHEDULED_EVENT_DELETE"
	EventGuildScheduledEventUserAdd          = "GUILD_SCHEDULED_EVENT_USER_ADD"
	EventGuildScheduledEventUserRemove       = "GUILD_SCHEDULED_EVENT_USER_REMOVE"
	EventIntegrationCreate                   = "INTEGRATION_CREATE"
	EventIntegrationUpdate                   = "INTEGRATION_UPDATE"
	EventIntegrationDelete                   = "INTEGRATION_DELETE"
	EventInteractionCreate                   = "INTERACTION_CREATE"
	EventInviteCreate                        = "INVITE_CREATE"
	EventInviteDelete                        = "INVITE_DELETE"
	EventMessageCreate                       = "MESSAGE_CREATE"
	EventMessageUpdate                       = "MESSAGE_UPDATE"
	EventMessageDelete                       = "MESSAGE_DELETE"
	EventMessageDeleteBulk                   = "MESSAGE_DELETE_BULK"
	EventMessageReactionAdd                  = "MESSAGE_REACTION_ADD"
	EventMessageReactionRemove               = "MESSAGE_REACTION_REMOVE"
	EventMessageReactionRemoveAll            = "MESSAGE_REACTION_REMOVE_ALL"
	EventMessageReactionRemoveEmoji          = "MESSAGE_REACTION_REMOVE_EMOJI"
	EventPresenceUpdate                      = "PRESENCE_UPDATE"
	EventStageInstanceCreate                 = "STAGE_INSTANCE_CREATE"
	EventStageInstanceUpdate                 = "STAGE_INSTANCE_UPDATE"
	EventStageInstanceDelete                 = "STAGE_INSTANCE_DELETE"
	EventTypingStart                         = "TYPING_START"
	EventUserUpdate                          = "USER_UPDATE"
	EventVoiceStateUpdate                    = "VOICE_STATE_UPDATE"
	EventVoiceServerUpdate                   = "VOICE_SERVER_UPDATE"
	EventWebhooksUpdate                      = "WEBHOOKS_UPDATE"
	EventMessagePollVoteAdd                  = "MESSAGE_POLL_VOTE_ADD"
	EventMessagePollVoteRemove               = "MESSAGE_POLL_VOTE_REMOVE"
)

// ReadyEvent represents the READY gateway event.
type ReadyEvent struct {
	// V is the gateway version.
	V int `json:"v"`
	// User is the current user.
	User *types.User `json:"user"`
	// Guilds are the unavailable guilds.
	Guilds []UnavailableGuild `json:"guilds"`
	// SessionID is the session ID.
	SessionID string `json:"session_id"`
	// ResumeGatewayURL is the URL to use for resuming.
	ResumeGatewayURL string `json:"resume_gateway_url"`
	// Shard is the shard information.
	Shard *[2]int `json:"shard,omitempty"`
	// Application contains partial application information.
	Application *PartialApplication `json:"application"`
}

// UnavailableGuild represents an unavailable guild in READY.
type UnavailableGuild struct {
	ID          types.Snowflake `json:"id"`
	Unavailable bool            `json:"unavailable"`
}

// PartialApplication represents partial application data in READY.
type PartialApplication struct {
	ID    types.Snowflake `json:"id"`
	Flags int             `json:"flags"`
}

// MessageCreateEvent represents the MESSAGE_CREATE gateway event.
type MessageCreateEvent struct {
	types.Message
	// GuildID is the guild ID (when in a guild).
	GuildID *types.Snowflake `json:"guild_id,omitempty"`
	// Member is the guild member who sent the message.
	Member *types.GuildMember `json:"member,omitempty"`
}

// InteractionCreateEvent represents the INTERACTION_CREATE gateway event.
type InteractionCreateEvent struct {
	types.Interaction
}

// TypingStartEvent represents the TYPING_START gateway event.
type TypingStartEvent struct {
	// ChannelID is the channel ID.
	ChannelID types.Snowflake `json:"channel_id"`
	// GuildID is the guild ID.
	GuildID *types.Snowflake `json:"guild_id,omitempty"`
	// UserID is the user ID.
	UserID types.Snowflake `json:"user_id"`
	// Timestamp is the Unix timestamp.
	Timestamp int64 `json:"timestamp"`
	// Member is the guild member.
	Member *types.GuildMember `json:"member,omitempty"`
}

// MessageReactionAddEvent represents the MESSAGE_REACTION_ADD gateway event.
type MessageReactionAddEvent struct {
	// UserID is the user who reacted.
	UserID types.Snowflake `json:"user_id"`
	// ChannelID is the channel ID.
	ChannelID types.Snowflake `json:"channel_id"`
	// MessageID is the message ID.
	MessageID types.Snowflake `json:"message_id"`
	// GuildID is the guild ID.
	GuildID *types.Snowflake `json:"guild_id,omitempty"`
	// Member is the guild member.
	Member *types.GuildMember `json:"member,omitempty"`
	// Emoji is the emoji used.
	Emoji types.PartialEmoji `json:"emoji"`
	// MessageAuthorID is the author of the message.
	MessageAuthorID *types.Snowflake `json:"message_author_id,omitempty"`
	// Burst indicates a super reaction.
	Burst bool `json:"burst"`
	// BurstColors are the colors for super reactions.
	BurstColors []string `json:"burst_colors,omitempty"`
	// Type is the reaction type.
	Type int `json:"type"`
}

// ParseEvent parses raw event data into a typed event struct.
func ParseEvent(eventType string, data []byte) (interface{}, error) {
	var event interface{}

	switch eventType {
	case EventReady:
		event = &ReadyEvent{}
	case EventMessageCreate:
		event = &MessageCreateEvent{}
	case EventInteractionCreate:
		event = &InteractionCreateEvent{}
	case EventTypingStart:
		event = &TypingStartEvent{}
	case EventMessageReactionAdd:
		event = &MessageReactionAddEvent{}
	case EventChannelCreate, EventChannelUpdate, EventChannelDelete:
		event = &types.Channel{}
	case EventGuildMemberAdd, EventGuildMemberUpdate:
		event = &types.GuildMember{}
	case EventUserUpdate:
		event = &types.User{}
	case EventMessageUpdate:
		// Skip JSON unmarshaling for MessageUpdate to avoid recursion
		// Discord echoes back messages with components which triggers MarshalJSON
		event = &types.Message{}
	default:
		// Return raw data for unknown events
		var raw map[string]interface{}
		if err := json.Unmarshal(data, &raw); err != nil {
			return nil, err
		}
		return raw, nil
	}

	if eventType != EventMessageUpdate {
		if err := json.Unmarshal(data, event); err != nil {
			return nil, err
		}
	}

	return event, nil
}
