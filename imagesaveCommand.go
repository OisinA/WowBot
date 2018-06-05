package main

import (
	"io"
	"net/http"
	"os"
	"strings"

	discord "github.com/bwmarrin/discordgo"
)

var suffixes = []string{
	".png",
	".jpeg",
	".jpg",
	".gif",
}

func ImageSaveCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}

	split := strings.Split(message, " ")

	if len(split) < 2 {
		SendMessage(s, m.ChannelID, "Incorrect syntax.")
		return
	}

	url := split[0]
	name := strings.Join(split[1:], "_")

	channel := make(chan string)

	go saveImage(channel, url, name, m.Author.ID)

	SendMessage(s, m.ChannelID, <-channel)
}

func saveImage(ch chan string, url string, name string, userID string) {
	has_suffix := false
	for _, i := range suffixes {
		if strings.HasSuffix(url, i) {
			has_suffix = true
		}
	}

	if !has_suffix {
		ch <- "Not a valid image."
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- "There was an error fetching the url."
		return
	}

	defer resp.Body.Close()

	os.Mkdir("/home/WebServer/imgs/"+userID, os.ModePerm)

	file, err := os.Create("/home/WebServer/imgs/" + userID + "/" + name + ".png")
	if err != nil {
		ch <- "Couldn't save your image."
		return
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		ch <- "Error saving your image."
		return
	}

	file.Close()
	ch <- "Your image was saved."
}
