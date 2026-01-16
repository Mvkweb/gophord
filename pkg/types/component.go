// Package types provides Discord API type definitions.
//
// This package contains all the core data structures used to interact
// with the Discord API, including Message Components V2 types.
package types

// ComponentType represents the type of a Discord message component.
type ComponentType int

const (
	// ComponentTypeActionRow is a container for other components (type 1).
	ComponentTypeActionRow ComponentType = 1
	// ComponentTypeButton is an interactive button (type 2).
	ComponentTypeButton ComponentType = 2
	// ComponentTypeStringSelect is a select menu with string options (type 3).
	ComponentTypeStringSelect ComponentType = 3
	// ComponentTypeTextInput is a text input field for modals (type 4).
	ComponentTypeTextInput ComponentType = 4
	// ComponentTypeUserSelect is a select menu for users (type 5).
	ComponentTypeUserSelect ComponentType = 5
	// ComponentTypeRoleSelect is a select menu for roles (type 6).
	ComponentTypeRoleSelect ComponentType = 6
	// ComponentTypeMentionableSelect is a select menu for users and roles (type 7).
	ComponentTypeMentionableSelect ComponentType = 7
	// ComponentTypeChannelSelect is a select menu for channels (type 8).
	ComponentTypeChannelSelect ComponentType = 8
	// ComponentTypeSection is a layout component for content with accessory (type 9).
	ComponentTypeSection ComponentType = 9
	// ComponentTypeTextDisplay is a content component for markdown text (type 10).
	ComponentTypeTextDisplay ComponentType = 10
	// ComponentTypeThumbnail is a content component for small images (type 11).
	ComponentTypeThumbnail ComponentType = 11
	// ComponentTypeMediaGallery is a content component for displaying media (type 12).
	ComponentTypeMediaGallery ComponentType = 12
	// ComponentTypeFile is a content component for file attachments (type 13).
	ComponentTypeFile ComponentType = 13
	// ComponentTypeSeparator is a layout component for visual division (type 14).
	ComponentTypeSeparator ComponentType = 14
	// ComponentTypeContainer is a layout component for grouping components (type 17).
	ComponentTypeContainer ComponentType = 17
)

// String returns the string representation of a ComponentType.
func (c ComponentType) String() string {
	switch c {
	case ComponentTypeActionRow:
		return "ActionRow"
	case ComponentTypeButton:
		return "Button"
	case ComponentTypeStringSelect:
		return "StringSelect"
	case ComponentTypeTextInput:
		return "TextInput"
	case ComponentTypeUserSelect:
		return "UserSelect"
	case ComponentTypeRoleSelect:
		return "RoleSelect"
	case ComponentTypeMentionableSelect:
		return "MentionableSelect"
	case ComponentTypeChannelSelect:
		return "ChannelSelect"
	case ComponentTypeSection:
		return "Section"
	case ComponentTypeTextDisplay:
		return "TextDisplay"
	case ComponentTypeThumbnail:
		return "Thumbnail"
	case ComponentTypeMediaGallery:
		return "MediaGallery"
	case ComponentTypeFile:
		return "File"
	case ComponentTypeSeparator:
		return "Separator"
	case ComponentTypeContainer:
		return "Container"
	default:
		return "Unknown"
	}
}

// ButtonStyle represents the visual style of a button.
type ButtonStyle int

const (
	// ButtonStylePrimary is the most important action (blurple, style 1).
	ButtonStylePrimary ButtonStyle = 1
	// ButtonStyleSecondary is for alternative actions (grey, style 2).
	ButtonStyleSecondary ButtonStyle = 2
	// ButtonStyleSuccess is for positive confirmation (green, style 3).
	ButtonStyleSuccess ButtonStyle = 3
	// ButtonStyleDanger is for destructive actions (red, style 4).
	ButtonStyleDanger ButtonStyle = 4
	// ButtonStyleLink navigates to a URL (grey, style 5).
	ButtonStyleLink ButtonStyle = 5
	// ButtonStylePremium is for purchase actions (style 6).
	ButtonStylePremium ButtonStyle = 6
)

