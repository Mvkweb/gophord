# Agent Guidelines for Gophord

This document outlines patterns and conventions for working on the gophord project, particularly for the example bots.

## Project Structure

### Examples Directory

```
examples/
├── simple_bot/           # Main example bot - comprehensive feature demonstrations
│   ├── main.go          # Entry point, minimal setup
│   ├── .env             # Environment variables
│   │
│   ├── commands/        # Prefix command handlers (!ping, !kick, etc.)
│   │   ├── base.go              # Type definitions and interfaces
│   │   ├── ping.go              # !ping command
│   │   ├── hello.go             # !hello command with buttons
│   │   ├── help.go             # !help command
│   │   ├── moderation.go       # !kick, !purge commands
│   │   ├── webhooks.go         # !webhook command
│   │   ├── messages.go         # !react, !pin commands
│   │   ├── components_demo.go  # !components, !container, !gallery, !section, !fileupload
│   │   └── modal.go            # Modal demo handler
│   │
│   ├── interactions/    # Slash commands & button/component handlers
│   │   ├── slash_commands.go   # Slash command definitions
│   │   └── components.go       # Button/select interaction handlers
│   │
│   ├── modals/         # Modal form handlers
│   │   ├── handler.go          # Modal submit router
│   │   └── file_upload.go     # File upload modal logic
│   │
│   ├── events/         # Gateway event handlers
│   │   ├── ready.go            # OnReady handler
│   │   └── message_create.go   # Message routing to commands
│   │
│   └── utils/          # Shared utilities
│       ├── constants.go       # Command names, IDs, messages
│       └── helpers.go         # Helper functions
│
├── other_examples/     # Standalone examples for specific features
```

## Adding New Features

### Adding a New Prefix Command (!command)

1. **Create a new file** in `commands/` (e.g., `commands/mycommand.go`)
2. **Implement the handler function**:

```go
package commands

import (
    "context"
    "log"

    "github.com/Mvkweb/gophord/pkg/client"
    "github.com/Mvkweb/gophord/pkg/gateway"
)

func HandleMyCommand(ctx context.Context, bot *client.Client, event *gateway.MessageCreateEvent) {
    // Your command logic here
    _, err := bot.SendMessage(ctx, event.ChannelID, "Hello from my command!")
    if err != nil {
        log.Printf("Failed to send message: %v", err)
    }
}
```

3. **Register the command** in `events/message_create.go`:
   - Add a case in the switch statement
   - Follow the pattern: `case command == "!mycommand": commands.HandleMyCommand(ctx, bot, event)`

### Adding a New Slash Command

1. **Add command definition** in `interactions/slash_commands.go`:
   - Add to the `GetSlashCommands()` slice

2. **Add handler** in `interactions/components.go`:
   - Add a case in `HandleSlashCommand()` function
   - Create a helper function if complex

### Adding a New Button/Component Handler

1. **Add button constant** in `utils/constants.go`:
   - Add to appropriate const section (ButtonXXX)

2. **Add handler** in `interactions/components.go`:
   - Add case in `HandleButtonClick()` function

### Adding a New Modal

1. **Create handler file** in `modals/` (e.g., `modals/mymodal.go`)
2. **Implement two functions**:
   - `HandleMyModal(ctx, bot, interaction)` - Shows the modal
   - `HandleMyModalSubmit(ctx, bot, interaction)` - Processes submission

3. **Register in router** in `modals/handler.go`:
   - Add case in `HandleModalSubmit()` switch

### Adding a New Event Handler

1. **Create file** in `events/` (e.g., `events/myevent.go`)
2. **Create setup function**:
```go
package events

func OnMyEvent(bot *client.Client) {
    bot.OnMyEvent(func(event *gateway.MyEventEvent) {
        // Handle event
    })
}
```
3. **Call from main.go** in the event setup section

## Code Style Guidelines

### Comments

- Use clear, human-readable comments
- Explain **what** the code does, not obvious **how**
- Include usage examples for command handlers
- Add doc comments (//) for public functions

### Naming Conventions

- **Handlers**: `HandleCommandName` (PascalCase)
- **Constants**: Descriptive, use Caps with underscores (e.g., `ButtonGreetBack`)
- **Files**: Use underscores, lowercase (e.g., `my_command.go`)

### Error Handling

- Log errors with context using `log.Printf`
- Don't silently ignore errors
- Provide user feedback when possible

### Imports

- Group imports: standard library, external packages, internal packages
- Use aliases when needed to avoid conflicts

## Testing Features

When adding new features to simple_bot:

1. **Test as standalone** first - create a temp example if needed
2. **Add to simple_bot** - integrate into appropriate handler file
3. **Test both prefix and slash** - if applicable to both command types
4. **Update help** - add to !help command output

## Best Practices

1. **Keep handlers focused** - One function = one responsibility
2. **Reuse code** - Extract common logic to utils/
3. **Use constants** - Don't hardcode strings, put them in utils/constants.go
4. **Handle errors gracefully** - Don't crash on bad user input
5. **Follow existing patterns** - Match the style of surrounding code
