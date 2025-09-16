package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// messageReactionRemove is called whenever a reaction is removed from a message
func messageReactionRemove(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
	// Ignore all reactions removed by the bot itself
	if m.UserID == s.State.User.ID {
		return
	}

	// log the emoji reactions removals
	user, _ := s.User(m.UserID)
	message, _ := s.ChannelMessage(m.ChannelID, m.MessageID)
	channel, _ := s.Channel(m.ChannelID)
	fmt.Printf("%s removed %s from %s's message \"%s\" in %s\n",
		user.Username, m.Emoji.Name, message.Author, message.Content, channel.Name)
}
