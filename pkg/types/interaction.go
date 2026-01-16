// Package types provides Discord API type definitions.
package types

// InteractionType represents the type of an interaction.
type InteractionType int

const (
	// InteractionTypePing is a ping interaction (type 1).
	InteractionTypePing InteractionType = 1
	// InteractionTypeApplicationCommand is a slash command (type 2).
	InteractionTypeApplicationCommand InteractionType = 2
	// InteractionTypeMessageComponent is a component interaction (type 3).
	InteractionTypeMessageComponent InteractionType = 3
	// InteractionTypeApplicationCommandAutocomplete is autocomplete (type 4).
	InteractionTypeApplicationCommandAutocomplete InteractionType = 4
	// InteractionTypeModalSubmit is a modal submit (type 5).
	InteractionTypeModalSubmit InteractionType = 5
)

// String returns the string representation of an InteractionType.
func (i InteractionType) String() string {
	switch i {
	case InteractionTypePing:
		return "Ping"
	case InteractionTypeApplicationCommand:
		return "ApplicationCommand"
	case InteractionTypeMessageComponent:
		return "MessageComponent"
	case InteractionTypeApplicationCommandAutocomplete:
		return "ApplicationCommandAutocomplete"
	case InteractionTypeModalSubmit:
		return "ModalSubmit"
	default:
		return "Unknown"
	}
}

// InteractionCallbackType represents the type of an interaction response.
type InteractionCallbackType int

const (
	// InteractionCallbackTypePong responds to a ping (type 1).
	InteractionCallbackTypePong InteractionCallbackType = 1
	// InteractionCallbackTypeChannelMessageWithSource responds with a message (type 4).
	InteractionCallbackTypeChannelMessageWithSource InteractionCallbackType = 4
	// InteractionCallbackTypeDeferredChannelMessageWithSource defers and shows loading (type 5).
	InteractionCallbackTypeDeferredChannelMessageWithSource InteractionCallbackType = 5
	// InteractionCallbackTypeDeferredUpdateMessage defers for component updates (type 6).
	InteractionCallbackTypeDeferredUpdateMessage InteractionCallbackType = 6
	// InteractionCallbackTypeUpdateMessage updates the component message (type 7).
	InteractionCallbackTypeUpdateMessage InteractionCallbackType = 7
	// InteractionCallbackTypeApplicationCommandAutocompleteResult returns autocomplete results (type 8).
	InteractionCallbackTypeApplicationCommandAutocompleteResult InteractionCallbackType = 8
	// InteractionCallbackTypeModal opens a modal (type 9).
	InteractionCallbackTypeModal InteractionCallbackType = 9
	// InteractionCallbackTypePremiumRequired shows premium required (type 10, deprecated).
	InteractionCallbackTypePremiumRequired InteractionCallbackType = 10
	// InteractionCallbackTypeLaunchActivity launches an activity (type 12).
	InteractionCallbackTypeLaunchActivity InteractionCallbackType = 12
)

// Interaction represents a Discord interaction.
type Interaction struct {
	// ID is the interaction's unique ID.
	ID Snowflake `json:"id"`
	// ApplicationID is the application's ID.
	ApplicationID Snowflake `json:"application_id"`
	// Type is the type of interaction.
	Type InteractionType `json:"type"`
	// Data contains the interaction data payload.
	Data *InteractionData `json:"data,omitempty"`
	// GuildID is the guild where the interaction was triggered.
	GuildID *Snowflake `json:"guild_id,omitempty"`
	// Channel is the channel where the interaction was triggered.
	Channel *Channel `json:"channel,omitempty"`
	// ChannelID is the channel ID where the interaction was triggered.
	ChannelID *Snowflake `json:"channel_id,omitempty"`
	// Member is the guild member who triggered the interaction.
	Member *GuildMember `json:"member,omitempty"`
	// User is the user who triggered the interaction.
	User *User `json:"user,omitempty"`
	// Token is the continuation token for responding.
	Token string `json:"token"`
	// Version is the interaction version (always 1).
	Version int `json:"version"`
	// Message is the message the component was attached to.
	Message *Message `json:"message,omitempty"`
	// AppPermissions is the app's permissions in the channel.
	AppPermissions string `json:"app_permissions,omitempty"`
	// Locale is the user's locale.
	Locale string `json:"locale,omitempty"`
	// GuildLocale is the guild's locale.
	GuildLocale string `json:"guild_locale,omitempty"`
	// Entitlements are the monetization entitlements.
	Entitlements []Entitlement `json:"entitlements,omitempty"`
	// AuthorizingIntegrationOwners maps integration types to IDs.
	AuthorizingIntegrationOwners map[string]Snowflake `json:"authorizing_integration_owners,omitempty"`
	// Context is the interaction context.
	Context *InteractionContext `json:"context,omitempty"`
}

// InteractionData represents the data payload of an interaction.
type InteractionData struct {
	// ID is the command ID (for application commands).
	ID Snowflake `json:"id,omitempty"`
	// Name is the command name (for application commands).
	Name string `json:"name,omitempty"`
	// Type is the command type (for application commands).
	Type int `json:"type,omitempty"`
	// Resolved contains resolved users, roles, channels, etc.
	Resolved *ResolvedData `json:"resolved,omitempty"`
	// Options are the command options (for application commands).
	Options []ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	// GuildID is the guild ID for guild commands.
	GuildID *Snowflake `json:"guild_id,omitempty"`
	// TargetID is the target ID (for user/message commands).
	TargetID *Snowflake `json:"target_id,omitempty"`

	// CustomID is the component's custom ID (for components/modals).
	CustomID string `json:"custom_id,omitempty"`
	// ComponentType is the type of component (for component interactions).
	ComponentType ComponentType `json:"component_type,omitempty"`
	// Values are the selected values (for select menus).
	Values []string `json:"values,omitempty"`

	// Components are the modal components (for modal submissions).
	Components ComponentList `json:"components,omitempty"`
}

