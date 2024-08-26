package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func messageReactionAdd(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	// Ignore all reactions added by the bot itself
	if m.UserID == s.State.User.ID {
		return
	}

	if m.Emoji.Name == "🐸" {
		s.ChannelMessageSend(m.ChannelID, "Ribbit!")
	} else {
		// write out the emoji reactions
		fmt.Printf("%v reacted with %v\n", m.UserID, m.Emoji.Name)
	}

}
