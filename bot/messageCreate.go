package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/togdon/frogbot/internal/responses"

	"github.com/bwmarrin/discordgo"
)

// messageCreate is called whenever a message is created
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Log all messages received
	channel, _ := s.Channel(m.ChannelID)
	if channel.Name != "" {
		fmt.Printf("%s wrote: \"%s\" in %s\n", m.Author, m.Content, channel.Name)
	} else {
		fmt.Printf("%s DM'd: \"%s\"\n", m.Author, m.Content)
	}

	// Parse the message for mentions...
	for _, user := range m.Mentions {
		// ... and if the mention is to the bot...
		if user.ID == s.State.User.ID {
			response := responses.MentionsResponse(m.Content)
			s.ChannelMessageSend(m.ChannelID, response)
		}
	}

	// If the message is a request to "frog me"
	frogme := regexp.MustCompile(`(?i)frog me`)
	fff := regexp.MustCompile(`(?i)(fun\b|frog[s]?\b|fact\b)`)
	if frogme.MatchString(m.Content) {
		// send a random frog image
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		frog_file := frogs[r1.Intn(len(frogs))]
		f, err := os.Open("frogs/" + frog_file)
		if err != nil {
			panic(fmt.Sprintf("Error opening file: %v", err))
		}
		s.ChannelFileSend(m.ChannelID, frog_file, f)
	} else if fff.MatchString(m.Content) {
		// the mesage contains "fun", "frog[s]?", or "fact", so...
		// send a random frog fact
		match := strings.ToLower(fff.FindString(m.Content))
		response := responses.FunFrogFact()
		message := ""
		switch match {
		case "fun":
			message = "**Fun** frog fact: "
		case "frog", "frogs":
			message = "Fun **frog** fact: "
		case "fact":
			message = "Fun frog **fact**: "
		default:
			message = fmt.Sprintf("||Oh noes... You got here through \"%s\"|| Here's a fun frog fact: ", match)
		}
		message += response
		s.ChannelMessageSend(m.ChannelID, message)
	}

	// If the message is in all caps...
	yelling := regexp.MustCompile(`^\P{L}*\p{Lu}\P{Ll}*$`)
	if yelling.MatchString(m.Content) && m.Content != "LOL" && m.Content != "WTF" {
		response := responses.YellingResponse(m.Content, m.Author.ID)
		fmt.Printf("LLM Response: %v", response)
		s.ChannelMessageSend(m.ChannelID, response)
	}
}
