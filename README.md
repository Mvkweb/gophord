# Gophord

<p align="center">
  <a href="https://go.dev">
    <img src="https://cdn.jsdelivr.net/npm/@intergrav/devins-badges@3.3.1/assets/cozy/built-with/go_vector.svg" alt="Built with Go" height="48">
  </a>
  <a href="https://github.com/Mvkweb/gophord">
    <img src="https://cdn.jsdelivr.net/npm/@intergrav/devins-badges@3.3.1/assets/cozy/documentation/website_vector.svg" alt="Documentation" height="48">
  </a>
</p>

<p align="center">
  A fast, idiomatic Go library for the Discord API.<br>
  Clean REST and gateway APIs, optimized JSON with <code>bytedance/sonic</code>, and full support for Discord Message Components V2.
</p>

<p align="center">
  <em>⚠️ Work in progress. Built primarily to power <a href="https://evelith.dev">Evelith</a>.</em>
</p>

## Why Gophord?

Gophord is built for Go projects that want a clean API surface without giving up speed. It stays close to Discord's model, avoids code generation, and is structured for long-term maintenance.

- High-performance JSON encoding and decoding with `bytedance/sonic`
- Discord API v10+ support across REST and gateway workflows
- Full Message Components V2 support, including buttons, selects, containers, and sections
- Idiomatic Go package design with a straightforward developer experience
- Runnable examples for common bot patterns

## Installation

```bash
go get github.com/Mvkweb/gophord
```

## Authentication

Most examples read the bot token from `DISCORD_BOT_TOKEN`.

```bash
# PowerShell
$env:DISCORD_BOT_TOKEN="your_bot_token_here"

# CMD
set DISCORD_BOT_TOKEN=your_bot_token_here

# Linux / macOS
export DISCORD_BOT_TOKEN="your_bot_token_here"
```

Create or manage your bot token in the [Discord Developer Portal](https://discord.com/developers/applications).

## Quick Start

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/Mvkweb/gophord/pkg/gateway"
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_BOT_TOKEN is not set")
	}

	gw := gateway.New(token)

	if err := gw.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}
	defer gw.Close()

	log.Println("connected to Discord")

	for event := range gw.Events() {
		log.Printf("received %T", event)
	}
}
```

For REST requests, use `pkg/client`.

## Examples

The `examples` directory contains small, runnable programs:

- `simple_bot` for a minimal bot setup
- `message_components` for interactive component handling
- `event_handler` for structured gateway event processing

Run an example with:

```bash
cd examples/simple_bot
go run main.go
```

## Package Overview

- `pkg/types` contains Discord type definitions, including Message Components V2
- `pkg/client` provides the REST API client
- `pkg/gateway` handles the Discord gateway connection
- `pkg/json` wraps JSON operations using `bytedance/sonic`
- `pkg/rest` contains lower-level request and response helpers

## Performance

Gophord uses `bytedance/sonic` for high-throughput JSON handling, with benchmark results around **2,468 MB/s** for encoding and **2,521 MB/s** for decoding.

## Documentation

API documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/Mvkweb/gophord).

## Contributing

Contributions are welcome. Please keep changes idiomatic, tested, and aligned with the existing package structure.

## License

MIT. See [LICENSE](LICENSE).