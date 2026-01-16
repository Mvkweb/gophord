// Package types provides Discord API type definitions.
package types

// MessageFlags represents bitfield flags for messages.
type MessageFlags int

const (
	// MessageFlagCrossposted indicates the message has been published.
	MessageFlagCrossposted MessageFlags = 1 << 0
	// MessageFlagIsCrosspost indicates the message originated from another channel.
	MessageFlagIsCrosspost MessageFlags = 1 << 1
	// MessageFlagSuppressEmbeds suppresses embeds for the message.
	MessageFlagSuppressEmbeds MessageFlags = 1 << 2
	// MessageFlagSourceMessageDeleted indicates the source message was deleted.
	MessageFlagSourceMessageDeleted MessageFlags = 1 << 3
	// MessageFlagUrgent indicates an urgent message from Discord.
	MessageFlagUrgent MessageFlags = 1 << 4
	// MessageFlagHasThread indicates the message has an associated thread.
	MessageFlagHasThread MessageFlags = 1 << 5
	// MessageFlagEphemeral indicates the message is only visible to the invoker.
	MessageFlagEphemeral MessageFlags = 1 << 6
	// MessageFlagLoading indicates the message is a deferred interaction response.
	MessageFlagLoading MessageFlags = 1 << 7
	// MessageFlagFailedToMentionRoles indicates role mention failure.
	MessageFlagFailedToMentionRoles MessageFlags = 1 << 8
	// MessageFlagSuppressNotifications indicates no notifications for the message.
	MessageFlagSuppressNotifications MessageFlags = 1 << 12
	// MessageFlagIsVoiceMessage indicates a voice message.
	MessageFlagIsVoiceMessage MessageFlags = 1 << 13
	// MessageFlagIsComponentsV2 indicates Components V2 is enabled (32768).
	MessageFlagIsComponentsV2 MessageFlags = 1 << 15
)

// MessageType represents the type of a message.
type MessageType int

const (
	// MessageTypeDefault is a regular message.
	MessageTypeDefault MessageType = 0
	// MessageTypeRecipientAdd indicates a recipient was added to a group DM.
	MessageTypeRecipientAdd MessageType = 1
	// MessageTypeRecipientRemove indicates a recipient was removed from a group DM.
	MessageTypeRecipientRemove MessageType = 2
	// MessageTypeCall indicates a call.
	MessageTypeCall MessageType = 3
	// MessageTypeChannelNameChange indicates a channel name change.
	MessageTypeChannelNameChange MessageType = 4
	// MessageTypeChannelIconChange indicates a channel icon change.
	MessageTypeChannelIconChange MessageType = 5
	// MessageTypeChannelPinnedMessage indicates a message was pinned.
	MessageTypeChannelPinnedMessage MessageType = 6
	// MessageTypeUserJoin indicates a user joined the guild.
	MessageTypeUserJoin MessageType = 7
	// MessageTypeGuildBoost indicates a user boosted the guild.
	MessageTypeGuildBoost MessageType = 8
	// MessageTypeGuildBoostTier1 indicates the guild reached boost tier 1.
	MessageTypeGuildBoostTier1 MessageType = 9
	// MessageTypeGuildBoostTier2 indicates the guild reached boost tier 2.
	MessageTypeGuildBoostTier2 MessageType = 10
	// MessageTypeGuildBoostTier3 indicates the guild reached boost tier 3.
	MessageTypeGuildBoostTier3 MessageType = 11
	// MessageTypeChannelFollowAdd indicates a channel follow was added.
	MessageTypeChannelFollowAdd MessageType = 12
	// MessageTypeGuildDiscoveryDisqualified indicates discovery disqualification.
	MessageTypeGuildDiscoveryDisqualified MessageType = 14
	// MessageTypeGuildDiscoveryRequalified indicates discovery requalification.
	MessageTypeGuildDiscoveryRequalified MessageType = 15
	// MessageTypeGuildDiscoveryGracePeriodInitialWarning is an initial discovery warning.
	MessageTypeGuildDiscoveryGracePeriodInitialWarning MessageType = 16
	// MessageTypeGuildDiscoveryGracePeriodFinalWarning is a final discovery warning.
	MessageTypeGuildDiscoveryGracePeriodFinalWarning MessageType = 17
	// MessageTypeThreadCreated indicates a thread was created.
	MessageTypeThreadCreated MessageType = 18
	// MessageTypeReply is a reply to another message.
	MessageTypeReply MessageType = 19
	// MessageTypeChatInputCommand is a slash command response.
	MessageTypeChatInputCommand MessageType = 20
	// MessageTypeThreadStarterMessage is the first message in a public thread.
	MessageTypeThreadStarterMessage MessageType = 21
	// MessageTypeGuildInviteReminder is a server invite reminder.
	MessageTypeGuildInviteReminder MessageType = 22
	// MessageTypeContextMenuCommand is a context menu command response.
	MessageTypeContextMenuCommand MessageType = 23
	// MessageTypeAutoModerationAction is an auto-moderation action.
	MessageTypeAutoModerationAction MessageType = 24
	// MessageTypeRoleSubscriptionPurchase is a role subscription purchase.
	MessageTypeRoleSubscriptionPurchase MessageType = 25
	// MessageTypeInteractionPremiumUpsell is a premium upsell.
	MessageTypeInteractionPremiumUpsell MessageType = 26
	// MessageTypeStageStart indicates a stage started.
	MessageTypeStageStart MessageType = 27
	// MessageTypeStageEnd indicates a stage ended.
	MessageTypeStageEnd MessageType = 28
	// MessageTypeStageSpeaker indicates a stage speaker update.
	MessageTypeStageSpeaker MessageType = 29
	// MessageTypeStageTopic indicates a stage topic change.
	MessageTypeStageTopic MessageType = 31
	// MessageTypeGuildApplicationPremiumSubscription is an app subscription purchase.
	MessageTypeGuildApplicationPremiumSubscription MessageType = 32
	// MessageTypePurchaseNotification is a purchase notification.
	MessageTypePurchaseNotification MessageType = 44
)

