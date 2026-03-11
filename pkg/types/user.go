// Package types provides Discord API type definitions.
package types

// User represents a Discord user.
type User struct {
	// ID is the user's snowflake ID.
	ID Snowflake `json:"id"`
	// Username is the user's username (not unique).
	Username string `json:"username"`
	// Discriminator is the user's 4-digit tag (deprecated, use display name).
	Discriminator string `json:"discriminator"`
	// GlobalName is the user's display name (if set).
	GlobalName string `json:"global_name,omitempty"`
	// Avatar is the user's avatar hash.
	Avatar string `json:"avatar,omitempty"`
	// Bot indicates whether the user is a bot.
	Bot bool `json:"bot,omitempty"`
	// System indicates whether the user is an official Discord system user.
	System bool `json:"system,omitempty"`
	// MFAEnabled indicates whether 2FA is enabled.
	MFAEnabled bool `json:"mfa_enabled,omitempty"`
	// Banner is the user's banner hash.
	Banner string `json:"banner,omitempty"`
	// AccentColor is the user's banner color as an integer.
	AccentColor *int `json:"accent_color,omitempty"`
	// Locale is the user's chosen language.
	Locale string `json:"locale,omitempty"`
	// Verified indicates whether the user's email is verified.
	Verified bool `json:"verified,omitempty"`
	// Email is the user's email address.
	Email string `json:"email,omitempty"`
	// Flags are the user's flags.
	Flags int `json:"flags,omitempty"`
	// PremiumType is the type of Nitro subscription.
	PremiumType int `json:"premium_type,omitempty"`
	// PublicFlags are the user's public flags.
	PublicFlags int `json:"public_flags,omitempty"`
	// AvatarDecoration is the user's avatar decoration data.
	AvatarDecoration *AvatarDecoration `json:"avatar_decoration_data,omitempty"`
}

// AvatarDecoration represents avatar decoration data.
type AvatarDecoration struct {
	// Asset is the decoration asset hash.
	Asset string `json:"asset"`
	// SKUID is the SKU ID of the decoration.
	SKUID Snowflake `json:"sku_id"`
}

// GuildMember represents a member of a guild.
type GuildMember struct {
	// User is the user object.
	User *User `json:"user,omitempty"`
	// Nick is the user's guild nickname.
	Nick string `json:"nick,omitempty"`
	// Avatar is the guild-specific avatar hash.
	Avatar string `json:"avatar,omitempty"`
	// Roles contains the role IDs assigned to the member.
	Roles []Snowflake `json:"roles"`
	// JoinedAt is when the user joined the guild.
	JoinedAt Timestamp `json:"joined_at"`
	// PremiumSince is when the user started boosting.
	PremiumSince *Timestamp `json:"premium_since,omitempty"`
	// Deaf indicates whether the user is deafened in voice.
	Deaf bool `json:"deaf"`
	// Mute indicates whether the user is muted in voice.
	Mute bool `json:"mute"`
	// Flags are the guild member flags.
	Flags int `json:"flags"`
	// Pending indicates whether the user has passed membership screening.
	Pending bool `json:"pending,omitempty"`
	// Permissions is the total permissions (when in interaction).
	Permissions string `json:"permissions,omitempty"`
	// CommunicationDisabledUntil is the timeout expiry timestamp.
	CommunicationDisabledUntil *Timestamp `json:"communication_disabled_until,omitempty"`
	// AvatarDecorationData is the guild-specific avatar decoration.
	AvatarDecorationData *AvatarDecoration `json:"avatar_decoration_data,omitempty"`
}

// PartialEmoji represents a partial emoji used in components.
type PartialEmoji struct {
	// ID is the emoji ID (nil for unicode emoji).
	ID *Snowflake `json:"id,omitempty"`
	// Name is the emoji name (unicode character or custom emoji name).
	Name string `json:"name"`
	// Animated indicates whether the emoji is animated.
	Animated bool `json:"animated,omitempty"`
}

// Emoji represents a full emoji object.
type Emoji struct {
	// ID is the emoji ID.
	ID *Snowflake `json:"id"`
	// Name is the emoji name.
	Name string `json:"name"`
	// Roles are the roles allowed to use this emoji.
	Roles []Snowflake `json:"roles,omitempty"`
	// User is the user who created this emoji.
	User *User `json:"user,omitempty"`
	// RequireColons indicates whether the emoji must be wrapped in colons.
	RequireColons bool `json:"require_colons,omitempty"`
	// Managed indicates whether the emoji is managed by an integration.
	Managed bool `json:"managed,omitempty"`
	// Animated indicates whether the emoji is animated.
	Animated bool `json:"animated,omitempty"`
	// Available indicates whether the emoji can be used.
	Available bool `json:"available,omitempty"`
}

// ModifyCurrentUserParams contains parameters for modifying the current user (bot).
type ModifyCurrentUserParams struct {
	// Username is the user's username (2-32 characters, 1-32 if bots).
	Username string `json:"username,omitempty"`
	// Avatar is the user's avatar (base64 encoded image, max 128x128).
	Avatar string `json:"avatar,omitempty"`
	// Bio is the user's "about me" field (max 190 characters, bots max 190).
	Bio string `json:"bio,omitempty"`
}
