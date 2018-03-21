package main

import (
	discord "github.com/bwmarrin/discordgo"
	"golang.org/x/net/html"
	"net/http"
)

func ImgCommand(s *discord.Session, m *discord.MessageCreate, message string) {
	if m.Author.Bot {
		return
	}
	ch := make(chan string)
	go GetImage(ch)
	img := <-ch
	s.ChannelMessageSendEmbed(m.ChannelID, &discord.MessageEmbed{
		Title:       "Image",
		Description: message,
		Image: &discord.MessageEmbedImage{
			URL: img,
		},
	})
}

func GetImage(ch chan string) {
	resp, _ := http.Get("https://apod.nasa.gov/apod/astropix.html")
	z := html.NewTokenizer(resp.Body)
OUTERLOOP:
	for {
		iter := z.Next()
		switch {
		case iter == html.ErrorToken:
			break OUTERLOOP
		case iter == html.StartTagToken:
			t := z.Token()

			isImage := t.Data == "img"
			if isImage {
				for _, a := range t.Attr {
					if a.Key == "src" {
						ch <- "https://apod.nasa.gov/apod/" + a.Val
					}
				}
			}
		}
	}
	ch <- ""
}
