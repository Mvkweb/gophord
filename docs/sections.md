# Sections

> Create content with thumbnail or button accessories.

## Overview

Sections pair text content with an accessory on the right side. The accessory can be an image (Thumbnail) or an interactive element (Button).

## Detailed Usage

You can create sections by instantiating a `types.Section` directly.

```go
package main

import (
	"context"
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/types"
)

func sendSectionDemo(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
	// Scenario A: A Product-style section with an image thumbnail
	productSection := &types.Section{
		Components: []types.Component{
			&types.TextDisplay{
				Content: "**Gophord Pro** 🚀\nHigh-performance Discord API wrapper\n• Blazing fast JSON parsing\n• Native Components V2",
			},
		},
		Accessory: &types.Thumbnail{
			Media: types.UnfurledMediaItem{URL: "https://picsum.photos/80/80?random=2"},
		},
	}

	// Scenario B: An interactive row with an action button
	buttonSection := &types.Section{
		Components: []types.Component{
			&types.TextDisplay{Content: "**With Button**\nThis section has an action button on the right."},
		},
		Accessory: &types.Button{
			Style:    types.ButtonStylePrimary,
			CustomID: "section_action",
			Label:    "Action",
		},
	}

	_, err := bot.SendMessageWithComponents(ctx, channelID, []types.Component{productSection, buttonSection})
	if err != nil {
		log.Printf("Failed to send sections: %v", err)
	}
}
```

## API Reference

### Struct Definition

`types.Section` is constructed manually.

| Field | Type | Required? | Description |
|---|---|---|---|
| `Components` | `[]types.Component` | No* | Content on the left. Must contain at least a `TextDisplay`. |
| `Accessory` | `types.Component` | No | Content on the right. Usually a `*types.Button` or `*types.Thumbnail`. |

## Related

- [Components V2](components-v2.md)
- [Containers](containers.md)