// ApplicationCommandInteractionDataOption represents a command option value.
type ApplicationCommandInteractionDataOption struct {
	// Name is the option name.
	Name string `json:"name"`
	// Type is the option type.
	Type int `json:"type"`
	// Value is the option value.
	Value interface{} `json:"value,omitempty"`
	// Options are the nested options (for subcommands/groups).
	Options []ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	// Focused indicates whether this option is being autocompleted.
	Focused bool `json:"focused,omitempty"`
}

// InteractionContext represents the context of an interaction.
type InteractionContext int

const (
	// InteractionContextGuild is a guild context.
	InteractionContextGuild InteractionContext = 0
	// InteractionContextBotDM is a bot DM context.
	InteractionContextBotDM InteractionContext = 1
	// InteractionContextPrivateChannel is a private channel context.
	InteractionContextPrivateChannel InteractionContext = 2
)

// Entitlement represents a monetization entitlement.
type Entitlement struct {
	// ID is the entitlement ID.
	ID Snowflake `json:"id"`
	// SKUID is the SKU ID.
	SKUID Snowflake `json:"sku_id"`
	// ApplicationID is the application ID.
	ApplicationID Snowflake `json:"application_id"`
	// UserID is the user ID (if user subscription).
	UserID *Snowflake `json:"user_id,omitempty"`
	// Type is the entitlement type.
	Type int `json:"type"`
	// Deleted indicates whether the entitlement was deleted.
	Deleted bool `json:"deleted"`
	// StartsAt is when the entitlement starts.
	StartsAt *Timestamp `json:"starts_at,omitempty"`
	// EndsAt is when the entitlement ends.
	EndsAt *Timestamp `json:"ends_at,omitempty"`
	// GuildID is the guild ID (if guild subscription).
	GuildID *Snowflake `json:"guild_id,omitempty"`
	// Consumed indicates whether the entitlement was consumed.
	Consumed bool `json:"consumed,omitempty"`
}

// InteractionResponse represents a response to an interaction.
type InteractionResponse struct {
	// Type is the response type.
	Type InteractionCallbackType `json:"type"`
	// Data is the response data.
	Data *InteractionCallbackData `json:"data,omitempty"`
}

// InteractionCallbackData represents the data in an interaction response.
type InteractionCallbackData struct {
	// TTS indicates whether the message is text-to-speech.
	TTS bool `json:"tts,omitempty"`
	// Content is the message content (max 2000 characters).
	Content string `json:"content,omitempty"`
	// Embeds are the message embeds (max 10).
	Embeds []Embed `json:"embeds,omitempty"`
	// AllowedMentions controls mentions in the message.
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	// Flags are the message flags.
	Flags MessageFlags `json:"flags,omitempty"`
	// Components are the message components.
	Components ComponentList `json:"components,omitempty"`
	// Attachments are the message attachments.
	Attachments []Attachment `json:"attachments,omitempty"`
	// Poll is the message poll.
	Poll *Poll `json:"poll,omitempty"`

	// Title is the modal title (for modal responses).
	Title string `json:"title,omitempty"`
	// CustomID is the modal custom ID (for modal responses).
	CustomID string `json:"custom_id,omitempty"`

	// Choices are autocomplete results (for autocomplete responses).
	Choices []ApplicationCommandOptionChoice `json:"choices,omitempty"`
}

// AllowedMentions controls which mentions parse in a message.
type AllowedMentions struct {
	// Parse is an array of allowed mention types.
	Parse []string `json:"parse,omitempty"`
	// Roles are specific role IDs to mention.
	Roles []Snowflake `json:"roles,omitempty"`
	// Users are specific user IDs to mention.
	Users []Snowflake `json:"users,omitempty"`
	// RepliedUser indicates whether to mention the replied user.
	RepliedUser bool `json:"replied_user,omitempty"`
}

// ApplicationCommandOptionChoice represents an autocomplete choice.
type ApplicationCommandOptionChoice struct {
	// Name is the choice name (max 100 characters).
	Name string `json:"name"`
	// NameLocalizations are localized names.
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	// Value is the choice value.
	Value interface{} `json:"value"`
}

// TextInput is a modal text input component.
type TextInput struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// CustomID is the developer-defined identifier.
	CustomID string `json:"custom_id"`
	// Style is the text input style (1=short, 2=paragraph).
	Style int `json:"style"`
	// Label is the label for the input (max 45 characters). Not used when inside InputContainer.
	Label string `json:"label,omitempty"`
	// MinLength is the minimum input length (0-4000).
	MinLength *int `json:"min_length,omitempty"`
	// MaxLength is the maximum input length (1-4000).
	MaxLength *int `json:"max_length,omitempty"`
	// Required indicates whether the input is required.
	Required *bool `json:"required,omitempty"`
	// Value is the pre-filled value (max 4000 characters).
	Value string `json:"value,omitempty"`
	// Placeholder is the placeholder text (max 100 characters).
	Placeholder string `json:"placeholder,omitempty"`
}

// Type returns ComponentTypeTextInput.
func (t *TextInput) Type() ComponentType {
	return ComponentTypeTextInput
}
