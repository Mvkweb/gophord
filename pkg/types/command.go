// Package types provides Discord API type definitions.
package types

// ApplicationCommandType represents the type of an application command.
type ApplicationCommandType int

const (
	// ApplicationCommandTypeChatInput is a slash command (type 1).
	ApplicationCommandTypeChatInput ApplicationCommandType = 1
	// ApplicationCommandTypeUser is a user context menu command (type 2).
	ApplicationCommandTypeUser ApplicationCommandType = 2
	// ApplicationCommandTypeMessage is a message context menu command (type 3).
	ApplicationCommandTypeMessage ApplicationCommandType = 3
)

// ApplicationCommandOptionType represents the type of an option.
type ApplicationCommandOptionType int

const (
	// ApplicationCommandOptionTypeSubCommand is a subcommand (type 1).
	ApplicationCommandOptionTypeSubCommand ApplicationCommandOptionType = 1
	// ApplicationCommandOptionTypeSubCommandGroup is a subcommand group (type 2).
	ApplicationCommandOptionTypeSubCommandGroup ApplicationCommandOptionType = 2
	// ApplicationCommandOptionTypeString is a string option (type 3).
	ApplicationCommandOptionTypeString ApplicationCommandOptionType = 3
	// ApplicationCommandOptionTypeInteger is an integer option (type 4).
	ApplicationCommandOptionTypeInteger ApplicationCommandOptionType = 4
	// ApplicationCommandOptionTypeBoolean is a boolean option (type 5).
	ApplicationCommandOptionTypeBoolean ApplicationCommandOptionType = 5
	// ApplicationCommandOptionTypeUser is a user option (type 6).
	ApplicationCommandOptionTypeUser ApplicationCommandOptionType = 6
	// ApplicationCommandOptionTypeChannel is a channel option (type 7).
	ApplicationCommandOptionTypeChannel ApplicationCommandOptionType = 7
	// ApplicationCommandOptionTypeRole is a role option (type 8).
	ApplicationCommandOptionTypeRole ApplicationCommandOptionType = 8
	// ApplicationCommandOptionTypeMentionable is a mentionable option (type 9).
	ApplicationCommandOptionTypeMentionable ApplicationCommandOptionType = 9
	// ApplicationCommandOptionTypeNumber is a double option (type 10).
	ApplicationCommandOptionTypeNumber ApplicationCommandOptionType = 10
	// ApplicationCommandOptionTypeAttachment is an attachment option (type 11).
	ApplicationCommandOptionTypeAttachment ApplicationCommandOptionType = 11
)

// ApplicationCommand represents a command that can be executed in Discord.
type ApplicationCommand struct {
	// ID is the unique ID of the command.
	ID Snowflake `json:"id,omitempty"`
	// Type is the type of command (default 1).
	Type ApplicationCommandType `json:"type,omitempty"`
	// ApplicationID is the unique ID of the parent application.
	ApplicationID Snowflake `json:"application_id,omitempty"`
	// GuildID is the guild ID (if guild command).
	GuildID *Snowflake `json:"guild_id,omitempty"`
	// Name is the name of the command (1-32 chars).
	Name string `json:"name"`
	// NameLocalizations are localized names.
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	// Description is the description of the command (1-100 chars).
	Description string `json:"description"`
	// DescriptionLocalizations are localized descriptions.
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	// Options are the parameters for the command.
	Options []ApplicationCommandOption `json:"options,omitempty"`
	// DefaultMemberPermissions are the default permissions required.
	DefaultMemberPermissions *string `json:"default_member_permissions,omitempty"`
	// DMPermission indicates whether the command is available in DMs (deprecated).
	DMPermission *bool `json:"dm_permission,omitempty"`
	// Nsfw indicates whether the command is age-restricted.
	Nsfw bool `json:"nsfw,omitempty"`
	// Version is an autoincrementing version identifier.
	Version Snowflake `json:"version,omitempty"`
}

// ApplicationCommandOption represents a parameter for a command.
type ApplicationCommandOption struct {
	// Type is the option type.
	Type ApplicationCommandOptionType `json:"type"`
	// Name is the option name (1-32 chars).
	Name string `json:"name"`
	// NameLocalizations are localized names.
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	// Description is the option description (1-100 chars).
	Description string `json:"description"`
	// DescriptionLocalizations are localized descriptions.
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	// Required indicates whether the option is required.
	Required bool `json:"required,omitempty"`
	// Choices are the choices for string/int/number options.
	Choices []ApplicationCommandOptionChoice `json:"choices,omitempty"`
	// Options are nested options (for subcommands).
	Options []ApplicationCommandOption `json:"options,omitempty"`
	// ChannelTypes restricts channel types for channel options.
	ChannelTypes []ChannelType `json:"channel_types,omitempty"`
	// MinValue is the minimum value for int/number options.
	MinValue *float64 `json:"min_value,omitempty"`
	// MaxValue is the maximum value for int/number options.
	MaxValue *float64 `json:"max_value,omitempty"`
	// MinLength is the minimum string length.
	MinLength *int `json:"min_length,omitempty"`
	// MaxLength is the maximum string length.
	MaxLength *int `json:"max_length,omitempty"`
	// Autocomplete indicates whether autocomplete is enabled.
	Autocomplete bool `json:"autocomplete,omitempty"`
}

// ApplicationCommandOptionChoice represents a pre-determined choice.
// Note: This is already defined in interaction.go as ApplicationCommandOptionChoice
// but we might want to consolidate or alias it if needed.
// For now, I will assume the one in interaction.go is sufficient,
// but since this file is new, I will check interaction.go first to avoid duplicates
// or simply reuse it. The one in interaction.go is:
/*
type ApplicationCommandOptionChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Value             interface{}       `json:"value"`
}
*/
// I will rely on the type defined in interaction.go.

// CreateApplicationCommandParams contains parameters for creating a command.
type CreateApplicationCommandParams struct {
	// Name is the name of the command.
	Name string `json:"name"`
	// NameLocalizations are localized names.
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	// Description is the description of the command.
	Description string `json:"description,omitempty"`
	// DescriptionLocalizations are localized descriptions.
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	// Options are the parameters for the command.
	Options []ApplicationCommandOption `json:"options,omitempty"`
	// DefaultMemberPermissions are the default permissions required.
	DefaultMemberPermissions *string `json:"default_member_permissions,omitempty"`
	// DMPermission indicates whether the command is available in DMs.
	DMPermission *bool `json:"dm_permission,omitempty"`
	// Type is the type of command.
	Type *ApplicationCommandType `json:"type,omitempty"`
	// Nsfw indicates whether the command is age-restricted.
	Nsfw *bool `json:"nsfw,omitempty"`
}
