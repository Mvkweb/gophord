// Package rest provides a REST client for the Discord API.
package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Mvkweb/gophord/pkg/json"
	"github.com/Mvkweb/gophord/pkg/types"
)

// Channels

// GetChannel returns a channel by ID.
func (c *Client) GetChannel(ctx context.Context, channelID types.Snowflake) (*types.Channel, error) {
	path := fmt.Sprintf("/channels/%s", channelID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var channel types.Channel
	if err := json.Unmarshal(data, &channel); err != nil {
		return nil, fmt.Errorf("unmarshal channel: %w", err)
	}

	return &channel, nil
}

// ModifyChannel modifies a channel's settings.
func (c *Client) ModifyChannel(ctx context.Context, channelID types.Snowflake, params interface{}) (*types.Channel, error) {
	path := fmt.Sprintf("/channels/%s", channelID)
	data, err := c.Request(ctx, http.MethodPatch, path, params)
	if err != nil {
		return nil, err
	}

	var channel types.Channel
	if err := json.Unmarshal(data, &channel); err != nil {
		return nil, fmt.Errorf("unmarshal channel: %w", err)
	}

	return &channel, nil
}

// DeleteChannel deletes a channel.
func (c *Client) DeleteChannel(ctx context.Context, channelID types.Snowflake) error {
	path := fmt.Sprintf("/channels/%s", channelID)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// Messages

// CreateMessageParams contains parameters for creating a message.
type CreateMessageParams struct {
	// Content is the message content (up to 2000 characters).
	Content string `json:"content,omitempty"`
	// TTS indicates whether this is a TTS message.
	TTS bool `json:"tts,omitempty"`
	// Embeds are the rich embeds (up to 10).
	Embeds []types.Embed `json:"embeds,omitempty"`
	// AllowedMentions controls mention behavior.
	AllowedMentions *types.AllowedMentions `json:"allowed_mentions,omitempty"`
	// MessageReference is the message to reply to.
	MessageReference *types.MessageReference `json:"message_reference,omitempty"`
	// Components are the message components.
	Components types.ComponentList `json:"components,omitempty"`
	// StickerIDs are sticker IDs to send (up to 3).
	StickerIDs []types.Snowflake `json:"sticker_ids,omitempty"`
	// Flags are message flags (e.g., MessageFlagIsComponentsV2).
	Flags types.MessageFlags `json:"flags,omitempty"`
	// Nonce is used for optimistic message sending.
	Nonce string `json:"nonce,omitempty"`
	// EnforceNonce indicates whether to check for duplicate nonces.
	EnforceNonce bool `json:"enforce_nonce,omitempty"`
	// Poll is a poll to send with the message.
	Poll *types.Poll `json:"poll,omitempty"`
}

// CreateMessage sends a message to a channel.
func (c *Client) CreateMessage(ctx context.Context, channelID types.Snowflake, params *CreateMessageParams) (*types.Message, error) {
	path := fmt.Sprintf("/channels/%s/messages", channelID)
	data, err := c.Request(ctx, http.MethodPost, path, params)
	if err != nil {
		return nil, err
	}

	var message types.Message
	if err := json.Unmarshal(data, &message); err != nil {
		return nil, fmt.Errorf("unmarshal message: %w", err)
	}

	return &message, nil
}

// GetMessage returns a message by ID.
func (c *Client) GetMessage(ctx context.Context, channelID, messageID types.Snowflake) (*types.Message, error) {
	path := fmt.Sprintf("/channels/%s/messages/%s", channelID, messageID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var message types.Message
	if err := json.Unmarshal(data, &message); err != nil {
		return nil, fmt.Errorf("unmarshal message: %w", err)
	}

	return &message, nil
}

// GetMessagesParams contains query parameters for GetMessages.
type GetMessagesParams struct {
	Around types.Snowflake `json:"around,omitempty"`
	Before types.Snowflake `json:"before,omitempty"`
	After  types.Snowflake `json:"after,omitempty"`
	Limit  int             `json:"limit,omitempty"`
}

// GetMessages returns a list of messages in a channel.
func (c *Client) GetMessages(ctx context.Context, channelID types.Snowflake, params *GetMessagesParams) ([]types.Message, error) {
	path := fmt.Sprintf("/channels/%s/messages", channelID)
	query := make(map[string]string)
	if params != nil {
		if params.Around != 0 {
			query["around"] = params.Around.String()
		}
		if params.Before != 0 {
			query["before"] = params.Before.String()
		}
		if params.After != 0 {
			query["after"] = params.After.String()
		}
		if params.Limit > 0 {
			query["limit"] = fmt.Sprintf("%d", params.Limit)
		}
	}

	data, err := c.RequestWithQuery(ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return nil, err
	}

	var messages []types.Message
	if err := json.Unmarshal(data, &messages); err != nil {
		return nil, fmt.Errorf("unmarshal messages: %w", err)
	}

	return messages, nil
}

// EditMessageParams contains parameters for editing a message.
type EditMessageParams struct {
	// Content is the new message content.
	Content *string `json:"content,omitempty"`
	// Embeds are the new embeds.
	Embeds *[]types.Embed `json:"embeds,omitempty"`
	// Flags are the new message flags.
	Flags *types.MessageFlags `json:"flags,omitempty"`
	// AllowedMentions controls mention behavior.
	AllowedMentions *types.AllowedMentions `json:"allowed_mentions,omitempty"`
	// Components are the new message components.
	Components *types.ComponentList `json:"components,omitempty"`
	// Attachments are the new attachments.
	Attachments *[]types.Attachment `json:"attachments,omitempty"`
}

// EditMessage edits a message.
func (c *Client) EditMessage(ctx context.Context, channelID, messageID types.Snowflake, params *EditMessageParams) (*types.Message, error) {
	path := fmt.Sprintf("/channels/%s/messages/%s", channelID, messageID)
	data, err := c.Request(ctx, http.MethodPatch, path, params)
	if err != nil {
		return nil, err
	}

	var message types.Message
	if err := json.Unmarshal(data, &message); err != nil {
		return nil, fmt.Errorf("unmarshal message: %w", err)
	}

	return &message, nil
}

// DeleteMessage deletes a message.
func (c *Client) DeleteMessage(ctx context.Context, channelID, messageID types.Snowflake) error {
	path := fmt.Sprintf("/channels/%s/messages/%s", channelID, messageID)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// BulkDeleteMessages deletes multiple messages at once.
func (c *Client) BulkDeleteMessages(ctx context.Context, channelID types.Snowflake, messageIDs []types.Snowflake) error {
	path := fmt.Sprintf("/channels/%s/messages/bulk-delete", channelID)
	params := map[string][]types.Snowflake{"messages": messageIDs}
	_, err := c.Request(ctx, http.MethodPost, path, params)
	return err
}

// Reactions

// CreateReaction adds a reaction to a message.
func (c *Client) CreateReaction(ctx context.Context, channelID, messageID types.Snowflake, emoji string) error {
	path := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/@me", channelID, messageID, emoji)
	_, err := c.Request(ctx, http.MethodPut, path, nil)
	return err
}

// DeleteOwnReaction removes the current user's reaction.
func (c *Client) DeleteOwnReaction(ctx context.Context, channelID, messageID types.Snowflake, emoji string) error {
	path := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/@me", channelID, messageID, emoji)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// DeleteUserReaction removes another user's reaction.
func (c *Client) DeleteUserReaction(ctx context.Context, channelID, messageID, userID types.Snowflake, emoji string) error {
	path := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/%s", channelID, messageID, emoji, userID)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// Pins

// PinMessage pins a message in a channel.
func (c *Client) PinMessage(ctx context.Context, channelID, messageID types.Snowflake) error {
	path := fmt.Sprintf("/channels/%s/pins/%s", channelID, messageID)
	_, err := c.Request(ctx, http.MethodPut, path, nil)
	return err
}

// UnpinMessage unpins a message in a channel.
func (c *Client) UnpinMessage(ctx context.Context, channelID, messageID types.Snowflake) error {
	path := fmt.Sprintf("/channels/%s/pins/%s", channelID, messageID)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// GetChannelPins returns the pinned messages in a channel.
func (c *Client) GetChannelPins(ctx context.Context, channelID types.Snowflake) ([]types.Message, error) {
	path := fmt.Sprintf("/channels/%s/pins", channelID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var messages []types.Message
	if err := json.Unmarshal(data, &messages); err != nil {
		return nil, fmt.Errorf("unmarshal pins: %w", err)
	}

	return messages, nil
}

// Interactions

// CreateInteractionResponse responds to an interaction.
func (c *Client) CreateInteractionResponse(ctx context.Context, interactionID types.Snowflake, interactionToken string, response *types.InteractionResponse) error {
	path := fmt.Sprintf("/interactions/%s/%s/callback", interactionID, interactionToken)
	_, err := c.Request(ctx, http.MethodPost, path, response)
	return err
}

// GetOriginalInteractionResponse gets the original interaction response.
func (c *Client) GetOriginalInteractionResponse(ctx context.Context, applicationID types.Snowflake, interactionToken string) (*types.Message, error) {
	path := fmt.Sprintf("/webhooks/%s/%s/messages/@original", applicationID, interactionToken)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var message types.Message
	if err := json.Unmarshal(data, &message); err != nil {
		return nil, fmt.Errorf("unmarshal message: %w", err)
	}

	return &message, nil
}

// EditOriginalInteractionResponse edits the original interaction response.
func (c *Client) EditOriginalInteractionResponse(ctx context.Context, applicationID types.Snowflake, interactionToken string, params *EditMessageParams) (*types.Message, error) {
	path := fmt.Sprintf("/webhooks/%s/%s/messages/@original", applicationID, interactionToken)
	data, err := c.Request(ctx, http.MethodPatch, path, params)
	if err != nil {
		return nil, err
	}

	var message types.Message
	if err := json.Unmarshal(data, &message); err != nil {
		return nil, fmt.Errorf("unmarshal message: %w", err)
	}

	return &message, nil
}

// DeleteOriginalInteractionResponse deletes the original interaction response.
func (c *Client) DeleteOriginalInteractionResponse(ctx context.Context, applicationID types.Snowflake, interactionToken string) error {
	path := fmt.Sprintf("/webhooks/%s/%s/messages/@original", applicationID, interactionToken)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// CreateFollowupMessage creates a followup message for an interaction.
func (c *Client) CreateFollowupMessage(ctx context.Context, applicationID types.Snowflake, interactionToken string, params *CreateMessageParams) (*types.Message, error) {
	path := fmt.Sprintf("/webhooks/%s/%s", applicationID, interactionToken)
	data, err := c.Request(ctx, http.MethodPost, path, params)
	if err != nil {
		return nil, err
	}

	var message types.Message
	if err := json.Unmarshal(data, &message); err != nil {
		return nil, fmt.Errorf("unmarshal message: %w", err)
	}

	return &message, nil
}

// Application Commands

// CreateGlobalApplicationCommand creates a global slash command.
func (c *Client) CreateGlobalApplicationCommand(ctx context.Context, applicationID types.Snowflake, params *types.CreateApplicationCommandParams) (*types.ApplicationCommand, error) {
	path := fmt.Sprintf("/applications/%s/commands", applicationID)
	data, err := c.Request(ctx, http.MethodPost, path, params)
	if err != nil {
		return nil, err
	}

	var command types.ApplicationCommand
	if err := json.Unmarshal(data, &command); err != nil {
		return nil, fmt.Errorf("unmarshal command: %w", err)
	}

	return &command, nil
}

// GetGlobalApplicationCommands returns all global slash commands.
func (c *Client) GetGlobalApplicationCommands(ctx context.Context, applicationID types.Snowflake) ([]types.ApplicationCommand, error) {
	path := fmt.Sprintf("/applications/%s/commands", applicationID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var commands []types.ApplicationCommand
	if err := json.Unmarshal(data, &commands); err != nil {
		return nil, fmt.Errorf("unmarshal commands: %w", err)
	}

	return commands, nil
}

// GetGlobalApplicationCommand returns a global slash command by ID.
func (c *Client) GetGlobalApplicationCommand(ctx context.Context, applicationID, commandID types.Snowflake) (*types.ApplicationCommand, error) {
	path := fmt.Sprintf("/applications/%s/commands/%s", applicationID, commandID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var command types.ApplicationCommand
	if err := json.Unmarshal(data, &command); err != nil {
		return nil, fmt.Errorf("unmarshal command: %w", err)
	}

	return &command, nil
}

// EditGlobalApplicationCommand edits a global slash command.
func (c *Client) EditGlobalApplicationCommand(ctx context.Context, applicationID, commandID types.Snowflake, params *types.CreateApplicationCommandParams) (*types.ApplicationCommand, error) {
	path := fmt.Sprintf("/applications/%s/commands/%s", applicationID, commandID)
	data, err := c.Request(ctx, http.MethodPatch, path, params)
	if err != nil {
		return nil, err
	}

	var command types.ApplicationCommand
	if err := json.Unmarshal(data, &command); err != nil {
		return nil, fmt.Errorf("unmarshal command: %w", err)
	}

	return &command, nil
}

// DeleteGlobalApplicationCommand deletes a global slash command.
func (c *Client) DeleteGlobalApplicationCommand(ctx context.Context, applicationID, commandID types.Snowflake) error {
	path := fmt.Sprintf("/applications/%s/commands/%s", applicationID, commandID)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// BulkOverwriteGlobalApplicationCommands overwrites all global slash commands.
func (c *Client) BulkOverwriteGlobalApplicationCommands(ctx context.Context, applicationID types.Snowflake, commands []types.CreateApplicationCommandParams) ([]types.ApplicationCommand, error) {
	path := fmt.Sprintf("/applications/%s/commands", applicationID)
	data, err := c.Request(ctx, http.MethodPut, path, commands)
	if err != nil {
		return nil, err
	}

	var newCommands []types.ApplicationCommand
	if err := json.Unmarshal(data, &newCommands); err != nil {
		return nil, fmt.Errorf("unmarshal commands: %w", err)
	}

	return newCommands, nil
}

// CreateGuildApplicationCommand creates a guild slash command.
func (c *Client) CreateGuildApplicationCommand(ctx context.Context, applicationID, guildID types.Snowflake, params *types.CreateApplicationCommandParams) (*types.ApplicationCommand, error) {
	path := fmt.Sprintf("/applications/%s/guilds/%s/commands", applicationID, guildID)
	data, err := c.Request(ctx, http.MethodPost, path, params)
	if err != nil {
		return nil, err
	}

	var command types.ApplicationCommand
	if err := json.Unmarshal(data, &command); err != nil {
		return nil, fmt.Errorf("unmarshal command: %w", err)
	}

	return &command, nil
}

// GetGuildApplicationCommands returns all guild slash commands.
func (c *Client) GetGuildApplicationCommands(ctx context.Context, applicationID, guildID types.Snowflake) ([]types.ApplicationCommand, error) {
	path := fmt.Sprintf("/applications/%s/guilds/%s/commands", applicationID, guildID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var commands []types.ApplicationCommand
	if err := json.Unmarshal(data, &commands); err != nil {
		return nil, fmt.Errorf("unmarshal commands: %w", err)
	}

	return commands, nil
}

// GetGuildApplicationCommand returns a guild slash command by ID.
func (c *Client) GetGuildApplicationCommand(ctx context.Context, applicationID, guildID, commandID types.Snowflake) (*types.ApplicationCommand, error) {
	path := fmt.Sprintf("/applications/%s/guilds/%s/commands/%s", applicationID, guildID, commandID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var command types.ApplicationCommand
	if err := json.Unmarshal(data, &command); err != nil {
		return nil, fmt.Errorf("unmarshal command: %w", err)
	}

	return &command, nil
}

// EditGuildApplicationCommand edits a guild slash command.
func (c *Client) EditGuildApplicationCommand(ctx context.Context, applicationID, guildID, commandID types.Snowflake, params *types.CreateApplicationCommandParams) (*types.ApplicationCommand, error) {
	path := fmt.Sprintf("/applications/%s/guilds/%s/commands/%s", applicationID, guildID, commandID)
	data, err := c.Request(ctx, http.MethodPatch, path, params)
	if err != nil {
		return nil, err
	}

	var command types.ApplicationCommand
	if err := json.Unmarshal(data, &command); err != nil {
		return nil, fmt.Errorf("unmarshal command: %w", err)
	}

	return &command, nil
}

// DeleteGuildApplicationCommand deletes a guild slash command.
func (c *Client) DeleteGuildApplicationCommand(ctx context.Context, applicationID, guildID, commandID types.Snowflake) error {
	path := fmt.Sprintf("/applications/%s/guilds/%s/commands/%s", applicationID, guildID, commandID)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// BulkOverwriteGuildApplicationCommands overwrites all guild slash commands.
func (c *Client) BulkOverwriteGuildApplicationCommands(ctx context.Context, applicationID, guildID types.Snowflake, commands []types.CreateApplicationCommandParams) ([]types.ApplicationCommand, error) {
	path := fmt.Sprintf("/applications/%s/guilds/%s/commands", applicationID, guildID)
	data, err := c.Request(ctx, http.MethodPut, path, commands)
	if err != nil {
		return nil, err
	}

	var newCommands []types.ApplicationCommand
	if err := json.Unmarshal(data, &newCommands); err != nil {
		return nil, fmt.Errorf("unmarshal commands: %w", err)
	}

	return newCommands, nil
}

// Users

// GetCurrentUser returns the current user.
func (c *Client) GetCurrentUser(ctx context.Context) (*types.User, error) {
	data, err := c.Request(ctx, http.MethodGet, "/users/@me", nil)
	if err != nil {
		return nil, err
	}

	var user types.User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, fmt.Errorf("unmarshal user: %w", err)
	}

	return &user, nil
}

// GetUser returns a user by ID.
func (c *Client) GetUser(ctx context.Context, userID types.Snowflake) (*types.User, error) {
	path := fmt.Sprintf("/users/%s", userID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var user types.User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, fmt.Errorf("unmarshal user: %w", err)
	}

	return &user, nil
}

// Gateway

// GetGateway returns the WebSocket URL for the gateway.
func (c *Client) GetGateway(ctx context.Context) (string, error) {
	data, err := c.Request(ctx, http.MethodGet, "/gateway", nil)
	if err != nil {
		return "", err
	}

	var response struct {
		URL string `json:"url"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return "", fmt.Errorf("unmarshal gateway: %w", err)
	}

	return response.URL, nil
}

// GetGatewayBot returns the WebSocket URL and shard info for bot gateway.
func (c *Client) GetGatewayBot(ctx context.Context) (*GatewayBotResponse, error) {
	data, err := c.Request(ctx, http.MethodGet, "/gateway/bot", nil)
	if err != nil {
		return nil, err
	}

	var response GatewayBotResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unmarshal gateway bot: %w", err)
	}

	return &response, nil
}

// GatewayBotResponse contains gateway information for bots.
type GatewayBotResponse struct {
	// URL is the WebSocket URL.
	URL string `json:"url"`
	// Shards is the recommended number of shards.
	Shards int `json:"shards"`
	// SessionStartLimit contains session start limit info.
	SessionStartLimit SessionStartLimit `json:"session_start_limit"`
}

// SessionStartLimit contains session start limit information.
type SessionStartLimit struct {
	// Total is the total number of session starts allowed.
	Total int `json:"total"`
	// Remaining is the remaining number of session starts.
	Remaining int `json:"remaining"`
	// ResetAfter is milliseconds until the limit resets.
	ResetAfter int `json:"reset_after"`
	// MaxConcurrency is the max concurrent identify requests.
	MaxConcurrency int `json:"max_concurrency"`
}
