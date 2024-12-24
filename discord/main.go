package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("BOT_TOKEN env must be set!")
	}

	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up and running!")
		RegisterCommands(s, r)
		log.Println("Command registration complete!")
	})

	bot.AddHandler(RouteHandlers)

	err = bot.Open()
	if err != nil {
		log.Fatalf("Error opening Websocket connection: %v", err)
	}
	defer bot.Close()

	select {}
}
