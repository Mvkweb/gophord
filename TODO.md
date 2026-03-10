# Gophord Roadmap & TODO

This document outlines the current feature gaps in `gophord` compared to the Discord API v10+ and the planned roadmap for future development.

## 🚀 Priority: Components V2 Completion

Support for the latest Discord message components is a core goal of `gophord`.

- [ ] **Full Select Menu V2 Support**
  - [x] String Select (Type 3)
  - [x] User Select (Type 5)
  - [x] Role Select (Type 6)
  - [x] Mentionable Select (Type 7)
  - [x] Channel Select (Type 8)
  - [ ] Default values support for all Select Menus.
- [ ] **Premium Button Handling**
  - [x] Implement `ButtonStylePremium` (Type 6).
  - [ ] Add helper for fetching/validating SKUs for Premium Buttons.
- [ ] **File Upload Component**
  - [x] Implement `FileUpload` (Type 16).
  - [x] Add example demonstrating file uploads in components/modals.
- [ ] **Section & Media Gallery Enhancements**
  - [x] Basic Section support.
  - [x] Basic Media Gallery support.
  - [ ] Support for more complex layouts within Sections.

## 🛠️ REST API Coverage

Many core Discord endpoints are still missing in the REST client.

- [ ] **Guild Management**
  - [ ] Create/Edit/Delete Guild.
  - [ ] Fetch Guilds list.
  - [ ] Guild Widget & Vanity URL management.
  - [ ] MFA & Pruning endpoints.
- [ ] **Channel & Thread Management**
  - [ ] Permission Overwrites (CRUD).
  - [ ] Invite Management (CRUD).
  - [ ] Full Thread Support (Start, Join, Leave, Add/Remove Member, List Active/Archived).
  - [ ] Forum Channel Support (Post creation, Tags management).
- [ ] **Content Management**
  - [ ] Guild Emoji CRUD.
  - [ ] Guild Sticker CRUD.
  - [ ] Guild Scheduled Events CRUD.
  - [ ] Guild Soundboard Management (Sounds CRUD).
- [ ] **Stage Management**
  - [ ] Stage Instance CRUD.
- [ ] **Safety & Moderation**
  - [ ] Auto-moderation Rules & Action Execution.
  - [ ] Audit Log querying.
- [ ] **Monetization (Premium Features)**
  - [ ] SKU & Entitlement management.
- [ ] **Polls**
  - [ ] Full support for Message Polls (parsing and creation).

## 🔧 API Updates

Staying current with Discord's evolving API.

- [ ] **2025-2026 Permission Changes**
  - [ ] Handle new separate permissions: `PIN_MESSAGES`, `BYPASS_SLOWMODE`, `CREATE_GUILD_EXPRESSIONS`, `CREATE_EVENTS`
  - [ ] Note: `Create Guild` is now heavily restricted for apps
- [ ] **Components V2 Flag Handling**
  - [ ] Proper `IS_COMPONENTS_V2` message flag handling (disables traditional content/embeds)
  - [ ] Newer modal components (Labels, String Selects in modals, File Upload in modals, Radio Groups, Checkbox Groups)
- [ ] **Error Handling**
  - [ ] Better handling of new error response formats
- [ ] **Rate Limiting**
  - [ ] Per-guild rate limiting strategies
  - [ ] Smart sharding support

## 📡 Gateway & Events

Expanding event coverage to handle complex bot workflows.

### High Priority Events (Complex/Important)

- [ ] **Auto Moderation Events**
  - [ ] `AUTO_MODERATION_RULE_CREATE` - Auto Moderation rule was created
  - [ ] `AUTO_MODERATION_RULE_UPDATE` - Auto Moderation rule was updated
  - [ ] `AUTO_MODERATION_RULE_DELETE` - Auto Moderation rule was deleted
  - [ ] `AUTO_MODERATION_ACTION_EXECUTION` - Auto Moderation rule was triggered and action executed

- [ ] **Thread Events**
  - [ ] `THREAD_CREATE` - New thread created
  - [ ] `THREAD_UPDATE` - Thread was updated
  - [ ] `THREAD_DELETE` - Thread was deleted
  - [ ] `THREAD_LIST_SYNC` - Thread list sync for threads in channels
  - [ ] `THREAD_MEMBER_UPDATE` - Current user's thread member object updated
  - [ ] `THREAD_MEMBERS_UPDATE` - Thread members were updated

- [ ] **Guild Scheduled Events**
  - [ ] `GUILD_SCHEDULED_EVENT_CREATE` - Scheduled event created
  - [ ] `GUILD_SCHEDULED_EVENT_UPDATE` - Scheduled event updated
  - [ ] `GUILD_SCHEDULED_EVENT_DELETE` - Scheduled event deleted
  - [ ] `GUILD_SCHEDULED_EVENT_USER_ADD` - User subscribed to event
  - [ ] `GUILD_SCHEDULED_EVENT_USER_REMOVE` - User unsubscribed from event

- [ ] **Poll Events**
  - [ ] `MESSAGE_POLL_VOTE_ADD` - User voted on a poll
  - [ ] `MESSAGE_POLL_VOTE_REMOVE` - User removed vote from poll

- [ ] **Voice Events**
  - [ ] `VOICE_STATE_UPDATE` - User joined, left, or moved voice channel
  - [ ] `VOICE_SERVER_UPDATE` - Guild's voice server was updated

