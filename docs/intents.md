# Intents

> Configure gateway intents to receive specific events.

## Overview

Discord requires bots to specify "Intents" when connecting to explicitly declare which events they want to receive. This saves bandwidth and processing power. Some events containing sensitive data (like message content or presence updates) require "Privileged Intents" which must be enabled in your Discord Developer Portal.

## Detailed Usage

When creating your client, use `client.WithIntents()` to specify the bitwise OR of all intents your bot needs.

```go
package main

import (
	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/types"
)

func createBotWithIntents(token string) *client.Client {
	// 1. Sensible defaults (guilds, guild messages, reactions, DMs)
	bot := client.New(token, client.WithIntents(types.IntentsDefault))

	// 2. Custom combination using bitwise OR (|)
	customIntents := types.IntentGuilds |
		types.IntentGuildMessages |
		types.IntentGuildMessageReactions |
		types.IntentMessageContent // NOTE: Privileged!

	botWithCustom := client.New(token, client.WithIntents(customIntents))

	// 3. All non-privileged intents
	botAll := client.New(token, client.WithIntents(types.IntentsAll))

	// 4. Including privileged intents
	botPrivileged := client.New(token,
		client.WithIntents(types.IntentsAll|types.IntentsPrivileged))

	return botPrivileged
}
```

## API Reference

### Intent Constants

All intents are defined as `types.IntentFlags` constants. Combine them with the `|` operator.

| Intent Flag | Description |
|---|---|
| `types.IntentGuilds` | Guild create/update/delete, role/channel events |
| `types.IntentGuildModeration` | Ban add/remove events |
| `types.IntentGuildEmojisAndStickers` | Emoji/sticker update events |
| `types.IntentGuildIntegrations` | Integration events |
| `types.IntentGuildWebhooks` | Webhook update events |
| `types.IntentGuildInvites` | Invite create/delete events |
| `types.IntentGuildVoiceStates` | Voice state update events |
| `types.IntentGuildMessages` | Message create/update/delete in guilds |
| `types.IntentGuildMessageReactions` | Reaction add/remove in guilds |
| `types.IntentGuildMessageTyping` | Typing start in guilds |
| `types.IntentDirectMessages` | DM message events |
| `types.IntentDirectMessageReactions` | DM reaction events |
| `types.IntentDirectMessageTyping` | DM typing events |
| `types.IntentGuildScheduledEvents` | Scheduled event events |
| `types.IntentAutoModerationConfiguration` | Auto-mod rule events |
| `types.IntentAutoModerationExecution` | Auto-mod execution events |
| `types.IntentGuildMessagePolls` | Poll events in guilds |
| `types.IntentDirectMessagePolls` | Poll events in DMs |

### Privileged Intents

The following intents **must** be explicitly toggled on in the "Bot" section of the Discord Developer Portal before your bot can connect using them.

| Intent Flag | Description |
|---|---|
| `types.IntentGuildMembers` | Member add/remove/update |
| `types.IntentGuildPresences` | User presence and status updates |
| `types.IntentMessageContent` | Access to the actual text content of messages in guilds |

### Predefined Intent Groups

| Variable | Description |
|---|---|
| `types.IntentsDefault` | Sensible default containing guilds, messages, reactions, and DMs. |
| `types.IntentsAll` | All non-privileged intents combined. |
| `types.IntentsPrivileged` | All privileged intents combined. |

## Related

- [Session & Client](session.md) — How to pass `WithIntents` to your client.
- [Events](events.md) — Reference for which events you will receive.
