package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/togdon/frogbot/internal/responses"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Log all messages received
	fmt.Printf("%v (%v) wrote: %v\n", m.Author, m.Author.ID, m.Content)

	// Parse the message for mentions...
	for _, user := range m.Mentions {
		// ... and if the mention is to the bot...
		if user.ID == s.State.User.ID {
			response := responses.MentionsResponse(m.Content)
			s.ChannelMessageSend(m.ChannelID, response)
		}
	}

	// If the message is the request to "frog me"
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

	// If the message is in all caps...
	yelling := regexp.MustCompile(`^\P{L}*\p{Lu}\P{Ll}*$`)
	if yelling.MatchString(m.Content) && m.Content != "LOL" && m.Content != "WTF" {
		response := responses.YellingResponse(m.Content, m.Author.ID)
		fmt.Printf("LLM Response: %v", response)
		s.ChannelMessageSend(m.ChannelID, response)
	}
}
