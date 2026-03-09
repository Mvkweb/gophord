// Package client provides a high-level Discord client combining REST and Gateway.
//
// This package provides a unified interface for interacting with Discord,
// handling both HTTP REST requests and WebSocket gateway events.
package client

import (
	"context"
	"fmt"
	"log"

	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/json"
	"github.com/Mvkweb/gophord/pkg/rest"
	"github.com/Mvkweb/gophord/pkg/types"
)

// Client is a high-level Discord client.
type Client struct {
	// Token is the bot token.
	token string
	// REST is the REST API client.
	REST *rest.Client
	// Gateway is the gateway client.
	Gateway *gateway.Client
	// User is the current bot user (populated after Ready).
	User *types.User
	// ApplicationID is the application ID (populated after Ready).
	ApplicationID types.Snowflake

	// Gateway options
	gatewayOpts []gateway.ClientOption

	// Event handlers
	handlers map[string][]EventHandler
}

// EventHandler is a function that handles gateway events.
type EventHandler func(event interface{})

// ClientOption is a function that configures a Client.
type ClientOption func(*Client)

// WithIntents sets the gateway intents.
func WithIntents(intents types.IntentFlags) ClientOption {
	return func(c *Client) {
		c.gatewayOpts = append(c.gatewayOpts, gateway.WithIntents(intents))
	}
}

// WithMobileStatus sets whether to identify as a mobile client.
func WithMobileStatus(enabled bool) ClientOption {
	return func(c *Client) {
		c.gatewayOpts = append(c.gatewayOpts, gateway.WithMobileStatus(enabled))
	}
}

// New creates a new Discord client.
func New(token string, opts ...ClientOption) *Client {
	c := &Client{
		token:       token,
		REST:        rest.New(token),
		handlers:    make(map[string][]EventHandler),
		gatewayOpts: make([]gateway.ClientOption, 0),
	}

	for _, opt := range opts {
		opt(c)
	}

	// Create gateway with accumulated options
	c.Gateway = gateway.New(token, c.gatewayOpts...)

	return c
}

// On registers an event handler for the specified event type.
func (c *Client) On(eventType string, handler EventHandler) {
	c.handlers[eventType] = append(c.handlers[eventType], handler)
}

// OnReady registers a handler for the READY event.
func (c *Client) OnReady(handler func(*gateway.ReadyEvent)) {
	c.On(gateway.EventReady, func(event interface{}) {
		if e, ok := event.(*gateway.ReadyEvent); ok {
			handler(e)
		}
	})
}

// OnMessageCreate registers a handler for MESSAGE_CREATE events.
func (c *Client) OnMessageCreate(handler func(*gateway.MessageCreateEvent)) {
	c.On(gateway.EventMessageCreate, func(event interface{}) {
		if e, ok := event.(*gateway.MessageCreateEvent); ok {
			handler(e)
		}
	})
}

// OnInteractionCreate registers a handler for INTERACTION_CREATE events.
func (c *Client) OnInteractionCreate(handler func(*gateway.InteractionCreateEvent)) {
	c.On(gateway.EventInteractionCreate, func(event interface{}) {
		if e, ok := event.(*gateway.InteractionCreateEvent); ok {
			handler(e)
		}
	})
}

// Connect connects to the Discord gateway and starts handling events.
func (c *Client) Connect(ctx context.Context) error {
	// Connect to gateway
	if err := c.Gateway.Connect(ctx); err != nil {
		return fmt.Errorf("connect gateway: %w", err)
	}

	// Start event loop
	go c.eventLoop(ctx)

	return nil
}

// eventLoop processes gateway events.
func (c *Client) eventLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case event := <-c.Gateway.Events():
			c.handleEvent(event)
		case err := <-c.Gateway.Errors():
			log.Printf("Gateway error: %v", err)
		}
	}
}

