package main

import (
	"fmt"
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

func UserStatsCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}

	split := strings.Split(message, " ")

	if len(split) != 1 {
		s.ChannelMessageSend(m.ChannelID, "Incorrect usage.")
		return
	}

	user := split[0]
	s.ChannelMessageSend(m.ChannelID, getUserStats(s, m, user))
}

func getUserStats(s *discord.Session, m *discord.MessageCreate, username string) string {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return "Something went wrong."
	}

	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		return "Something went wrong."
	}

	for _, u := range g.Members {
		if "<@"+u.User.ID+">" == strings.Replace(username, "!", "", 1) {
			printStats(c.GuildID, s, m, u)
			return ""
		}
	}
	return "User not found."
}

func printStats(guildID string, s *discord.Session, m *discord.MessageCreate, user *discord.Member) {
	fields := []*discord.MessageEmbedField{}
	if user.Nick != "" {
		fields = append(fields, &discord.MessageEmbedField{
			Name:  "Nickname",
			Value: user.Nick,
		})
	}
	fields = append(fields, &discord.MessageEmbedField{
		Name:  "Joined",
		Value: strings.Split(user.JoinedAt, "T")[0],
	})
	roles := []string{}
	for _, r := range user.Roles {
		r, err := s.State.Role(guildID, r)
		if err != nil {
			continue
		}
		roles = append(roles, r.Name)
	}
	fields = append(fields, &discord.MessageEmbedField{
		Name:  "Roles",
		Value: fmt.Sprint(roles),
	})
	s.ChannelMessageSendEmbed(m.ChannelID, &discord.MessageEmbed{
		Title: user.User.String(),
		Thumbnail: &discord.MessageEmbedThumbnail{
			URL: user.User.AvatarURL("32"),
		},
		Fields: fields,
	})
}
