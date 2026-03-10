# Components V2

> Interactive UI elements inside messages.

## Overview

Components V2 represent the modern Discord UI: Action Rows, Buttons, Select Menus, Text Inputs, Containers, and Sections. In Gophord, every component implements the `types.Component` interface, meaning you can nest them appropriately in lists.

## Top-Level Layout: Action Rows

Every interactive component MUST be wrapped inside an `ActionRow`. You cannot attach a raw button to a message; you attach an Action Row that contains buttons.

A single message can have up to 5 Action Rows.
A single Action Row can have up to:
- 5 Buttons OR
- 1 Select Menu

## The Component Interface

`types.Component` is an interface with a `Type()` method. This allows slices like `types.ComponentList` to hold a mix of `*types.ActionRow`, `*types.Button`, and `*types.StringSelect`.

```go
package main

import (
	"log"

	"github.com/Mvkweb/gophord/pkg/rest"
	"github.com/Mvkweb/gophord/pkg/types"
)

func sendComponentsMessage(restClient *rest.Client, channelID string) {
	button1 := &types.Button{
		Style:    types.ButtonStylePrimary,
		Label:    "Accept",
		CustomID: "btn_accept",
	}

	button2 := &types.Button{
		Style:    types.ButtonStyleDanger,
		Label:    "Decline",
		CustomID: "btn_decline",
	}

	row := &types.ActionRow{
		Components: types.ComponentList{button1, button2},
	}

	_, err := restClient.CreateMessage(channelID, rest.CreateMessageParams{
		Content:    "Do you accept the terms?",
		Components: types.ComponentList{row},
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
```

## Parsing Polymorphic JSON

Gophord handles the polymorphic JSON unmarshalling for you. When you receive a `types.Message` from an event (like `gateway.MessageCreateEvent`), the `Components` slice will be properly instantiated as `*types.ActionRow` objects containing `*types.Button`s or `*types.StringSelect`s.

You can safely type-assert them:

```go
func parseComponents(msg *types.Message) {
	for _, comp := range msg.Components {
		if row, ok := comp.(*types.ActionRow); ok {
			for _, inner := range row.Components {
				switch c := inner.(type) {
				case *types.Button:
					log.Printf("Found button: %s", c.CustomID)
				case *types.StringSelect:
					log.Printf("Found select: %s", c.CustomID)
				}
			}
		}
	}
}
```

## Component Types Reference

| Component Type | Struct | Description |
|---|---|---|
| `types.ComponentTypeActionRow` | `*types.ActionRow` | Top-level container. Required wrapper. |
| `types.ComponentTypeButton` | `*types.Button` | Clickable buttons. |
| `types.ComponentTypeStringSelect` | `*types.StringSelect` | Dropdown menus for selecting text. |
| `types.ComponentTypeText` | `*types.TextInput` | Input fields (only valid in Modals). |
| `types.ComponentTypeUserSelect` | `*types.UserSelect` | Auto-populated dropdown of users. |
| `types.ComponentTypeRoleSelect` | `*types.RoleSelect` | Auto-populated dropdown of roles. |
| `types.ComponentTypeMentionableSelect` | `*types.MentionableSelect` | Auto-populated dropdown of users AND roles. |
| `types.ComponentTypeChannelSelect` | `*types.ChannelSelect` | Auto-populated dropdown of channels. |

## Sub-Pages

- [Buttons](buttons.md)
- [Select Menus](select-menus.md)
- [Containers](containers.md)
- [Sections](sections.md)
- [Media Galleries](media-galleries.md)
- [Modals](modals.md)