// String returns the string representation of a ButtonStyle.
func (b ButtonStyle) String() string {
	switch b {
	case ButtonStylePrimary:
		return "Primary"
	case ButtonStyleSecondary:
		return "Secondary"
	case ButtonStyleSuccess:
		return "Success"
	case ButtonStyleDanger:
		return "Danger"
	case ButtonStyleLink:
		return "Link"
	case ButtonStylePremium:
		return "Premium"
	default:
		return "Unknown"
	}
}

// SeparatorSpacing represents the padding size for separators.
type SeparatorSpacing int

const (
	// SeparatorSpacingSmall is small padding (1).
	SeparatorSpacingSmall SeparatorSpacing = 1
	// SeparatorSpacingLarge is large padding (2).
	SeparatorSpacingLarge SeparatorSpacing = 2
)

// Component is the interface that all message components implement.
type Component interface {
	// Type returns the component type.
	Type() ComponentType
}

// ComponentList is a slice of components that supports polymorphic JSON unmarshalling.
type ComponentList []Component

// ActionRow is a top-level layout component that contains other components.
// It can hold up to 5 buttons or a single select menu.
type ActionRow struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// Components contains the child components (buttons or a single select).
	Components ComponentList `json:"components"`
}

// Type returns ComponentTypeActionRow.
func (a *ActionRow) Type() ComponentType {
	return ComponentTypeActionRow
}

// Button is an interactive component for user actions.
type Button struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// Style determines the button's visual appearance.
	Style ButtonStyle `json:"style"`
	// Label is the text displayed on the button (max 80 characters).
	Label string `json:"label,omitempty"`
	// Emoji is the partial emoji displayed on the button.
	Emoji *PartialEmoji `json:"emoji,omitempty"`
	// CustomID is the developer-defined identifier (max 100 characters).
	// Required for non-link, non-premium buttons.
	CustomID string `json:"custom_id,omitempty"`
	// URL is the link for link-style buttons (max 512 characters).
	URL string `json:"url,omitempty"`
	// SKUID is the identifier for premium buttons.
	SKUID *Snowflake `json:"sku_id,omitempty"`
	// Disabled indicates whether the button is disabled.
	Disabled bool `json:"disabled,omitempty"`
}

// Type returns ComponentTypeButton.
func (b *Button) Type() ComponentType {
	return ComponentTypeButton
}

// SelectOption represents an option in a string select menu.
type SelectOption struct {
	// Label is the user-facing name (max 100 characters).
	Label string `json:"label"`
	// Value is the developer-defined value (max 100 characters).
	Value string `json:"value"`
	// Description is additional text (max 100 characters).
	Description string `json:"description,omitempty"`
	// Emoji is the partial emoji for the option.
	Emoji *PartialEmoji `json:"emoji,omitempty"`
	// Default indicates if this option is selected by default.
	Default bool `json:"default,omitempty"`
}

// StringSelect is a select menu with predefined string options.
type StringSelect struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// CustomID is the developer-defined identifier (max 100 characters).
	CustomID string `json:"custom_id"`
	// Options contains the available choices (max 25).
	Options []SelectOption `json:"options"`
	// Placeholder is the text shown when nothing is selected (max 150 characters).
	Placeholder string `json:"placeholder,omitempty"`
	// MinValues is the minimum number of selections (0-25, default 1).
	MinValues *int `json:"min_values,omitempty"`
	// MaxValues is the maximum number of selections (1-25, default 1).
	MaxValues *int `json:"max_values,omitempty"`
	// Disabled indicates whether the select is disabled.
	Disabled bool `json:"disabled,omitempty"`
}

// Type returns ComponentTypeStringSelect.
func (s *StringSelect) Type() ComponentType {
	return ComponentTypeStringSelect
}

// SelectDefaultValue represents a default value for auto-populated selects.
type SelectDefaultValue struct {
	// ID is the snowflake ID of the user, role, or channel.
	ID Snowflake `json:"id"`
	// Type indicates what ID represents: "user", "role", or "channel".
	Type string `json:"type"`
}

