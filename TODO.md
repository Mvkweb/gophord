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
  - [ ] Add example demonstrating file uploads in components/modals.
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
- [ ] **Safety & Moderation**
  - [ ] Auto-moderation Rules & Action Execution.
  - [ ] Audit Log querying.
- [ ] **Monetization (Premium Features)**
  - [ ] SKU & Entitlement management.
- [ ] **Polls**
  - [ ] Full support for Message Polls (parsing and creation).

## 📡 Gateway & Events

Expanding event coverage to handle complex bot workflows.

- [ ] **Event Parsing Coverage**
  - [ ] Implement structs and `ParseEvent` cases for:
    - [ ] `AUTO_MODERATION_*`
    - [ ] `THREAD_*`
    - [ ] `GUILD_AUDIT_LOG_ENTRY_CREATE`
    - [ ] `GUILD_SCHEDULED_EVENT_*`
    - [ ] `MESSAGE_POLL_VOTE_*`
    - [ ] `VOICE_STATE_UPDATE` & `VOICE_SERVER_UPDATE`
    - [ ] `PRESENCE_UPDATE`
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
