package main

import (
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

type command struct {
	Name        string
	Description string
	Show        bool
	Execute     func(*discord.Session, *discord.MessageCreate, string)
}

var commands = make(map[string]command)

func RegisterCommands() {
	commands["~help"] = command{"help", "List of commands.", true, HelpCommand}
	commands["~img"] = command{"img", "Used to recall images saved using ~save.\nUsage: ~img [name]", true, ImgCommand}
	commands["~lyrics"] = command{"lyrics", "Get the lyrics to a song.\nUsage: ~lyrics [artist]-[song]", true, LyricCommand}
	commands["~yesorno"] = command{"yesorno", "Allow the bot to decide your fate with this command.\nUsage: ~yesorno", true, YesOrNoCommand}
	commands["~dadjoke"] = command{"dadjoke", "Enjoy a fun dadjoke.\nUsage: ~dadjoke", true, DadJokeCommand}
	commands["~save"] = command{"save", "Save an image, that you can recall with ~img.\nUsage: ~save [url] [name]", true, ImageSaveCommand}
	commands["~ow"] = command{"ow", "Fetch your Overwatch statistics with this command.\nUsage: ~ow [battletag]", true, OverwatchCommand}
	commands["~stats"] = command{"stats", "Fetch the stats of a user in your discord server.\nUsage: ~stats [user]", true, UserStatsCommand}
	commands["~status"] = command{"status", "", false, SetStatusCommand}
}

func ParseCommands(s *discord.Session, m *discord.MessageCreate) {
	if IsBetaUserConnected(s, m) && !beta {
		return
	}
	if m.Author.Bot {
		return
	}

	content := m.Message.Content
	split := strings.Split(content, " ")
	command := strings.ToLower(split[0])

	returned, ok := commands[command]
	if !ok {
		return
	}

	returned.Execute(s, m, strings.Join(split[1:], " "))
	return
}

func IsBetaUserConnected(s *discord.Session, m *discord.MessageCreate) bool {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return false
	}

	_, err = s.GuildMember(c.GuildID, "452445290365059072")
	if err != nil {
		return false
	}
	return true
}

func HelpCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if len(message) == 0 {
		keys := make([]string, 0, len(commands))
		for k := range commands {
			if commands[k].Show {
				keys = append(keys, k)
			}
		}
		s.ChannelMessageSend(m.ChannelID, "To use a command, send a message as follows: ~[command]\nTo find out more about a command, use ~help [command].\nAvailable commands: "+strings.Join(keys, ", "))
	} else {
		c, ok := commands["~"+message]
		if !ok {
			s.ChannelMessageSend(m.ChannelID, "Command not found.")
			return
		}
		s.ChannelMessageSend(m.ChannelID, "~"+message+"\nDescription: "+c.Description)
	}
}

func SayCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	s.ChannelMessageSend(m.ChannelID, m.Author.Username+" - "+message)
}
