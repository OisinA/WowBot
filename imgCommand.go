package main

import (
	discord "github.com/bwmarrin/discordgo"
)

func ImgCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discord.MessageEmbed {
		Title: "Image",
		Description: message,
		Image: &discord.MessageEmbedImage {
			URL: "https://picsum.photos/200/300/?random",
		},
	})
}