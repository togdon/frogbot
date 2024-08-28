package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// messageReactionAdd is called whenever a reaction is added to a message
func messageReactionAdd(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	// Ignore all reactions added by the bot itself
	if m.UserID == s.State.User.ID {
		return
	}

	if m.Emoji.Name == "🐸" {
		s.ChannelMessageSend(m.ChannelID, "Ribbit!")
	} else {
		// log the emoji reactions
		user, _ := s.User(m.UserID)
		message, _ := s.ChannelMessage(m.ChannelID, m.MessageID)
		channel, _ := s.Channel(m.ChannelID)
		fmt.Printf("%s reacted with %s to %s's message \"%s\" in %s\n",
			user.Username, m.Emoji.Name, message.Author, message.Content, channel.Name)
	}

}
