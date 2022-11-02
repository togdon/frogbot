package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var frogs []string

func main() {

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("No token provided.")
	}

	// Create a new Discord session using the provided bot token.
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(fmt.Sprintf("Error creating Discord session: %v", err))
	}

	// We need information about guilds and messages
	bot.Identify.Intents = discordgo.IntentsAll

	// Open the websocket and begin listening.
	err = bot.Open()
	if err != nil {
		panic(fmt.Sprintf("Error opening Discord session: %v", err))
	}

	files, err := os.ReadDir("frogs/")
	if err != nil {
		panic(fmt.Sprintf("Error opening folder: %v", err))
	}

	for _, file := range files {
		frogs = append(frogs, file.Name())
	}

	bot.AddHandler(messageCreate)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Printf("Frogbot is now running; I currently know about %v frogs. Press CTRL-C to exit.\n", len(frogs))
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	for _, user := range m.Mentions {
		if user.ID == s.State.User.ID {

			bil := regexp.MustCompile(`(?i)what is best in life`)

			if bil.MatchString(m.Content) {

				//To [verb] your [noun], to see them [pparticiple] [preposition] you, and
				//to [sense] the [singular] of their [kinfolk].

				s1 := rand.NewSource(time.Now().UnixNano())
				r1 := rand.New(s1)

				verb := []string{"abolish", "blot out", "crush", "decimate", "demolish", "eradicate", "erase", "exterminate", "extinguish", "massacre", "obliterate", "quash", "quell", "raze", "slaughter", "take out", "wipe out"}
				noun := []string{"adversaries", "antagonists", "assailants", "attackers", "competitors", "detractors", "enemies", "foes", "invaders", "opponents", "opposition", "rivals"}
				pparticiple := []string{"driven", "consumed", "forced", "herded", "possessed"}
				preposition := []string{"above", "across", "among", "at", "before", "beneath", "beside", "between", "by", "down", "in", "in front of", "on", "over", "through", "to", "up", "with"}
				sense := []string{"hear", "see", "smell", "taste", "touch"}
				singular := []string{"lamentations", "complaints", "dirges", "elegies", "keenings", "laments", "moanings", "mournings", "requiems", "sobs", "tears", "ululations", "wails"}
				kinfolk := []string{"fuckbois", "women", "children", "cats", "dogs", "Great Aunt Mildred", "Crazy Uncle Ernie"}

				response := fmt.Sprintf("To %s your %s, to see them %s %s you, and to %s the %s of their %s.", verb[r1.Intn(len(verb))], noun[r1.Intn(len(noun))], pparticiple[r1.Intn(len(pparticiple))], preposition[r1.Intn(len(preposition))], sense[r1.Intn(len(sense))], singular[r1.Intn(len(singular))], kinfolk[r1.Intn(len(kinfolk))])

				s.ChannelMessageSend(m.ChannelID, response)
			} else {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hello <@%v>! It's me %v, ready to serve!", m.Author.ID, user.Username))
			}
		}
	}

	fmt.Printf("%v (%v) wrote: %v\n", m.Author, m.Author.ID, m.Content)

	frogme := regexp.MustCompile(`(?i)Frog me`)

	if frogme.MatchString(m.Content) {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		frog_file := frogs[r1.Intn(len(frogs))]
		f, err := os.Open("frogs/" + frog_file)
		if err != nil {
			panic(fmt.Sprintf("Error opening file: %v", err))
		}

		s.ChannelFileSend(m.ChannelID, frog_file, f)
	}

	// Upper case only
	yelling := regexp.MustCompile(`^\P{L}*\p{Lu}\P{Ll}*$`)

	if yelling.MatchString(m.Content) && m.Content != "LOL" && m.Content != "WTF" {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hey <@%v>, there's no need to yell.", m.Author.ID))
	}

	// If the message is "ping" reply with "Pong!"
	// if m.Content == "ping" {
	// 	s.ChannelMessageSend(m.ChannelID, "Pong!")
	// }
	// If the message is "pong" reply with "Ping!"
	// if m.Content == "pong" {
	// 	s.ChannelMessageSend(m.ChannelID, "Ping!")
	// }

	// More thoughts:
	// https://huggingface.co/facebook/blenderbot-400M-distill?text=Hey+my+name+is+Julien%21+How+are+you%3F
	// https://github.com/bwmarrin/discordgo/blob/master/examples/slash_commands/main.go
	// https://github.com/montanaflynn/meme-generator/blob/master/main.go

}