// handleEvent dispatches an event to registered handlers.
func (c *Client) handleEvent(event *gateway.Event) {
	// Parse the event
	parsed, err := gateway.ParseEvent(event.Type, event.Data)
	if err != nil {
		log.Printf("Failed to parse event %s: %v", event.Type, err)
		return
	}

	// Handle READY specially to populate client data
	if event.Type == gateway.EventReady {
		if ready, ok := parsed.(*gateway.ReadyEvent); ok {
			c.User = ready.User
			c.ApplicationID = ready.Application.ID
		}
	}

	// Dispatch to handlers
	handlers, ok := c.handlers[event.Type]
	if !ok {
		return
	}

	for _, handler := range handlers {
		go handler(parsed)
	}
}

// Close closes the client connection.
func (c *Client) Close() error {
	return c.Gateway.Close()
}

// SendMessage sends a message to a channel.
func (c *Client) SendMessage(ctx context.Context, channelID types.Snowflake, content string) (*types.Message, error) {
	return c.REST.CreateMessage(ctx, channelID, &rest.CreateMessageParams{
		Content: content,
	})
}

// SendMessageSilent sends a message without triggering a notification (@silent behavior).
func (c *Client) SendMessageSilent(ctx context.Context, channelID types.Snowflake, content string) (*types.Message, error) {
	return c.REST.CreateMessage(ctx, channelID, &rest.CreateMessageParams{
		Content: content,
		Flags:   types.MessageFlagSuppressNotifications,
	})
}

// SendMessageWithComponents sends a message with components V2.
func (c *Client) SendMessageWithComponents(ctx context.Context, channelID types.Snowflake, components types.ComponentList) (*types.Message, error) {
	return c.REST.CreateMessage(ctx, channelID, &rest.CreateMessageParams{
		Components: components,
		Flags:      types.MessageFlagIsComponentsV2,
	})
}

// RespondToInteraction responds to an interaction.
func (c *Client) RespondToInteraction(ctx context.Context, interaction *types.Interaction, response *types.InteractionResponse) error {
	return c.REST.CreateInteractionResponse(ctx, interaction.ID, interaction.Token, response)
}

// RespondWithMessage responds to an interaction with a message.
func (c *Client) RespondWithMessage(ctx context.Context, interaction *types.Interaction, content string) error {
	return c.RespondToInteraction(ctx, interaction, &types.InteractionResponse{
		Type: types.InteractionCallbackTypeChannelMessageWithSource,
		Data: &types.InteractionCallbackData{
			Content: content,
		},
	})
}

// RespondWithEphemeral responds to an interaction with an ephemeral message (visible only to the user).
func (c *Client) RespondWithEphemeral(ctx context.Context, interaction *types.Interaction, content string) error {
	return c.RespondToInteraction(ctx, interaction, &types.InteractionResponse{
		Type: types.InteractionCallbackTypeChannelMessageWithSource,
		Data: &types.InteractionCallbackData{
			Content: content,
			Flags:   types.MessageFlagEphemeral,
		},
	})
}

// ShowModal shows a modal to the user.
func (c *Client) ShowModal(ctx context.Context, interaction *types.Interaction, title, customID string, components types.ComponentList) error {
	return c.RespondToInteraction(ctx, interaction, &types.InteractionResponse{
		Type: types.InteractionCallbackTypeModal,
		Data: &types.InteractionCallbackData{
			Title:      title,
			CustomID:   customID,
			Components: components,
			Flags:      types.MessageFlagIsComponentsV2,
		},
	})
}

// RespondWithComponents responds to an interaction with components V2.
func (c *Client) RespondWithComponents(ctx context.Context, interaction *types.Interaction, components types.ComponentList) error {
	return c.RespondToInteraction(ctx, interaction, &types.InteractionResponse{
		Type: types.InteractionCallbackTypeChannelMessageWithSource,
		Data: &types.InteractionCallbackData{
			Components: components,
			Flags:      types.MessageFlagIsComponentsV2,
		},
	})
}

