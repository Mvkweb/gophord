// Package rest provides a REST client for the Discord API.
package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Mvkweb/gophord/pkg/json"
	"github.com/Mvkweb/gophord/pkg/types"
)

// CreateGuildParams contains parameters for creating a guild.
type CreateGuildParams struct {
	Name                        string                      `json:"name"`
	Region                      string                      `json:"region,omitempty"`
	Icon                        string                      `json:"icon,omitempty"`
	VerificationLevel           int                         `json:"verification_level,omitempty"`
	DefaultMessageNotifications int                         `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       int                         `json:"explicit_content_filter,omitempty"`
	Roles                       []types.Role                `json:"roles,omitempty"`
	Channels                    []types.CreateChannelParams `json:"channels,omitempty"`
	AFKChannelID                types.Snowflake             `json:"afk_channel_id,omitempty"`
	AFKTimeout                  int                         `json:"afk_timeout,omitempty"`
	SystemChannelID             types.Snowflake             `json:"system_channel_id,omitempty"`
	SystemChannelFlags          int                         `json:"system_channel_flags,omitempty"`
}

// CreateGuild creates a new guild.
func (c *Client) CreateGuild(ctx context.Context, params *CreateGuildParams) (*types.Guild, error) {
	data, err := c.Request(ctx, http.MethodPost, "/guilds", params)
	if err != nil {
		return nil, err
	}

	var guild types.Guild
	if err := json.Unmarshal(data, &guild); err != nil {
		return nil, fmt.Errorf("unmarshal guild: %w", err)
	}

	return &guild, nil
}

// GetGuild returns a guild by ID.
func (c *Client) GetGuild(ctx context.Context, guildID types.Snowflake) (*types.Guild, error) {
	path := fmt.Sprintf("/guilds/%s", guildID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var guild types.Guild
	if err := json.Unmarshal(data, &guild); err != nil {
		return nil, fmt.Errorf("unmarshal guild: %w", err)
	}

	return &guild, nil
}

// GetGuildPreview returns a guild preview by ID.
func (c *Client) GetGuildPreview(ctx context.Context, guildID types.Snowflake) (*types.GuildPreview, error) {
	path := fmt.Sprintf("/guilds/%s/preview", guildID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var preview types.GuildPreview
	if err := json.Unmarshal(data, &preview); err != nil {
		return nil, fmt.Errorf("unmarshal guild preview: %w", err)
	}

	return &preview, nil
}

// ModifyGuildParams contains parameters for modifying a guild.
type ModifyGuildParams struct {
	Name                        string          `json:"name,omitempty"`
	Region                      string          `json:"region,omitempty"`
	VerificationLevel           *int            `json:"verification_level,omitempty"`
	DefaultMessageNotifications *int            `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       *int            `json:"explicit_content_filter,omitempty"`
	AFKChannelID                types.Snowflake `json:"afk_channel_id,omitempty"`
	AFKTimeout                  int             `json:"afk_timeout,omitempty"`
	Icon                        string          `json:"icon,omitempty"`
	OwnerID                     types.Snowflake `json:"owner_id,omitempty"`
	Splash                      string          `json:"splash,omitempty"`
	DiscoverySplash             string          `json:"discovery_splash,omitempty"`
	Banner                      string          `json:"banner,omitempty"`
	SystemChannelID             types.Snowflake `json:"system_channel_id,omitempty"`
	SystemChannelFlags          int             `json:"system_channel_flags,omitempty"`
	RulesChannelID              types.Snowflake `json:"rules_channel_id,omitempty"`
	PublicUpdatesChannelID      types.Snowflake `json:"public_updates_channel_id,omitempty"`
	PreferredLocale             string          `json:"preferred_locale,omitempty"`
	Features                    []string        `json:"features,omitempty"`
	Description                 string          `json:"description,omitempty"`
	PremiumProgressBarEnabled   bool            `json:"premium_progress_bar_enabled,omitempty"`
}

