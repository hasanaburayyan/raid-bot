package commands

import (
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/rand"
)

var LouCommand = &discordgo.ApplicationCommand{
	Name:        "lou",
	Description: "Tell you everything we know about Lou",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "random_lou_fact",
			Description: "The random fact I should remember about Lou",
			Type:        discordgo.ApplicationCommandOptionString,
		},
	},
}

var LouFacts []string

func Init() {
	log.Panicln("Initializing Lou facts")
	// Load facts from file
	facts, err := os.ReadFile("lou_facts.txt")
	if err != nil {
		log.Printf("Error reading file: %v", err)
	}
	LouFacts = strings.Split(string(facts), "\n")
}

func HandleLouCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	for _, option := range i.ApplicationCommandData().Options {
		if option.Name == "random_lou_fact" && option.StringValue() != "" {
			LouFacts = append(LouFacts, option.StringValue())
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Alright Ill remember that",
				},
			})

			if err != nil {
				log.Printf("Error responding to interaction: %v", err)
			}

			// Write all facts to a file clear file if it exists
			if _, err := os.Stat("lou_facts.txt"); err == nil {
				os.Remove("lou_facts.txt")
			}

			file, err := os.Create("lou_facts.txt")
			if err != nil {
				log.Printf("Error creating file: %v", err)
			}
			defer file.Close()

			for _, fact := range LouFacts {
				file.WriteString(fact + "\n")
			}

			return
		}
	}

	response := "He is the gayest mother fucker we know"
	if len(LouFacts) > 0 {
		response = LouFacts[rand.Intn(len(LouFacts))]
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
