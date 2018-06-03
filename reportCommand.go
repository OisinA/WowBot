package main

import (
	discord "github.com/bwmarrin/discordgo"
)

func ReportCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}

	ch, err := s.UserChannelCreate("280103141469585408")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Sorry, something went wrong :(")
		return
	}

	s.ChannelMessageSend(ch.ID, m.Author.String()+" > "+message)
	s.ChannelMessageSend(m.ChannelID, "Your report was sent.")
}
