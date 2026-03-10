// Package interactions provides slash command and component interaction handlers.
// This file handles button clicks, select menu choices, and other message component interactions.

package interactions

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Mvkweb/gophord/examples/simple_bot/commands"
	"github.com/Mvkweb/gophord/examples/simple_bot/modals"
	"github.com/Mvkweb/gophord/examples/simple_bot/utils"
	"github.com/Mvkweb/gophord/pkg/client"
	"github.com/Mvkweb/gophord/pkg/gateway"
	"github.com/Mvkweb/gophord/pkg/types"
)

// HandleSlashCommand processes slash command interactions.
// Routes to appropriate handler based on command name.
func HandleSlashCommand(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	if interaction.Data == nil {
		log.Println("Interaction data is nil for application command")
		return
	}

	switch interaction.Data.Name {
	case utils.CmdHello:
		handleHelloSlash(ctx, bot, interaction)
	case utils.CmdEphemeral:
		handleEphemeralSlash(ctx, bot, interaction)
	case utils.CmdModal:
		commands.HandleModalDemo(ctx, bot, interaction)
	case utils.CmdGallery:
		handleGallerySlash(ctx, bot, interaction)
	case utils.CmdSection:
		handleSectionSlash(ctx, bot, interaction)
	case utils.CmdFileUpload:
		modals.HandleFileUploadModal(ctx, bot, interaction)
	case utils.CmdPremium:
		handlePremiumSlash(ctx, bot, interaction)
	}
}

// HandleButtonClick processes button interaction callbacks.
// Routes to appropriate response based on button custom_id.
func HandleButtonClick(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	if interaction.Data == nil {
		log.Println("Interaction data is nil for message component")
		return
	}

	customID := interaction.Data.CustomID
	var response string
	var isEphemeral bool

	switch customID {
	// Greeting buttons
	case utils.ButtonGreetBack:
		userName := utils.GetUserName(interaction)
		response = fmt.Sprintf("Hello to you too, %s! 😊", userName)
	case utils.ButtonShowInfo:
		response = utils.BotInfo
		isEphemeral = true // Make info ephemeral to keep chat clean
	case utils.ButtonSayGoodbye:
		response = "Goodbye! See you next time! 👋"

	// Component demo buttons
	case utils.ButtonPrimary, utils.ButtonSecondary, utils.ButtonSuccess, utils.ButtonDanger:
		response = fmt.Sprintf("You clicked the **%s** button!", strings.TrimPrefix(customID, "btn_"))

	// Container buttons
	case utils.ButtonContainerAction:
		response = "Action taken! ✅"
	case utils.ButtonContainerDismiss:
		response = "Dismissed! 👍"
		isEphemeral = true

	// File upload
	case utils.ButtonOpenFileModal:
		modals.HandleFileUploadModal(ctx, bot, interaction)
		return

	default:
		response = fmt.Sprintf("Button clicked: `%s`", customID)
	}

	// Send the response
	var err error
	if isEphemeral {
		err = bot.RespondWithEphemeral(ctx, interaction, response)
	} else {
		err = bot.RespondWithMessage(ctx, interaction, response)
	}

	if err != nil {
		log.Printf("Failed to respond to interaction: %v", err)
	}
}

// HandleSelectMenu processes select menu (dropdown) interactions.
// Currently routes to modal submit handler since selects are primarily in modals.
func HandleSelectMenu(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	// Select menus in modals are handled by the modal submit handler
	// This is here for future expansion if selects are added to messages
	log.Printf("Select menu interaction: %s", interaction.Data.CustomID)
}

// handleHelloSlash handles the /hello slash command.
func handleHelloSlash(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	userName := utils.GetUserName(interaction)

	components := client.NewComponentBuilder().
		AddTextDisplay(fmt.Sprintf("# Hello, %s! 👋", userName)).
		AddTextDisplay("Welcome to the **gophord** Discord library demo!").
		AddSeparator(true, types.SeparatorSpacingSmall).
		AddTextDisplay("Click a button below to try out the interactive features:").
		AddActionRow(
			client.NewPrimaryButton(utils.ButtonGreetBack, "Greet Back"),
			client.NewSuccessButton(utils.ButtonShowInfo, "Show Info"),
			client.NewDangerButton(utils.ButtonSayGoodbye, "Goodbye"),
		).
		Build()

	err := bot.RespondWithComponents(ctx, interaction, components)
	if err != nil {
		log.Printf("Failed to respond to /hello interaction: %v", err)
	}
}

// handleEphemeralSlash handles the /ephemeral slash command.
func handleEphemeralSlash(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	err := bot.RespondWithEphemeral(ctx, interaction, "This is an ephemeral message! Only you can see this. 👀")
	if err != nil {
		log.Printf("Failed to respond to /ephemeral interaction: %v", err)
	}
}

// handleGallerySlash handles the /gallery slash command.
func handleGallerySlash(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# Media Gallery Demo").
		AddTextDisplay("Media galleries display multiple images in a grid.").
		AddMediaGallery(
			types.MediaGalleryItem{
				Media:       types.UnfurledMediaItem{URL: "https://http.cat/200.jpg"},
				Description: "HTTP 200 OK",
			},
			types.MediaGalleryItem{
				Media:       types.UnfurledMediaItem{URL: "https://http.cat/404.jpg"},
				Description: "HTTP 404 Not Found",
			},
		).
		Build()

	err := bot.RespondWithComponents(ctx, interaction, components)
	if err != nil {
		log.Printf("Failed to respond to /gallery interaction: %v", err)
	}
}

// handleSectionSlash handles the /section slash command.
func handleSectionSlash(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	components := client.NewComponentBuilder().
		AddTextDisplay("# Section Demo").
		AddSection(
			"Here is an example of a **Section** layout component, accompanied by a Thumbnail accessory.",
			&types.Thumbnail{
				Media:       types.UnfurledMediaItem{URL: "https://http.cat/418.jpg"},
				Description: "I'm a teapot",
			},
		).
		Build()

	err := bot.RespondWithComponents(ctx, interaction, components)
	if err != nil {
		log.Printf("Failed to respond to /section interaction: %v", err)
	}
}

// handlePremiumSlash handles the /premium slash command.
func handlePremiumSlash(ctx context.Context, bot *client.Client, interaction *types.Interaction) {
	if interaction.ChannelID == nil {
		return
	}

	// Delegate to the commands handler which has the env var logic
	commands.HandlePremium(ctx, bot, &gateway.MessageCreateEvent{
		Message: types.Message{
			ChannelID: *interaction.ChannelID,
		},
	})
}