// DeferInteraction defers an interaction response (shows "thinking").
func (c *Client) DeferInteraction(ctx context.Context, interaction *types.Interaction) error {
	return c.RespondToInteraction(ctx, interaction, &types.InteractionResponse{
		Type: types.InteractionCallbackTypeDeferredChannelMessageWithSource,
	})
}

// UpdateInteractionMessage updates the original interaction message.
func (c *Client) UpdateInteractionMessage(ctx context.Context, interaction *types.Interaction, content string) (*types.Message, error) {
	return c.REST.EditOriginalInteractionResponse(ctx, c.ApplicationID, interaction.Token, &rest.EditMessageParams{
		Content: &content,
	})
}

// RegisterGlobalCommand registers a global slash command.
func (c *Client) RegisterGlobalCommand(ctx context.Context, command types.CreateApplicationCommandParams) (*types.ApplicationCommand, error) {
	if c.ApplicationID == 0 {
		// Try to fetch current user to get application ID if not already set (e.g. before Ready)
		user, err := c.REST.GetCurrentUser(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get current user to determine application ID: %w", err)
		}
		c.ApplicationID = user.ID
	}
	return c.REST.CreateGlobalApplicationCommand(ctx, c.ApplicationID, &command)
}

// RegisterGuildCommand registers a guild slash command.
func (c *Client) RegisterGuildCommand(ctx context.Context, guildID types.Snowflake, command types.CreateApplicationCommandParams) (*types.ApplicationCommand, error) {
	if c.ApplicationID == 0 {
		// Try to fetch current user to get application ID if not already set
		user, err := c.REST.GetCurrentUser(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get current user to determine application ID: %w", err)
		}
		c.ApplicationID = user.ID
	}
	return c.REST.CreateGuildApplicationCommand(ctx, c.ApplicationID, guildID, &command)
}

// BulkRegisterGlobalCommands overwrites all global commands.
func (c *Client) BulkRegisterGlobalCommands(ctx context.Context, commands []types.CreateApplicationCommandParams) ([]types.ApplicationCommand, error) {
	if c.ApplicationID == 0 {
		user, err := c.REST.GetCurrentUser(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get current user to determine application ID: %w", err)
		}
		c.ApplicationID = user.ID
	}
	return c.REST.BulkOverwriteGlobalApplicationCommands(ctx, c.ApplicationID, commands)
}

// Guild & Member Management

// KickMember kicks a member from a guild.
func (c *Client) KickMember(ctx context.Context, guildID, userID types.Snowflake) error {
	return c.REST.RemoveGuildMember(ctx, guildID, userID)
}

// BanMember bans a member from a guild.
func (c *Client) BanMember(ctx context.Context, guildID, userID types.Snowflake, deleteMessageSeconds int) error {
	return c.REST.CreateGuildBan(ctx, guildID, userID, deleteMessageSeconds)
}

// UnbanMember removes a ban from a guild.
func (c *Client) UnbanMember(ctx context.Context, guildID, userID types.Snowflake) error {
	return c.REST.RemoveGuildBan(ctx, guildID, userID)
}

// AddMemberRole adds a role to a guild member.
func (c *Client) AddMemberRole(ctx context.Context, guildID, userID, roleID types.Snowflake) error {
	return c.REST.AddGuildMemberRole(ctx, guildID, userID, roleID)
}

// RemoveMemberRole removes a role from a guild member.
func (c *Client) RemoveMemberRole(ctx context.Context, guildID, userID, roleID types.Snowflake) error {
	return c.REST.RemoveGuildMemberRole(ctx, guildID, userID, roleID)
}

// GetMember returns a guild member.
func (c *Client) GetMember(ctx context.Context, guildID, userID types.Snowflake) (*types.GuildMember, error) {
	return c.REST.GetGuildMember(ctx, guildID, userID)
}

// Webhooks

