package command

import "github.com/bwmarrin/discordgo"

var (
	initialCommands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping the bot",
		},
	}
	initialHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
				},
			})
		},
	}
)

func GetInitialCommands() []*discordgo.ApplicationCommand {
	return initialCommands
}

func GetInitialHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return initialHandlers
}
