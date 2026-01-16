// Package types provides Discord API type definitions.
package types

// Guild represents a Discord guild (server).
type Guild struct {
	// ID is the guild ID.
	ID Snowflake `json:"id"`
	// Name is the guild name (2-100 characters).
	Name string `json:"name"`
	// Icon is the icon hash.
	Icon string `json:"icon,omitempty"`
	// IconHash is the icon hash, returned when in the template object.
	IconHash string `json:"icon_hash,omitempty"`
	// Splash is the splash hash.
	Splash string `json:"splash,omitempty"`
	// DiscoverySplash is the discovery splash hash.
	DiscoverySplash string `json:"discovery_splash,omitempty"`
	// Owner indicates whether the current user is the owner.
	Owner bool `json:"owner,omitempty"`
	// OwnerID is the user ID of the owner.
	OwnerID Snowflake `json:"owner_id"`
	// Permissions is the total permissions for the user in the guild (excludes overwrites).
	Permissions string `json:"permissions,omitempty"`
	// Region is the voice region (deprecated).
	Region string `json:"region,omitempty"`
	// AFKChannelID is the ID of the AFK channel.
	AFKChannelID *Snowflake `json:"afk_channel_id,omitempty"`
	// AFKTimeout is the AFK timeout in seconds.
	AFKTimeout int `json:"afk_timeout"`
	// WidgetEnabled indicates whether the server widget is enabled.
	WidgetEnabled bool `json:"widget_enabled,omitempty"`
	// WidgetChannelID is the channel ID for the server widget.
	WidgetChannelID *Snowflake `json:"widget_channel_id,omitempty"`
	// VerificationLevel is the verification level required for the guild.
	VerificationLevel int `json:"verification_level"`
	// DefaultMessageNotifications is the default message notification level.
	DefaultMessageNotifications int `json:"default_message_notifications"`
	// ExplicitContentFilter is the explicit content filter level.
	ExplicitContentFilter int `json:"explicit_content_filter"`
	// Roles are the roles in the guild.
	Roles []Role `json:"roles"`
	// Emojis are the custom emojis in the guild.
	Emojis []Emoji `json:"emojis"`
	// Features are the enabled guild features.
	Features []string `json:"features"`
	// MFALevel is the required MFA level for the guild.
	MFALevel int `json:"mfa_level"`
	// ApplicationID is the application ID of the guild creator if it is bot-created.
	ApplicationID *Snowflake `json:"application_id,omitempty"`
	// SystemChannelID is the ID of the system channel.
	SystemChannelID *Snowflake `json:"system_channel_id,omitempty"`
	// SystemChannelFlags are the system channel flags.
	SystemChannelFlags int `json:"system_channel_flags"`
	// RulesChannelID is the ID of the rules channel (community guilds).
	RulesChannelID *Snowflake `json:"rules_channel_id,omitempty"`
	// MaxPresences is the maximum number of presences (null = default 25000).
	MaxPresences *int `json:"max_presences,omitempty"`
	// MaxMembers is the maximum number of members.
	MaxMembers int `json:"max_members,omitempty"`
	// VanityURLCode is the vanity URL code.
	VanityURLCode string `json:"vanity_url_code,omitempty"`
	// Description is the guild description (community guilds).
	Description string `json:"description,omitempty"`
	// Banner is the banner hash.
	Banner string `json:"banner,omitempty"`
	// PremiumTier is the server boost level.
	PremiumTier int `json:"premium_tier"`
	// PremiumSubscriptionCount is the number of boosts.
	PremiumSubscriptionCount int `json:"premium_subscription_count,omitempty"`
	// PreferredLocale is the preferred locale of the guild.
	PreferredLocale string `json:"preferred_locale"`
	// PublicUpdatesChannelID is the ID of the channel where admin notices are sent.
	PublicUpdatesChannelID *Snowflake `json:"public_updates_channel_id,omitempty"`
	// MaxVideoChannelUsers is the maximum number of users in a video channel.
	MaxVideoChannelUsers int `json:"max_video_channel_users,omitempty"`
	// MaxStageVideoChannelUsers is the max users in stage video.
	MaxStageVideoChannelUsers int `json:"max_stage_video_channel_users,omitempty"`
	// ApproximateMemberCount is the approximate number of members.
	ApproximateMemberCount int `json:"approximate_member_count,omitempty"`
	// ApproximatePresenceCount is the approximate number of non-offline members.
	ApproximatePresenceCount int `json:"approximate_presence_count,omitempty"`
	// WelcomeScreen is the welcome screen of the guild.
	WelcomeScreen *WelcomeScreen `json:"welcome_screen,omitempty"`
	// NSFWLevel is the guild NSFW level.
	NSFWLevel int `json:"nsfw_level"`
	// Stickers are the custom stickers in the guild.
	Stickers []Sticker `json:"stickers,omitempty"`
	// PremiumProgressBarEnabled indicates whether the boost progress bar is enabled.
	PremiumProgressBarEnabled bool `json:"premium_progress_bar_enabled"`
	// SafetyAlertsChannelID is the ID of the safety alerts channel.
	SafetyAlertsChannelID *Snowflake `json:"safety_alerts_channel_id,omitempty"`
}