// Message represents a Discord message.
type Message struct {
	// ID is the message's snowflake ID.
	ID Snowflake `json:"id"`
	// ChannelID is the channel ID where the message was sent.
	ChannelID Snowflake `json:"channel_id"`
	// Author is the message author.
	Author *User `json:"author,omitempty"`
	// Content is the message content (max 2000 characters, 4000 for Nitro).
	Content string `json:"content,omitempty"`
	// Timestamp is when the message was sent.
	Timestamp Timestamp `json:"timestamp"`
	// EditedTimestamp is when the message was last edited.
	EditedTimestamp *Timestamp `json:"edited_timestamp,omitempty"`
	// TTS indicates whether the message is text-to-speech.
	TTS bool `json:"tts"`
	// MentionEveryone indicates whether the message mentions @everyone.
	MentionEveryone bool `json:"mention_everyone"`
	// Mentions contains the users mentioned in the message.
	Mentions []User `json:"mentions"`
	// MentionRoles contains the role IDs mentioned in the message.
	MentionRoles []Snowflake `json:"mention_roles"`
	// MentionChannels contains the channels mentioned in the message.
	MentionChannels []ChannelMention `json:"mention_channels,omitempty"`
	// Attachments contains the message attachments.
	Attachments []Attachment `json:"attachments"`
	// Embeds contains the message embeds (max 10).
	Embeds []Embed `json:"embeds"`
	// Reactions contains the message reactions.
	Reactions []Reaction `json:"reactions,omitempty"`
	// Nonce is used for message send validation.
	Nonce interface{} `json:"nonce,omitempty"`
	// Pinned indicates whether the message is pinned.
	Pinned bool `json:"pinned"`
	// WebhookID is the webhook ID if sent by a webhook.
	WebhookID *Snowflake `json:"webhook_id,omitempty"`
	// Type is the message type.
	Type MessageType `json:"type"`
	// Activity is the message activity (for Rich Presence invites).
	Activity *MessageActivity `json:"activity,omitempty"`
	// Application is the application associated with the message.
	Application *Application `json:"application,omitempty"`
	// ApplicationID is the application ID for interactions.
	ApplicationID *Snowflake `json:"application_id,omitempty"`
	// Flags are the message flags.
	Flags MessageFlags `json:"flags,omitempty"`
	// MessageReference references another message.
	MessageReference *MessageReference `json:"message_reference,omitempty"`
	// MessageSnapshots contains message snapshots for forwarded messages.
	MessageSnapshots []MessageSnapshot `json:"message_snapshots,omitempty"`
	// ReferencedMessage is the referenced message (for replies).
	ReferencedMessage *Message `json:"referenced_message,omitempty"`
	// InteractionMetadata is metadata for interaction-generated messages.
	InteractionMetadata *MessageInteractionMetadata `json:"interaction_metadata,omitempty"`
	// Interaction is the interaction that triggered the message (deprecated).
	Interaction *MessageInteraction `json:"interaction,omitempty"`
	// Thread is the thread started from this message.
	Thread *Channel `json:"thread,omitempty"`
	// Components contains the message components.
	// NOTE: We don't unmarshal components in MESSAGE events to avoid
	// recursion with MarshalJSON methods. Components are only needed for
	// sending messages via REST API.
	Components ComponentList `json:"components,omitempty"`
	// StickerItems contains the sticker items sent with the message.
	StickerItems []StickerItem `json:"sticker_items,omitempty"`
	// Stickers contains the stickers sent with the message (deprecated).
	Stickers []Sticker `json:"stickers,omitempty"`
	// Position is the approximate position in the thread.
	Position int `json:"position,omitempty"`
	// RoleSubscriptionData is for role subscription messages.
	RoleSubscriptionData *RoleSubscriptionData `json:"role_subscription_data,omitempty"`
	// Resolved contains resolved data for auto-populated selects.
	Resolved *ResolvedData `json:"resolved,omitempty"`
	// Poll is the poll attached to the message.
	Poll *Poll `json:"poll,omitempty"`
	// Call is the call info for the message.
	Call *MessageCall `json:"call,omitempty"`
}

