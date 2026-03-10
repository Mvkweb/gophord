# Embeds

> Create rich formatted messages with titles, images, and fields.

## Overview

Embeds are rich blocks of content that can contain colored borders, images, thumbnails, authors, footers, and grid-like fields. A single message can contain up to 10 embeds.

## Detailed Usage

Embeds are built using the `types.Embed` struct and its associated sub-structs. There are no builder methods in the core types, so you populate the structs directly.

```go
package main

import (
	"log"

	"github.com/Mvkweb/gophord/pkg/rest"
	"github.com/Mvkweb/gophord/pkg/types"
)

func sendEmbedMessage(restClient *rest.Client, channelID string) {
	embed := types.Embed{
		Title:       "Gophord Documentation",
		Description: "Learning how to build embeds with Gophord.",
		URL:         "https://github.com/Mvkweb/gophord",
		Color:       0x00ADD8, // Go blue (hex)
		Author: &types.EmbedAuthor{
			Name: "Gopher",
			IconURL: "https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png",
		},
		Thumbnail: &types.EmbedThumbnail{
			URL: "https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png",
		},
		Fields: []types.EmbedField{
			{
				Name:   "Version",
				Value:  "v1.0.0",
				Inline: true,
			},
			{
				Name:   "API",
				Value:  "v10",
				Inline: true,
			},
			{
				Name:   "Performance",
				Value:  "Uses bytedance/sonic for ultra-fast JSON parsing.",
				Inline: false, // Breaks to a new line
			},
		},
		Footer: &types.EmbedFooter{
			Text: "Documentation Generated",
		},
		Timestamp: "2023-10-27T12:00:00Z", // ISO8601 string
	}

	_, err := restClient.CreateMessage(channelID, rest.CreateMessageParams{
		Content: "Here is an example embed:",
		Embeds:  []types.Embed{embed},
	})
	if err != nil {
		log.Fatalf("Error sending embed: %v", err)
	}
}
```

## API Reference

### Core Struct

| Struct | Description |
|---|---|
| `types.Embed` | The top-level embed object. |

### Sub-Components

| Struct | Attached properties | Description |
|---|---|---|
| `types.EmbedAuthor` | `Name`, `URL`, `IconURL` | Displayed at the very top of the embed. |
| `types.EmbedFooter` | `Text`, `IconURL` | Displayed at the very bottom. |
| `types.EmbedField` | `Name`, `Value`, `Inline` | Grid-like key-value pairs inside the embed body. |
| `types.EmbedImage` | `URL` | Large image displayed at the bottom of the embed content. |
| `types.EmbedThumbnail` | `URL` | Small image displayed to the top-right of the descriptions. |

*Note: The `types.EmbedVideo` and `types.EmbedProvider` structs exist in the Discord API but cannot be sent by bots. They are only present on embeds attached automatically by Discord for user-sent links.*

## Limits and Constraints

When sending embeds, ensure you stay within Discord's limits (otherwise the API will return an HTTP 400 error):

- **Embed Title**: 256 characters max
- **Embed Description**: 4096 characters max
- **Fields**: Up to 25 per embed
- **Field Name**: 256 characters max
- **Field Value**: 1024 characters max
- **Footer Text**: 2048 characters max
- **Author Name**: 256 characters max
- **Total Characters**: The sum of all characters in all embeds in a single message cannot exceed 6000.

## Related

- [Messages](messages.md) — How to attach these embeds to standard messages.
- [Interactions](interactions.md) — You can respond to interactions using `RespondWithEmbeds`.