// CreateWebhook creates a new webhook in a channel.
func (c *Client) CreateWebhook(ctx context.Context, channelID types.Snowflake, name string) (*types.Webhook, error) {
	return c.REST.CreateWebhook(ctx, channelID, &types.CreateWebhookParams{
		Name: name,
	})
}

// ExecuteWebhook executes a webhook by ID and token.
func (c *Client) ExecuteWebhook(ctx context.Context, webhookID types.Snowflake, token string, content string) error {
	_, err := c.REST.ExecuteWebhook(ctx, webhookID, token, false, &types.ExecuteWebhookParams{
		Content: content,
	})
	return err
}

// Reactions & Pins

// React adds an emoji reaction to a message.
func (c *Client) React(ctx context.Context, channelID, messageID types.Snowflake, emoji string) error {
	return c.REST.CreateReaction(ctx, channelID, messageID, emoji)
}

// Unreact removes the bot's emoji reaction from a message.
func (c *Client) Unreact(ctx context.Context, channelID, messageID types.Snowflake, emoji string) error {
	return c.REST.DeleteOwnReaction(ctx, channelID, messageID, emoji)
}

// Pin pins a message in a channel.
func (c *Client) Pin(ctx context.Context, channelID, messageID types.Snowflake) error {
	return c.REST.PinMessage(ctx, channelID, messageID)
}

// Unpin unpins a message in a channel.
func (c *Client) Unpin(ctx context.Context, channelID, messageID types.Snowflake) error {
	return c.REST.UnpinMessage(ctx, channelID, messageID)
}

// ComponentBuilder provides fluent API for building message components.

type ComponentBuilder struct {
	components types.ComponentList
}

// NewComponentBuilder creates a new component builder.
func NewComponentBuilder() *ComponentBuilder {
	return &ComponentBuilder{
		components: make(types.ComponentList, 0),
	}
}

// AddTextDisplay adds a text display component.
func (b *ComponentBuilder) AddTextDisplay(content string) *ComponentBuilder {
	b.components = append(b.components, &types.TextDisplay{Content: content})
	return b
}

// AddActionRow adds an action row with components.
func (b *ComponentBuilder) AddActionRow(components ...types.Component) *ComponentBuilder {
	list := make(types.ComponentList, len(components))
	copy(list, components)
	b.components = append(b.components, &types.ActionRow{Components: list})
	return b
}

// AddSeparator adds a separator component.
func (b *ComponentBuilder) AddSeparator(divider bool, spacing types.SeparatorSpacing) *ComponentBuilder {
	b.components = append(b.components, &types.Separator{
		Divider: &divider,
		Spacing: spacing,
	})
	return b
}

// AddContainer wraps components in a container with an accent color.
func (b *ComponentBuilder) AddContainer(accentColor int, children ...types.Component) *ComponentBuilder {
	b.components = append(b.components, &types.Container{
		AccentColor: &accentColor,
		Components:  children,
	})
	return b
}

// AddMediaGallery adds a media gallery component.
func (b *ComponentBuilder) AddMediaGallery(items ...types.MediaGalleryItem) *ComponentBuilder {
	b.components = append(b.components, &types.MediaGallery{Items: items})
	return b
}

// AddSection adds a section with text and an accessory.
func (b *ComponentBuilder) AddSection(textContent string, accessory types.Component) *ComponentBuilder {
	b.components = append(b.components, &types.Section{
		Components: types.ComponentList{&types.TextDisplay{Content: textContent}},
		Accessory:  accessory,
	})
	return b
}

// AddLabel adds a label component (Type 18) for modals.
func (b *ComponentBuilder) AddLabel(label *types.Label) *ComponentBuilder {
	b.components = append(b.components, label)
	return b
}

// Build returns the built components.
func (b *ComponentBuilder) Build() types.ComponentList {
	return b.components
}

// ToJSON returns the components as JSON bytes.
func (b *ComponentBuilder) ToJSON() ([]byte, error) {
	return json.Marshal(b.components)
}