// UserSelect is a select menu that auto-populates with server users.
type UserSelect struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// CustomID is the developer-defined identifier (max 100 characters).
	CustomID string `json:"custom_id"`
	// Placeholder is the text shown when nothing is selected (max 150 characters).
	Placeholder string `json:"placeholder,omitempty"`
	// DefaultValues contains pre-selected users.
	DefaultValues []SelectDefaultValue `json:"default_values,omitempty"`
	// MinValues is the minimum number of selections (0-25, default 1).
	MinValues *int `json:"min_values,omitempty"`
	// MaxValues is the maximum number of selections (1-25, default 1).
	MaxValues *int `json:"max_values,omitempty"`
	// Disabled indicates whether the select is disabled.
	Disabled bool `json:"disabled,omitempty"`
}

// Type returns ComponentTypeUserSelect.
func (u *UserSelect) Type() ComponentType {
	return ComponentTypeUserSelect
}

// RoleSelect is a select menu that auto-populates with server roles.
type RoleSelect struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// CustomID is the developer-defined identifier (max 100 characters).
	CustomID string `json:"custom_id"`
	// Placeholder is the text shown when nothing is selected (max 150 characters).
	Placeholder string `json:"placeholder,omitempty"`
	// DefaultValues contains pre-selected roles.
	DefaultValues []SelectDefaultValue `json:"default_values,omitempty"`
	// MinValues is the minimum number of selections (0-25, default 1).
	MinValues *int `json:"min_values,omitempty"`
	// MaxValues is the maximum number of selections (1-25, default 1).
	MaxValues *int `json:"max_values,omitempty"`
	// Disabled indicates whether the select is disabled.
	Disabled bool `json:"disabled,omitempty"`
}

// Type returns ComponentTypeRoleSelect.
func (r *RoleSelect) Type() ComponentType {
	return ComponentTypeRoleSelect
}

// MentionableSelect is a select menu for both users and roles.
type MentionableSelect struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// CustomID is the developer-defined identifier (max 100 characters).
	CustomID string `json:"custom_id"`
	// Placeholder is the text shown when nothing is selected (max 150 characters).
	Placeholder string `json:"placeholder,omitempty"`
	// DefaultValues contains pre-selected users and roles.
	DefaultValues []SelectDefaultValue `json:"default_values,omitempty"`
	// MinValues is the minimum number of selections (0-25, default 1).
	MinValues *int `json:"min_values,omitempty"`
	// MaxValues is the maximum number of selections (1-25, default 1).
	MaxValues *int `json:"max_values,omitempty"`
	// Disabled indicates whether the select is disabled.
	Disabled bool `json:"disabled,omitempty"`
}

// Type returns ComponentTypeMentionableSelect.
func (m *MentionableSelect) Type() ComponentType {
	return ComponentTypeMentionableSelect
}

// ChannelSelect is a select menu that auto-populates with channels.
type ChannelSelect struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// CustomID is the developer-defined identifier (max 100 characters).
	CustomID string `json:"custom_id"`
	// ChannelTypes filters the available channel types.
	ChannelTypes []ChannelType `json:"channel_types,omitempty"`
	// Placeholder is the text shown when nothing is selected (max 150 characters).
	Placeholder string `json:"placeholder,omitempty"`
	// DefaultValues contains pre-selected channels.
	DefaultValues []SelectDefaultValue `json:"default_values,omitempty"`
	// MinValues is the minimum number of selections (0-25, default 1).
	MinValues *int `json:"min_values,omitempty"`
	// MaxValues is the maximum number of selections (1-25, default 1).
	MaxValues *int `json:"max_values,omitempty"`
	// Disabled indicates whether the select is disabled.
	Disabled bool `json:"disabled,omitempty"`
}

// Type returns ComponentTypeChannelSelect.
func (c *ChannelSelect) Type() ComponentType {
	return ComponentTypeChannelSelect
}

// Section is a layout component associating content with an accessory.
type Section struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// Components contains 1-3 text display components.
	Components ComponentList `json:"components"`
	// Accessory is a button or thumbnail associated with the content.
	Accessory Component `json:"accessory"`
}