- [ ] **Presence & Relationships**
  - [ ] `PRESENCE_UPDATE` - User presence was updated
  - [ ] `GUILD_MEMBERS_CHUNK` - Chunk of guild members received (large guilds)

### Additional Events (For Completeness)

- [ ] **Entitlement/Monetization Events**
  - [ ] `ENTITLEMENT_CREATE` - Entitlement was created
  - [ ] `ENTITLEMENT_UPDATE` - Entitlement was updated
  - [ ] `ENTITLEMENT_DELETE` - Entitlement was deleted
  - [ ] `SUBSCRIPTION_CREATE` - Premium subscription created
  - [ ] `SUBSCRIPTION_UPDATE` - Premium subscription updated
  - [ ] `SUBSCRIPTION_DELETE` - Premium subscription deleted

- [ ] **Soundboard Events**
  - [ ] `GUILD_SOUNDBOARD_SOUND_CREATE` - Soundboard sound created
  - [ ] `GUILD_SOUNDBOARD_SOUND_UPDATE` - Soundboard sound updated
  - [ ] `GUILD_SOUNDBOARD_SOUND_DELETE` - Soundboard sound deleted
  - [ ] `GUILD_SOUNDBOARD_SOUNDS_UPDATE` - Soundboard sounds bulk updated

- [ ] **Stage Instance Events**
  - [ ] `STAGE_INSTANCE_CREATE` - Stage instance created
  - [ ] `STAGE_INSTANCE_UPDATE` - Stage instance updated
  - [ ] `STAGE_INSTANCE_DELETE` - Stage instance deleted

- [ ] **Integration Events**
  - [ ] `INTEGRATION_CREATE` - Integration created
  - [ ] `INTEGRATION_UPDATE` - Integration updated
  - [ ] `INTEGRATION_DELETE` - Integration deleted

- [ ] **Invite Events**
  - [ ] `INVITE_CREATE` - Invite created
  - [ ] `INVITE_DELETE` - Invite deleted

- [ ] **Channel Events**
  - [ ] `CHANNEL_PINS_UPDATE` - Message pinned/unpinned
  - [ ] `WEBHOOKS_UPDATE` - Webhook created/updated/deleted

- [ ] **User & Application Events**
  - [ ] `TYPING_START` - User started typing
  - [ ] `USER_UPDATE` - User properties changed
  - [ ] `APPLICATION_COMMAND_PERMISSIONS_UPDATE` - Command permissions updated

- [ ] **Message Reaction Events**
  - [ ] `MESSAGE_REACTION_REMOVE` - Reaction removed from message
  - [ ] `MESSAGE_REACTION_REMOVE_ALL` - All reactions removed
  - [ ] `MESSAGE_REACTION_REMOVE_EMOJI` - Specific emoji reactions removed

- [ ] **Guild Events**
  - [ ] `GUILD_AUDIT_LOG_ENTRY_CREATE` - Audit log entry created
  - [ ] `GUILD_CREATE` - Guild created/available
  - [ ] `GUILD_UPDATE` - Guild updated
  - [ ] `GUILD_DELETE` - Guild unavailable/deleted
  - [ ] `GUILD_BAN_ADD` - Member banned
  - [ ] `GUILD_BAN_REMOVE` - Member unbanned
  - [ ] `GUILD_EMOJIS_UPDATE` - Emojis updated
  - [ ] `GUILD_STICKERS_UPDATE` - Stickers updated
  - [ ] `GUILD_MEMBER_REMOVE` - Member left
  - [ ] `GUILD_ROLE_CREATE` - Role created
  - [ ] `GUILD_ROLE_UPDATE` - Role updated
  - [ ] `GUILD_ROLE_DELETE` - Role deleted
  - [ ] `GUILD_INTEGRATIONS_UPDATE` - Integrations updated

- [ ] **Voice Channel Effects**
  - [ ] `VOICE_CHANNEL_EFFECT_SEND` - Soundboard/emoji effect in voice

- [ ] **Gateway Reliability**
  - [ ] Improved shard management and identification.
  - [ ] Automatic reconnection polishing.

## 📦 Types & Models

Keeping types up-to-date with Discord's JSON schema.

- [ ] **Missing Models**
  - [ ] `AuditLog`
  - [ ] `AutoModerationRule`
  - [ ] `ScheduledEvent`
  - [ ] `Entitlement` & `SKU`
  - [ ] `ThreadMetadata`
  - [ ] `Poll`
  - [ ] `Subscription`
  - [ ] `SoundboardSound`
  - [ ] `StageInstance`
- [ ] **JSON Performance**
  - [x] `bytedance/sonic` integration for fast serialization.
  - [ ] Custom `MarshalJSON` benchmarks for critical paths.

## 📚 Documentation & Examples

- [ ] **Advanced Examples**
  - [ ] `examples/modals`: Complex multi-step modals.
  - [ ] `examples/threads`: Managing active threads and forum posts.
  - [ ] `examples/moderation`: Implementing auto-mod responses.
- [ ] **Contributor Guide**
  - [ ] Create `CONTRIBUTING.md`.
  - [ ] Document internal architecture and design patterns (Builder pattern, REST rate limiting).

## ✅ Recently Completed

- [x] **Gateway Shutdown Fix**: Resolved panic when closing the client.
- [x] **Components V2 Intro**: Implemented Sections, Separators, Containers, and Text Displays.
- [x] **Module Path Update**: Migrated to `github.com/Mvkweb/gophord`.
- [x] **Documentation Badge**: Modernized README.
- [x] **Global/Guild Slash Commands**: Robust registration system.
