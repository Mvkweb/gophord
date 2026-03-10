# Containers

> Group components with accent colors and visual organization.

## Overview

Containers group components together and draw attention with a colored accent bar on their left edge. They are used exclusively as a visual grouping tool to make your messages look more structured. 

## Detailed Usage

```go
package main

import (
	"context"
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/types"
)

func sendContainerDemo(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
	// 1. Define an accent color (integer hex)
	successColor := 0x57F287 // Green

	// 2. Create the container
	container := &types.Container{
		AccentColor: &successColor,
		Components: types.ComponentList{
			&types.TextDisplay{Content: "✅ **Success!**\nOperation completed successfully!"},
			&types.ActionRow{
				Components: types.ComponentList{
					client.NewPrimaryButton("container_confirm", "Okay"),
				},
			},
		},
	}

	// 3. Send
	_, err := bot.SendMessageWithComponents(ctx, channelID, types.ComponentList{container})
	if err != nil {
		log.Printf("Failed to send container: %v", err)
	}
}
```

## API Reference

### Struct Definition

`types.Container` is constructed using manual structs.

| Field | Type | Required? | Description |
|---|---|---|---|
| `AccentColor` | `*int` | No | A hexadecimal color integer (e.g. `0x5865F2`). A pointer is used to differentiate between no color and `0x000000` (black). |
| `Components` | `types.ComponentList` | Yes | What components go inside this container. |

## Related

- [Components V2](components-v2.md)
- [Sections](sections.md)
