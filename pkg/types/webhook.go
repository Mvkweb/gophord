// Package types provides Discord API type definitions.
package types

// WebhookType represents the type of a webhook.
type WebhookType int

const (
	// WebhookTypeIncoming is an incoming webhook (type 1).
	WebhookTypeIncoming WebhookType = 1
	// WebhookTypeChannelFollower is a channel follower webhook (type 2).
	WebhookTypeChannelFollower WebhookType = 2
	// WebhookTypeApplication is an application webhook (type 3).
	WebhookTypeApplication WebhookType = 3
)

// Webhook represents a Discord webhook.
type Webhook struct {
	// ID is the unique ID of the webhook.
	ID Snowflake `json:"id"`
	// Type is the type of the webhook.
	Type WebhookType `json:"type"`
	// GuildID is the guild ID this webhook belongs to.
	GuildID *Snowflake `json:"guild_id,omitempty"`
	// ChannelID is the channel ID this webhook belongs to.
	ChannelID *Snowflake `json:"channel_id,omitempty"`
	// User is the user that created the webhook.
	User *User `json:"user,omitempty"`
	// Name is the default name of the webhook.
	Name string `json:"name,omitempty"`
	// Avatar is the default avatar hash of the webhook.
	Avatar string `json:"avatar,omitempty"`
	// Token is the secure token of the webhook (returned for Incoming Webhooks).
	Token string `json:"token,omitempty"`
	// ApplicationID is the bot/application that created this webhook.
	ApplicationID *Snowflake `json:"application_id,omitempty"`
	// SourceGuild is the partial guild of the followed channel (follower webhooks).
	SourceGuild *Guild `json:"source_guild,omitempty"`
	// SourceChannel is the partial channel of the followed channel (follower webhooks).
	SourceChannel *Channel `json:"source_channel,omitempty"`
	// URL is the url used for executing the webhook.
	URL string `json:"url,omitempty"`
}

// CreateWebhookParams contains parameters for creating a webhook.
type CreateWebhookParams struct {
	// Name is the name of the webhook (1-80 chars).
	Name string `json:"name"`
	// Avatar is the image for the default webhook avatar.
	Avatar string `json:"avatar,omitempty"`
}

// ExecuteWebhookParams contains parameters for executing a webhook.
type ExecuteWebhookParams struct {
	// Content is the message contents (up to 2000 characters).
	Content string `json:"content,omitempty"`
	// Username overrides the default username of the webhook.
	Username string `json:"username,omitempty"`
	// AvatarURL overrides the default avatar of the webhook.
	AvatarURL string `json:"avatar_url,omitempty"`
	// TTS indicates whether this is a TTS message.
	TTS bool `json:"tts,omitempty"`
	// Embeds are rich embeds (up to 10).
	Embeds []Embed `json:"embeds,omitempty"`
	// AllowedMentions controls mention behavior.
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	// Components are message components.
	Components ComponentList `json:"components,omitempty"`
	// Files are the contents of the file being sent.
	// Files []File `json:"-"` // Handled separately in multipart/form-data
	// PayloadJSON is the JSON encoded body of the request (for multipart).
	PayloadJSON string `json:"payload_json,omitempty"`
	// Attachments are attachment objects with filename and description.
	Attachments []Attachment `json:"attachments,omitempty"`
	// Flags are message flags (only SUPPRESS_EMBEDS can be set).
	Flags MessageFlags `json:"flags,omitempty"`
	// ThreadName is the name of the thread to create (for forum channels).
	ThreadName string `json:"thread_name,omitempty"`
}
