package main

import (
	discord "github.com/bwmarrin/discordgo"
	"strings"
)

type command struct {
	Name    string
	Execute func(*discord.Session, *discord.MessageCreate, string)
}

var commands = make(map[string]command)

func RegisterCommands() {
	commands["~help"] = command{"help", HelpCommand}
	commands["~img"] = command{"img", ImgCommand}
	commands["~say"] = command{"say", SayCommand}
	commands["~lyrics"] = command{"lyrics", LyricCommand}
	commands["~yesorno"] = command{"yesorno", YesOrNoCommand}
	commands["~dadjoke"] = command{"dadjoke", DadJokeCommand}
	commands["~save"] = command{"save", ImageSaveCommand}
}

func ParseCommands(s *discord.Session, m *discord.MessageCreate) {
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

func HelpCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}
	s.ChannelMessageSend(m.ChannelID, "To use a command, send a message as follows: ~[command]\nAvailable commands: "+strings.Join(keys, ", "))
}

func SayCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	s.ChannelMessageSend(m.ChannelID, m.Author.Username+" - "+message)
}
