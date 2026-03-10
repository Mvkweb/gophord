# Media Galleries

> Display multiple images in a responsive grid layout.

## Overview

Media Galleries display multiple images in an automatic grid layout natively. This avoids spamming a channel with multiple attachments or needing a collage builder.

## Detailed Usage

Images must use valid URLs pointing either to the open internet (via Discord's CDN proxy) or as an attachment in the same message.

```go
package main

import (
	"context"
	"log"

	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/types"
)

func sendGalleryDemo(ctx context.Context, bot *client.Client, channelID types.Snowflake) {
	// Add media gallery with multiple items
	gallery := &types.MediaGallery{
		Items: []types.MediaGalleryItem{
			{
				Media:       types.UnfurledMediaItem{URL: "https://picsum.photos/400/300?random=1"},
				Description: "Random landscape image 1",
			},
			{
				Media:       types.UnfurledMediaItem{URL: "https://picsum.photos/400/300?random=2"},
				Description: "Random landscape image 2",
			},
			{
				Media:       types.UnfurledMediaItem{URL: "https://picsum.photos/400/300?random=3"},
				Description: "Random landscape image 3",
			},
		},
	}

	_, err := bot.SendMessageWithComponents(ctx, channelID, types.ComponentList{gallery})
	if err != nil {
		log.Printf("Failed to send gallery: %v", err)
	}
}
```

## API Reference

### Struct Definitions

| Struct | Description |
|---|---|
| `types.MediaGallery` | Top level component grouping images. Contains exactly 1 field: `Items []types.MediaGalleryItem` |
| `types.MediaGalleryItem` | An individual item representing one grid cell. |
| `types.UnfurledMediaItem` | Specifically required by `Media` inside an item. Provides the `URL`. |

### Fields inside `MediaGalleryItem`

| Field | Type | Description |
|---|---|---|
| `Media` | `types.UnfurledMediaItem` | Contains the `URL`. |
| `Description` | `string` | Alt text for accessibility. |

## Related

- [Components V2](components-v2.md)
