# Select Menus

> Dropdown menus for selecting text, users, roles, or channels.

## Overview

Select Menus allow users to pick one or more options from a dropdown. Like buttons, they must be wrapped in an `ActionRow`. However, **an `ActionRow` can only contain exactly one Select Menu.**

Gophord supports 4 flavors of Select Menus:
- **String Select**: You provide an explicit list of hardcoded `types.SelectOption` choices.
- **User Select**: Discord auto-populates the list with server members.
- **Role Select**: Discord auto-populates the list with server roles.
- **Channel Select**: Discord auto-populates the list with channels (filterable by type).

## Detailed Usage

Use `client.ComponentBuilder` (or manual structs) to send selects.

```go
package main

import (
	"context"
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/types"
)

func sendSelectsDemo(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
	components := client.NewComponentBuilder().
		// 1. String Select - Custom options
		AddActionRow(&types.StringSelect{
			CustomID:    "language_select",
			Placeholder: "Choose your favorite language",
			Options: []types.SelectOption{
				{Label: "Go", Value: "go", Description: "Simple and efficient", Emoji: &types.PartialEmoji{Name: "🐹"}},
				{Label: "Rust", Value: "rust", Description: "Memory safety", Emoji: &types.PartialEmoji{Name: "🦀"}},
			},
		}).
		// 2. User Select - Auto-populated with server members
		AddActionRow(&types.UserSelect{
			CustomID:    "user_select",
			Placeholder: "Select a user",
		}).
		// 3. Channel Select - Filtered to text channels only
		AddActionRow(&types.ChannelSelect{
			CustomID:     "channel_select",
			Placeholder:  "Select a text channel",
			ChannelTypes: []types.ChannelType{types.ChannelTypeGuildText},
		}).
		Build()

	_, err := bot.SendMessageWithComponents(ctx, channelID, components)
	if err != nil {
		log.Printf("Failed to send selects: %v", err)
	}
}
```

## Handling Selections

When a user makes a choice, Discord fires a `MessageComponent` Interaction. The chosen values are stored in the slice `interaction.Data.Values`.

```go
func handleSelectInteraction(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	customID := interaction.Data.CustomID
	
	// Values is a []string. Even for user/role selects, it contains Snowflake IDs as strings.
	values := interaction.Data.Values

	switch customID {
	case "language_select":
		// Handle string choices
		bot.RespondWithMessage(ctx, interaction, "You chose: " + values[0])
	case "user_select":
		// Handle user pick 
		bot.RespondWithMessage(ctx, interaction, "You pinged <@!" + values[0] + ">")
	}
}
```

## API Reference

### Component Structs

| Struct | Property | Description |
|---|---|---|
| `types.StringSelect` | `Options []types.SelectOption` | Provide 1-25 choices manually. |
| `types.UserSelect` | `DefaultValues []types.SelectDefaultValue` | Users. |
| `types.RoleSelect` | `DefaultValues []types.SelectDefaultValue` | Roles. |
| `types.ChannelSelect` | `ChannelTypes []types.ChannelType` | Channels. Can restrict to voice/text/etc. |
| `types.MentionableSelect` | `DefaultValues []types.SelectDefaultValue` | Users AND Roles. |

### Common Properties

All selects share these optional properties:

| Field | Type | Description |
|---|---|---|
| `CustomID` | `string` | **Required.** Identifier returned in the interaction. (Max 100 char) |
| `Placeholder` | `string` | Text shown if nothing is selected. (Max 150 char) |
| `MinValues` | `*int` | Minimum items a user must select. Defaults to 1 (ptr). |
| `MaxValues` | `*int` | Maximum items a user can select. Max 25. Defaults to 1 (ptr). |
| `Disabled` | `bool` | Disables the select menu. |

## Related

- [Components V2 Overview](components-v2.md) — Top level concepts like Action Rows.
- [Interactions](interactions.md) — Responding to the interaction event.
