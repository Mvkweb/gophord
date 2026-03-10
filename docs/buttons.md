# Buttons

> Clickable buttons that trigger interactions or open links.

## Overview

Buttons must be placed inside an `ActionRow`. You can have up to 5 buttons per Action Row. There are 5 primary styles you can use: Primary, Secondary, Success, Danger, and Link.

## Detailed Usage

Using the `client.ComponentBuilder` makes it easy to construct complex button layouts without dealing with nested structs directly.

```go
package main

import (
    "context"
    "log"

    "github.com/Mvkweb/gophord/pkg/client"
    "github.com/Mvkweb/gophord/pkg/types"
)

func sendButtonsDemo(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
    components := client.NewComponentBuilder().
        AddTextDisplay("# Button Styles Demo").
        AddSeparator(true, types.SeparatorSpacingSmall).
        AddActionRow(
            // Primary button (Blurple) - requires CustomID and Label
            client.NewPrimaryButton("btn_primary", "Primary"),
            // Secondary button (Gray)
            client.NewSecondaryButton("btn_secondary", "Secondary"),
            // Success button (Green)
            client.NewSuccessButton("btn_success", "Success"),
            // Danger button (Red)
            client.NewDangerButton("btn_danger", "Danger"),
        ).
        AddActionRow(
            // Link button - Navigates instead of sending an interaction back
            client.NewLinkButton("https://github.com/Mvkweb/gophord", "GitHub"),
            // Creating a custom button struct for Emoji/Disabled states
            &types.Button{
                Style:    types.ButtonStylePrimary,
                CustomID: "btn_emoji",
                Label:    "With Emoji",
                Emoji:    &types.PartialEmoji{Name: "🚀"},
            },
            &types.Button{
                Style:    types.ButtonStyleSecondary,
                CustomID: "btn_disabled",
                Label:    "Disabled",
                Disabled: true,
            },
        ).
        Build()

    _, err := bot.SendMessageWithComponents(ctx, channelID, components)
    if err != nil {
        log.Printf("Failed to send buttons: %v", err)
    }
}
```

## Handling Clicks

When a user clicks a button (except a Link button), Discord sends an `InteractionTypeMessageComponent`.

```go
func handleButtonClick(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
    // The CustomID is inside interaction.Data
    customID := interaction.Data.CustomID

    switch customID {
    case "btn_primary":
        bot.RespondWithMessage(ctx, interaction, "You clicked the Primary button!")
    case "btn_danger":
        bot.RespondWithMessage(ctx, interaction, "Danger zone! ⚠️")
    }
}
```

## API Reference

### ComponentBuilder Methods

| Method | Signature | Description |
|---|---|---|
| `AddActionRow` | `func (b *ComponentBuilder) AddActionRow(components ...types.Component) *ComponentBuilder` | Adds a row containing the provided components. |

### Button Constructors (Helper Methods)

| Method | Signature | Generates |
|---|---|---|
| `NewPrimaryButton` | `func NewPrimaryButton(customID, label string) *types.Button` | `ButtonStylePrimary` |
| `NewSecondaryButton` | `func NewSecondaryButton(customID, label string) *types.Button` | `ButtonStyleSecondary` |
| `NewSuccessButton` | `func NewSuccessButton(customID, label string) *types.Button` | `ButtonStyleSuccess` |
| `NewDangerButton` | `func NewDangerButton(customID, label string) *types.Button` | `ButtonStyleDanger` |
| `NewLinkButton` | `func NewLinkButton(url, label string) *types.Button` | `ButtonStyleLink` |

### Struct Definition

`types.Button` contains the following configurable fields:

| Field | Type | Required? | Notes |
|---|---|---|---|
| `Style` | `types.ButtonStyle` | Yes | 1-6 |
| `CustomID` | `string` | Yes* | *Not allowed for Link buttons, required for all others. |
| `Label`| `string` | No* | *Required if `Emoji` is nil. Max 80 chars. |
| `URL` | `string` | Yes* | *Only for Link buttons. |
| `Emoji` | `*types.PartialEmoji` | No | Displayed alongside or instead of the label. |
| `Disabled` | `bool` | No | Defaults to false. |

## Related

- [Components V2 Overview](components-v2.md) — Top level concepts like Action Rows.
- [Interactions](interactions.md) — Responding to the interaction event.
