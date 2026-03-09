// Package types provides Discord API type definitions.
package types

// ChannelType represents the type of a Discord channel.
type ChannelType int

const (
	// ChannelTypeGuildText is a text channel within a server (0).
	ChannelTypeGuildText ChannelType = 0
	// ChannelTypeDM is a direct message between users (1).
	ChannelTypeDM ChannelType = 1
	// ChannelTypeGuildVoice is a voice channel within a server (2).
	ChannelTypeGuildVoice ChannelType = 2
	// ChannelTypeGroupDM is a direct message between multiple users (3).
	ChannelTypeGroupDM ChannelType = 3
	// ChannelTypeGuildCategory is an organizational category (4).
	ChannelTypeGuildCategory ChannelType = 4
	// ChannelTypeGuildAnnouncement is a channel users can follow (5).
	ChannelTypeGuildAnnouncement ChannelType = 5
	// ChannelTypeAnnouncementThread is a thread in an announcement channel (10).
	ChannelTypeAnnouncementThread ChannelType = 10
	// ChannelTypePublicThread is a public thread (11).
	ChannelTypePublicThread ChannelType = 11
	// ChannelTypePrivateThread is a private thread (12).
	ChannelTypePrivateThread ChannelType = 12
	// ChannelTypeGuildStageVoice is a voice channel for events (13).
	ChannelTypeGuildStageVoice ChannelType = 13
	// ChannelTypeGuildDirectory is a hub channel listing (14).
	ChannelTypeGuildDirectory ChannelType = 14
	// ChannelTypeGuildForum is a forum channel (15).
	ChannelTypeGuildForum ChannelType = 15
	// ChannelTypeGuildMedia is a media channel (16).
	ChannelTypeGuildMedia ChannelType = 16
)

// String returns the string representation of a ChannelType.
func (c ChannelType) String() string {
	switch c {
	case ChannelTypeGuildText:
		return "GuildText"
	case ChannelTypeDM:
		return "DM"
	case ChannelTypeGuildVoice:
		return "GuildVoice"
	case ChannelTypeGroupDM:
		return "GroupDM"
	case ChannelTypeGuildCategory:
		return "GuildCategory"
	case ChannelTypeGuildAnnouncement:
		return "GuildAnnouncement"
	case ChannelTypeAnnouncementThread:
		return "AnnouncementThread"
	case ChannelTypePublicThread:
		return "PublicThread"
	case ChannelTypePrivateThread:
		return "PrivateThread"
	case ChannelTypeGuildStageVoice:
		return "GuildStageVoice"
	case ChannelTypeGuildDirectory:
		return "GuildDirectory"
	case ChannelTypeGuildForum:
		return "GuildForum"
	case ChannelTypeGuildMedia:
		return "GuildMedia"
	default:
		return "Unknown"
	}
}

// Channel represents a Discord channel.
type Channel struct {
	// ID is the channel's unique snowflake ID.
	ID Snowflake `json:"id"`
	// Type is the type of channel.
	Type ChannelType `json:"type"`
	// GuildID is the ID of the guild (may be missing for DMs).
	GuildID *Snowflake `json:"guild_id,omitempty"`
	// Position is the sorting position of the channel.
	Position int `json:"position,omitempty"`
	// PermissionOverwrites contains explicit permission overwrites.
	PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	// Name is the channel name (1-100 characters).
	Name string `json:"name,omitempty"`
	// Topic is the channel topic (0-4096 characters for forum/media, 0-1024 for others).
	Topic string `json:"topic,omitempty"`
	// NSFW indicates whether the channel is NSFW.
	NSFW bool `json:"nsfw,omitempty"`
	// LastMessageID is the ID of the last message (or thread for forum/media).
	LastMessageID *Snowflake `json:"last_message_id,omitempty"`
	// Bitrate is the voice channel bitrate in bits.
	Bitrate int `json:"bitrate,omitempty"`
	// UserLimit is the voice channel user limit (0 = no limit).
	UserLimit int `json:"user_limit,omitempty"`
	// RateLimitPerUser is the slowmode delay in seconds (0-21600).
	RateLimitPerUser int `json:"rate_limit_per_user,omitempty"`
	// Recipients contains the DM recipients.
	Recipients []User `json:"recipients,omitempty"`
	// Icon is the icon hash for group DMs.
	Icon string `json:"icon,omitempty"`
	// OwnerID is the creator ID for group DMs or threads.
	OwnerID *Snowflake `json:"owner_id,omitempty"`
	// ApplicationID is the application ID for bot-created group DMs.
	ApplicationID *Snowflake `json:"application_id,omitempty"`
	// Managed indicates whether the group DM is managed by an application.
	Managed bool `json:"managed,omitempty"`
	// ParentID is the parent category or channel ID.
	ParentID *Snowflake `json:"parent_id,omitempty"`
	// LastPinTimestamp is when the last pin occurred.
	LastPinTimestamp *Timestamp `json:"last_pin_timestamp,omitempty"`
	// RTCRegion is the voice region ID.
	RTCRegion string `json:"rtc_region,omitempty"`
	// VideoQualityMode is the camera quality mode (1=auto, 2=full).
	VideoQualityMode int `json:"video_quality_mode,omitempty"`
	// MessageCount is the approximate message count in threads.
	MessageCount int `json:"message_count,omitempty"`
	// MemberCount is the approximate member count in threads.
	MemberCount int `json:"member_count,omitempty"`
	// ThreadMetadata contains thread-specific fields.
	ThreadMetadata *ThreadMetadata `json:"thread_metadata,omitempty"`
	// Member is the thread member object for the current user.
	Member *ThreadMember `json:"member,omitempty"`
	// DefaultAutoArchiveDuration is the default archive duration for threads.
	DefaultAutoArchiveDuration int `json:"default_auto_archive_duration,omitempty"`
	// Permissions is the computed permissions for the invoking user.
	Permissions string `json:"permissions,omitempty"`
	// Flags are the channel flags combined as a bitfield.
	Flags int `json:"flags,omitempty"`
	// TotalMessageSent is the total messages ever sent in threads.
	TotalMessageSent int `json:"total_message_sent,omitempty"`
	// AvailableTags are the available tags for forum/media channels.
	AvailableTags []ForumTag `json:"available_tags,omitempty"`
	// AppliedTags are the IDs of tags applied to forum/media threads.
	AppliedTags []Snowflake `json:"applied_tags,omitempty"`
	// DefaultReactionEmoji is the default reaction emoji for forum posts.
	DefaultReactionEmoji *DefaultReaction `json:"default_reaction_emoji,omitempty"`
	// DefaultThreadRateLimitPerUser is the default slowmode for new threads.
	DefaultThreadRateLimitPerUser int `json:"default_thread_rate_limit_per_user,omitempty"`
	// DefaultSortOrder is the default sort order for forum posts.
	DefaultSortOrder *int `json:"default_sort_order,omitempty"`
	// DefaultForumLayout is the default layout for forum channels.
	DefaultForumLayout int `json:"default_forum_layout,omitempty"`
}

