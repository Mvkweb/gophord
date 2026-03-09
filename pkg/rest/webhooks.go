// Package rest provides a REST client for the Discord API.
package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Mvkweb/gophord/pkg/json"
	"github.com/Mvkweb/gophord/pkg/types"
)

// CreateWebhook creates a new webhook.
func (c *Client) CreateWebhook(ctx context.Context, channelID types.Snowflake, params *types.CreateWebhookParams) (*types.Webhook, error) {
	path := fmt.Sprintf("/channels/%s/webhooks", channelID)
	data, err := c.Request(ctx, http.MethodPost, path, params)
	if err != nil {
		return nil, err
	}

	var webhook types.Webhook
	if err := json.Unmarshal(data, &webhook); err != nil {
		return nil, fmt.Errorf("unmarshal webhook: %w", err)
	}

	return &webhook, nil
}

// GetChannelWebhooks returns a list of webhooks for a channel.
func (c *Client) GetChannelWebhooks(ctx context.Context, channelID types.Snowflake) ([]types.Webhook, error) {
	path := fmt.Sprintf("/channels/%s/webhooks", channelID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var webhooks []types.Webhook
	if err := json.Unmarshal(data, &webhooks); err != nil {
		return nil, fmt.Errorf("unmarshal webhooks: %w", err)
	}

	return webhooks, nil
}

// GetGuildWebhooks returns a list of webhooks for a guild.
func (c *Client) GetGuildWebhooks(ctx context.Context, guildID types.Snowflake) ([]types.Webhook, error) {
	path := fmt.Sprintf("/guilds/%s/webhooks", guildID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var webhooks []types.Webhook
	if err := json.Unmarshal(data, &webhooks); err != nil {
		return nil, fmt.Errorf("unmarshal webhooks: %w", err)
	}

	return webhooks, nil
}

// GetWebhook returns a webhook by ID.
func (c *Client) GetWebhook(ctx context.Context, webhookID types.Snowflake) (*types.Webhook, error) {
	path := fmt.Sprintf("/webhooks/%s", webhookID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var webhook types.Webhook
	if err := json.Unmarshal(data, &webhook); err != nil {
		return nil, fmt.Errorf("unmarshal webhook: %w", err)
	}

	return &webhook, nil
}

// GetWebhookWithToken returns a webhook by ID and token (does not require authentication).
func (c *Client) GetWebhookWithToken(ctx context.Context, webhookID types.Snowflake, token string) (*types.Webhook, error) {
	path := fmt.Sprintf("/webhooks/%s/%s", webhookID, token)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var webhook types.Webhook
	if err := json.Unmarshal(data, &webhook); err != nil {
		return nil, fmt.Errorf("unmarshal webhook: %w", err)
	}

	return &webhook, nil
}

// ExecuteWebhook executes a webhook.
func (c *Client) ExecuteWebhook(ctx context.Context, webhookID types.Snowflake, token string, wait bool, params *types.ExecuteWebhookParams) (*types.Message, error) {
	path := fmt.Sprintf("/webhooks/%s/%s", webhookID, token)
	if wait {
		path += "?wait=true"
	}

	data, err := c.Request(ctx, http.MethodPost, path, params)
	if err != nil {
		return nil, err
	}

	if !wait {
		return nil, nil
	}

	var message types.Message
	if err := json.Unmarshal(data, &message); err != nil {
		return nil, fmt.Errorf("unmarshal message: %w", err)
	}

	return &message, nil
}

// DeleteWebhook deletes a webhook.
func (c *Client) DeleteWebhook(ctx context.Context, webhookID types.Snowflake) error {
	path := fmt.Sprintf("/webhooks/%s", webhookID)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// DeleteWebhookWithToken deletes a webhook with token.
func (c *Client) DeleteWebhookWithToken(ctx context.Context, webhookID types.Snowflake, token string) error {
	path := fmt.Sprintf("/webhooks/%s/%s", webhookID, token)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}