// ChannelMention represents a mentioned channel.
type ChannelMention struct {
	// ID is the channel ID.
	ID Snowflake `json:"id"`
	// GuildID is the guild ID containing the channel.
	GuildID Snowflake `json:"guild_id"`
	// Type is the channel type.
	Type ChannelType `json:"type"`
	// Name is the channel name.
	Name string `json:"name"`
}

// Attachment represents a message attachment.
type Attachment struct {
	// ID is the attachment ID.
	ID Snowflake `json:"id"`
	// Filename is the name of the attached file.
	Filename string `json:"filename"`
	// Title is the attachment title.
	Title string `json:"title,omitempty"`
	// Description is the attachment description (max 1024 characters).
	Description string `json:"description,omitempty"`
	// ContentType is the MIME type.
	ContentType string `json:"content_type,omitempty"`
	// Size is the file size in bytes.
	Size int `json:"size"`
	// URL is the attachment source URL.
	URL string `json:"url"`
	// ProxyURL is the proxied URL.
	ProxyURL string `json:"proxy_url"`
	// Height is the image height (if image).
	Height *int `json:"height,omitempty"`
	// Width is the image width (if image).
	Width *int `json:"width,omitempty"`
	// Ephemeral indicates whether the attachment is ephemeral.
	Ephemeral bool `json:"ephemeral,omitempty"`
	// DurationSecs is the audio duration in seconds (voice messages).
	DurationSecs *float64 `json:"duration_secs,omitempty"`
	// Waveform is the base64-encoded audio waveform (voice messages).
	Waveform string `json:"waveform,omitempty"`
	// Flags are the attachment flags.
	Flags int `json:"flags,omitempty"`
}

