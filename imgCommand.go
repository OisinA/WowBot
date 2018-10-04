package main

import (
	"io/ioutil"
	"net/http"

	discord "github.com/bwmarrin/discordgo"
)

func ImgCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}
	if len(message) < 1 {
		SendMessage(s, m.ChannelID, "No image specified.")
		return
	}
	url := "/home/images/" + m.Author.ID + "/" + message + ".png"
	f, err := os.Open(url)
	if err != nil {
		SendMessage(s, m.ChannelID, "No image found.")
		return
	}
	defer f.Close()
	ms := &dg.MessageSend{
		Embed: &dg.MessageEmbed{
			Image: &dg.MessageEmbedImage{
				URL: "attachment://" + url,
			},
		},
		Files: []*dg.File{
			&dg.File{
				Name:   url,
				Reader: f,
			},
		},
	}
	s.ChannelMessageSendComplex(m.ChannelID, ms)
}

func DoesExist(ch chan string, url string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- "error"
		return
	}

	defer resp.Body.Close()

	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- "error"
		return
	}

	if string(read) == "404 page not found\n" {
		ch <- "error"
		return
	}

	ch <- "found"
}
