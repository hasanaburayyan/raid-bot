package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/hasanaburayyan/raid-bot/discord/commands"
)

var Commands = []*discordgo.ApplicationCommand{
	commands.PingCommand,
	commands.LouCommand,
}

func RegisterCommands(session *discordgo.Session, r *discordgo.Ready) {
	for _, guild := range r.Guilds {
		for _, command := range Commands {
			_, err := session.ApplicationCommandCreate(session.State.User.ID, guild.ID, command)
			if err != nil {
				log.Printf("Error creating command for guild %s: %v", guild.ID, err)
			} else {
				log.Printf("Ping command registered successfully in guild: %s", guild.ID)
			}
		}
	}
}

func RouteHandlers(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		switch i.ApplicationCommandData().Name {
		case "lou":
			commands.HandleLouCommand(s, i)
		case "ping":
			commands.HandlePingCommand(s, i)
		}
	}
}
