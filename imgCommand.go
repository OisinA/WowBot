package main

import (
	discord "github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
)

func ImgCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}
	url := "http://oisinaylward.me/imgs/" + m.Author.ID + "/" + message + ".png"
	channel := make(chan string)
	go DoesExist(channel, url)
	exists := <-channel
	if exists == "error" {
		s.ChannelMessageSend(m.ChannelID, "No image found.")
		return
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discord.MessageEmbed{
		Title: message,
		Image: &discord.MessageEmbedImage{
			URL: "http://oisinaylward.me/imgs/" + m.Author.ID + "/" + message + ".png",
		},
	})
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