// WelcomeScreen represents a guild welcome screen.
type WelcomeScreen struct {
	// Description is the server description shown in the welcome screen.
	Description string `json:"description,omitempty"`
	// WelcomeChannels are the channels shown in the welcome screen.
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"`
}

// WelcomeScreenChannel represents a channel in the welcome screen.
type WelcomeScreenChannel struct {
	// ChannelID is the channel ID.
	ChannelID Snowflake `json:"channel_id"`
	// Description is the description shown for the channel.
	Description string `json:"description"`
	// EmojiID is the emoji ID if specific emoji is used.
	EmojiID *Snowflake `json:"emoji_id,omitempty"`
	// EmojiName is the emoji name if unicode emoji is used.
	EmojiName string `json:"emoji_name,omitempty"`
}

// Role represents a Discord role.
type Role struct {
	// ID is the role ID.
	ID Snowflake `json:"id"`
	// Name is the role name.
	Name string `json:"name"`
	// Color is the role color as an integer.
	Color int `json:"color"`
	// Hoist indicates whether the role is hoisted.
	Hoist bool `json:"hoist"`
	// Icon is the role icon hash.
	Icon string `json:"icon,omitempty"`
	// UnicodeEmoji is the unicode emoji for the role.
	UnicodeEmoji string `json:"unicode_emoji,omitempty"`
	// Position is the role position.
	Position int `json:"position"`
	// Permissions is the permission bitfield.
	Permissions string `json:"permissions"`
	// Managed indicates whether the role is managed by an integration.
	Managed bool `json:"managed"`
	// Mentionable indicates whether the role is mentionable.
	Mentionable bool `json:"mentionable"`
	// Tags are the role tags.
	Tags *RoleTags `json:"tags,omitempty"`
	// Flags are the role flags.
	Flags int `json:"flags"`
}

// RoleTags contains role tag information.
type RoleTags struct {
	// BotID is the bot ID for bot roles.
	BotID *Snowflake `json:"bot_id,omitempty"`
	// IntegrationID is the integration ID for managed roles.
	IntegrationID *Snowflake `json:"integration_id,omitempty"`
	// PremiumSubscriber indicates a booster role (null if true).
	PremiumSubscriber interface{} `json:"premium_subscriber,omitempty"`
	// SubscriptionListingID is the subscription listing ID.
	SubscriptionListingID *Snowflake `json:"subscription_listing_id,omitempty"`
	// AvailableForPurchase indicates purchasable role (null if true).
	AvailableForPurchase interface{} `json:"available_for_purchase,omitempty"`
	// GuildConnections indicates guild connections role (null if true).
	GuildConnections interface{} `json:"guild_connections,omitempty"`
}

// GuildPreview represents a guild preview.
type GuildPreview struct {
	// ID is the guild ID.
	ID Snowflake `json:"id"`
	// Name is the guild name.
	Name string `json:"name"`
	// Icon is the icon hash.
	Icon string `json:"icon,omitempty"`
	// Splash is the splash hash.
	Splash string `json:"splash,omitempty"`
	// DiscoverySplash is the discovery splash hash.
	DiscoverySplash string `json:"discovery_splash,omitempty"`
	// Emojis are the custom emojis.
	Emojis []Emoji `json:"emojis"`
	// Features are the enabled guild features.
	Features []string `json:"features"`
	// ApproximateMemberCount is the approximate number of members.
	ApproximateMemberCount int `json:"approximate_member_count"`
	// ApproximatePresenceCount is the approximate number of non-offline members.
	ApproximatePresenceCount int `json:"approximate_presence_count"`
	// Description is the guild description.
	Description string `json:"description,omitempty"`
	// Stickers are the custom stickers.
	Stickers []Sticker `json:"stickers"`
}

// Ban represents a guild ban.
type Ban struct {
	// Reason is the reason for the ban.
	Reason string `json:"reason,omitempty"`
	// User is the banned user.
	User User `json:"user"`
}
