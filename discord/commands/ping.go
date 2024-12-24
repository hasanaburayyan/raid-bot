package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var PingCommand = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Just a health check!",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "response_message",
			Description: "The message you want the health check to return (default: Pong!)",
			Type:        discordgo.ApplicationCommandOptionString,
		},
	},
}

func HandlePingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var response = "Pong!"
	for _, option := range i.ApplicationCommandData().Options {
		if option.Name == "response_message" && option.StringValue() != "" {
			response = option.StringValue()
		}
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	})

	if err != nil {
		log.Printf("Error responding to interaction: %v", err)
	}
}
