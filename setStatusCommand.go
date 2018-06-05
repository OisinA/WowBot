package main

import (
	"log"

	discord "github.com/bwmarrin/discordgo"
)

func SetStatusCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot || m.Author.ID != "280103141469585408" {
		return
	}

	err := s.UpdateStatus(0, message)
	if err != nil {
		SendMessage(s, m.ChannelID, "Something went wrong.")
		log.Fatal(err)
	}

	SendMessage(s, m.ChannelID, "Status updated.")
}
