// Package gophord is a high-performance Discord API library for Go.
//
// # Overview
//
// Gophord provides a complete interface to the Discord API with:
//
//   - REST API client for HTTP endpoints
//   - Gateway client for real-time WebSocket events
//   - Full support for Message Components V2
//   - High-performance JSON using bytedance/sonic
//
// # Installation
//
//	go get github.com/gophord/gophord
//
// # Basic Usage
//
//	import (
//	    "github.com/gophord/gophord/pkg/client"
//	    "github.com/gophord/gophord/pkg/types"
//	)
//
//	func main() {
//	    bot := client.New("your-token")
//	    // ...
//	}
//
// # Components V2
//
// Discord's Components V2 system is fully supported. The following
// component types are available:
//
// Layout Components:
//   - ActionRow - Container for buttons or a select menu
//   - Section - Associates text content with an accessory
//   - Container - Groups components with an accent color
//   - Separator - Visual divider between components
//
// Content Components:
//   - TextDisplay - Markdown-formatted text
//   - Thumbnail - Small image accessory
//   - MediaGallery - Display multiple images
//   - File - Display file attachments
//
// Interactive Components:
//   - Button - Clickable button (various styles)
//   - StringSelect - Select from predefined options
//   - UserSelect - Select users
//   - RoleSelect - Select roles
//   - MentionableSelect - Select users or roles
//   - ChannelSelect - Select channels
//   - TextInput - Text input for modals
//
// # Documentation Standards
//
// This library follows Go documentation best practices and is designed
// to work seamlessly with godoc and external documentation generators.
// All exported types, functions, and methods include comprehensive
// documentation comments.
//
// For GitHub Codespace documentation generation, the library structure
// supports automatic documentation extraction from source comments.
package gophord