// ModifyGuild modifies a guild.
func (c *Client) ModifyGuild(ctx context.Context, guildID types.Snowflake, params *ModifyGuildParams) (*types.Guild, error) {
	path := fmt.Sprintf("/guilds/%s", guildID)
	data, err := c.Request(ctx, http.MethodPatch, path, params)
	if err != nil {
		return nil, err
	}

	var guild types.Guild
	if err := json.Unmarshal(data, &guild); err != nil {
		return nil, fmt.Errorf("unmarshal guild: %w", err)
	}

	return &guild, nil
}

// DeleteGuild deletes a guild.
func (c *Client) DeleteGuild(ctx context.Context, guildID types.Snowflake) error {
	path := fmt.Sprintf("/guilds/%s", guildID)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// GetGuildChannels returns a list of guild channels.
func (c *Client) GetGuildChannels(ctx context.Context, guildID types.Snowflake) ([]types.Channel, error) {
	path := fmt.Sprintf("/guilds/%s/channels", guildID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var channels []types.Channel
	if err := json.Unmarshal(data, &channels); err != nil {
		return nil, fmt.Errorf("unmarshal channels: %w", err)
	}

	return channels, nil
}

// CreateGuildChannel creates a new channel in a guild.
func (c *Client) CreateGuildChannel(ctx context.Context, guildID types.Snowflake, params *types.CreateChannelParams) (*types.Channel, error) {
	path := fmt.Sprintf("/guilds/%s/channels", guildID)
	data, err := c.Request(ctx, http.MethodPost, path, params)
	if err != nil {
		return nil, err
	}

	var channel types.Channel
	if err := json.Unmarshal(data, &channel); err != nil {
		return nil, fmt.Errorf("unmarshal channel: %w", err)
	}

	return &channel, nil
}

// ModifyGuildChannelPositionsParams contains parameters for modifying channel positions.
type ModifyGuildChannelPositionsParams struct {
	ID       types.Snowflake `json:"id"`
	Position *int            `json:"position,omitempty"`
	Lock     *bool           `json:"lock_permissions,omitempty"`
	ParentID types.Snowflake `json:"parent_id,omitempty"`
}

// ModifyGuildChannelPositions modifies the positions of channels in a guild.
func (c *Client) ModifyGuildChannelPositions(ctx context.Context, guildID types.Snowflake, params []ModifyGuildChannelPositionsParams) error {
	path := fmt.Sprintf("/guilds/%s/channels", guildID)
	_, err := c.Request(ctx, http.MethodPatch, path, params)
	return err
}

// GetGuildMember returns a guild member by user ID.
func (c *Client) GetGuildMember(ctx context.Context, guildID, userID types.Snowflake) (*types.GuildMember, error) {
	path := fmt.Sprintf("/guilds/%s/members/%s", guildID, userID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var member types.GuildMember
	if err := json.Unmarshal(data, &member); err != nil {
		return nil, fmt.Errorf("unmarshal member: %w", err)
	}

	return &member, nil
}

// ListGuildMembersParams contains parameters for listing guild members.
type ListGuildMembersParams struct {
	Limit int             `json:"limit,omitempty"`
	After types.Snowflake `json:"after,omitempty"`
}

// ListGuildMembers returns a list of guild members.
func (c *Client) ListGuildMembers(ctx context.Context, guildID types.Snowflake, params *ListGuildMembersParams) ([]types.GuildMember, error) {
	path := fmt.Sprintf("/guilds/%s/members", guildID)

	query := make(map[string]string)
	if params != nil {
		if params.Limit > 0 {
			query["limit"] = fmt.Sprintf("%d", params.Limit)
		}
		if params.After != 0 {
			query["after"] = params.After.String()
		}
	}

	data, err := c.RequestWithQuery(ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return nil, err
	}

	var members []types.GuildMember
	if err := json.Unmarshal(data, &members); err != nil {
		return nil, fmt.Errorf("unmarshal members: %w", err)
	}

	return members, nil
}

// AddGuildMemberParams contains parameters for adding a guild member.
type AddGuildMemberParams struct {
	AccessToken string            `json:"access_token"`
	Nick        string            `json:"nick,omitempty"`
	Roles       []types.Snowflake `json:"roles,omitempty"`
	Mute        bool              `json:"mute,omitempty"`
	Deaf        bool              `json:"deaf,omitempty"`
}

// AddGuildMember adds a user to the guild (requires OAuth2 access token).
func (c *Client) AddGuildMember(ctx context.Context, guildID, userID types.Snowflake, params *AddGuildMemberParams) (*types.GuildMember, error) {
	path := fmt.Sprintf("/guilds/%s/members/%s", guildID, userID)
	data, err := c.Request(ctx, http.MethodPut, path, params)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil // 204 No Content (already joined)
	}

	var member types.GuildMember
	if err := json.Unmarshal(data, &member); err != nil {
		return nil, fmt.Errorf("unmarshal member: %w", err)
	}

	return &member, nil
}

// ModifyGuildMemberParams contains parameters for modifying a guild member.
type ModifyGuildMemberParams struct {
	Nick                       *string           `json:"nick,omitempty"`
	Roles                      []types.Snowflake `json:"roles,omitempty"`
	Mute                       *bool             `json:"mute,omitempty"`
	Deaf                       *bool             `json:"deaf,omitempty"`
	ChannelID                  *types.Snowflake  `json:"channel_id,omitempty"`
	CommunicationDisabledUntil *types.Timestamp  `json:"communication_disabled_until,omitempty"`
	Flags                      *int              `json:"flags,omitempty"`
}

// ModifyGuildMember modifies a guild member.
func (c *Client) ModifyGuildMember(ctx context.Context, guildID, userID types.Snowflake, params *ModifyGuildMemberParams) (*types.GuildMember, error) {
	path := fmt.Sprintf("/guilds/%s/members/%s", guildID, userID)
	data, err := c.Request(ctx, http.MethodPatch, path, params)
	if err != nil {
		return nil, err
	}

	var member types.GuildMember
	if err := json.Unmarshal(data, &member); err != nil {
		return nil, fmt.Errorf("unmarshal member: %w", err)
	}

	return &member, nil
}

// ModifyCurrentUserNick modifies the current user's nickname in a guild.
func (c *Client) ModifyCurrentUserNick(ctx context.Context, guildID types.Snowflake, nick string) error {
	path := fmt.Sprintf("/guilds/%s/members/@me", guildID)
	_, err := c.Request(ctx, http.MethodPatch, path, map[string]string{"nick": nick})
	return err
}

// AddGuildMemberRole adds a role to a guild member.
func (c *Client) AddGuildMemberRole(ctx context.Context, guildID, userID, roleID types.Snowflake) error {
	path := fmt.Sprintf("/guilds/%s/members/%s/roles/%s", guildID, userID, roleID)
	_, err := c.Request(ctx, http.MethodPut, path, nil)
	return err
}

// RemoveGuildMemberRole removes a role from a guild member.
func (c *Client) RemoveGuildMemberRole(ctx context.Context, guildID, userID, roleID types.Snowflake) error {
	path := fmt.Sprintf("/guilds/%s/members/%s/roles/%s", guildID, userID, roleID)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// RemoveGuildMember kicks a member from the guild.
func (c *Client) RemoveGuildMember(ctx context.Context, guildID, userID types.Snowflake) error {
	path := fmt.Sprintf("/guilds/%s/members/%s", guildID, userID)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// GetGuildBans returns a list of guild bans.
func (c *Client) GetGuildBans(ctx context.Context, guildID types.Snowflake, limit int, before, after types.Snowflake) ([]types.Ban, error) {
	path := fmt.Sprintf("/guilds/%s/bans", guildID)
	query := make(map[string]string)
	if limit > 0 {
		query["limit"] = fmt.Sprintf("%d", limit)
	}
	if before != 0 {
		query["before"] = before.String()
	}
	if after != 0 {
		query["after"] = after.String()
	}

	data, err := c.RequestWithQuery(ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return nil, err
	}

	var bans []types.Ban
	if err := json.Unmarshal(data, &bans); err != nil {
		return nil, fmt.Errorf("unmarshal bans: %w", err)
	}

	return bans, nil
}

// GetGuildBan returns a guild ban by user ID.
func (c *Client) GetGuildBan(ctx context.Context, guildID, userID types.Snowflake) (*types.Ban, error) {
	path := fmt.Sprintf("/guilds/%s/bans/%s", guildID, userID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var ban types.Ban
	if err := json.Unmarshal(data, &ban); err != nil {
		return nil, fmt.Errorf("unmarshal ban: %w", err)
	}

	return &ban, nil
}

// CreateGuildBan adds a ban to the guild.
func (c *Client) CreateGuildBan(ctx context.Context, guildID, userID types.Snowflake, deleteMessageSeconds int) error {
	path := fmt.Sprintf("/guilds/%s/bans/%s", guildID, userID)
	params := map[string]int{"delete_message_seconds": deleteMessageSeconds}
	_, err := c.Request(ctx, http.MethodPut, path, params)
	return err
}

// RemoveGuildBan removes a ban from the guild.
func (c *Client) RemoveGuildBan(ctx context.Context, guildID, userID types.Snowflake) error {
	path := fmt.Sprintf("/guilds/%s/bans/%s", guildID, userID)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}

// GetGuildRoles returns a list of guild roles.
func (c *Client) GetGuildRoles(ctx context.Context, guildID types.Snowflake) ([]types.Role, error) {
	path := fmt.Sprintf("/guilds/%s/roles", guildID)
	data, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var roles []types.Role
	if err := json.Unmarshal(data, &roles); err != nil {
		return nil, fmt.Errorf("unmarshal roles: %w", err)
	}

	return roles, nil
}

// CreateGuildRoleParams contains parameters for creating a role.
type CreateGuildRoleParams struct {
	Name         string `json:"name,omitempty"`
	Permissions  string `json:"permissions,omitempty"`
	Color        int    `json:"color,omitempty"`
	Hoist        bool   `json:"hoist,omitempty"`
	Icon         string `json:"icon,omitempty"`
	UnicodeEmoji string `json:"unicode_emoji,omitempty"`
	Mentionable  bool   `json:"mentionable,omitempty"`
}

// CreateGuildRole creates a new role in the guild.
func (c *Client) CreateGuildRole(ctx context.Context, guildID types.Snowflake, params *CreateGuildRoleParams) (*types.Role, error) {
	path := fmt.Sprintf("/guilds/%s/roles", guildID)
	data, err := c.Request(ctx, http.MethodPost, path, params)
	if err != nil {
		return nil, err
	}

	var role types.Role
	if err := json.Unmarshal(data, &role); err != nil {
		return nil, fmt.Errorf("unmarshal role: %w", err)
	}

	return &role, nil
}

// ModifyGuildRolePositionParams contains parameters for modifying a role position.
type ModifyGuildRolePositionParams struct {
	ID       types.Snowflake `json:"id"`
	Position int             `json:"position,omitempty"`
}

// ModifyGuildRolePositions modifies the positions of roles in a guild.
func (c *Client) ModifyGuildRolePositions(ctx context.Context, guildID types.Snowflake, params []ModifyGuildRolePositionParams) ([]types.Role, error) {
	path := fmt.Sprintf("/guilds/%s/roles", guildID)
	data, err := c.Request(ctx, http.MethodPatch, path, params)
	if err != nil {
		return nil, err
	}

	var roles []types.Role
	if err := json.Unmarshal(data, &roles); err != nil {
		return nil, fmt.Errorf("unmarshal roles: %w", err)
	}

	return roles, nil
}

// ModifyGuildRole modifies a guild role.
func (c *Client) ModifyGuildRole(ctx context.Context, guildID, roleID types.Snowflake, params *CreateGuildRoleParams) (*types.Role, error) {
	path := fmt.Sprintf("/guilds/%s/roles/%s", guildID, roleID)
	data, err := c.Request(ctx, http.MethodPatch, path, params)
	if err != nil {
		return nil, err
	}

	var role types.Role
	if err := json.Unmarshal(data, &role); err != nil {
		return nil, fmt.Errorf("unmarshal role: %w", err)
	}

	return &role, nil
}

// DeleteGuildRole deletes a guild role.
func (c *Client) DeleteGuildRole(ctx context.Context, guildID, roleID types.Snowflake) error {
	path := fmt.Sprintf("/guilds/%s/roles/%s", guildID, roleID)
	_, err := c.Request(ctx, http.MethodDelete, path, nil)
	return err
}
