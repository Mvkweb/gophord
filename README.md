# Gophord

A high-performance, idiomatic Go library for interacting with the Discord API. Built with senior-level code quality, optimized JSON serialization using `bytedance/sonic`, and comprehensive Discord API v10+ support.

## Features

- **High Performance JSON**: Uses `bytedance/sonic` for blazingly fast JSON marshaling/unmarshaling
- **Message Components V2**: Full support for Discord's latest message components (buttons, selects, containers, sections, etc.)
- **REST API Client**: Complete REST client implementation with proper error handling
- **Gateway Support**: WebSocket-based gateway connection with event handling
- **Idiomatic Go**: Follows Go best practices, no code generation required
- **Production Ready**: Designed for external documentation generation and long-term maintenance

## Installation

```bash
go get github.com/gophord/gophord
```

## Configuration

The bot uses the `DISCORD_BOT_TOKEN` environment variable for authentication.

### Getting Your Bot Token

1. Go to [Discord Developer Portal](https://discord.com/developers/applications)
2. Create or select your application
3. Click "Bot" in the sidebar
4. Copy the token (click "Reset Token" if you don't see one)

### Setting Up Your Token

**Option 1: Environment Variable (Recommended)**

```bash
# Windows PowerShell
$env:DISCORD_BOT_TOKEN="your_bot_token_here"

# Windows CMD
set DISCORD_BOT_TOKEN=your_bot_token_here

# Linux/Mac
export DISCORD_BOT_TOKEN="your_bot_token_here"
```

**Option 2: Command-Line Flag**

```bash
go run examples/simple_bot/main.go -token your_bot_token_here
```

**Option 3: Using a .env file**

Create a `.env` file in the project root:

```env
DISCORD_BOT_TOKEN=your_bot_token_here
```

Then load it before running (requires a .env loader like `joho/godotenv`):

```bash
# Install godotenv
go get github.com/joho/godotenv/v4

# Create a runner script that loads .env first
echo 'package main; import "github.com/joho/godotenv/v4"; func main() { godotenv.Load(); // your bot code }' > run.go
```

### Running Examples

```bash
# Simple bot example
cd examples/simple_bot
go run main.go

# Message components example
cd examples/message_components
go run main.go -token your_token
```

## Quick Start

```go
package main

import (
	"context"
	"log"

	"github.com/gophord/gophord/pkg/client"
	"github.com/gophord/gophord/pkg/gateway"
)

func main() {
	// Create a new Discord client
	botToken := "your-bot-token-here"
	
	discordClient := client.New(botToken)
	
	// Create a gateway connection
	gwClient := gateway.New(botToken)
	ctx := context.Background()
	
	// Connect to Discord
	err := gwClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer gwClient.Close()
	
	// Handle events
	for event := range gwClient.Events() {
		log.Printf("Event: %v", event)
	}
}
```

## Documentation

Documentation is automatically generated from the source code and published to the GitHub CodeSpace for easy access and maintenance.

### Key Packages

- **`pkg/types`**: Core Discord type definitions including Message Components V2
- **`pkg/client`**: REST API client for Discord endpoints
- **`pkg/gateway`**: WebSocket gateway connection handling
- **`pkg/json`**: JSON marshaling/unmarshaling utilities using bytedance/sonic
- **`pkg/rest`**: REST request/response handling

## Examples

See the `examples` directory for complete working examples:

- **`simple_bot`**: Basic bot setup and command handling
- **`message_components`**: Interactive message components demonstration
- **`event_handler`**: Comprehensive event handling example

## Performance

Benchmarks show Gophord achieves **2,468 MB/s** encoding and **2,521 MB/s** decoding throughput using bytedance/sonic.

## License

MIT License - See LICENSE file for details

## Contributing

Contributions are welcome! Please ensure code follows Go best practices and includes appropriate tests.