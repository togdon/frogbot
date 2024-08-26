package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var frogs []string

func main() {

	// Load all the frog files into memory
	files, err := os.ReadDir("frogs/")
	if err != nil {
		panic(fmt.Sprintf("Error opening folder: %v", err))
	}
	for _, file := range files {
		frogs = append(frogs, file.Name())
	}

	// Load the Discord token from the environment
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("No token provided.")
	}

	// Create a new Discord session using the provided bot token.
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(fmt.Sprintf("Error creating Discord session: %v", err))
	}

	// Grab information about guilds and messages
	bot.Identify.Intents = discordgo.IntentsAll

	// Open the websocket and begin listening
	err = bot.Open()
	if err != nil {
		panic(fmt.Sprintf("Error opening Discord session: %v", err))
	}

	// Add handlers for the bot for message creation and reaction creation
	bot.AddHandler(messageCreate)
	bot.AddHandler(messageReactionAdd)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Printf("Frogbot is now running; I currently know about %v frogs. Press CTRL-C to exit.\n", len(frogs))
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
}