// CreateChannelParams contains parameters for creating a channel.
type CreateChannelParams struct {
	Name                 string                `json:"name"`
	Type                 ChannelType           `json:"type,omitempty"`
	Topic                string                `json:"topic,omitempty"`
	Bitrate              int                   `json:"bitrate,omitempty"`
	UserLimit            int                   `json:"user_limit,omitempty"`
	RateLimitPerUser     int                   `json:"rate_limit_per_user,omitempty"`
	Position             int                   `json:"position,omitempty"`
	PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	ParentID             Snowflake             `json:"parent_id,omitempty"`
	NSFW                 bool                  `json:"nsfw,omitempty"`
}

// PermissionOverwrite represents a channel permission overwrite.
type PermissionOverwrite struct {
	// ID is the role or user ID.
	ID Snowflake `json:"id"`
	// Type is the overwrite type (0=role, 1=member).
	Type int `json:"type"`
	// Allow is the bitwise value of allowed permissions.
	Allow string `json:"allow"`
	// Deny is the bitwise value of denied permissions.
	Deny string `json:"deny"`
}

// ThreadMetadata contains thread-specific fields.
type ThreadMetadata struct {
	// Archived indicates whether the thread is archived.
	Archived bool `json:"archived"`
	// AutoArchiveDuration is the archive duration in minutes.
	AutoArchiveDuration int `json:"auto_archive_duration"`
	// ArchiveTimestamp is when the thread's archive status last changed.
	ArchiveTimestamp Timestamp `json:"archive_timestamp"`
	// Locked indicates whether the thread is locked.
	Locked bool `json:"locked"`
	// Invitable indicates whether non-moderators can add users (private threads).
	Invitable bool `json:"invitable,omitempty"`
	// CreateTimestamp is when the thread was created.
	CreateTimestamp *Timestamp `json:"create_timestamp,omitempty"`
}

// ThreadMember represents a user's thread membership.
type ThreadMember struct {
	// ID is the thread ID.
	ID *Snowflake `json:"id,omitempty"`
	// UserID is the user ID.
	UserID *Snowflake `json:"user_id,omitempty"`
	// JoinTimestamp is when the user joined the thread.
	JoinTimestamp Timestamp `json:"join_timestamp"`
	// Flags are user-specific thread settings.
	Flags int `json:"flags"`
	// Member is the guild member object.
	Member *GuildMember `json:"member,omitempty"`
}

// ForumTag represents a tag available in forum/media channels.
type ForumTag struct {
	// ID is the tag's snowflake ID.
	ID Snowflake `json:"id"`
	// Name is the tag name (0-20 characters).
	Name string `json:"name"`
	// Moderated indicates if only moderators can apply this tag.
	Moderated bool `json:"moderated"`
	// EmojiID is the emoji ID (if custom emoji).
	EmojiID *Snowflake `json:"emoji_id"`
	// EmojiName is the emoji name (if unicode emoji).
	EmojiName string `json:"emoji_name"`
}

// DefaultReaction represents the default reaction emoji for forum posts.
type DefaultReaction struct {
	// EmojiID is the emoji ID (if custom emoji).
	EmojiID *Snowflake `json:"emoji_id"`
	// EmojiName is the emoji name (if unicode emoji).
	EmojiName string `json:"emoji_name"`
}