// Button helper functions

// NewPrimaryButton creates a primary button.
func NewPrimaryButton(customID, label string) *types.Button {
	return &types.Button{
		Style:    types.ButtonStylePrimary,
		CustomID: customID,
		Label:    label,
	}
}

// NewSecondaryButton creates a secondary button.
func NewSecondaryButton(customID, label string) *types.Button {
	return &types.Button{
		Style:    types.ButtonStyleSecondary,
		CustomID: customID,
		Label:    label,
	}
}

// NewSuccessButton creates a success button.
func NewSuccessButton(customID, label string) *types.Button {
	return &types.Button{
		Style:    types.ButtonStyleSuccess,
		CustomID: customID,
		Label:    label,
	}
}

// NewDangerButton creates a danger button.
func NewDangerButton(customID, label string) *types.Button {
	return &types.Button{
		Style:    types.ButtonStyleDanger,
		CustomID: customID,
		Label:    label,
	}
}

// NewLinkButton creates a link button.
func NewLinkButton(url, label string) *types.Button {
	return &types.Button{
		Style: types.ButtonStyleLink,
		URL:   url,
		Label: label,
	}
}

// NewTextInput creates a new text input component.
func NewTextInput(customID, label string, style int, opts ...func(*types.TextInput)) *types.TextInput {
	ti := &types.TextInput{
		CustomID: customID,
		Label:    label,
		Style:    style,
	}
	for _, opt := range opts {
		opt(ti)
	}
	return ti
}

// WithMinLength sets the minimum length for a text input.
func WithMinLength(min int) func(*types.TextInput) {
	return func(ti *types.TextInput) {
		ti.MinLength = &min
	}
}

// WithMaxLength sets the maximum length for a text input.
func WithMaxLength(max int) func(*types.TextInput) {
	return func(ti *types.TextInput) {
		ti.MaxLength = &max
	}
}

// WithPlaceholder sets the placeholder text for a text input.
func WithPlaceholder(placeholder string) func(*types.TextInput) {
	return func(ti *types.TextInput) {
		ti.Placeholder = placeholder
	}
}

// WithRequired sets whether a text input is required.
func WithRequired(required bool) func(*types.TextInput) {
	return func(ti *types.TextInput) {
		ti.Required = &required
	}
}

// NewStringSelect creates a new string select menu.
func NewStringSelect(customID string, options ...types.SelectOption) *types.StringSelect {
	return &types.StringSelect{
		CustomID: customID,
		Options:  options,
	}
}

// NewLabel creates a new label (Type 18) for modals.
func NewLabel(label, description string, component types.Component) *types.Label {
	return &types.Label{
		Label:       label,
		Description: description,
		Component:   component,
	}
}

// NewUserSelect creates a new user select menu for modals.
func NewUserSelect(customID string) *types.UserSelect {
	return &types.UserSelect{
		CustomID: customID,
	}
}

// NewRoleSelect creates a new role select menu for modals.
func NewRoleSelect(customID string) *types.RoleSelect {
	return &types.RoleSelect{
		CustomID: customID,
	}
}

// NewMentionableSelect creates a new mentionable select menu (users + roles) for modals.
func NewMentionableSelect(customID string) *types.MentionableSelect {
	return &types.MentionableSelect{
		CustomID: customID,
	}
}

// NewChannelSelect creates a new channel select menu for modals.
func NewChannelSelect(customID string) *types.ChannelSelect {
	return &types.ChannelSelect{
		CustomID: customID,
	}
}

// NewFileUpload creates a new file upload component for modals.
func NewFileUpload(customID string) *types.FileUpload {
	return &types.FileUpload{
		CustomID: customID,
	}
}

// NewTextDisplay creates a new text display component for modals.
func NewTextDisplay(content string) *types.TextDisplay {
	return &types.TextDisplay{
		Content: content,
	}
}
