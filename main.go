package main

import (
	"flag"
	"fmt"
	"log"
	command "main/commands"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token          string // bot token
	GuildID        string // test guild id
	RemoveCommands bool   // remove commands on graceful shutdown
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&GuildID, "g", "", "Guild ID")
	flag.BoolVar(&RemoveCommands, "r", false, "Remove commands on graceful shutdown")
	flag.Parse()
}

func main() {
	session, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	handlers := command.GetInitialHandlers()
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := handlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		}
	})

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = session.Open()
	if err != nil {
		log.Fatalf("error opening connection: %v", err)
	}
	defer session.Close()

	log.Printf("Registering commands...")
	commands := command.GetInitialCommands()
	registeredCommands := make([]*discordgo.ApplicationCommand, 0)
	for _, command := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, GuildID, command)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", command.Name, err)
		}
		registeredCommands = append(registeredCommands, cmd)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-sc

	if RemoveCommands {
		log.Println("Removing commands...")

		for _, v := range registeredCommands {
			err := session.ApplicationCommandDelete(session.State.User.ID, GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}
}