// Embed represents a rich embed.
type Embed struct {
	// Title is the embed title.
	Title string `json:"title,omitempty"`
	// Type is the embed type (always "rich" for webhook/bot embeds).
	Type string `json:"type,omitempty"`
	// Description is the embed description.
	Description string `json:"description,omitempty"`
	// URL is the embed URL.
	URL string `json:"url,omitempty"`
	// Timestamp is the embed timestamp.
	Timestamp string `json:"timestamp,omitempty"`
	// Color is the embed color code.
	Color int `json:"color,omitempty"`
	// Footer is the embed footer.
	Footer *EmbedFooter `json:"footer,omitempty"`
	// Image is the embed image.
	Image *EmbedImage `json:"image,omitempty"`
	// Thumbnail is the embed thumbnail.
	Thumbnail *EmbedThumbnail `json:"thumbnail,omitempty"`
	// Video is the embed video.
	Video *EmbedVideo `json:"video,omitempty"`
	// Provider is the embed provider.
	Provider *EmbedProvider `json:"provider,omitempty"`
	// Author is the embed author.
	Author *EmbedAuthor `json:"author,omitempty"`
	// Fields are the embed fields (max 25).
	Fields []EmbedField `json:"fields,omitempty"`
}

// EmbedFooter represents an embed footer.
type EmbedFooter struct {
	// Text is the footer text.
	Text string `json:"text"`
	// IconURL is the footer icon URL.
	IconURL string `json:"icon_url,omitempty"`
	// ProxyIconURL is the proxied icon URL.
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// EmbedImage represents an embed image.
type EmbedImage struct {
	// URL is the image source URL.
	URL string `json:"url"`
	// ProxyURL is the proxied URL.
	ProxyURL string `json:"proxy_url,omitempty"`
	// Height is the image height.
	Height int `json:"height,omitempty"`
	// Width is the image width.
	Width int `json:"width,omitempty"`
}

// EmbedThumbnail represents an embed thumbnail.
type EmbedThumbnail struct {
	// URL is the thumbnail source URL.
	URL string `json:"url"`
	// ProxyURL is the proxied URL.
	ProxyURL string `json:"proxy_url,omitempty"`
	// Height is the thumbnail height.
	Height int `json:"height,omitempty"`
	// Width is the thumbnail width.
	Width int `json:"width,omitempty"`
}

// EmbedVideo represents an embed video.
type EmbedVideo struct {
	// URL is the video source URL.
	URL string `json:"url,omitempty"`
	// ProxyURL is the proxied URL.
	ProxyURL string `json:"proxy_url,omitempty"`
	// Height is the video height.
	Height int `json:"height,omitempty"`
	// Width is the video width.
	Width int `json:"width,omitempty"`
}

// EmbedProvider represents an embed provider.
type EmbedProvider struct {
	// Name is the provider name.
	Name string `json:"name,omitempty"`
	// URL is the provider URL.
	URL string `json:"url,omitempty"`
}

// EmbedAuthor represents an embed author.
type EmbedAuthor struct {
	// Name is the author name.
	Name string `json:"name"`
	// URL is the author URL.
	URL string `json:"url,omitempty"`
	// IconURL is the author icon URL.
	IconURL string `json:"icon_url,omitempty"`
	// ProxyIconURL is the proxied icon URL.
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// EmbedField represents an embed field.
type EmbedField struct {
	// Name is the field name.
	Name string `json:"name"`
	// Value is the field value.
	Value string `json:"value"`
	// Inline indicates whether the field is inline.
	Inline bool `json:"inline,omitempty"`
}

// Reaction represents a message reaction.
type Reaction struct {
	// Count is the number of times this emoji was reacted.
	Count int `json:"count"`
	// CountDetails contains burst and normal reaction counts.
	CountDetails ReactionCountDetails `json:"count_details"`
	// Me indicates whether the current user reacted.
	Me bool `json:"me"`
	// MeBurst indicates whether the current user super-reacted.
	MeBurst bool `json:"me_burst"`
	// Emoji is the reaction emoji.
	Emoji PartialEmoji `json:"emoji"`
	// BurstColors are the colors for super reactions.
	BurstColors []string `json:"burst_colors"`
}

// ReactionCountDetails contains reaction count breakdown.
type ReactionCountDetails struct {
	// Burst is the super reaction count.
	Burst int `json:"burst"`
	// Normal is the normal reaction count.
	Normal int `json:"normal"`
}

// MessageActivity represents a message activity.
type MessageActivity struct {
	// Type is the activity type.
	Type int `json:"type"`
	// PartyID is the party ID from Rich Presence.
	PartyID string `json:"party_id,omitempty"`
}

// Application represents a Discord application.
type Application struct {
	// ID is the application ID.
	ID Snowflake `json:"id"`
	// Name is the application name.
	Name string `json:"name"`
	// Icon is the application icon hash.
	Icon string `json:"icon,omitempty"`
	// Description is the application description.
	Description string `json:"description"`
	// CoverImage is the cover image hash.
	CoverImage string `json:"cover_image,omitempty"`
}

// MessageReference references another message.
type MessageReference struct {
	// Type is the reference type (0=default, 1=forward).
	Type int `json:"type,omitempty"`
	// MessageID is the referenced message ID.
	MessageID *Snowflake `json:"message_id,omitempty"`
	// ChannelID is the referenced channel ID.
	ChannelID *Snowflake `json:"channel_id,omitempty"`
	// GuildID is the referenced guild ID.
	GuildID *Snowflake `json:"guild_id,omitempty"`
	// FailIfNotExists indicates whether to error if reference doesn't exist.
	FailIfNotExists *bool `json:"fail_if_not_exists,omitempty"`
}

// MessageSnapshot represents a message snapshot (for forwards).
type MessageSnapshot struct {
	// Message is the partial message object.
	Message *Message `json:"message"`
}

// MessageInteractionMetadata contains interaction metadata.
type MessageInteractionMetadata struct {
	// ID is the interaction ID.
	ID Snowflake `json:"id"`
	// Type is the interaction type.
	Type InteractionType `json:"type"`
	// UserID is the ID of the user who triggered the interaction.
	UserID Snowflake `json:"user_id"`
	// User is the user who triggered the interaction.
	User *User `json:"user,omitempty"`
	// AuthorizingIntegrationOwners maps integration types to IDs.
	AuthorizingIntegrationOwners map[string]Snowflake `json:"authorizing_integration_owners"`
	// OriginalResponseMessageID is the original response message ID.
	OriginalResponseMessageID *Snowflake `json:"original_response_message_id,omitempty"`
	// TargetUser is the target user for user commands.
	TargetUser *User `json:"target_user,omitempty"`
	// TargetMessageID is the target message for message commands.
	TargetMessageID *Snowflake `json:"target_message_id,omitempty"`
	// TriggeringInteractionMetadata is the metadata of the triggering interaction.
	TriggeringInteractionMetadata *MessageInteractionMetadata `json:"triggering_interaction_metadata,omitempty"`
}

// MessageInteraction represents an interaction that triggered a message (deprecated).
type MessageInteraction struct {
	// ID is the interaction ID.
	ID Snowflake `json:"id"`
	// Type is the interaction type.
	Type InteractionType `json:"type"`
	// Name is the interaction name.
	Name string `json:"name"`
	// User is the user who invoked the interaction.
	User User `json:"user"`
	// Member is the guild member who invoked the interaction.
	Member *GuildMember `json:"member,omitempty"`
}

// StickerItem represents a sticker item.
type StickerItem struct {
	// ID is the sticker ID.
	ID Snowflake `json:"id"`
	// Name is the sticker name.
	Name string `json:"name"`
	// FormatType is the sticker format type.
	FormatType int `json:"format_type"`
}

// Sticker represents a full sticker object.
type Sticker struct {
	// ID is the sticker ID.
	ID Snowflake `json:"id"`
	// PackID is the sticker pack ID.
	PackID *Snowflake `json:"pack_id,omitempty"`
	// Name is the sticker name.
	Name string `json:"name"`
	// Description is the sticker description.
	Description string `json:"description,omitempty"`
	// Tags is the autocomplete/suggestion tags.
	Tags string `json:"tags"`
	// Type is the sticker type (1=standard, 2=guild).
	Type int `json:"type"`
	// FormatType is the sticker format type.
	FormatType int `json:"format_type"`
	// Available indicates whether the sticker can be used.
	Available bool `json:"available,omitempty"`
	// GuildID is the guild ID for guild stickers.
	GuildID *Snowflake `json:"guild_id,omitempty"`
	// User is the user who uploaded the sticker.
	User *User `json:"user,omitempty"`
	// SortValue is the standard sticker sort order.
	SortValue int `json:"sort_value,omitempty"`
}

// RoleSubscriptionData contains role subscription purchase data.
type RoleSubscriptionData struct {
	// RoleSubscriptionListingID is the subscription listing ID.
	RoleSubscriptionListingID Snowflake `json:"role_subscription_listing_id"`
	// TierName is the subscription tier name.
	TierName string `json:"tier_name"`
	// TotalMonthsSubscribed is the total months subscribed.
	TotalMonthsSubscribed int `json:"total_months_subscribed"`
	// IsRenewal indicates whether this is a renewal.
	IsRenewal bool `json:"is_renewal"`
}

// Poll represents a poll attached to a message.
type Poll struct {
	// Question is the poll question.
	Question PollMedia `json:"question"`
	// Answers are the poll answers.
	Answers []PollAnswer `json:"answers"`
	// Expiry is when the poll expires.
	Expiry *Timestamp `json:"expiry,omitempty"`
	// AllowMultiselect indicates whether multiple selections are allowed.
	AllowMultiselect bool `json:"allow_multiselect"`
	// LayoutType is the poll layout type.
	LayoutType int `json:"layout_type"`
	// Results are the poll results.
	Results *PollResults `json:"results,omitempty"`
}

// PollMedia represents poll question/answer media.
type PollMedia struct {
	// Text is the text content.
	Text string `json:"text,omitempty"`
	// Emoji is the emoji.
	Emoji *PartialEmoji `json:"emoji,omitempty"`
}

// PollAnswer represents a poll answer.
type PollAnswer struct {
	// AnswerID is the answer ID.
	AnswerID int `json:"answer_id"`
	// PollMedia is the answer content.
	PollMedia PollMedia `json:"poll_media"`
}

// PollResults represents poll results.
type PollResults struct {
	// IsFinalized indicates whether the poll is finalized.
	IsFinalized bool `json:"is_finalized"`
	// AnswerCounts contains the vote counts per answer.
	AnswerCounts []PollAnswerCount `json:"answer_counts"`
}

// PollAnswerCount represents a poll answer's vote count.
type PollAnswerCount struct {
	// ID is the answer ID.
	ID int `json:"id"`
	// Count is the vote count.
	Count int `json:"count"`
	// MeVoted indicates whether the current user voted for this.
	MeVoted bool `json:"me_voted"`
}

// MessageCall represents call information in a message.
type MessageCall struct {
	// Participants are the user IDs in the call.
	Participants []Snowflake `json:"participants"`
	// EndedTimestamp is when the call ended.
	EndedTimestamp *Timestamp `json:"ended_timestamp,omitempty"`
}

// ResolvedData contains resolved entities for interactions.
type ResolvedData struct {
	// Users maps user IDs to user objects.
	Users map[Snowflake]*User `json:"users,omitempty"`
	// Members maps user IDs to partial member objects.
	Members map[Snowflake]*GuildMember `json:"members,omitempty"`
	// Roles maps role IDs to role objects.
	Roles map[Snowflake]*Role `json:"roles,omitempty"`
	// Channels maps channel IDs to partial channel objects.
	Channels map[Snowflake]*Channel `json:"channels,omitempty"`
	// Messages maps message IDs to partial message objects.
	Messages map[Snowflake]*Message `json:"messages,omitempty"`
	// Attachments maps attachment IDs to attachment objects.
	Attachments map[Snowflake]*Attachment `json:"attachments,omitempty"`
}