// Type returns ComponentTypeSection.
func (s *Section) Type() ComponentType {
	return ComponentTypeSection
}

// TextDisplay is a content component for displaying markdown text.
type TextDisplay struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// Content is the markdown text to display.
	Content string `json:"content"`
}

// Type returns ComponentTypeTextDisplay.
func (t *TextDisplay) Type() ComponentType {
	return ComponentTypeTextDisplay
}

// UnfurledMediaItem represents media referenced by URL.
type UnfurledMediaItem struct {
	// URL supports arbitrary URLs and attachment:// references.
	URL string `json:"url"`
	// ProxyURL is the proxied URL (populated by Discord).
	ProxyURL string `json:"proxy_url,omitempty"`
	// Height is the media height in pixels (populated by Discord).
	Height *int `json:"height,omitempty"`
	// Width is the media width in pixels (populated by Discord).
	Width *int `json:"width,omitempty"`
	// ContentType is the MIME type (populated by Discord).
	ContentType string `json:"content_type,omitempty"`
	// AttachmentID is the uploaded attachment ID (populated by Discord).
	AttachmentID *Snowflake `json:"attachment_id,omitempty"`
}

// Thumbnail is a content component for displaying small images.
type Thumbnail struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// Media is the URL or attachment reference.
	Media UnfurledMediaItem `json:"media"`
	// Description is alt text for accessibility (max 1024 characters).
	Description string `json:"description,omitempty"`
	// Spoiler indicates whether to blur the image.
	Spoiler bool `json:"spoiler,omitempty"`
}

// Type returns ComponentTypeThumbnail.
func (t *Thumbnail) Type() ComponentType {
	return ComponentTypeThumbnail
}

// MediaGalleryItem represents a single item in a media gallery.
type MediaGalleryItem struct {
	// Media is the URL or attachment reference.
	Media UnfurledMediaItem `json:"media"`
	// Description is alt text for accessibility (max 1024 characters).
	Description string `json:"description,omitempty"`
	// Spoiler indicates whether to blur the media.
	Spoiler bool `json:"spoiler,omitempty"`
}

// MediaGallery is a content component for displaying 1-10 media items.
type MediaGallery struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// Items contains the gallery media (1-10 items).
	Items []MediaGalleryItem `json:"items"`
}

// Type returns ComponentTypeMediaGallery.
func (m *MediaGallery) Type() ComponentType {
	return ComponentTypeMediaGallery
}

// File is a content component for displaying file attachments.
type File struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// File references the attachment (attachment:// protocol only).
	File UnfurledMediaItem `json:"file"`
	// Spoiler indicates whether to blur the file preview.
	Spoiler bool `json:"spoiler,omitempty"`
	// Name is the filename (populated by Discord).
	Name string `json:"name,omitempty"`
	// Size is the file size in bytes (populated by Discord).
	Size int64 `json:"size,omitempty"`
}

// Type returns ComponentTypeFile.
func (f *File) Type() ComponentType {
	return ComponentTypeFile
}

// Separator is a layout component for visual division between components.
type Separator struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// Divider indicates whether to show a visual line (default true).
	Divider *bool `json:"divider,omitempty"`
	// Spacing is the padding size (1=small, 2=large, default 1).
	Spacing SeparatorSpacing `json:"spacing,omitempty"`
}

// Type returns ComponentTypeSeparator.
func (s *Separator) Type() ComponentType {
	return ComponentTypeSeparator
}

// Container is a layout component for visually grouping components.
type Container struct {
	// ID is an optional identifier for the component.
	ID *int `json:"id,omitempty"`
	// Components contains the child components.
	Components ComponentList `json:"components"`
	// AccentColor is the RGB color for the accent bar (0x000000 to 0xFFFFFF).
	AccentColor *int `json:"accent_color,omitempty"`
	// Spoiler indicates whether to blur the container contents.
	Spoiler bool `json:"spoiler,omitempty"`
}

// Type returns ComponentTypeContainer.
func (c *Container) Type() ComponentType {
	return ComponentTypeContainer
}
